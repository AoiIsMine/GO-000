## task：
基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。

## 大致思路
利用errgroup创建goroutine,一个跑http server,一个监听linux signal信号事件
main监听errgroup wait是否关闭

执行ctrl+c命令,让监听signal事件触发,close chan写入数据,并return nil(即通知errgroup此协程已处理完毕)  
http server收到close chan数据,立即执行关闭操作,return nil  
(http server goroutine关闭前,先处理一些业务代码,但是利用了context做了定时器,超时就不等业务代码处理完,立即return关闭http server)  
最终main,group.Wait()触发,实现整体退出
***
## 参考
[golang之信号处理(Signal)](https://zhuanlan.zhihu.com/p/128953024)
```go
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
```
***