package main

import "fmt"

func updateslice(s []int) {
	s[0] = 100
}


func main() {

	// 10为len（s1）长度，32为cap空间长度
	s1 := make([]int, 10, 32)
	fmt.Println(s1, len(s1), cap(s1))
	s2 := []int{0,1,2,3,4}

	//// 将s2拷贝到s1中 s1 = [0 1 2 3 4 0 0 0 0 0]
	copy(s1, s2)
	//fmt.Println(s1, s2)

	//// 删除一个selice中的元素，通过切片组合的形式达到删除效果，[4:]...代表下标为4之后的所有元素
	//s1 = append(s1[:3], s1[4:]...)
	//fmt.Println(s1)

	// 掐头去尾
	fmt.Println(s1)
	fmt.Println("Popping from front")
	front := s1[0]
	s1 = s1[1:]
	fmt.Println(front, "\n", s1)

	fmt.Println("Popping from back")
	tail := s1[len(s1)-1]
	s1 = s1[:len(s1)-1]
	fmt.Println(tail, "\n", s1)

	//var s []int
	//for i := 0; i < 10; i++ {
	//	s = append(s, 2 * i +1)
	//}
	//fmt.Println(s)

	//arr := [...]int{0,1,2,3,4,5,6,7}
	////fmt.Println(arr[2:6])
	////fmt.Println(arr[:6])
	////fmt.Println(arr[2:])
	////fmt.Println(arr[:])
	//s1 := arr[:]
	//s2 := arr[:]
	//fmt.Println(s1)
	//fmt.Println(s2)
	//
	//updateslice(s1)
	//fmt.Println(s1)
	//fmt.Println(s2)
	//fmt.Println(arr)
}
