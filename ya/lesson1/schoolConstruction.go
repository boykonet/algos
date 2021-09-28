package main

import (
	"math"
	"fmt"
	"os"
)

func main() {
	var nums, v, indx, result int32

	fmt.Fscan(os.Stdin, &nums)

	indx = int32(math.Round(float64(nums - 1) / 2.0))

	for i := int32(0); i < indx + 1; i++ {
		fmt.Fscan(os.Stdin, &v)
		if i == indx {
			result = v
		}
	}
	fmt.Println(result)
}
