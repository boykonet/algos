package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}

func lenNumb(num int) int {
	l := 0

	for num != 0 {
		num = num / 10
		l++
	}
	return l
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	lenNumber := lenNumb(x)
	m := make(map[int]int, lenNumber)

	for indx := 0; indx < lenNumber; indx++ {
		m[lenNumber - indx - 1] = x % 10
		x = x / 10
	}

	for i, j := 0, lenNumber - 1; i <= j; i, j = i + 1, j - 1 {
		if m[i] != m[j] {
			return false
		}
	}
    return true
}
