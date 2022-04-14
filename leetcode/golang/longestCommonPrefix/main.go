package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
}

func longestCommonPrefix(s []string) string {
	res := ""

	sort.Strings(s)
	firstString := s[0]
	lastString := s[len(s) - 1]
	for i, j := 0, 0; i < len(firstString) && j < len(lastString); i, j = i + 1, j + 1 {
		if firstString[i] != lastString[j] {
			break
		}
		res += string(firstString[i])
	}
	return res
}
