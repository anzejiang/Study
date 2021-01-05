package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//func variableZeroValue() {
//	var a int
//	var s string
//	fmt.Printf("%d %q\n", a, s)
//	//fmt.Println(a, s)
//}

// 在函数外定义变量必须使用var关键字
// 函数中可以使用“:=”进行定义变量
// GO没有全局变量，在函数外部定义变量相当于实在main包内定义变量
var (
	name = "anzejiang"
	age = 18
	address = "shanxi"
	a, b = 111, "hello"
)

// 变量赋值
func azj()  {
	var a, b int = 4, 3
	var s string = "abc"
	fmt.Println("azj", a, b, s)
}

// 定义变量“:=”
func azjone()  {
	a := 3
	b := 4
	c := "aaa"
	fmt.Println("azjone", a, b, c)
}

// 定义变量
func azjtow() {
	a, b, c := 1, 2, 3
	d, e, f := "d", "e", "f"
	fmt.Println("azjtow", a, b, c ,d, e, f)
}


// 内建变量类型
// bool, string类型
// 整数类型，无符号类型(u)int, 有无符号类型int
// (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
// byte, rune(字符型32字节，类似数据库的char类型)
// float32, float64, complex64, complex128

// 欧拉公式
func euler() {
	c := cmplx.Pow(math.E, 1i * math.Pi) + 1
	fmt.Println("euler:", c)
	d := cmplx.Exp(1i*math.Pi) + 1
	fmt.Println("euler:", d)
}

// float浮点数
// 类型转换是强制的，没有隐式转换
func float() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Print(c)
}


// 常量的定义,关键字const
func consts() {
	const (
		filename = "abc.tst"
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b + b))
	fmt.Printf(filename, c)
}
// 使用常量定义枚举类型
func enums() {
	const (
		cpp = iota // iota 自增，从0开始一次向下赋值
		_		   // _ 表示跳过该次赋值
		python
		php
		ruby
		javascript
	)
	const (
		// b, kb, mb, gb, tb, pb
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
	fmt.Println("enums>", cpp, javascript,  python, php, ruby)
}


// 执行函数main
func main() {
	enums()
	consts()
	fmt.Println("Hello word")
	//variableZeroValue()
	azj()
	azjone()
	azjtow()
	fmt.Println(name, age, address, a, b)

	c := 3 + 4i	// '4i'表示复数，4i=2， 写成4 * i “go会把i当成变量”
	fmt.Println(cmplx.Abs(c))
	//euler()
	//float()
}