package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

/**
  cond.wait()方法之前一定要加锁，cond.Wait()会释放锁
  被唤醒不等于检查条件被满足
  cond.Wait()放在for循环里，因为存在虚假唤醒的情况
*/

func main() {

	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {

			time.Sleep(time.Duration(rand.Int63n(10)+1) * 500 * time.Millisecond)

			//加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()
			log.Printf("运动员%d 已准备就绪", i)
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		c.Wait()
		//log.Printf("裁判员被唤醒一次,ready: %d", ready)
	}

	c.L.Unlock()

	log.Printf("所有运动员已就绪,ready: %d", ready)

}
