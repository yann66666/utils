// Author: yann
// Date: 2020/5/25 6:26 下午
// Desc: or_donw 模式, 利用反射实现.

package sync_util

import "reflect"

func OrDone(channels ...<-chan interface{}) <-chan interface{} {
	//特殊情况，只有0个或者1个
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	done := make(chan interface{})
	go func() {
		defer close(done)
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		reflect.Select(cases)
	}()

	return done
}
