package main

import (
	"fmt"
	"unicode/utf8"
)

// 使用range便利pos， rune
// 使用rtf8.RuneCountInString获得字符数量
// 使用len获得字节长度
// 使用[]byte获得字节


func main() {
	s := "Yes我爱慕课网!"
	fmt.Println(s)

	for _, b := range []byte(s){
		fmt.Printf("%X ", b) // utf-8编码
	}
	fmt.Println()

	for i, ch := range s{	// ch is a rune
		fmt.Printf("(%d %X) ", i, ch) // unicode编码
	}
	fmt.Println()

	fmt.Println("rune count:", utf8.RuneCountInString(s))

	byths := []byte(s)	// 获取s的字节
	for len(byths) > 0{
		ch, size  := utf8.DecodeRune(byths)
		byths = byths[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s){
		fmt.Printf("(%d %c) ", i, ch)
	}
}
