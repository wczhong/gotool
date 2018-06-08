package main

import (
	"fmt"
	"sync"
)

var zf []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func PrintZF(zfc chan int, nc chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 0; i < 26; i++ {
		select {
		case <-zfc:
			fmt.Println(zf[i])
			nc <- 1
		}
	}

}

func PrintN(zfc chan int, nc chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 26; i++ {
		select {
		case <-nc:
			fmt.Println(i)
			zfc <- 1
		}
	}

}

func main() {
	wg := &sync.WaitGroup{}
	// 1 防止: 两个协程，其中一个结束，另一个扔向通道中写入数据，导致死锁
	zfc := make(chan int, 1)
	nc := make(chan int, 1)
	wg.Add(1)
	go PrintZF(zfc, nc, wg)
	wg.Add(1)
	go PrintN(zfc, nc, wg)
	zfc <- 1
	wg.Wait()
	close(zfc)
	close(nc)
}
