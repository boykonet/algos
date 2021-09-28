package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func LeftBinarySearch(l, r, num int, array []int) int {
	for l < r {
		m := (l + r) / 2
		if array[m] >= num {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func RightBinarySearch(l, r, num int, array []int) int {
	for l < r {
		m := (l + r + 1) / 2
		if array[m] <= num {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}

func main() {
	var countNums, countRequest, num int
	var nums []int

	reader := bufio.NewReaderSize(os.Stdin, 200000)

	fmt.Fscan(reader, &countNums)
	for i := 0; i < countNums; i++ {
		fmt.Fscan(reader, &num)
		nums = append(nums, num)
	}

	fmt.Fscan(reader, &countRequest)

	result := make([][]int, countRequest)

	begin := 0
	end := len(nums) - 1
	for i := 0; i < countRequest; i++ {
		fmt.Fscan(reader, &num)
		lefti := LeftBinarySearch(begin, end, num, nums)
		righti := RightBinarySearch(begin, end, num, nums)
		if nums[lefti] == num {
			result[i] = append(result[i], lefti + 1)
		} else {
			result[i] = append(result[i], 0)
		}
		if nums[righti] == num {
			result[i] = append(result[i], righti + 1)
		} else {
			result[i] = append(result[i], 0)
		}
	}

	for i := 0; i < countRequest; i++ {
		fmt.Println(strings.Trim(fmt.Sprint(result[i]), "[]"))
	}
}

