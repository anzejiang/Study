package main

import "fmt"

func platArry(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr{
		fmt.Println(i, v)
	}
}


func main() {
	var arr1 [5]int  // 使用var关键字定义的数组为五个空数组
	//arr2 := [3]int{1, 2, 3}	// 使用:=定义的数组需要赋予初值
	arr3 := [...]int{5, 6, 7, 8, 9} // 不确定有数组的位数用[...]代表
	//var grid [4][5]int	// 四行五列，四个数组，每个数组有5个值
	//fmt.Println(arr1, arr2, arr3, grid)
	//for i:= 0; i < len(arr3); i++{
	//	fmt.Println(arr3[i])
	//}

	// 使用range关键子，可以获得数组的下标和值，【i=下标，v=值】
	//for i, v := range arr3{
	//	fmt.Println(i, v)
	//}

	platArry(&arr1)
	fmt.Println("~~~~~")
	platArry(&arr3)
	fmt.Println("~~~~~")
	fmt.Println(arr1, arr3)
}
