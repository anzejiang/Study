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
	//	fmt.Println(err)
	//}else {
	//	fmt.Printf("%s\n", contents)
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