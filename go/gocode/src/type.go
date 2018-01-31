package main

import (
	"fmt"
	"unsafe"
)

func variableDefine() {
	// 1.1 变量
	println("---1.1 变量---")
	/*
	   var x int
	   var f float32 = 1.6
	   var s = "abc"
	*/

	//可一次定义多个变量
	var x1, y1, z1 int
	println(x1, y1, z1)
	var s, n = "abc", 123
	var (
		a int
		b float32
	)
	println(a, b)
	var xx = "fang"
	println(xx)
	// ":=" only can be used in function
	x := 123
	println(x, s)
	n, s = 0x1234, "hello, world!"
	println(x, s, n)
	// 编译器会将未使用的局部变量当做错误
	// 但是未使用的局部常量不会被编译器当做错误
	// xx := 222

	// 重新赋值的同名变量地址相同与定义新同名变量的地址不同
	// 注意：变量不能在函数外进行赋值，调用等操作
	// example
	s, y := "abc", 20
	println(&s, y)
	s, y = "hello", 30 // 重新赋值: 与前 s 在同一层次的代码块中，且有新的变量被定义。
	println(&s, y)
	{
		s, z := 1000, 30 // 定义新同名变量: 不在同⼀一层次代码块
		println(&s, z)
	}
	// 多变量赋值时，先计算所有相关值，然后再从左到右依次赋值。
	var data, i = [3]int{0, 1, 2}, 0
	// (i = 0) -> (i = 2), (data[0] = 100)
	i, data[i] = 2, 100
	println(i, data[0])
}

func constType() {
	// 1.2 常量
	println("---1.2 常量---")
	// 常量值必须是编译器可以预定的数字，字符串，布尔值
	const x2, y2 int = 1, 2 //多常量初始化
	const s2 = "hello, world!"
	const (
		a, b      = 10, 100
		c    bool = false
	)
	println(x2, y2, a, b, c)

}

func quoteType() {
	// 1.4引用类型
	println("---1.4引用类型---")
	a := []int{0, 0, 0} //提供初始化表达式。
	a[1] = 10
	// make会被编译器翻译成具体的创建函数，由其分配内存和初始化成员结构，返回对象而非指针
	b := make([]int, 3) //makeslice
	b[1] = 10
	c := new([]int)
	// c[0] = 10 //Error: Invalide operation
	println(a[0], b[0], c)
}

func typeConversion() {
	// 1.5类型转换
	//不支持隐式类型转换，即便是从窄向宽转换也不行
	println("---1.5类型转换---")
	var b3 byte = 100
	//var n int = b //Error: can not use b as type int in assignment
	// 显式转换使用括号避免优先级错误
	var n1 int = int(b3)
	println(n1)
	// 不能将其它类型当做布尔值来使用
}

func stringType() {
	// 1.6 字符串
	// 要修改字符串，可先将其转换成 []rune 或 []byte，完成后再转换为 string。⽆无论哪种转 换，都会重新分配内存，并复制字节数组。
	println("---1.6 字符串---")
	s := "abcd"
	bs := []byte(s)
	bs[1] = 'B'
	println(string(bs))
	u := "电脑"
	us := []rune(u)
	us[1] = '话'
	println(string(us))
	println("---用 for 循环遍历字符串时，也有 byte 和 rune 两种方式---")
	s = "abc汉字"
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%c,", s[i])
	}
	fmt.Println()
	for _, r := range s { // rune
		fmt.Printf("%c,", r)
	}
	fmt.Println()
}

func pointerType() {
	/*
			1.7 指针
		    支持指针类型 *T，指针的指针 **T，以及包含包名前缀的 *<package>.T。
			•默认值 nil，没有 NULL 常量。
			•操作符 "&" 取变量地址，"*" 透过指针访问⺫⽬目标对象。
			•不支持指针运算，不支持 "->" 运算符，直接⽤用 "." 访问⺫⽬目标成员。
	*/
	println("---1.7 指针---")
	type data struct{ a int }
	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("%p, %v\n", p, p.a) // 直接用指针访问目标对象成员，无须转换。
	/*
		不能对指针做加减法等运算。
	*/
	//p++ // Error: invalid operation: p += 1 (mismatched types *int and int)

	/*
		可以在 unsafe.Pointer 和任意类型指针间进⾏行转换。
	*/
	x := 0x12345678
	pp := unsafe.Pointer(&x) // *int -> Pointer
	n := (*[4]byte)(pp)      // Pointer -> *[4]byte
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	}
	fmt.Println()
	/*
		返回局部变量指针是安全的，编译器会根据需要将其分配在 GC Heap 上。
	*/
	fmt.Println(test())
}

func test() *int {
	x := 100
	return &x // 在堆上分配 x 内存。但在内联时，也可能直接分配在目标栈。
}

func customDefine() {
	/*
		1.8 自定义类型
		可将类型分为命名和未命名两大类。
		命名类型包括 bool、int、string 等，而 array、 slice、map 等和具体元素类型、长度等有关，属于未命名类型。

		具有相同声明的未命名类型被视为同一类型。

		•具有相同基类型的指针。
		•具有相同元素类型和长度的 array。
		•具有相同元素类型的 slice。
		•具有相同键值类型的 map。
		•具有相同元素类型和传送⽅方向的 channel。
		•具有相同字段序列 (字段名、类型、标签、顺序) 的匿名 struct。
		•签名相同 (参数和返回值，不包括参数名称) 的 function。
		•方法集相同 (方法名、方法签名相同，和次序⽆无关) 的 interface。
	*/
	fmt.Printf("---1.8 自定义类型---\n")
	var a struct {
		x int `a`
	}
	var b struct {
		x int `ab`
	}
	// b = a // cannot use a (type struct { x int "a"  }) as type struct { x int "ab"  } in assignment b = a
	fmt.Println(a.x, b.x)
	/*
		可用 type 在全局或函数内定义新类型。
		新类型不是原类型的别名，除拥有相同数据存储结构外，它们之间没有任何关系，不会持 有原类型任何信息。除⾮非⺫⽬目标类型是未命名类型，否则必须显式转换。
	*/
	type bigint int32
	var x bigint = 100
	println(x)
}

func main() {
	variableDefine()
	constType()
	quoteType()
	typeConversion()
	stringType()
	pointerType()
	customDefine()
}
