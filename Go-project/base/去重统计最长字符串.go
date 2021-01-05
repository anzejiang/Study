package main

import "fmt"

func lengthOFNonRepeatingSubStr(s string) int  {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start{
			start = lastI +1
		}
		if i - start + 1 > maxLength{
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	fmt.Println()
	return  maxLength
}

func main() {
	fmt.Print(lengthOFNonRepeatingSubStr("asdfasdfasdfasdfasdf"))
	fmt.Print(lengthOFNonRepeatingSubStr("abcdef"))
	fmt.Print(lengthOFNonRepeatingSubStr("aa"))
	fmt.Print(lengthOFNonRepeatingSubStr("abca"))
	fmt.Print(lengthOFNonRepeatingSubStr("我爱慕课网"))
	fmt.Print(lengthOFNonRepeatingSubStr("一二三二一"))

}
