package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChan1(t *testing.T) {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	ch4 := make(chan struct{})

	for i := 1; i <= 4; i++ {
		go func(i int) {
			for {
				select {
				case <-ch1:
					fmt.Println(1)
					time.Sleep(1 * time.Second)
					ch2 <- struct{}{}
				case <-ch2:
					fmt.Println(2)
					time.Sleep(1 * time.Second)
					ch3 <- struct{}{}
				case <-ch3:
					fmt.Println(3)
					time.Sleep(1 * time.Second)
					ch4 <- struct{}{}
				case <-ch4:
					fmt.Println(4)
					time.Sleep(1 * time.Second)
					ch1 <- struct{}{}
				}
			}

		}(i)
	}

	ch1 <- struct{}{}
	select {}
}
