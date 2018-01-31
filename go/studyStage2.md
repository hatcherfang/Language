下周任务:   
1. 好好理解goroutine, 以及如何同步   
2. 理解reflect   
3. 试用IDE, 看看哪个好用   
推荐的试用IDE有: Eclipse, IntelliJ IDEA, Visual Studio Code, GoWorks   
文档我放在https://wiki.jarvis.trendmicro.com/display/DDWI/golang_study   

另外, 第3周开始, 我们会开始准备一个feature, 我准备开始做一个类似iwsaas的auth daemon, 使用go实现, 因此, 需要一部分人参与进来. 这部分工作量会比较大, 需要占用比较多的课余时间. 另外, 我们还要准备go study的ppt, 作为我们study group的输出, 这部分的工作量比较小. 请大家凭自己的兴趣和时间报名这两项任务的一项.   

## goroutine  
eg:  
gocode/src/withGoroutine.go  
gocode/src/withoutGoroutine.go  

### IDE 试用  
### 理解reflect  
1. python set string as function by using `eval`  
2. go and python can also be conpiled to source file  
3. go reflect is same with `eval` but reflect is more powerful. 
- For example, It can get the string define from a number.  
4. pstack tool to check status of goroutine.   
5. coroutine example:    
- scanner/iws_include/task   
- scanner/lib_daemon/daemonbase/task   
- TmTaskBase.cpp, TmHandleTaskThread.cpp, TmHandleTaskIOThread.cpp  
6. why use coroutine?  
- coroutine is easier than callback function for engineer     

1.基本背景介绍, 关键字, 运算符和优先级, 表达式, 基本数据类型
2.定义: 变量, 常量, 类型, 函数等
3.语句: 判断, 循环, switch, select等
4.复杂数据类型: slice, map, chan, String等
5.复杂数据类型: struct, interface
6.预定义函数和常量的基本使用, 以及基本的库方法
7.GoRoutine原理和使用
8.reflect的使用

#### Reference  
- [Win+vsCode+golang 配置](http://www.jianshu.com/p/b8810bbbd068)  
- [VSCode+Golang开发环境](http://dreamlikes.cn/archives/717)  
- [Visual Studio Code中配置GO开发环境 ](https://studygolang.com/articles/6590)  
