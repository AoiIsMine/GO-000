package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"

	"golang.org/x/sync/errgroup"
)

/**
context需要结合errgroup使用吗

*/
func main() {
	fmt.Println("main start")

	g := new(errgroup.Group)
	close := make(chan int)
	ctx := context.Background()

	g.Go(func() error {
		err := listenSignal(close)
		return err
	})
	g.Go(func() error {
		err := serverStart(close)
		return err
	})

	if err := g.Wait(); err == nil {
		fmt.Println("mian over")
	}

	//TODO,附加上,关闭后处理任务,超时情况,强制关闭
	// //10s超时会触发ctx.Done(),cancel即取消定时器
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	// go listenSignal(close)
	// <-close
	fmt.Println("main over")
}

func listenSignal(close chan int) error {
	c := make(chan os.Signal)
	signal.Notify(c)
	fmt.Println("listenSignal start...")

	switch s := <-c; s {
	case syscall.SIGINT, syscall.SIGTERM:
		fmt.Println("receive close signal =", s)
		serverExit(close)
		return errors.New("receive close signal")
	default:
		//TODO,会导致signal goroutine失效
		fmt.Println("unknow signal =", s)
		return nil
	}
}

func serverExit(close chan int) {
	close <- 1
}

//超时 级联退出,
func exitTimer() {

}

func serverStart(close chan int) error {
	// // http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	// httpServer := http.Server{
	// 	Addr:    ":8888",
	// 	Handler: http.DefaultServeMux,
	// }

	// //err:too many arguments to conversion to http.HandlerFunc
	// http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("receive http request")
	// })
	// http.HandlerFunc("/", func() {
	// 	fmt.Println("receive http request")
	// })
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "just another http server...")
	})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("http server start err", err)
		//TODO启动err
		close <- 1
		return errors.New("http server start err")
	}

	fmt.Println("http server start")
	<-close
	return errors.New("server close")
}
func serverHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("receive http request")
}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"golang.org/x/sync/errgroup"
// )

// func main() {
// 	fmt.Println("main start")

// 	g := new(errgroup.Group)

// 	g.Go(func() error {
// 		listenSignal(close)

// 	})

// 	ctx := context.Background()

// 	// http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {})
// 	httpServer := http.Server{
// 		Addr:    ":8888",
// 		Handler: http.DefaultServeMux,
// 	}

// 	//10s超时会触发ctx.Done(),cancel即取消定时器
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

// 	close := make(chan int)
// 	go listenSignal(close)
// 	<-close
// 	fmt.Println("main over")
// }

// func listenSignal(close chan int) {
// 	c := make(chan os.Signal)
// 	signal.Notify(c)
// 	fmt.Println("listenSignal start...")

// 	switch s := <-c; s {
// 	case syscall.SIGINT, syscall.SIGTERM:
// 		fmt.Println("receive close signal =", s)
// 		serverExit(close)
// 	default:
// 		fmt.Println("unknow signal =", s)
// 	}

// 	// s := <-c
// 	// //ctrl+c s=interrupt
// 	// //kill -9不会输出下面的打印,但是打印Killed
// 	// //kill 不会输出下面的打印,但是打印Terminated
// 	// fmt.Println("end... s=", s)
// }

// func serverExit(close chan int) {
// 	close <- 1
// }

// //超时 级联退出,
// func exitTimer() {

// }

// func serverStart(close chan int) {
// 	// //err:too many arguments to conversion to http.HandlerFunc
// 	// http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	fmt.Println("receive http request")
// 	// })
// 	// http.HandlerFunc("/", func() {
// 	// 	fmt.Println("receive http request")
// 	// })
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "just another http server...")
// 	})

// 	err := http.ListenAndServe(":8888", nil)
// 	if err != nil {
// 		fmt.Println("http server start err", err)
// 		//TODO启动err
// 		close <- 1
// 		return
// 	}

// 	fmt.Println("http server start")
// 	<-close

// }
// func serverHandler(rw http.ResponseWriter, req *http.Request) {
// 	fmt.Println("receive http request")
// }

// /**
// http 要 defer close?
// goroutine的recover?

// 开启信号监听
// 开启http server
// 关闭,触发新号,所有关闭
// http server拖时,触发超时强制

// */
// /**
// 涉及
// goroutine
// ch
// sync.WatiGroup

// sync.errgroup(需要安装`go get -u golang.org/x/sync`)
// https://pkg.go.dev/golang.org/x/sync@v0.0.0-20201207232520-09787c993a3a/errgroup

// context.WithTimeout
// https://golang.google.cn/pkg/context/#pkg-examples
// https://blog.csdn.net/yzf279533105/article/details/107292247

// os signal
// https://zhuanlan.zhihu.com/p/128953024

// net http
// https://go-zh.org/pkg/net/http/#pkg-overview
// https://studygolang.com/articles/15826?utm_medium=referral
// */
