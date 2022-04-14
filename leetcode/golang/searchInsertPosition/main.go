package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))

	fmt.Println(searchInsert([]int{1}, 7))
	fmt.Println(searchInsert([]int{1}, 1))
}

func searchInsert(nums []int, target int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
		if nums[i] < target {
			res = i + 1
		} else {
			return res
		}
	}
	return res
}
