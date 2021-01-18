package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func writefile(filename string)  {
	var a [5]int

	file, err := os.Create(filename)
	if err != nil {
		log.Print(err.Error())
	}
	defer file.Close()

	f := bufio.NewWriter(file)
	defer f.Flush()
	for i := range a{
		fmt.Fprintln(f, i)
	}
}

func readfiletow(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fmt.Println(string(content))
}

func main() {
	//writefile("abc.txt")
	//readfiletow("abc.txt")

}
