package main

import (
	"sync/atomic"
	"unsafe"
)

type LKQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

type node struct {
	value interface{}
	next  unsafe.Pointer
}

func NewLKQueue() *LKQueue {
	n := unsafe.Pointer(&node{})
	return &LKQueue{
		head: n,
		tail: n,
	}
}

// 入队
func (q *LKQueue) Enqueue(v interface{}) {
	n := &node{
		value: v,
	}
	for {
		tail := load(&q.tail)
		next := load(&tail.next)
		if tail == load(&q.tail) {
			if next == nil { //没有元素入队
				if cas(&tail.next, next, n) { //增加到队尾
					cas(&q.tail, tail, n) //入队成功，移动尾巴指针
					return
				}
			} else {
				cas(&q.tail, tail, next)
			}
		}
	}
}

// 出队，没有元素返回nil
func (q *LKQueue) Dequeue() interface{} {
	for {
		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&head.next)
		if head == load(&q.head) {
			if head == tail {
				if next == nil {
					return nil
				}
				//只是尾指针还没来得及调整，尝试调整它指向下一个
				cas(&q.tail, tail, next)
			} else {
				//读取队列的数据
				v := next.value
				if cas(&q.head, head, next) {
					return v
				}
			}
		}

	}
}

// 将unsafe.Ponter原子加载为*node
func load(p *unsafe.Pointer) *node {
	return (*node)(atomic.LoadPointer(p))
}

// 封装cas，避免直接将*node转换为unsafe.Pointer
func cas(p *unsafe.Pointer, old, new *node) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
