package main

import (
	"bufio"
	"fmt"
	"os"
)

func testone(filename string) {
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func main() {
	testone("abc.txt")
}
