package main

import (
	"fmt"
)

func main() {
	arr1 := append([]int{}, 1)
	num := removeDuplicates(arr1)
	fmt.Println(num, arr1)

	arr2 := append([]int{}, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4)
	num = removeDuplicates(arr2)
	fmt.Println(num, arr2)
}

func removeDuplicates(nums []int) int {
	currNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if currNum == nums[i] {
			nums[i] = 101
		} else {
			currNum = nums[i]
		}
	}

	count := 0
	for i := 0; i < len(nums); {
		if count < len(nums) && nums[count] == 101 {
			for i < len(nums) && nums[i] == 101 {
				i++
			}
			if i >= len(nums) {
				break
			}
			nums[count], nums[i] = nums[i], nums[count]
		}
		i++
		count++
	}
	nums = nums[:count]
	return count
}
