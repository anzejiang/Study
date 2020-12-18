package main

import "fmt"

func updateslice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	//fmt.Println(arr[2:6])
	//fmt.Println(arr[:6])
	//fmt.Println(arr[2:])
	//fmt.Println(arr[:])
	s1 := arr[:]
	s2 := arr[:]
	fmt.Println(s1)
	fmt.Println(s2)

	updateslice(s1)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(arr)
}
