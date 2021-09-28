package main

import (
	"fmt"
	"os"
)

func main() {
	var first, last, year int64

	mdays := [12]int64{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	fmt.Fscan(os.Stdin, &first, &last, &year)
	if first <= 12 && last <= 12 && first != last || first > 12 && last > 12 {
		fmt.Println(0)
		return
	}
	if first <= 12 && last > 12 {
		if last > mdays[first - 1] {
			fmt.Println(0)
			return
		}
	} else if last <= 12 && first > 12 {
		if first > mdays[last - 1] {
			fmt.Println(0)
			return
		}
	}
	fmt.Println(1)
}
