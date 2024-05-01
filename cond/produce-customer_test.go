package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var cond sync.Cond

func produce(ch chan<- int, num int) {

	for {
		cond.L.Lock()
		//产片区满了，等到消费者消费
		for len(ch) == 3 {
			cond.Wait()
		}

		t := rand.Intn(100)
		ch <- t
		fmt.Printf("%dth producer produce t = %d,len(chan) = %d\n", num, t, len(ch))
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(1000 * time.Millisecond)
	}

}

func customer(ch <-chan int, num int) {
	for {
		cond.L.Lock()
		for len(ch) == 0 {
			cond.Wait()
		}

		t := <-ch
		fmt.Printf("%dth customer custom t = %d,len(chan) = %d\n", num, t, len(ch))
		cond.L.Unlock()
		cond.Signal()

		time.Sleep(500 * time.Millisecond)
	}
}

func TestCond(t *testing.T) {
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	quit := make(chan bool)
	//产品区 使用channel模拟
	product := make(chan int, 3)

	//创建互斥锁和条件变量
	cond.L = new(sync.Mutex)

	//5个消费者
	for i := 0; i < 5; i++ {
		go produce(product, i)
	}
	//3个生产者
	for i := 0; i < 3; i++ {
		go customer(product, i)
	}

	//主协程阻塞 不结束
	<-quit
}
