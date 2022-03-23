package main

import (
	"os"
	"fmt"
)

func max(arr []int) int {
	len_arr := len(arr)
	max := arr[0]

	for i := 1; i < len_arr; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {
	var arr []int
	var data, v, count int

	fmt.Fscan(os.Stdin, &data)
	for i := 0; i < data; i++ {
		fmt.Fscan(os.Stdin, &v)
		arr = append(arr, v)
		count += v
	}
	fmt.Println(count - max(arr))
}
