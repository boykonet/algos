package main

import (
	"fmt"
)

func main() {
	arr1 := append([]int{}, 3, 2, 2, 3)
	num := removeElement(arr1, 3)
	fmt.Println(num, arr1)

	arr2 := append([]int{}, 0, 1, 2, 2, 3, 0, 4, 2)
	num = removeElement(arr2, 2)
	fmt.Println(num, arr2)

}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = -1
		}
	}

	count := 0
	for i := 0; i < len(nums); {
		if count < len(nums) && nums[count] == -1 {
			for i < len(nums) && nums[i] == -1 {
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
