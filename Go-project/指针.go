package main

import "fmt"
// 值传递
func swap(a, b int) {
	a, b = b, a
}
// 引用传递，
func swapone(a, b *int) {
	*a, *b = *b, *a
}

// 相当于引用传递，修改源参数内容
func swaptow(a, b int) (int, int)  {
	return b, a
}

func main() {
	a, b := 3, 4
	//swap(a, b)
	//fmt.Println(a, b)
	//
	//// 引用指针【&】
	//swapone(&a, &b)
	//fmt.Println(a, b)

	a, b = swaptow(a, b)
	fmt.Println(a, b)
}
