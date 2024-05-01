package main

import (
	"fmt"
	"sync"
)

/**
  临界区：程序中的一部分会被并发访问或者修改，为了避免并发访问导致未知的结果，
  这部分程序需要被保护起来，这部分程序就叫临界区
  CAS:将给定值与内存中的值进行比较，如果相等，将新值替换为内存中的值
  mutex不允许任何一个goroutine落下，永远没有机会获得锁，不抛弃，不放弃是他的宗旨，而且
  他也尽可能让等待更长的goroutine更有机会获得锁
*/
/**
  mutex的常见错误：
  1.mutex.lock()和mutex.Unlock()不是成对出现的
  2.copy已经使用的mutex
  3.mutex不是可重入锁，
*/

func main() {
	var cnt = 0
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mutex.Lock()
				cnt++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}
