package main

import (
	"bufio"
	"fmt"
	"go-Study/Go-project/azj/fib"
	"os"
)

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	defer write.Flush()

	f := fib.Fibonacci()
	for i := 0;i < 20; i++{
	fmt.Fprintln(write, f())
	//fmt.Println(f())
	}
}

func main() {
	writeFile("fib.txt")
}
