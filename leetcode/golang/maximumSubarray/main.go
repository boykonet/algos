package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{-1}))
	fmt.Println(maxSubArray([]int{-2, -1}))
	fmt.Println(maxSubArray([]int{-1, -2}))
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := 0

	for i := 0; i < len(nums); i++ {
		currSum += nums[i]
		if currSum > maxSum {
			maxSum = currSum
		}
		if currSum < 0 {
			currSum = 0
		}
	}
	return maxSum
}
