package main

import (
	"fmt"
	"os"
	"bufio"
)

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	var sizeArray int
	var num, max, sum, min_sum int64
	// var l, r int

	reader := bufio.NewReaderSize(os.Stdin, 100000)

	fmt.Fscan(reader, &sizeArray)

	array := make([]int64, sizeArray)
	// prefixSum := make([]int64, sizeArray + 1)

	for i := 0; i < sizeArray; i++ {
		// if i < sizeArray {
		fmt.Fscan(reader, &num)
		array[i] = num
		// }
		// if i > 0 {
		// 	prefixSum[i] = array[i - 1] + prefixSum[i - 1]
		// }
	}
	fmt.Println(array)
	fmt.Println(prefixSum)

	max = array[0]
	for r := 0; r < sizeArray; r++ {
		fmt.Println()
		fmt.Println("1. sum =", sum, "max =", max, "min_sum =", min_sum)
		sum += array[r]
		max = Max(max, sum - min_sum)
		min_sum = Min(min_sum, sum)
		fmt.Println("2. sum =", sum, "max =", max, "min_sum =", min_sum)
	}

	fmt.Println(max)
}

