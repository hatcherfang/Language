## js basic knowledge
- js match upperlower case
- 由于JavaScript这个设计缺陷，不要使用==比较，始终坚持使用===比较
- 字符串是以单引号\'或双引号\"括起来的任意文本\"\'
- 数组是一组按顺序排列的集合，集合的每个值称为元素。JavaScript的数组可以包括任意数据类型
- JavaScript对象的键都是字符串类型，值可以是任意数据类型。  
```
    var person = {  
    name: 'Bob',  
    age: 20,  
    tags: ['js', 'web', 'mobile'],  
    city: 'Beijing',  
    hasCar: true,  
    zipcode: null  

};
```
## 要获取一个对象的属性，我们用对象变量.属性名的方式  
- 变量在JavaScript中就是用一个变量名表示，变量名是大小写英文、数字、$和_的组合，且不能用数字开头。变量名也不能是JavaScript的关键字，如if、while等。申明一个变量用var语句
- JavaScript在设计之初，为了方便初学者学习，并不强制要求用var申明变量。这个设计错误带来了严重的后果：如果一个变量没有通过var申明就被使用，那么该变量就自动被申明为全局变量,在同一个页面的不同的JavaScript文件中，如果都不用var申明，恰好都使用了变量i，将造成变量i互相影响，产生难以调试的错误结果。
使用var申明的变量则不是全局变量，它的范围被限制在该变量被申明的函数体内（函数的概念将稍后讲解），同名变量在不同的函数体内互不冲突。
- 需要特别注意的是，字符串是不可变的，如果对字符串的某个索引赋值，不会有任何错误，但是，也没有任何效果
- 在JavaScript中，用var申明的变量实际上是有作用域的。
如果一个变量在函数体内部申明，则该变量的作用域为整个函数体，在函数体外不可引用该变量：
- 为了解决块级作用域，ES6引入了新的关键字let，用let替代var可以申明一个块级作用域的变量：
```
'use strict';  

function foo() {  
    var sum = 0;  
    for (let i=0; i<100; i++) {  
        sum += i;  
    
    }  
    i += 1; // SyntaxError  

}  
```
## 常量：ES6标准引入了新的关键字const来定义常量，const与let都具有块级作用域：
```
'use strict';

const PI = 3.14;
PI = 3; // 某些浏览器不报错，但是无效果！
PI; // 3.14
```
## 高阶函数英文叫Higher-order function。那么什么是高阶函数？

JavaScript的函数其实都指向某个变量。既然变量可以指向函数，函数的参数能接收变量，那么一个函数就可以接收另一个函数作为参数，这种函数就称之为高阶函数。编写高阶函数，就是让函数的参数能够接收别的函数。
- map: map方法定义在JavaScript的Array中,我们调用Array的map()方法，传入我们自己的函数,得到了一个新的Array作为结果。
- reduce: Array的reduce()把一个函数作用  
在这个Array的[x1, x2, x3...]上，这个函数必须接收两个参数，
reduce()把结果继续和序列的下一个元素做累积计算，其效果就是：
`[x1, x2, x3, x4].reduce(f) = f(f(f(x1, x2), x3), x4)`
- filter:也是一个常用的操作，它用于把Array的某些元素过滤掉，然后返回剩下的元素。和map()类似，Array的filter()也接收一个函数。和map()不同的是，filter()把传入的函数依次作用于每个元素，然后根据返回值是true还是false决定保留还是丢弃该元素。
- 闭包：闭包有非常强大的功能。举个栗子： 
在面向对象的程序设计语言里，比如Java和C++，要在对象内部封装一个私有变量，可以用private修饰一个成员变量。  
在没有class机制，只有函数的语言里，借助闭包，同样可以封装一个私有变量。  
在返回的对象中，实现了一个闭包，该闭包携带了局部变量x，并且，从外部代码根本无法访问到变量x。换句话说，闭包就是携带状态的函数，并且它的状态可以完全对外隐藏起来。  
闭包还可以把多参数的函数变成单参数的函数  
- 箭头函数相当于匿名函数，并且简化了函数定义。箭头函数有两种格式，一种像上面的，只包含一个表达式，连{ ...  }和return都省略掉了。还有一种可以包含多条语句，这时候就不能省略{ ...  }和return。
```
x => {
	if (x > 0) {
        return x * x;
    
	}
	else {
        return - x * x;
    
	}

}
```
- 如果参数不是一个，就需要用括号()括起来
```
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
```
- 如果要返回一个对象，就要注意，如果是单表达式，这么写的话会报错：
```
// SyntaxError:
x => { foo: x  }
```
- 因为和函数体的{ ...  }有语法冲突，所以要改为：
```
// ok:
x => ({ foo: x  })
```
- typeof操作符获取对象的类型
```
typeof 123; // 'number'
typeof NaN; // 'number'
typeof 'str'; // 'string'
typeof true; // 'boolean'
typeof undefined; // 'undefined'
typeof Math.abs; // 'function'
typeof null; // 'object'
typeof []; // 'object'
typeof {}; // 'object'
```
- 包装对象
总结一下，有这么几条规则需要遵守：  

```
不要使用new Number()、new Boolean()、new String()创建包装对象；  

用parseInt()或parseFloat()来转换任意类型到number；  

用String()来转换任意类型到string，或者直接调用某个对象的toString()方法；  

通常不必把任意类型转换为boolean再判断，因为可以直接写if (myVar) {...}；  

typeof操作符可以判断出number、boolean、string、function和undefined；  

判断Array要使用Array.isArray(arr)；  

判断null请使用myVar === null；  

判断某个全局变量是否存在用`typeof window.myVar === 'undefined'；`  

函数内部判断某个变量是否存在用`typeof myVar === 'undefined'。`  

最后有细心的同学指出，任何对象都有toString()方法吗？null和undefined就没有！确实如此，这两个特殊值要除外，虽然null还伪装成了object类型。  
```
## Date 
- 你可能观察到了一个非常非常坑爹的地方，就是JavaScript的月份范围用整数表示是0~11，0表示一月，1表示二月……，所以要表示6月，我们传入的是5！这绝对是JavaScript的设计者当时脑抽了一下，但是现在要修复已经不可能了。
##  RegExp
1. 用\d可以匹配一个数字;\w可以匹配一个字母或数字; .可以匹配任意字符; 要匹配变长的字符，在正则表达式中，用*表示任意个字符（包括0个）;用+表示至少一个字符;用?表示0个或1个字符; 用{n}表示n个字符，用{n,m}表示n-m个字符;
2. eg:
我们来从左到右解读一下：  
- \d{3}表示匹配3个数字，例如'010'；  
- \s可以匹配一个空格（也包括Tab等空白符），所以\s+表示至少有一个空格，例如匹配' '，'\t\t'等；  
- \d{3,8}表示3-8个数字，例如'1234567'  
综合起来，上面的正则表达式可以匹配以任意个空格隔开的带区号的电话号码。  
如果要匹配'010-12345'这样的号码呢？由于'-'是特殊字符，在正则表达式中，要用'\'转义，所以，上面的正则是\d{3}\-\d{3,8}。  
但是，仍然无法匹配'010 - 12345'，因为带有空格。所以我们需要更复杂的匹配方式。  
3. 要做更精确地匹配，可以用[]表示范围，比如：
- [0-9a-zA-Z\_]可以匹配一个数字、字母或者下划线；
- [0-9a-zA-Z\_]+可以匹配至少由一个数字、字母或者下划线组成的字符串，比如'a100'，'0_Z'，'js2015'等等；
- [a-zA-Z\_\$][0-9a-zA-Z\_\$]*可以匹配由字母或下划线、$开头，后接任意个由一个数字、字母或者下划线、$组成的字符串，也就是JavaScript允许的变量名；
- [a-zA-Z\_\$][0-9a-zA-Z\_\$]{0, 19}更精确地限制了变量的长度是1-20个字符（前面1个字符+后面最多19个字符）。
- A|B可以匹配A或B，所以(J|j)ava(S|s)cript可以匹配'JavaScript'、'Javascript'、'javaScript'或者'javascript'。
- ^表示行的开头，^\d表示必须以数字开头。
- $表示行的结束，\d$表示必须以数字结束。
- 你可能注意到了，js也可以匹配'jsp'，但是加上^js$就变成了整行匹配，就只能匹配'js'了。
4. 切分字符串
```
用正则表达式切分字符串比用固定的字符更灵活，请看正常的切分代码：

'a b   c'.split(' '); // ['a', 'b', '', '', 'c']
嗯，无法识别连续的空格，用正则表达式试试：

'a b   c'.split(/\s+/); // ['a', 'b', 'c']
无论多少个空格都可以正常分割。加入,试试：

'a,b, c  d'.split(/[\s\,]+/); // ['a', 'b', 'c', 'd']
再加入;试试：

'a,b;; c  d'.split(/[\s\,\;]+/); // ['a', 'b', 'c', 'd']
如果用户输入了一组标签，下次记得用正则表达式来把不规范的输入转化成正确的数组。
```
5. 分组
```
除了简单地判断是否匹配之外，正则表达式还有提取子串的强大功能。用()表示的就是要提取的分组（Group）。比如：

^(\d{3})-(\d{3,8})$分别定义了两个组，可以直接从匹配的字符串中提取出区号和本地号码：

var re = /^(\d{3})-(\d{3,8})$/;
re.exec('010-12345'); // ['010-12345', '010', '12345']
re.exec('010 12345'); // null
如果正则表达式中定义了组，就可以在RegExp对象上用exec()方法提取出子串来。

exec()方法在匹配成功后，会返回一个Array，第一个元素是正则表达式匹配到的整个字符串，后面的字符串表示匹配成功的子串。

exec()方法在匹配失败时返回null。

提取子串非常有用。来看一个更凶残的例子：

var re = /^(0[0-9]|1[0-9]|2[0-3]|[0-9])\:(0[0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9]|[0-9])\:(0[0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9]|[0-9])$/;
re.exec('19:05:30'); // ['19:05:30', '19', '05', '30']
这个正则表达式可以直接识别合法的时间。但是有些时候，用正则表达式也无法做到完全验证，比如识别日期：

var re = /^(0[1-9]|1[0-2]|[0-9])-(0[1-9]|1[0-9]|2[0-9]|3[0-1]|[0-9])$/;
对于'2-30'，'4-31'这样的非法日期，用正则还是识别不了，或者说写出来非常困难，这时就需要程序配合识别了。
```
6. 贪婪匹配
```
需要特别指出的是，正则匹配默认是贪婪匹配，也就是匹配尽可能多的字符。举例如下，匹配出数字后面的0：

var re = /^(\d+)(0*)$/;
re.exec('102300'); // ['102300', '102300', '']
由于\d+采用贪婪匹配，直接把后面的0全部匹配了，结果0*只能匹配空字符串了。

必须让\d+采用非贪婪匹配（也就是尽可能少匹配），才能把后面的0匹配出来，加个?就可以让\d+采用非贪婪匹配：

var re = /^(\d+?)(0*)$/;
re.exec('102300'); // ['102300', '1023', '00']
```
7. 全局搜索
```
JavaScript的正则表达式还有几个特殊的标志，最常用的是g，表示全局匹配：

var r1 = /test/g;
// 等价于:
var r2 = new RegExp('test', 'g');
全局匹配可以多次执行exec()方法来搜索一个匹配的字符串。当我们指定g标志后，每次运行exec()，正则表达式本身会更新lastIndex属性，表示上次匹配到的最后索引：

var s = 'JavaScript, VBScript, JScript and ECMAScript';
var re=/[a-zA-Z]+Script/g;

// 使用全局匹配:
re.exec(s); // ['JavaScript']
re.lastIndex; // 10

re.exec(s); // ['VBScript']
re.lastIndex; // 20

re.exec(s); // ['JScript']
re.lastIndex; // 29

re.exec(s); // ['ECMAScript']
re.lastIndex; // 44

re.exec(s); // null，直到结束仍没有匹配到
全局匹配类似搜索，因此不能使用/^...$/，那样只会最多匹配一次。

正则表达式还可以指定i标志，表示忽略大小写，m标志，表示执行多行匹配。
```
## JSON
- 在JSON中，一共就这么几种数据类型：
```
number：和JavaScript的number完全一致；
boolean：就是JavaScript的true或false；
string：就是JavaScript的string；
null：就是JavaScript的null；
array：就是JavaScript的Array表示方式——[]；
object：就是JavaScript的{ ...  }表示方式。
以及上面的任意组合。

并且，JSON还定死了字符集必须是UTF-8，表示多语言就没有问题了。为了统一解析，JSON的字符串规定必须用双引号""，Object的键也必须用双引号""。
```
- 序列化
让我们先把小明这个对象序列化成JSON格式的字符串：  
```
var xiaoming = {
    name: '小明',
    age: 14,
    gender: true,
    height: 1.65,
    grade: null,
    'middle-school': '\"W3C\" Middle School',
    skills: ['JavaScript', 'Java', 'Python', 'Lisp']

};

JSON.stringify(xiaoming); // '{"name":"小明","age":14,"gender":true,"height":1.65,"grade":null,"middle-school":"\"W3C\" Middle School","skills":["JavaScript","Java","Python","Lisp"]}'
```
- 反序列化  
拿到一个JSON格式的字符串，我们直接用JSON.parse()把它变成一个JavaScript对象：  
```
JSON.parse('[1,2,3,true]'); // [1, 2, 3, true]
JSON.parse('{"name":"小明","age":14}'); // Object {name: '小明', age: 14}
JSON.parse('true'); // true
JSON.parse('123.45'); // 123.45
```
## 面向对象编程
- 构造函数  
在JavaScript中，可以用关键字new来调用这个函数，并返回一个对象,如果不写new，这就是一个普通函数，它返回undefined。但是，如果写了new，它就变成了一个构造函数，它绑定的this指向新创建的对象，并默认返回this，也就是说，不需要在最后写return this;。
```
function Student(name) {
    this.name = name;
    this.hello = function () {
        alert('Hello, ' + this.name + '!');
    
    }

}
var xiaoming = new Student('小明');
xiaoming.name; // '小明'
xiaoming.hello(); // Hello, 小明!
var xiaohong = new Student('小红');
```
xiaoming和xiaohong各自的hello是一个函数，但它们是两个不同的函数，虽然函数名称和代码都是相同的！  
如果我们通过new Student()创建了很多对象，这些对象的hello函数实际上只需要共享同一个函数就可以了，这样可以节省很多内存。  
要让创建的对象共享一个hello函数，根据对象的属性查找原则，我们只要把hello函数移动到xiaoming、xiaohong这些对象共同的原型上就可以了，也就是Student.prototype：  
修改代码如下:  
```
function Student(name) {
    this.name = name;

}

Student.prototype.hello = function () {
    alert('Hello, ' + this.name + '!');

};
```
- 忘记写new怎么办  
如果一个函数被定义为用于创建对象的构造函数，但是调用时忘记了写new怎么办？
在strict模式下，this.name = name将报错，因为this绑定为undefined，在非strict模式下，this.name = name不报错，因为this绑定为window，于是无意间创建了全局变量name，并且返回undefined，这个结果更糟糕。
所以，调用构造函数千万不要忘记写new。为了区分普通函数和构造函数，按照约定，构造函数首字母应当大写，而普通函数首字母应当小写，这样，一些语法检查工具如jslint将可以帮你检测到漏写的new.
最后，我们还可以编写一个createStudent()函数，在内部封装所有的new操作。一个常用的编程模式像这样：  
```
function Student(props) {
    this.name = props.name || '匿名'; // 默认值为'匿名'
    this.grade = props.grade || 1; // 默认值为1

}

Student.prototype.hello = function () {
    alert('Hello, ' + this.name + '!');

};

function createStudent(props) {
    return new Student(props || {})

}
```
这个createStudent()函数有几个巨大的优点：一是不需要new来调用，二是参数非常灵活，可以不传，也可以这么传：  
```
var xiaoming = createStudent({
    name: '小明'
});

xiaoming.grade; // 1
```
- class继承
如果用新的class关键字来编写Student，可以这样写：  
```
class Student {
	constructor(name) {
        this.name = name;
    
	}

	hello() {
        alert('Hello, ' + this.name + '!');
    
	}
}
```
比较一下就可以发现，class的定义包含了构造函数constructor和定义在原型对象上的函数hello()（注意没有function关键字），这样就避免了Student.prototype.hello = function () {...}这样分散的代码。  
最后，创建一个Student对象代码和前面章节完全一样：  
```
var xiaoming = new Student('小明');
xiaoming.hello();
```
- class继承
用class定义对象的另一个巨大的好处是继承更方便了。想一想我们从Student派生一个PrimaryStudent需要编写的代码量。现在，原型继承的中间对象，原型对象的构造函数等等都不需要考虑了，直接通过extends来实现：  
```
class PrimaryStudent extends Student {
	constructor(name, grade) {
        super(name); // 记得用super调用父类的构造方法!
        this.grade = grade;
    
	}

	myGrade() {
        alert('I am at grade ' + this.grade);
    
	}
}
```
注意PrimaryStudent的定义也是class关键字实现的，而extends则表示原型链对象来自Student。子类的构造函数可能会与父类不太相同，例如，PrimaryStudent需要name和grade两个参数，并且需要通过super(name)来调用父类的构造函数，否则父类的name属性无法正常初始化。  

PrimaryStudent已经自动获得了父类Student的hello方法，我们又在子类中定义了新的myGrade方法。  

ES6引入的class和原有的JavaScript原型继承有什么区别呢？实际上它们没有任何区别，class的作用就是让JavaScript引擎去实现原来需要我们自己编写的原型链代码。简而言之，用class的好处就是极大地简化了原型链代码。  

你一定会问，class这么好用，能不能现在就用上？  

现在用还早了点，因为不是所有的主流浏览器都支持ES6的class。如果一定要现在就用上，就需要一个工具把class代码转换为传统的prototype代码，可以试试Babel这个工具。  
## 浏览器  
- 在编写JavaScript的时候，就要充分考虑到浏览器的差异，尽量让同一份JavaScript代码能运行在不同的浏览器中。  
### 浏览器对象
- window  
window对象不但充当全局作用域，而且表示浏览器窗口。 
window对象有innerWidth和innerHeight属性，可以获取浏览器窗口的内部宽度和高度。内部宽高是指除去菜单栏、工具栏、边框等占位元素后，用于显示网页的净宽高。  
兼容性：IE<=8不支持。  
```
'use strict';
// 可以调整浏览器窗口大小试试:
alert('window inner size: ' + window.innerWidth + ' x ' + window.innerHeight);
```
- navigator  
navigator对象表示浏览器的信息，最常用的属性包括：
```
navigator.appName：浏览器名称；
navigator.appVersion：浏览器版本；
navigator.language：浏览器设置的语言；
navigator.platform：操作系统类型；
navigator.userAgent：浏览器设定的User-Agent字符串。
```
- 短路运算符||
```
var s = 1===1 || 'fang';
console.log(s);
```
- location  
location对象表示当前页面的URL信息。例如，一个完整的URL：  
```
http://www.example.com:8080/path/index.html?a=1&b=2#TOP
```
可以用location.href获取。要获得URL各个部分的值，可以这么写：  
```
location.protocol; // 'http'
location.host; // 'www.example.com'
location.port; // '8080'
location.pathname; // '/path/index.html'
location.search; // '?a=1&b=2'
location.hash; // 'TOP'
```
要加载一个新页面，可以调用location.assign()。如果要重新加载当前页面，调用location.reload()方法非常方便。  
```
'use strict';
if (confirm('重新加载当前页' + location.href + '?')) {
    location.reload();

} else {
    location.assign('/discuss'); // 设置一个新的URL地址 /discuss is the pathname

}
```
- document  
document对象表示当前页面。由于HTML在浏览器中以DOM形式表示为树形结构，document对象就是整个DOM树的根节点。  
document的title属性是从HTML文档中的<title>xxx</title>读取的，但是可以动态改变：  
要查找DOM树的某个节点，需要从document对象开始查找。最常用的查找是根据ID和Tag Name。  
用document对象提供的getElementById()和getElementsByTagName()可以按ID获得一个DOM节点和按Tag名称获得一组DOM节点.  

document对象还有一个cookie属性，可以获取当前页面的Cookie。  
Cookie是由服务器发送的key-value标示符。因为HTTP协议是无状态的，但是服务器要区分到底是哪个用户发过来的请求，就可以用Cookie来区分。当一个用户成功登录后，服务器发送一个Cookie给浏览器，例如user=ABC123XYZ(加密的字符串)...，此后，浏览器访问该网站时，会在请求头附上这个Cookie，服务器根据Cookie即可区分出用户。  
Cookie还可以存储网站的一些设置，例如，页面显示的语言等等。  
JavaScript可以通过document.cookie读取到当前页面的Cookie：  
```
document.cookie; // 'v=123; remember=true; prefer=zh'
```
由于JavaScript能读取到页面的Cookie，而用户的登录信息通常也存在Cookie中，这就造成了巨大的安全隐患，这是因为在HTML页面中引入第三方的JavaScript代码是允许的：  
```
<!-- 当前页面在wwwexample.com -->
<html>
    <head>
        <script src="http://www.foo.com/jquery.js"></script>
    </head>
    ...
</html>
```
如果引入的第三方的JavaScript中存在恶意代码，则www.foo.com网站将直接获取到www.example.com网站的用户登录信息。  
为了解决这个问题，服务器在设置Cookie时可以使用httpOnly，设定了httpOnly的Cookie将不能被JavaScript读取。这个行为由浏览器实现，主流浏览器均支持httpOnly选项，IE从IE6 SP1开始支持。  
为了确保安全，服务器端在设置Cookie时，应该始终坚持使用httpOnly。  
- XSS就是在别人的应用程序中恶意执行一段js以窃取用户的cookie   
refer:`http://blog.csdn.net/zzzmmmkkk/article/details/10862949`  
refer:`http://ulricqin.com/post/httponly-cookie/`  
- history  
history对象保存了浏览器的历史记录，JavaScript可以调用history对象的back()或forward ()，相当于用户点击了浏览器的“后退”或“前进”按钮。  
这个对象属于历史遗留对象，对于现代Web页面来说，由于大量使用AJAX和页面交互，简单粗暴地调用history.back()可能会让用户感到非常愤怒。  
新手开始设计Web页面时喜欢在登录页登录成功时调用history.back()，试图回到登录前的页面。这是一种错误的方法。  
**任何情况，你都不应该使用history这个对象了。**  
### 操作DOM
由于HTML文档被浏览器解析后就是一棵DOM树，要改变HTML的结构，就需要通过JavaScript来操作DOM。  
始终记住DOM是一个树形结构。操作一个DOM节点实际上就是这么几个操作：  
```
更新：更新该DOM节点的内容，相当于更新了该DOM节点表示的HTML的内容；

遍历：遍历该DOM节点下的子节点，以便进行进一步操作；

添加：在该DOM节点下新增一个子节点，相当于动态增加了一个HTML节点；

删除：将该节点从HTML中删除，相当于删掉了该DOM节点的内容以及它包含的所有子节点。
```
在操作一个DOM节点前，我们需要通过各种方式先拿到这个DOM节点。最常用的方法是document.getElementById()和document.getElementsByTagName()，以及CSS选择器document.getElementsByClassName()。  

由于ID在HTML文档中是唯一的，所以document.getElementById()可以直接定位唯一的一个DOM节点。document.getElementsByTagName()和document.getElementsByClassName()总是返回一组DOM节点。要精确地选择DOM，可以先定位父节点，再从父节点开始选择，以缩小范围。  
例如：  
```
// 返回ID为'test'的节点：
var test = document.getElementById('test');

// 先定位ID为'test-table'的节点，再返回其内部所有tr节点：
var trs = document.getElementById('test-table').getElementsByTagName('tr');

// 先定位ID为'test-div'的节点，再返回其内部所有class包含red的节点：
var reds = document.getElementById('test-div').getElementsByClassName('red');

// 获取节点test下的所有直属子节点:
var cs = test.children;

// 获取节点test下第一个、最后一个子节点：
var first = test.firstElementChild;
var last = test.lastElementChild;
```

html example:  
```
<div id="test-div">
<div class="c-red">
    <p id="test-p">JavaScript</p>
    <p>Java</p>
  </div>
  <div class="c-red c-green">
    <p>Python</p>
    <p>Ruby</p>
    <p>Swift</p>
  </div>
  <div class="c-green">
    <p>Scheme</p>
    <p>Haskell</p>
  </div>
</div>
<script>
   var js = document.getElementById("test-p"); 
   //alert(js.innerHTML);
   var lan = document.getElementsByClassName("c-red c-green")[0].children;
   var strlan = '';
   var i = 0;
   for (i=0; i<lan.length; i++){
       strlan = strlan + ' ' + lan[i].innerText;
   
   }
   //alert(strlan);
   var cgreen = document.getElementsByClassName("c-green")[0].lastElementChild;
   alert(cgreen.innerText);
   for (i=0; i<cgreen.length; i++){
      alert(cgreen[i].innerText);
   
   }
</script>

结论：document.getElementsByClassName("c-red c-green")和document.getElementsByClassName("c-green")比较得出
getElementsByClassName是模糊匹配并不是唯一，唯一标识的是document.getElementById

```
- js中innerHTML与innerText的用法与区别  
```
用法：

<div id="test">
   <span style="color:red">test1</span> test2
</div>

在JS中可以使用：
test.innerHTML:
　　也就是从对象的起始位置到终止位置的全部内容,包括Html标签。
　　上例中的test.innerHTML的值也就是“<span style="color:red">test1</span> test2 ”。
test.innerText: 
　　从起始位置到终止位置的内容, 但它去除Html标签 
　　上例中的test.innerText的值也就是“test1 test2”, 其中span标签去除了。
```
#### 更新DOM 
- 一种是修改innerHTML属性  
这个方式非常强大，不但可以修改一个DOM节点的文本内容，还可以直接通过HTML片段修改DOM节点内部的子树：
```
// 获取<p id="p-id">...</p>
var p = document.getElementById('p-id');
// 设置文本为abc:
p.innerHTML = 'ABC'; // <p id="p-id">ABC</p>
// 设置HTML:
p.innerHTML = 'ABC <span style="color:red">RED</span> XYZ';
// <p>...</p>的内部结构已修改
```
**用innerHTML时要注意，是否需要写入HTML。如果写入的字符串是通过网络拿到了，要注意对字符编码来避免XSS攻击。**  
- 第二种是修改innerText或textContent属性，这样可以自动对字符串进行HTML编码，保证无法设置任何HTML标签：  
```
// 获取<p id="p-id">...</p>
var p = document.getElementById('p-id');
// 设置文本:
p.innerText = '<script>alert("Hi")</script>';
// HTML被自动编码，无法设置一个<script>节点:
// <p id="p-id">&lt;script&gt;alert("Hi")&lt;/script&gt;</p>
```
两者的区别在于读取属性时，innerText不返回隐藏元素的文本，而textContent返回所有文本。另外注意IE<9不支持textContent。  
修改CSS也是经常需要的操作。DOM节点的style属性对应所有的CSS，可以直接获取或设置。因为CSS允许font-size这样的名称，但它并非JavaScript有效的属性名，所以需要在JavaScript中改写为驼峰式命名fontSize：  
```
// 获取<p id="p-id">...</p>
var p = document.getElementById('p-id');
// 设置CSS:
p.style.color = '#ff0000';
p.style.fontSize = '20px';
p.style.paddingTop = '2em';
```
#### 插入DOM 
- 一个是使用appendChild  
把一个子节点添加到父节点的最后一个子节点。例如：  
```
<!-- HTML结构 -->
<p id="js">JavaScript</p>
<div id="list">
    <p id="java">Java</p>
    <p id="python">Python</p>
    <p id="scheme">Scheme</p>
</div>
```
把<p id="js">JavaScript</p>添加到<div id="list">的最后一项：  
```
var
    js = document.getElementById('js'),
    list = document.getElementById('list');
list.appendChild(js);
```
现在，HTML结构变成了这样：
```
<!-- HTML结构 -->
<div id="list">
    <p id="java">Java</p>
    <p id="python">Python</p>
    <p id="scheme">Scheme</p>
    <p id="js">JavaScript</p>
</div>
```
**因为我们插入的js节点已经存在于当前的文档树，因此这个节点首先会从原先的位置删除，再插入到新的位置。**  

更多的时候我们会从零创建一个新的节点，然后插入到指定位置：  
```
var
    list = document.getElementById('list'),
    haskell = document.createElement('p');
haskell.id = 'haskell';
haskell.innerText = 'Haskell';
list.appendChild(haskell);
```
动态创建一个节点然后添加到DOM树中，可以实现很多功能。举个例子，下面的代码动态创建了一个<style>节点，然后把它添加到<head>节点的末尾，这样就动态地给文档添加了新的CSS定义：  
```
var d = document.createElement('style');
d.setAttribute('type', 'text/css');
d.innerHTML = 'p { color: red  }';
document.getElementsByTagName('head')[0].appendChild(d);
```
可以在Chrome的控制台执行上述代码，观察页面样式的变化.  
- insertBefore  
如果我们要把子节点插入到指定的位置怎么办？可以使用parentElement.insertBefore(newElement, referenceElement);，子节点会插入到referenceElement之前。  

还是以上面的HTML为例，假定我们要把Haskell插入到Python之前：  
```
<!-- HTML结构 -->
<div id="list">
    <p id="java">Java</p>
    <p id="python">Python</p>
    <p id="scheme">Scheme</p>
</div>
```
可以这么写：  
```
var
    list = document.getElementById('list'),
    ref = document.getElementById('python'),
    haskell = document.createElement('p');
haskell.id = 'haskell';
haskell.innerText = 'Haskell';
list.insertBefore(haskell, ref);
```
新的HTML结构如下：  
```
<!-- HTML结构 -->
<div id="list">
    <p id="java">Java</p>
    <p id="haskell">Haskell</p>
    <p id="python">Python</p>
    <p id="scheme">Scheme</p>
</div>
```
可见，使用insertBefore重点是要拿到一个“参考子节点”的引用。很多时候，需要循环一个父节点的所有子节点，可以通过迭代children属性实现：  
```
var
    i, c,
    list = document.getElementById('list');
for (i = 0; i < list.children.length; i++) {
    c = list.children[i]; // 拿到第i个子节点

}
```
#### 删除DOM
删除一个DOM节点就比插入要容易得多。  

要删除一个节点，首先要获得该节点本身以及它的父节点，然后，调用父节点的removeChild把自己删掉：  
```
// 拿到待删除节点:
var self = document.getElementById('to-be-removed');
// 拿到父节点:
var parent = self.parentElement;
// 删除:
var removed = parent.removeChild(self);
removed === self; // true
```
注意到删除后的节点虽然不在文档树中了，但其实它还在内存中，可以随时再次被添加到别的位置。  

当你遍历一个父节点的子节点并进行删除操作时，要注意，**children属性是一个只读属性，并且它在子节点变化时会实时更新。**   

例如，对于如下HTML结构：  
```
<div id="parent">
    <p>First</p>
    <p>Second</p>
</div>
```
当我们用如下代码删除子节点时：  
```
var parent = document.getElementById('parent');
parent.removeChild(parent.children[0]);
parent.removeChild(parent.children[1]); // <-- 浏览器报错
```
浏览器报错：parent.children[1]不是一个有效的节点。原因就在于，当<p>First</p>节点被删除后，parent.children的节点数量已经从2变为了1，索引[1]已经不存在了。  

因此，删除多个节点时，要注意children属性时刻都在变化。  

**Related Tutorials**  
[www.liaoxuefeng.com](https://www.liaoxuefeng.com/wiki/001434446689867b27157e896e74d51a89c25cc8b43bdb3000)  
