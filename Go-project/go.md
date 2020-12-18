## 定义变量

```

//func variableZeroValue() {
// var a int
// var s string
// fmt.Printf("%d %q\n", a, s)
// //fmt.Println(a, s)
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
      _        // _ 表示跳过该次赋值
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

```



## if、switch（case）条件控制

```
package main

import (
   "fmt"
   "io/ioutil"
)

// if的条件里可以赋值
// if条件里赋的值的变量作用域只在这个if判断中
func readfile() {
   const filename = "abc.txt"
   // 读取文件函数"ioutil.ReadFile"返回两个参数，一个文本内容，一个报错内容
   //contents, err := ioutil.ReadFile(filename)
   //if err != nil {
   // fmt.Println(err)
   //}else {
   // fmt.Printf("%s\n", contents)
   //}
   if contents, err := ioutil.ReadFile(filename);  err != nil {
      fmt.Println(err)
   }else {
      fmt.Printf("%s\n", contents)
   }
}

func grade(source int) string {
   g := ""
   switch {
   case source < 0 || source > 100:
      // panic函数，当执行该函数时，立即终止执行函数内的其他代码
      panic(fmt.Sprintf("Wrong score: %d", source))
   case source < 60:
      g = "F"
   case source < 80:
      g = "C"
   case source < 90:
      g = "B"
   case source < 100:
      g = "A"
   }
   return g
}

func main() {
   //readfile()
   fmt.Println(
      //grade(111),
      grade(65),
      grade(89),
      grade(95),
   )
}
```

## for循环

```
for格式
for 初始条件;结束条件;控制条件{
	内容
}
死循环：无线循环输出123
for {
fmt.println(123)
}

func fortest() {
   sum := 0
   for i := 1; i <=100; i++{
      sum += i
   }
   fmt.Println(sum)
}

func convertToBin(n int) string {
   result := ""
   for ; n > 0; n /= 2 {
      lsb := n % 2
      result = strconv.Itoa(lsb) + result
   }
   return result
}

func printFile(filename string) {
   file, err := os.Open(filename)
   if err != nil{
      panic(err)
   }
   scanner := bufio.NewScanner(file)
   for scanner.Scan() {
      fmt.Println(scanner.Text())
   }
}

func main() {
   //fortest()
   //fmt.Println(convertToBin(5))
   printFile("abc.txt")
}
```

## 指针、参数传递

```
值传递：引用参数时，copy一份到当前工作流中，不对源参数进行修改
引用传递：直接引用源参数，在原有参数上进行改动

GO语言参数传递只有一种方式：值传递
// 使用指针修改源参数
var arr1 [5]int 
arr3 := [...]int{5, 6, 7, 8, 9} 

func platArry(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr{
		fmt.Println(i, v)
	}
}

func main() {
	platArry(&arr1)
	platArry(&arr3)
	fmt.Println(arr1, arr3)
}
// 使用指针方式调用参数会改变原有值的内容【'*'指定，'&'调用】
>>>[100 0 0 0 0] [100 6 7 8 9]

```

## Go 数组(Range)

```
数组定义方式：
	//var arr1 [5]int  
	//arr2 := [3]int{1, 2, 3}	
	arr3 := [...]int{4, 5, 6, 7, 8, 9} 
	//var grid [4][5]int	
	//fmt.Println(arr1, arr2, arr3, grid)
	//for i:= 0; i < len(arr3); i++{
	//	fmt.Println(arr3[i])
	//}
	
数组是值类型：在函数调用数组时使用的是值传递，不会改变原有数组的内容


Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

arr3 := [...]int{4, 5, 6, 7, 8, 9} // 不确定有数组的位数用[...]代表
for i, v := range arr3{
		fmt.Println(i, v)
	}
```

## 切片

```
Slice(切片)
// []int代表slice， [5]int代表数组
func updateslice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[:]
	s2 := arr[:]
	fmt.Println(s1)
	fmt.Println(s2)
	
	updateslice(s1)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(arr)
}
------------------------
[0 1 2 3 4 5 6 7]
[0 1 2 3 4 5 6 7]
[100 1 2 3 4 5 6 7]
[100 1 2 3 4 5 6 7]
[100 1 2 3 4 5 6 7]

使用slice（切片）
更改源数据：
s1、s2均指向arr，修改s1的时候arr的值发生变更，同时指向arr的s2也会因为s1的修改发生改变。

```











