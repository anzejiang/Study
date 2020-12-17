package main

import "fmt"

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

func main() {
	if resout, err := aaa(3, 4, "99"); err !=nil{
		fmt.Println("erroe: ", err)
	}else {
		fmt.Println(resout)
	}

}
