package main

import (
	"fmt"
)

func aaa(a, b int, op string) (int, error) {
	switch op {
	case "+" :
		return a + b, nil
	case "-" :
		return a - b, nil
	case "*" :
		return a * b, nil
	case "/" :
		return a / b, nil
	default:
		return 0, fmt.Errorf(
		op)
	}
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers{
		s += numbers[i]  // [i]是序列号，代表参数的第几个值
	}
	return s
}

func main() {
	fmt.Println(sum(1,2,3,4,5))

	//if resout, err := aaa(3, 4, "99"); err !=nil{
	//	fmt.Println("erroe: ", err)
	//}else {
	//	fmt.Println(resout)
	//}

}
