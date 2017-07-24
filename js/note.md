"reference `https://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000`"  
- js match upperlower case
- 由于JavaScript这个设计缺陷，不要使用==比较，始终坚持使用===比较
- 字符串是以单引号\'或双引号\"括起来的任意文本\"\'
- 数组是一组按顺序排列的集合，集合的每个值称为元素。JavaScript的数组可以包括任意数据类型
- JavaScript对象的键都是字符串类型，值可以是任意数据类型。  
"`var person = {  
    name: 'Bob',  
    age: 20,  
    tags: ['js', 'web', 'mobile'],  
    city: 'Beijing',  
    hasCar: true,  
    zipcode: null  

};`"
要获取一个对象的属性，我们用对象变量.属性名的方式  
- 变量在JavaScript中就是用一个变量名表示，变量名是大小写英文、数字、$和_的组合，且不能用数字开头。变量名也不能是JavaScript的关键字，如if、while等。申明一个变量用var语句
- JavaScript在设计之初，为了方便初学者学习，并不强制要求用var申明变量。这个设计错误带来了严重的后果：如果一个变量没有通过var申明就被使用，那么该变量就自动被申明为全局变量,在同一个页面的不同的JavaScript文件中，如果都不用var申明，恰好都使用了变量i，将造成变量i互相影响，产生难以调试的错误结果。
使用var申明的变量则不是全局变量，它的范围被限制在该变量被申明的函数体内（函数的概念将稍后讲解），同名变量在不同的函数体内互不冲突。
- 需要特别注意的是，字符串是不可变的，如果对字符串的某个索引赋值，不会有任何错误，但是，也没有任何效果
- 在JavaScript中，用var申明的变量实际上是有作用域的。
如果一个变量在函数体内部申明，则该变量的作用域为整个函数体，在函数体外不可引用该变量：
- 为了解决块级作用域，ES6引入了新的关键字let，用let替代var可以申明一个块级作用域的变量：
"`
'use strict';  

function foo() {  
    var sum = 0;  
    for (let i=0; i<100; i++) {  
        sum += i;  
    
    }  
    i += 1; // SyntaxError  

}  
`"  
- 常量：ES6标准引入了新的关键字const来定义常量，const与let都具有块级作用域：
"`
'use strict';

const PI = 3.14;
PI = 3; // 某些浏览器不报错，但是无效果！
PI; // 3.14
`"
- 高阶函数英文叫Higher-order function。那么什么是高阶函数？

JavaScript的函数其实都指向某个变量。既然变量可以指向函数，函数的参数能接收变量，那么一个函数就可以接收另一个函数作为参数，这种函数就称之为高阶函数。编写高阶函数，就是让函数的参数能够接收别的函数。
- map: map方法定义在JavaScript的Array中,我们调用Array的map()方法，传入我们自己的函数,得到了一个新的Array作为结果。
- reduce: Array的reduce()把一个函数作用  
在这个Array的[x1, x2, x3...]上，这个函数必须接收两个参数，
reduce()把结果继续和序列的下一个元素做累积计算，其效果就是：
[x1, x2, x3, x4].reduce(f) = f(f(f(x1, x2), x3), x4)
- filter:也是一个常用的操作，它用于把Array的某些元素过滤掉，然后返回剩下的元素。和map()类似，Array的filter()也接收一个函数。和map()不同的是，filter()把传入的函数依次作用于每个元素，然后根据返回值是true还是false决定保留还是丢弃该元素。
- 闭包：闭包有非常强大的功能。举个栗子： 
在面向对象的程序设计语言里，比如Java和C++，要在对象内部封装一个私有变量，可以用private修饰一个成员变量。  
在没有class机制，只有函数的语言里，借助闭包，同样可以封装一个私有变量。  
在返回的对象中，实现了一个闭包，该闭包携带了局部变量x，并且，从外部代码根本无法访问到变量x。换句话说，闭包就是携带状态的函数，并且它的状态可以完全对外隐藏起来。  
闭包还可以把多参数的函数变成单参数的函数  
- 箭头函数相当于匿名函数，并且简化了函数定义。箭头函数有两种格式，一种像上面的，只包含一个表达式，连{ ...  }和return都省略掉了。还有一种可以包含多条语句，这时候就不能省略{ ...  }和return。
`x => {
	if (x > 0) {
        return x * x;
    
	}
	else {
        return - x * x;
    
	}

}`
- 如果参数不是一个，就需要用括号()括起来
`
// 两个参数:
(x, y) => x * x + y * y

// 无参数:
() => 3.14

// 可变参数:
(x, y, ...rest) => {
    var i, sum = x + y;
    for (i=0; i<rest.length; i++) {
        sum += rest[i];
    
    }
    return sum;

}
`
- 如果要返回一个对象，就要注意，如果是单表达式，这么写的话会报错：`// SyntaxError:
x => { foo: x  }`
- 因为和函数体的{ ...  }有语法冲突，所以要改为：
`// ok:
x => ({ foo: x  })`
- typeof操作符获取对象的类型
`typeof 123; // 'number'
typeof NaN; // 'number'
typeof 'str'; // 'string'
typeof true; // 'boolean'
typeof undefined; // 'undefined'
typeof Math.abs; // 'function'
typeof null; // 'object'
typeof []; // 'object'
typeof {}; // 'object'`
- 包装对象
总结一下，有这么几条规则需要遵守：  

不要使用new Number()、new Boolean()、new String()创建包装对象；  

用parseInt()或parseFloat()来转换任意类型到number；  

用String()来转换任意类型到string，或者直接调用某个对象的toString()方法；  

通常不必把任意类型转换为boolean再判断，因为可以直接写if (myVar) {...}；  

typeof操作符可以判断出number、boolean、string、function和undefined；  

判断Array要使用Array.isArray(arr)；  

判断null请使用myVar === null；  

判断某个全局变量是否存在用typeof window.myVar === 'undefined'；  

函数内部判断某个变量是否存在用typeof myVar === 'undefined'。  

最后有细心的同学指出，任何对象都有toString()方法吗？null和undefined就没有！确实如此，这两个特殊值要除外，虽然null还伪装成了object类型。  
- Date: 你可能观察到了一个非常非常坑爹的地方，就是JavaScript的月份范围用整数表示是0~11，0表示一月，1表示二月……，所以要表示6月，我们传入的是5！这绝对是JavaScript的设计者当时脑抽了一下，但是现在要修复已经不可能了。
