package main

import (
	"bufio"
	"os"
)

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 使用bufio写文件更快，先读到内存，使用flush刷新写到文件
	write := bufio.NewWriter(file)
	defer write.Flush()
}

func main() {

}
