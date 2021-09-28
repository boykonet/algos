package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	var v int
	var nums, res []int

	m := make(map[int]bool)

	Reader := bufio.NewReaderSize(os.Stdin, 100000)

	for ; ; {
		_, err := fmt.Fscan(Reader, &v)
		if err != nil {
			break
		}
		_, ok := m[v]
		if ok == true {
			m[v] = true
		} else {
			m[v] = false
			nums = append(nums, v)
		}
	}
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == false {
			res = append(res, nums[i])
		}
	}
	fmt.Println(strings.Trim(fmt.Sprint(res), "[]"))
}
