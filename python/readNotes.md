## 《python源码剖析》笔记    
#### python List  
  
#### python Dict    
python的实现是采用散列表(hash table)。解决冲突的方式是开放定址法。    
在删除元素时时实行**伪删除(dummy态)**操作，因为这样起到承前启后的作用，为了元素正确查找。    
同时把伪删除的entry 放在缓存池中，有新元素插入的时候，首先会查找缓存池，如果有空余会预先在缓存池中插入新元素。    
Dict是类似C++ STL中map(实现基于红黑树)的一种关联容器。    
而关联容器中一个(键,值)元素对称为一个entry或slot, 记为PyDictObject。    
```  
typedef struct {  
    Py_ssize_t me_hash; /*cached hash code of me_key*/  
    PyObject *me_key;  
    PyObject *me_value;  
} PyDictEntry;  
```  
PyDictObject中其实存放的都是PyObject*。**所以dict类型的变量在python中进行变量传递的时候，如果改变dict内的元素值，那么传递前的dict内的元素也会改变，因为这种传递类似引用传递**在PyDictEntry中, me_hash域存储的是me_key的散列值，利用    
一个域来记录这个散列值可以避免每次查询的时候都要重新计算一遍散列值。      
在python中，关联容器是通过PyDictObject对象来实现的。而一个PyDictObject对象实际上是一大堆entry的集合。    
在创建Dict时，会被分配一个entry个数为8的数组，这个值是经过大量实验得出的最佳值。它既不会太浪费内存空间，又能很好满足Python内部大量使用    
PyDictObject的环境的要求，不需要在使用的过程中再次调用malloc申请内存空间。    
当table的装载率大于等于2/3时，table的大小需要变动。规则如下：    
```  
mp->ma_used*(mp->ma_used>50000 ? 2 : 4)  
```  
即entry used态大于50000创建新的table的大小按照2倍来扩充，否则按照4倍来扩充。    
调整table大小也会向size 变小的方向调整。    
当将旧table数据copy到新table的时候，会忽略dummy状态的entry，因为新table中不需要伪删除来查找元素。copy完之后会把旧table空间释放。    
  
## [程序员必知的Python陷阱与缺陷列表](http://www.cnblogs.com/xybaby/p/7183854.html)    
我个人对陷阱的定义是这样的：代码看起来可以工作，但不是以你“想当然”的方式。  
如果一段代码直接出错，抛出了异常，我不认为这是陷阱。  
比如，Python程序员应该都遇到过的“UnboundLocalError”, 示例：  
```  
>>> a=1  
>>> def func():  
...     a+=1  
...     print a  
...  
>>> func()  
Traceback (most recent call last):  
File "<stdin>", line 1, in <module>  
File "<stdin>", line 2, in func  
UnboundLocalError: local variable 'a'   
referenced before assignment  
```  
对于“UnboundLocalError”,还有更高级的版本：    
```  
import random  
  
def func(ok):  
    if ok:  
        a = random.random()  
    else:  
        import random  
        a = random.randint(1, 10)  
    return a  
  
func(True)  
  
# UnboundLocalError: local variable   
# 'random' referenced before assignment  
```  
可能对于很多Python新手来说，这个Error让人摸不着头脑。但我认为这不算陷阱，  
因为这段代码一定会报错，而不是默默的以错误的方式运行。  
不怕真小人，就怕伪君子。我认为缺陷就好比伪君子。  
  
那么Python中哪些真正算得上陷阱呢？  
  
#### 第一：以mutable对象作为默认参数    
这个估计是最广为人知的了，Python和其他很多语言一样，提供了默认参数，默认参数确实是个好东西，  
可以让函数调用者忽略一些细节（比如GUI编程，Tkinter，QT），  
对于lambda表达式也非常有用。但是如果使用了可变对象作为默认参数，那么事情就不那么愉快了。  
```  
>>> def f(lst = []):  
...     lst.append(1)  
...     return lst  
...  
>>> f()  
[1]  
>>> f()  
[1, 1]  
```  
惊喜不惊喜？！究其原因，Python中一切都是对象，函数也不列外，默认参数只是函数的一个属性。  
而默认参数在函数定义的时候已经求值了。  
```  
Default parameter values are evaluated when the function definition is executed.  
```  
stackoverflow上有一个更适当的例子来说明默认参数是在定义的时候求值，而不是调用的时候。  
```  
>>> import time  
>>> def report(when=time.time()):  
...     return when  
...  
>>> report()  
1500113234.487932  
>>> report()  
1500113234.487932  
```  
Python docoment 给出了标准的解决办法：    
```  
A way around this is to use None as the default, and explicitly test for it in the body of the function.  
```  
```  
>>> def report(when=None):  
...     if when is None:  
...        when = time.time()  
...     return when  
...  
>>> report()  
1500113446.746997  
>>> report()  
1500113448.552873  
```  
#### 第二: x += y vs x = x + y(hatcher: python 2.7 验证失败)    
一般来说，二者是等价的，至少看起来是等价的（这也是陷阱的定义 — 看起来都OK，但不一定正确）。    
```  
>>> x=1;x += 1;print x  
2   
>>> x=1;x = x+1;print x  
2  
>>> x=[1];x+=[2];print x  
[1, 2]  
>>> x=[1];x=x+[2];print x  
[1, 2]  
```  
呃，被光速打脸了？    
```  
>>> x=[1];print id(x);x=x+[2];print id(x)   
4357132800  
4357132728  
>>> x=[1];print id(x);x+=[2];print id(x)  
4357132800  
4357132800  
```  
前者x指向一个新的对象，后者x在原来的对象是修改，当然，那种效果是正确的取决于应用场景。  
至少，得知道，二者有时候并不一样。  
#### 第三，神奇的小括号–()    
小括号（parenthese）在各种编程语言中都有广泛的应用，Python中，小括号还能表示元组（tuple）这一数据类型,   
元组是immutable的序列。  
```  
>>> a = (1, 2)  
>>> type(a)  
<type 'tuple'>  
>>> type(())  
<type 'tuple'>  
```  
但如果只有一个元素呢？    
```  
>>> a=(1)  
>>> type(a)  
<type 'int'>  
```  
神奇不神奇，如果要表示只有一个元素的元组，正确的姿势是:    
```  
>>> a=(1,)  
>>> type(a)   
<type 'tuple'>  
```  
#### 第四：生成一个元素是列表的列表    
这个有点像二维数组，当然生成一个元素是字典的列表也是可以的，更通俗的说，生成一个元素是可变对象的序列。  
  
很简单嘛：  
```  
>>> a= [[]] * 10  
>>> a  
[[], [], [], [], [], [], [], [], [], []]  
>>> a[0].append(10)  
>>> a[0]   
[10]  
```  
看起来很不错，简单明了，but    
```  
>>> a[1]  
[10]  
>>> a  
[[10], [10], [10], [10], [10],   
[10], [10], [10], [10], [10]]  
```  
我猜，这应该不是你预期的结果吧，究其原因，还是因为Python中list是可变对象，上述的写法大家都指向的同一个可变对象，正确的姿势：    
```  
>>> a = [[] for _ in xrange(10)]  
>>> a[0].append(10)  
>>> a  
[[10], [], [], [], [], [], [], [], [], []]  
```  
#### 第五，在访问列表的时候，修改列表    
列表（list）在python中使用非常广泛，当然经常会在访问列表的时候增加或者删除一些元素。  
比如，下面这个函数，试图删掉列表中为3的倍数的元素：  
```  
>>> def modify_lst(lst):  
...     for idx, elem in enumerate(lst):  
...         if elem % 3 == 0:  
...            del lst[idx]  
```  
测试一下，    
```  
>>> lst = [1,2,3,4,5,6]  
>>> modify_lst(lst)  
>>> lst  
[1, 2, 4, 5]  
```  
好像没什么错，不过这只是运气好    
```  
>>> lst = [1,2,3,6,5,4]  
>>> modify_lst(lst)  
>>> lst  
[1, 2, 6, 5, 4]  
```  
上面的例子中，6这个元素就没有被删除。如果在modify_lst函数中print idx， item就可以发现端倪：lst在变短，但idx是递增的，  
所以在上面出错的例子中，当3被删除之后，6变成了lst的第2个元素（从0开始）。  
在C++中，如果遍历容器的时候用迭代器删除元素，也会有同样的问题。  
  
如果逻辑比较简单，使用list comprehension是不错的主意。  
####  第六，闭包与lambda    
这个也是老生长谈的例子，在其他语言也有类似的情况。先看一个例子:    
```  
>>> def create_multipliers():  
...     return [lambda x:i*x for i in range(5)]  
...  
>>> for multiplier in create_multipliers():  
...     print multiplier(2)  
```  
create_multipliers函数的返回值时一个列表，列表的每一个元素都是一个函数 －－ 将输入参数x乘以一个倍数i的函数。  
预期的结果时0，2，4，6，8. 但结果是5个8，意外不意外。  
  
由于出现这个陷阱的时候经常使用了lambda，所以可能会认为是lambda的问题，但lambda表示不愿意背这个锅。  
问题的本质在与Python中的属性查找规则，LEGB（local，enclousing，global，bulitin），在上面的例子中，  
i就是在闭包作用域（enclousing），而Python的闭包是延迟绑定 ， 这意味着闭包中用到的变量的值，是在内部函数被调用时查询得到的。  
  
解决办法也很简单，那就是变闭包作用域为局部作用域。  
```  
>>> def create_multipliers():  
...     return [lambda x, i = i:i*x for i in range(5)]  
```  
#### 第七，定义__del__    
大多数计算机专业的同学可能都是先学的C、C++，构造、析构函数的概念应该都非常熟。于是，当切换到Python的时候，自然也想知道有没有相应的函数。  
比如，在C++中非常有名的RAII，即通过构造、析构来管理资源（如内存、文件描述符）的声明周期。  
那在Python中要达到同样的效果怎么做呢，即需要找到一个对象在销毁的时候一定会调用的函数，  
于是发现了__init__, __del__函数，可能简单写了两个例子发现确实也能工作。  
但事实上可能掉进了一个陷阱，在Python documnet是有描述的：  
```  
Circular references which are garbage are detected when the option cycle detector is enabled (it’s on by default),   
but can only be cleaned up if there are no Python-level __del__() methods involved.  
```  
简单来说，如果在循环引用中的对象定义了__del__，那么Python GC不能进行回收，因此，存在内存泄漏的风险    
#### 简单来说，如果在循环引用中的对象定义了__del__，那么Python GC不能进行回收，因此，存在内存泄漏的风险    
示例在stackoverflow的例子上稍作修改，假设现在有一个package叫mypackage，里面包含三个Python文件：mymodule.py, main.py, __init__.py。  
mymodule.py代码如下：  
```  
l = []  
class A(object):  
    pass  
...  
```  
main.py代码如下：    
```  
def add(x):  
    from mypackage import mymodule  
    mymodule.l.append(x)  
    print "updated list", mymodule.l, id(mymodule)  
  
def get():  
    import mymodule  
    print 'module in get', id(mymodule)  
    return mymodule.l  
  
if __name__ == '__main__':  
    import sys  
  
    sys.path.append('../')  
    add(1)  
  
    ret = get()  
    print "lets check", ret  
```  
运行 main.py，结果如下：    
```  
updated list [1] 4406700752  
module in get 4406700920  
lets check []  
```  
从运行结果可以看到，在add 和 get函数中import的mymodule不是同一个module，ID不同。  
当然，在Python2.7中，需要main.py的sys.path.append('../') 才能出现这样的效果。  
你可能会问，谁会写出这样的代码呢？事实上，在很多项目中，为了import的时候方便，  
会往sys.path加入一堆路径。那么在项目中，大家统一一种import方式就非常有必要了。  
#### 第九，Python升级    
Python3.x并不向后兼容，所以如果从2.x升级到3.x的时候得小心了，下面列举两点：  
  
在Python2.7中，range的返回值是一个列表；而在Python3.x中，返回的是一个range对象。  
  
map()、filter()、 dict.items()在Python2.7返回列表，而在3.x中返回迭代器。当然迭代器大多数都是比较好的选择，  
更加pythonic，但是也有缺点，就是只能遍历一次。在instagram的分享中，也提到因为这个导致的一个坑爹的bug。  
#### 第十，GIL    
以GIL结尾，因为GIL是Python中大家公认的缺陷！  
  
从其他语言过来的同学可能看到Python用threading模块，拿过来就用，结果发现效果不对啊，然后就会喷，什么鬼。  
#### 总结：    
毫无疑问的说，Python是非常容易上手，也非常强大的一门语言。Python非常灵活，可定制化很强。  
同时，也存在一些陷阱，搞清楚这些陷阱能够更好的掌握、使用这么语言。  
本文列举了一些Python中的一些缺陷，这是一份不完全列表，欢迎大家补充。    
  
## Reference    
- [程序员必知的Python陷阱与缺陷列表](http://www.cnblogs.com/xybaby/p/7183854.html)    
