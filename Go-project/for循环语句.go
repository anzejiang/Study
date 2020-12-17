package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// for循环
func fortest() {
	sum := 0
	for i := 1; i <=100; i++{
		sum += i
	}
	fmt.Println(sum)
}

//for二进制转换
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//for文件读取
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

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation: " + op)
	}
}

func div(a, b int) (int, int)  {
	return a + b, a * b

}

// 死循环
func azjstree() {
	for {
		fmt.Println(123)
	}
}

func main() {
	//fortest()
	//fmt.Println(convertToBin(5))
	//printFile("abc.txt")
	//fmt.Println(eval(3, 4, "*"))
	//fmt.Println(div(3, 4))
	azjstree()
}
