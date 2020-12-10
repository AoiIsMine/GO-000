package main

import (
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func main() {
	g := new(errgroup.Group)
	close := make(chan int)

	g.Go(func() error {
		err := listenSignal(close)
		return err
	})

	err := g.Wait()
	fmt.Println("111")
	if err == nil {
		fmt.Println("main over")
	} else {
		fmt.Println("222 err=", err)
	}

	// close := make(chan int)
	// go listenSignal(close)
	// select {
	// case <-close:
	// 	fmt.Println("main over")
	// }
}

func listenSignal(close chan int) error {
	c := make(chan os.Signal)
	signal.Notify(c)
	fmt.Println("listenSignal start...")

	select {
	case <-c:
		//收到信号,先不区分信号,统一视为关闭
		// serverExit(close)
		fmt.Println("listenSignal close...")
		// return nil
		// case <-close:
		// 	//收到关闭服务信号
		// 	fmt.Println("listenSignal close...")
		// 	return nil

		//有的说:chan的内部实现已经类似指针了。chan即使作为值传递拷贝一份给别人，两个chan还是会“指向”同一个物.
		close <- 1 //这个的问题,阻塞了,导致没有关掉
		//传入的是副本,??

		return nil //都有关掉
		// return errors.New("hahaha err") //都有关掉
	}

	// switch s := <-c; s {
	// case syscall.SIGINT, syscall.SIGTERM:
	// 	fmt.Println("receive close signal =", s)
	// 	// serverExit(close)
	// 	// return errors.New("receive close signal")
	// 	// return nil
	// default:
	// 	//TODO,会导致signal goroutine失效
	// 	fmt.Println("unknow signal =", s)
	// 	// return nil
	// }
}
