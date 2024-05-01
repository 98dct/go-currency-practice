package main

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked       = 1 << iota //加锁标识位置
	mutexWoken                    //唤醒标识位置
	mutexStarving                 //锁饥饿标识位置
	mutexWaitingShift             //标识waiter的起始标志位
)

type Mutex struct {
	sync.Mutex
}

func (m *Mutex) TryLock() bool {
	//如果能成功抢到锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexWoken|mutexLocked|mutexStarving) != 0 {
		return false
	}

	//尝试在竞争条件下加锁
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}
