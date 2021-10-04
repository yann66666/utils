// Author: yann
// Date: 2020/5/25 6:26 下午
// Desc: Mutex封装, 实现TryLock方法

package sync_util

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// Mutex定义的常量
const (
	mutexLocked   = 1 << iota // 加锁标识位置
	mutexWoken                // 唤醒标识位置
	mutexStarving             // 锁饥饿标识位置
)

type Mutex struct {
	sync.Mutex
}

// 尝试获取锁
func (m *Mutex) TryLock() bool {
	// 如果能成功抢到锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	// 如果锁的原始值包含下面三位,则直接返回false
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}

	// 尝试请求锁
	current := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, current)
}
