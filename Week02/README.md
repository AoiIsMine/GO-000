# 问题
比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
# 答
这里db.go代码里面进行了warp  

其实并不需要warp,按照分层的思想的话,dao不需要堆栈追踪  
(只要代码没问题,error都是db处理的error,所以上层根本不需要追踪到dao层)  
直接返回error(例如ErrNoRows)  
逻辑层有需要处理error,就warp然后,日志打印

***
# TODO，应该在哪里warp  
待回看视频


***
# 使用
* 修改db.go里面的mysql config,运行`./test`即可

***

# go error 小结
### 问题  
利用error处理错误场景,但是没有堆栈以及附加信息,查找处理不方便
### 解决  
可以用第三方库`github.com/pkg/errors`,主要接口:  
1. `errors.New(msg)`自带调用堆栈信息
2. `errors.WithMessage`
3. `errors.WithStack`
4. `errors.Wrap(err,msg)`附加堆栈加msg
5. `errors.Is(err,targetErr)`当warp后,无法再用==判断error相等,需要用`Is`

### error处理流程
<!-- 1. 底层,error生成处,`warp`1次即可  
(dao层的可以不用warp,直接返回error,因为都是db处理级别error,只要dao层逻辑代码没问题的话,这样堆栈根本不需要追踪到dao层)
2. 中间,不处理error的层,直接返回error即可
3. 上层,处理error后,打日志,不再返回error
> 或者, -->

### 其他小结
* `fmt.Printf("err = %+v", err) //%+v才有详细输出`

***
# 测试输出示例
```js
/**
1. 没有用errors warp输出示例:
mydb.Query start
query err =  sql: unknown driver "mysql" (forgotten import?)

mydb.Query start
query err =  Error 1054: Unknown column '？' in 'where clause'

mydb.Query start
query err = err no rows
query name =

mydb.Query start
query name = 测时名

2. 使用附带msg以及stack输出示例(在db query就warp)
mydb.Query start
no rows,err = sql: no rows in result set
query error
main.Query
	/home/xiao/job/project/gotest/task/week2-1/db.go:34
main.main
	/home/xiao/job/project/gotest/task/week2-1/main.go:11
runtime.main
	/usr/local/go/src/runtime/proc.go:204
runtime.goexit
	/usr/local/go/src/runtime/asm_amd64.s:1374

3. 在调用处warp后,打印error
no rows,err = sql: no rows in result set
111
main.main
	/home/xiao/job/project/gotest/task/week2-1/main.go:16
runtime.main
	/usr/local/go/src/runtime/proc.go:204
runtime.goexit
	/usr/local/go/src/runtime/asm_amd64.s:1374

*/
```