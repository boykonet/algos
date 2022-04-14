package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int][]int, 0)

	for indx, value := range nums {
		m[value] = append(m[value], indx)
	}

	for firstIndx, value := range nums {
		diff := target - value
		diffIndxes, ok := m[diff]
		if ok == true {
			if len(diffIndxes) == 1 && diff != value {
				return append([]int{firstIndx}, diffIndxes[0])
			}
			if len(diffIndxes) > 1 && diff == value {
				return append([]int{firstIndx}, diffIndxes[1])
			}
		}
	}

	return make([]int, 0)
}
