package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

func numbers(s string) []int {
	var nums []int

	for _, str := range strings.Fields(s) {
		i, err := strconv.Atoi(str)
		if err == nil {
			nums = append(nums, i)
		}
	}
	return nums
}

func main() {
	var data, arr []int
	left := -1
	right := -1
	center := -1

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	data = numbers(scan.Text())
	scan.Scan()
	arr = numbers(scan.Text())

	parity := data[0] % 2

	for i := 0; i < data[1]; i++ {
		if parity == 1 && arr[i] == data[0] / 2 {
			center = arr[i]
			fmt.Println(center)
			break
		} else {
			if arr[i] < data[0] / 2 {
				left = arr[i]
			} else if right == -1 {
				fmt.Println(left, arr[i])
				break
			}
		}
	}
}
