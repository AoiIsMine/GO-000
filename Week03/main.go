package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/sync/errgroup"
)

type TestHandler struct {
}

func (th *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("my test handler server http")
}

/**
TODO
context:
本来是传递ctx,内部用withCancel返回的cancel,但是没有效果,看来是子级ctx.cancel不会影响父级别ctx.Done()
所以,直接传递ctx和cancel,但是,好像不应该这么使用,还传递cancel,待看视频后修改

*/
func main() {
	fmt.Println("main start")

	g := new(errgroup.Group)
	// ctx := context.Background()
	var ser1 *http.Server
	var ser2 *http.Server
	ctx, cancel := context.WithCancel(context.Background())
	g.Go(func() error {
		err := listenSignal(ctx, cancel)
		return err
	})
	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(w, "Hello, world!\n")
	// }
	//服务1
	g.Go(func() error {
		ser1 = &http.Server{
			Addr:    ":8888",
			Handler: &TestHandler{}, //结尾没,也有问题
		}
		err := serverStart(ctx, cancel, ser1)
		return err
	})
	//服务2
	g.Go(func() error {
		ser2 = &http.Server{
			Addr: ":9999",
			// Handler: helloHandler,	//err
			Handler: &TestHandler{},
		}
		err := serverStart(ctx, cancel, ser2)
		return err
	})
	//统一关闭ser的协程
	//TODO考虑,ser开启失败,这里强制关闭会有问题吗,先判断空值
	g.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("进入统一关闭ser")
			ser1.Shutdown(ctx)
			ser2.Shutdown(ctx)

			//其他关闭server前的数据处理,加超时强制退出
			// ctx1, cancel1 := context.WithTimeout(ctx, time.Second*5) //好像不能再用done了的ctx当父级,go没跑,ctx1.Done()直接进入了
			ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second*3)
			go func() {
				for i := 0; i < 5; i++ {
					time.Sleep(time.Second * 1)
					fmt.Println("close wait ,do someting step=", (i + 1), " all step=5")
				}
				defer cancel1()
			}()
			select {
			case <-ctx1.Done():
				fmt.Println("ctx1 done")
				return nil
			}

			// return nil
		}

	})

	if err := g.Wait(); err == nil {
		fmt.Println("main over")
	}

}

func listenSignal(ctx context.Context, cancel context.CancelFunc) error {
	// _, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal)
	signal.Notify(c)
	fmt.Println("listenSignal start...")

	select {
	case <-c:
		fmt.Println("listenSignal close start...")
		//收到信号,先不区分信号,统一视为关闭
		cancel()
		fmt.Println("listenSignal close...")
		return nil
	case <-ctx.Done():
		//收到关闭服务信号
		fmt.Println("listenSignal close...")
		return nil
	}
}

// hello world, the web server
func helloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
func serverStart(ctx context.Context, cancel context.CancelFunc, ser *http.Server) error {
	fmt.Println("http server start")
	// _, cancel := context.WithCancel(ctx)
	fmt.Println("111")
	err := ser.ListenAndServe() //这里会阻塞,明明不是chan
	fmt.Println("222")
	if err != nil {
		fmt.Println("http server start err", err)
		cancel()
		return nil
	}
	return nil
	// select {
	// case <-close:
	// 	//收到关闭服务信号
	// 	fmt.Println("http server close begin")

	// 	//可能需要做一些处理再关闭,但是需要设置定时器,超时强制关闭
	// 	//这里还未看看context 配合 errgroup的文档,先这样处理,TODO优化超时代码
	// 	//10s超时会触发ctx.Done(),cancel即取消定时器
	// 	//伪代码
	// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// 	done := make(chan int)
	// 	go func() {
	// 		for i := 0; i < 5; i++ {
	// 			time.Sleep(time.Second * 1)
	// 			fmt.Println("do someting", i)
	// 			if i == 5 {
	// 				done <- 1
	// 				cancel()
	// 			}
	// 		}
	// 	}()

	// 	select {
	// 	case <-done:
	// 		fmt.Println("http server close over")
	// 		return nil
	// 	case <-ctx.Done():
	// 		fmt.Println("http server timeout force close")
	// 		return nil
	// 	}

	// }

}

//废弃代码2
// package main

// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"time"

// 	"golang.org/x/sync/errgroup"
// )

// /**
// TODO
// 1.context需要结合errgroup使用吗
// 2.chan需要指针方式传入吗?
// 3.close chan会多写报错吗?(没取)
// 4.http server可能关闭需要执行其他方式?

// 5.ch只写入不读会阻塞吗(容量内)
// 6.http listen 会阻塞吗

// 7.版本2不行,单用一个done chan,如何关闭多个服务
// */
// func main() {
// 	fmt.Println("main start")

// 	g := new(errgroup.Group)
// 	close := make(chan int)

// 	g.Go(func() error {
// 		err := listenSignal(close)
// 		return err
// 	})
// 	g.Go(func() error {
// 		err := serverStart(close)
// 		return err
// 	})

// 	if err := g.Wait(); err == nil {
// 		fmt.Println("main over")
// 	}

// }

// func listenSignal(close chan int) error {
// 	c := make(chan os.Signal)
// 	signal.Notify(c)
// 	fmt.Println("listenSignal start...")

// 	select {
// 	case <-c:
// 		fmt.Println("listenSignal close start...")
// 		//收到信号,先不区分信号,统一视为关闭
// 		close <- 1
// 		// serverExit(close)
// 		fmt.Println("listenSignal close...")
// 		return nil
// 	case <-close:
// 		//收到关闭服务信号
// 		fmt.Println("listenSignal close...")
// 		return nil
// 	}

// 	// switch s := <-c; s {
// 	// case syscall.SIGINT, syscall.SIGTERM:
// 	// 	fmt.Println("receive close signal =", s)
// 	// 	serverExit(close)
// 	// 	// return errors.New("receive close signal")
// 	// 	return nil
// 	// default:
// 	// 	//TODO,会导致signal goroutine失效
// 	// 	fmt.Println("unknow signal =", s)
// 	// 	return nil
// 	// }
// }

// func serverExit(close chan int) {
// 	close <- 1
// }

// // hello world, the web server
// func helloServer(w http.ResponseWriter, req *http.Request) {
// 	io.WriteString(w, "hello, world!\n")
// }
// func serverStart(close chan int) error {
// 	fmt.Println("111")
// 	http.HandleFunc("/hello", helloServer)
// 	fmt.Println("222")
// 	err := http.ListenAndServe(":8888", nil) //这里会阻塞,明明不是chan
// 	fmt.Println("333")
// 	if err != nil {
// 		fmt.Println("http server start err", err)
// 		//TODO启动err
// 		serverExit(close)
// 		return nil
// 	}
// 	fmt.Println("http server start")
// 	select {
// 	case <-close:
// 		//收到关闭服务信号
// 		fmt.Println("http server close begin")

// 		//可能需要做一些处理再关闭,但是需要设置定时器,超时强制关闭
// 		//这里还未看看context 配合 errgroup的文档,先这样处理,TODO优化超时代码
// 		//10s超时会触发ctx.Done(),cancel即取消定时器
// 		//伪代码
// 		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
// 		done := make(chan int)
// 		go func() {
// 			for i := 0; i < 5; i++ {
// 				time.Sleep(time.Second * 1)
// 				fmt.Println("do someting", i)
// 				if i == 5 {
// 					done <- 1
// 					cancel()
// 				}
// 			}
// 		}()

// 		select {
// 		case <-done:
// 			fmt.Println("http server close over")
// 			return nil
// 		case <-ctx.Done():
// 			fmt.Println("http server timeout force close")
// 			return nil
// 		}

// 	}

// }

//废弃代码1
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
