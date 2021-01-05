package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

func main() {
	f := fibonacci()
	fmt.Printf("%d\n", f())

	//arr := [...]int{0,1,2,3,4,5,6,7}
	//s2 := arr[:]
	////tail := s2[len(s2)-1]
	//s2 = s2[:len(s2)-1]
	////fmt.Println(tail, s2)
	//fmt.Println(s2)

}
