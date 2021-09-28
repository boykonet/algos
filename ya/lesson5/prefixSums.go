package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var sizeArray, numOfRequest, l, r int
	var num int64

	reader := bufio.NewReaderSize(os.Stdin, 100000)

	fmt.Fscan(reader, &sizeArray)
	fmt.Fscan(reader, &numOfRequest)

	array := make([]int64, sizeArray)
	prefixSum := make([]int64, sizeArray + 1)

	for i := 0; i <= sizeArray; i++ {
		if i < sizeArray {
			fmt.Fscan(reader, &num)
			array[i] = num
		}
		if i > 0 {
			prefixSum[i] = array[i - 1] + prefixSum[i - 1]
		}
	}
	fmt.Println(array)
	fmt.Println(prefixSum)

	for i := 0; i < numOfRequest; i++ {
		fmt.Fscan(reader, &l)
		fmt.Fscan(reader, &r)

		fmt.Println(prefixSum[r] - prefixSum[l - 1])
	}
}
