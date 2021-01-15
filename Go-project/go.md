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

*、&：*代表引用指针，&调用指针

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

## 切片Slice

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

## 字符串操作

```
// 函数：strings.

// 找出不重复最长的字符串，
package main

import "fmt"

func lengthOFNonRepeatingSubStr(s string) int  {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start{
			start = lastI +1
		}
		if i - start + 1 > maxLength{
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	fmt.Println()
	return  maxLength
}

func main() {
	fmt.Print(lengthOFNonRepeatingSubStr("asdfasdfasdfasdfasdf"))
	fmt.Print(lengthOFNonRepeatingSubStr("abcdef"))
	fmt.Print(lengthOFNonRepeatingSubStr("aa"))
	fmt.Print(lengthOFNonRepeatingSubStr("abca"))
	fmt.Print(lengthOFNonRepeatingSubStr("我爱慕课网"))
	fmt.Print(lengthOFNonRepeatingSubStr("一二三二一"))

}

// 字符串操作
// 英文一字节，中文三字节
package main

import (
	"fmt"
	"unicode/utf8"
)
// 使用range便利pos， rune
// 使用rtf8.RuneCountInString获得字符数量
// 使用len获得字节长度
// 使用[]byte获得字节

func main() {
	s := "Yes我爱慕课网!"
	fmt.Println(s)
	
	for _, b := range []byte(s){	// 获取s字符串的字节，进行遍历输出
		fmt.Printf("%X ", b) // utf-8编码
	}
	
	fmt.Println()
	
	for i, ch := range s{	// ch is a rune
		fmt.Printf("(%d %X) ", i, ch) // unicode编码
	}
	
	fmt.Println()
	// utf8.RuneCountInString获取字符串的长度，长度为：9
	fmt.Println("rune count:", utf8.RuneCountInString(s))
	
	byths := []byte(s)	// 获取s的字节
	for len(byths) > 0{
		ch, size  := utf8.DecodeRune(byths)
		byths = byths[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s){
		fmt.Printf("(%d %c) ", i, ch)
	}
}
```

## 面向对象

```
// 结构体：
数组与切片，只能存储同一类型的变量。若要存储多个类型的变量，就需要用到结构体，它是将多个容易类型的命令变量组合在一起的聚合数据类型。
每个变量都成为该结构体的成员变量。

一个目录代表一个包，一个目录下只可以有一个包
函数名称首字母大写代表可以被其他程序引用，小写代表不可以被别人引用

包：
扩充系统类型或者别人的类型
1、定义别名
2、使用组合
```

## 接口interface

```
duck typing
像鸭子走路，像鸭子叫（帐的像鸭子）,那么就是鸭子
描述事物的外部行为而非内部结构

接口的定义：
1、接口由使用者定义

interface是一个接口，具体功能、方法的实现由struct、方法进行实现
interface中包含了N个方法，interface中的方法

接口分为两部分：
	一、定义：定义接口interface，及在接口interface中定义方法
	二、实现：由strukt、函数具体实现
	例：
		定义：
		// 定义接口，及方法
		package main
		
		type Retriever interface {
			Get(url string) string
		}
		// 使用接口
		func download(r Retriever) string  {
			return r.Get("http://www.imooc.com")
		}
		// 调用接口
		func main() {
            var r Retriever		// 定义r 为接口
            r = mock.Retriever{"The is a fake Content"} // 相当于实例化接口
            fmt.Printf("%T %v\n", r, r)

            r = &real.Retriever{
                UserAgent: "Mozilla/5.0",
                TimeOut: time.Minute,
            }
            fmt.Printf("%T %v\n", r, r)
            //fmt.Println(download(r))
        }
		
		实现：
		//  实现接口
		package real
		
		type Retriever struct {
            UserAgent string
            TimeOut time.Duration
         }
		// 实现方法
		func (r *Retriever) Get(url string) string {
            resp, err := http.Get(url)
            if err != nil{
                panic(err)
            }
            result , err := httputil.DumpResponse(resp, true)
            resp.Body.Close()
            if err != nil {
                panic(err)
            }
            return string(result)
        }
		
	接口组合：
	一个接口中包含多个接口interface
		
```

## 函数式编程

```
函数是一等公民：变量、参数、返回值都可以是函数
"正统"函数式编程：
	1、不可变性：不能有状态，只有常量和函数
	2、函数只有一个参数

闭包：
// 这个就是一个闭包，闭包就是一个函数（变量的作用域只限于在其函数内部）
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 +... + %d = %d\n",
			i, a(i))
	}
}
```

## 错误处理和资源管理

```
资源管理：当程序打开数据库在读取数据的过程中，程序错误退出，及时关闭数据库连接
使用场景：打开文件/关闭、连接数据库/断开，等等

defer（资源管理）
	在函数执行完后，执行生效，多个defer从下往上执行，倒序执行
    func writeFile(filename string) {
        file, err := os.Create(filename)
        if err != nil{
            panic(err)
        }
        defer file.Close()

		// bufio，先写道内存中，使用flush刷新到文件内，比直接写文件效率高
        write := bufio.NewWriter(file)
        defer write.Flush()

        f := fib.Fibonacci()
        for i := 0;i < 20; i++{
        fmt.Fprintln(write, f())
        }
    }
    func main() {
        writeFile("fib.txt")

```

