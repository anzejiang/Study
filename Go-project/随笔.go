package main

import "fmt"

//func writeFile(filename string) {
//	file, err := os.Create(filename)
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//
//	// 使用bufio写文件更快，先读到内存，使用flush刷新写到文件
//	write := bufio.NewWriter(file)
//	defer write.Flush()
//}

//func tuzi() func() int {
//	a, b := 0, 1
//	return func() int {
//			a, b = b, a + b
//		return a
//	}
//}

func main() {
	//f := tuzi()
	//for i := 0; i < 20; i++ {
	//	//fmt.Println(tuzi()())
	//	fmt.Println(f())
	//}

	var a [5]int
	for i, _ := range a{
		fmt.Print(i, " ")
	}
}
