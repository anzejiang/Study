package test

import "fmt"

type Azj struct {
	Username string
	Age int
}


func (b Azj) Search(a string) (string, int) {
	fmt.Printf("The is a test templet. %s", a)
	return b.Username, b.Age
}

