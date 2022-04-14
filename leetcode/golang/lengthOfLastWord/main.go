package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))

}

func lengthOfLastWord(s string) int {
	counter := 0
	lastIndex := len(s) - 1

	for s[lastIndex] == 32 {
		lastIndex--
	}
	for ; lastIndex >= 0 && s[lastIndex] != 32; lastIndex-- {
		counter++
	}
	return counter
}
