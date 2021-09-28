package main

import (
	"fmt"
	"os"
)

func main() {
	var n, l, count int64

	fmt.Fscan(os.Stdin, &n, &l)

	count = l / n
	if count * n < l {
		count++
	}

	fmt.Println(count)
}
