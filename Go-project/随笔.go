package main

import "fmt"

func main() {

	arr := [...]int{0,1,2,3,4,5,6,7}
	s2 := arr[2:]
	//tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	//fmt.Println(tail, s2)
	fmt.Println(s2)

}
