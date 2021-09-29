// Author: yann
// Date: 2021/9/27 6:26 下午
// Desc: Once 封装 sync.Once用于接收带有返回错误参数的once.Do方法

package sync_util

import (
	"sync"
	"sync/atomic"
)

// Once 封装 sync.Once用于接收带有返回错误参数的once.Do方法
type Once struct {
	done uint32
	lock sync.Mutex
}

func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 0 { //如果为0 则执行 否则直接返回nil
		return o.doSlow(f)
	}
	return nil
}
func (o *Once) doSlow(f func() error) (err error) {
	o.lock.Lock()
	defer o.lock.Unlock() //预防并发,加锁
	if o.done == 0 {      //再次判断是否已经执行过
		err = f() //判断返回值是否为空, 为空则设置已经执行, 否则返回错误, 这样可以再次执行do方法
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
		return err
	}
	return nil
}
