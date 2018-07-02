package main

import (
	"fmt"
	"time"
)

type Result struct {
	Finish bool
}

func CallA(result chan Result, param interface{}) {
	rr := Result{}
	rr.Finish = true
	result <- rr
	fmt.Println("----Over A----")
}

func CallB(result chan Result, param interface{}) {
	rr := Result{}
	rr.Finish = true
	result <- rr
	fmt.Println("----Over B----")
}

func main() {
	resultA := make(chan Result, 1)
	resultB := make(chan Result, 1)
	param := 1
	overtime := false
	count := 0
	go CallA(resultA, param)
	go CallB(resultB, param)
	timer := time.NewTimer(1 * time.Second)
	for {
		select {
		case ar := <-resultA:
			count = count + 1
			fmt.Println(" ar.Finish ", ar.Finish)
		case br := <-resultB:
			count = count + 1
			fmt.Println(" br.Finish ", br.Finish)
		case <-timer.C:
			fmt.Println("---Over Time---")
			overtime = true
		}
		if overtime || count == 2 {
			break
		}
	}
	fmt.Println("---All Over---")
}
