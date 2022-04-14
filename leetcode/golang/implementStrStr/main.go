package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("hello", ""))
	fmt.Println(strStr("", "ll"))
	fmt.Println(strStr("", ""))
}

func strStr(haystack, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		j := 0
		for i + j < len(haystack) && j < len(needle) && haystack[i + j] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i
		}
		if len(haystack[i:]) < len(needle) {
			break
		}
	}
	return -1
}
