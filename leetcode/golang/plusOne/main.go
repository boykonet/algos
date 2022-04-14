package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(plusOne([]int{1, 2, 3}))
	// fmt.Println(plusOne([]int{0}))
	// fmt.Println(plusOne([]int{9}))
	// fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
}

func plusOne(digits []int) []int {
	var res []int

	lenDigits := len(digits)
	rem := 0

	for i := 0; i < lenDigits; i++ {
		res = append(res, digits[i])
	}

	num := res[len(res) - 1] + 1

	if num <= 9 {
		res[len(res) - 1] += 1
	} else {
		rem := num / 10
		for i := lenDigits - 2; i >= 0; i-- {
			rem = num / 10
			if rem != 
		}
		if rem != 0 {
			res = append([]int{rem}, res...)
		}
	}
	return res
}

// func plusOne(digits []int) []int {
// 	var num int64

// 	lenDigits := len(digits)
// 	res := make([]int, 0)

// 	for i := 0; i < lenDigits; i++ {
// 		pow := float64(lenDigits - i - 1)
// 		num += float64(digits[i]) * float64(math.Pow(10., pow))
// 		fmt.Println(num)
// 	}
// 	num += 1.
// 	for num != 0 {
// 		rem := int(num % 10)
// 		res = append([]int{rem}, res...)
// 		num = num / 10
// 	}
// 	return res
// }
