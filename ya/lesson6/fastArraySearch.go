package main

import (
	"fmt"
	"os"
	"bufio"
	"sort"
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
	var countNums, countRequest, num, l, r int
	var nums, result []int

	reader := bufio.NewReaderSize(os.Stdin, 200000)

	fmt.Fscan(reader, &countNums)
	for i := 0; i < countNums; i++ {
		fmt.Fscan(reader, &num)
		nums = append(nums, num)
	}
	sort.Ints(nums[:])

	fmt.Fscan(reader, &countRequest)

	begin := 0
	end := len(nums) - 1
	for i := 0; i < countRequest; i++ {
		fmt.Fscan(reader, &l, &r)
		lefti := LeftBinarySearch(begin, end, l, nums)
		righti := RightBinarySearch(begin, end, r, nums)
		if nums[lefti] >= l && nums[righti] <= r {
			result = append(result, righti - lefti + 1)
		} else {
			result = append(result, 0)
		}
	}

	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}
