package main

import (
	"fmt"
	"go-Study/Go-project/azj/test"
)

type Azj interface {
	Search(a string) string
}



func main() {
	aa := test.Azj{"azj", 18}
	//fmt.Println(aa)
	//fmt.Println(aa.Username)
	fmt.Println(aa.Search("ZZZZZ"))
}
