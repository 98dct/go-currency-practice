package main

import "fmt"

/**
  执行业务处理的goroutine不要通过共享内存的方式通信，
  而是通过channel通信的方式共享数据
  v,ok := <- ch 如果ok是false代表，chan已经关闭，而且缓冲区没有数据，此时第一个值是零值
  channel panic的情景：
  1.close为nil的chan
  2.send已经关闭的chan
  3.close已经关闭的chan
*/

func main() {

	var ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	fmt.Println(len(ch))
	fmt.Println(cap(ch))

	for v := range ch {
		fmt.Println(v)
		fmt.Println("len:", len(ch))
		fmt.Println("cap:", cap(ch))
	}

}
