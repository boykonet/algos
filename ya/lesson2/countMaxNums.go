package main

import (
	"os"
	"fmt"
)

func main() {
	var v, count, max int

	v = -1
	for i := 0; v != 0; i++ {
		if max < v {
			max = v
			count = 1
		} else if max == v {
			count++
		}
		fmt.Fscan(os.Stdin, &v)
	}
	fmt.Println(count)
}
