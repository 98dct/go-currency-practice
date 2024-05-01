package main

import (
	"github.com/petermattis/goid"
	"sync"
	"sync/atomic"
)

/**
  基于mutex封装的可重入锁
  可以解决代码重入和递归调用带来的死锁问题
*/

type ReentrantMutex struct {
	sync.Mutex
	Owner     int64 //当前持有锁的goroutine id
	Recursion int64 //goroutine可重入次数
}

func (r *ReentrantMutex) Lock() {
	gid := goid.Get()

	if atomic.LoadInt64(&r.Owner) == gid { //说明是重入
		r.Recursion++
		return
	}
	r.Mutex.Lock()
	//goroutine第一次调用记录下goroutine id 调用次数加1
	atomic.StoreInt64(&r.Owner, gid)
	r.Recursion = 1
}

func (r *ReentrantMutex) Unlock() {
	gid := goid.Get()
	//非持有锁的goroutine 不能释放锁
	if atomic.LoadInt64(&r.Owner) != gid {
		panic("wrong onwer!")
	}
	r.Recursion--
	if r.Recursion != 0 {
		return
	}

	atomic.StoreInt64(&r.Owner, -1)
	r.Mutex.Unlock()
}
