package main

import "fmt"

func main() {
	//定义map
	m := map[string]string{
		"name": "azj",
		"age": "18",
		"address": "shanxi",
	}

	var s []string
	for k, v := range m{
		s := append(s, ...)
	}

	fmt.Println(s)

	//m2 := make(map[string]int) // m2 == empty map
	//var m3 map[string]int	// m3 == nil
	//fmt.Println(m, m2, m3)

	//// 获取map中的元素
	//name := m["name"]
	//fmt.Println(name)
	//
	////判断map值是否存在，通过接收参数的形式进行判断，分别为true、false
	////map元素不存在
	//if age, ok := m["agg"]; ok != true{
	//	fmt.Println(ok)
	//}else {
	//	fmt.Println(age)
	//}
	//// map元素存在
	//if address, ok := m["address"]; ok != true {
	//	fmt.Println(ok)
	//}else {
	//	fmt.Println(address, ok)
	//}

	// 删除map元素
	//fmt.Println(m)
	//delete(m, "name")
	//fmt.Println(m)

 	//s1 := [...]string{"name:azj", "age:18"}
 	//fmt.Println(s1)

	// map中的元素是无序的，需要排序时手动把元素拿出来放到slice中进行排序
	//for k, v := range m{
	//	fmt.Println(k, v)
	//}

}
