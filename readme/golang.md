delve工具使用参考。下原文，包含了 dlv test等命令

**1、dlv attach**
这个相当于gdb -p 或者 gdb attach ，即跟踪一个正在运行的程序。这中用法也是很常见，对于一个后台程序，它已经运行很久了，此时你需要查看程序内部的一些状态，只能借助attach.

dlv attach $PID  ## 后面的进程的ID
1
**2、dlv debug**
dlv debug main/hundredwar.go ## 先编译，后启动调试 
1
**3、dlv exec**
dlv exec ./HundredsServer  ## 直接启动调试

dlv exec ./HundredsServer -- -port 8888 -c /home/config.xml  ## 后面加参数启动调试
1
2
3
与dlv debug区别就是，dlv debug 编译一个临时可执行文件，然后启动调试，类似与go run。

**4、dlv core**
用来调试core文件，但是想让go程序生成core文件，需要配置一个环境变量，默认go程序不会产生core文件。

export GOTRACEBACK=crash
1
只有定义这个环境变量，go程序才会coredump。如果在协程入口定义defer函数，然后recover也不会产生core文件。

go func() {
​		defer func() {
​			if r := recover(); r != nil {
​                fmt.Printf("panic error\n") 
​			}
​		}()
​		var p *int = nil
​		fmt.Printf("p=%d\n", *p) // 访问nil指责
​	}()  // 这个协程不会生产core文件

因为recover的作用就是捕获异常，然后进行错误处理，所以不会产生coredump，这个需要注意。这个也是golang的一大特色吧，捕获异常，避免程序coredump。

调试coredump文件
关于调试core文件，其实和C/C++差不多，最后都是找到发生的函数帧，定位到具体某一行代码。但是golang稍有不同，对于golang的core文件需要先定位到时哪一个goroutine发生了异常。

dlv core ./Server core.26945  ## 启动
————————————————
版权声明：本文为CSDN博主「KentZhang_」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/KentZhang_/java/article/details/84925878



#### go test文件的调试。 

直接dlv test 然后break 需要的函数就进入到函数调试( dlv test会调用全部函数)

b TestFunc也可以

如果需要单独调试某几个文件 

```
dlv test c7-reverse_test.go c7-reverse.go

go run 也是同理 也要家两个文件同时编译

locals局部变量
args 参数
```

 

golang vim配置... 有空再看

<https://tpaschalis.github.io/vim-go-setup/>



<https://octetz.com/docs/2019/2019-04-24-vim-as-a-go-ide/>





