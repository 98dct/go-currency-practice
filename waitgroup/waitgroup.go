package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mutex sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
}

func (c *Counter) Count() uint64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.count
}

func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	c.Incr()
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go worker(&counter, &wg)
	}

	wg.Wait()
	fmt.Println(counter.Count())
}
