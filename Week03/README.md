## 题目：
基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。

## 大致做法
写一个http server,注册linux signal信号事件  
ctrl+c要做到,基于errgroup让所有goroutine退出
***

***
## 参考
[golang之信号处理(Signal)](https://zhuanlan.zhihu.com/p/128953024)

***



️以上作业，要求提交到Github上面，Week03作业提交地址：
https://github.com/Go-000/Go-000/issues/69

请务必按照示例格式进行提交，不要复制其他同学的格式，以免格式错误无法抓取作业。
﻿
️Github使用教程：https://u.geekbang.org/lesson/51?article=294701