package main

import (
	"fmt"
)

func operatorDemo() {
	/*
		2.2 运算符
		运算符结合律全部从左到右
	*/
	/*
		    简单位运算演示。
			0110 &  1011 = 0010        AND      都为 1。
			0110 |  1011 = 1111        OR       至少一个为 1。
			0110 ^  1011 = 1101        XOR      只能一个为 1。
			0110 &^ 1011 = 0100        AND NOT  清除标志位。
	*/
	fmt.Println("简单位运算演示:")
	A := 6  // equal to binary 0110
	B := 11 // equal to binary 1011
	fmt.Println("0110 & 1011 = 0010 = ", A&B)
	fmt.Println("0110 | 1011 = 1111 = ", A|B)
	fmt.Println("0110 ^ 1011 = 1101 = ", A^B)
	fmt.Println("0110 &^ 1011 = 0100 = ", A&^B)
	/*
		标志位操作。
		a := 0
		a |= 1 << 2            // 0000100: 在 bit2 设置标志位。
		a |= 1 << 6            // 1000100: 在 bit6 设置标志位。
		a = a &^ (1 << 6)      // 0000100: 清除 bit6 标志位。
	*/
	fmt.Println("标志位操作:")
	a := 0
	a |= 1 << 2
	fmt.Println("a |= 1 << 2  ", "0000100", "=", a)
	a |= 1 << 6
	fmt.Println("a |= 1 << 6  ", "1000100", "=", a)
	a = a &^ (1 << 6)
	fmt.Println("a = a &^ (1 << 6)", "0000100", "=", a)
	/*
		不支持运算符重载。尤其需要注意，"++"、"--" 是语句而非表达式。
	*/
	n := 0
	p := &n
	// b := n++           // syntax error
	// if n++ == 1 {}     // syntax error
	// ++n                // syntax error
	n++
	*p++ // (*p)++
	fmt.Println(n, *p)
	/*
		没有 "~"，取反运算也用 "^"。
	*/
	x := 1
	fmt.Println(x, ^x) // 0001, -0010
}

func initValue() {
	/*
		初始化复合对象，必须使用类型标签，且左大括号必须在类型尾部。
	*/
	fmt.Println("初始化复合对象")
	// var a struct { x int  } = { 100  }   // syntax error
	// var b []int = { 1, 2, 3  }          // syntax error
	// c := struct {x int; y string}      // syntax error: unexpected semicolon or newline // { //  }
	var a = struct{ x int }{100}
	var b = []int{1, 2, 3}
	fmt.Println(a, b)
	/*
		初始化值以 "," 分隔。可以分多行，但后一行必须以 "," 或 "}" 结尾。
	*/
	c := []int{
		1,
		//2 // Error: need trailing comma before newline in composite literal
	}
	c = []int{
		1,
		2, // ok
	}
	d := []int{
		1,
		2} // ok
	fmt.Println(c, d)

}

func main() {
	operatorDemo()
	initValue()
}
