package main

import (
	"fmt"
)

func main() {
	close := make(chan int, 2)
	done := make(chan int)
	go listenSignal(close, done)
	select {
	case <-done:
		fmt.Println("33")
	}
}
func listenSignal(close, done chan int) {
	fmt.Println("111")
	close <- 1 //这个的问题,阻塞了,导致没有关掉
	fmt.Println("222")
	done <- 1
}
