### 函数  
- 前面带有减号(-) 的方法为实例方法，必须使用类的实例才可以调用的。对应的有+号， 代表是类的静态方法，不需要实例化即可调用  
eg:  
```
-(void) HelloWorld:(BOOL)ishelloworld{  
//干点啥  
 } 
```
### 消息  
消息的定义：向对象发送信息。  
消息是ios的运行时环境特有的机制。 和C++,Java下的类，或实例调用类或实例的方法类似。我这说的是类似，他们的机制实际上是有很大的差别。  
### Import  
例子：  
import "Class.h"  
import <Class.h>  
import <director/Class.h>  
这个和C++里的include ,java的import类似  
`#import`可保证头文件只被包含一次，而不论次命令实际上在那个文件中出现了多少次  
### Property 和Synthesize  
Property定义：@property 声明用于自动创建property属性变量的getter和setter  
Synthesize定义：@Synthesize声明实现了property属性变量的getter和setter。  

例子:  
在  interface：@property dataType variableName  
在  implementation:  synthesiz variableName  
### 头文件中的方法  
例子：  
-(returnType)method  
-(returnType)method:(dataType)param1  
-(returnType)method:(dataType)param1 withParam:(dataType)param2  

类似于:  
C/C++/Java  
returnType method()  
returnType method(param1)  
returnType method(param1,param2)  

### self  
指向自己的指针  
[self method]  
类似于：c++/java  
this.method();  

### 继承关系和接口实现  
例子：  
ClassA:ParentA  
ClassA:ParentA<Protocol>  
ClassA <Protocol>  

类似于:  
java:  
ClassA extends ParentA  
ClassA extends ParentA implements interface  
ClassA implements interface  
objective-c的 Protocol和c++、java的接口类似。  

### 空指针  
id obj = nil;  
NSString *hello = nil;  
nil相当与Java中的null;  
### id  
objective-c的和C++里的(void*)类似  
PS：Objective-C和Java一样，都有运行时环境，有内省的能力。Objective-C和java有很多不同的地方，在iOS系统里，Objective-C的内存需要自己管理，添加了ARC机制后编译器帮助了Objective-C  添加release释放的代码。而Java是通过垃圾回收器管理内存的。  
### Reference  
[十分钟让你明白Objective-C的语法（和Java、C++的对比）](http://blog.csdn.net/totogo2010/article/details/7632384)
[Objective-C 语法快速参考](http://blog.jobbole.com/66202/)  
