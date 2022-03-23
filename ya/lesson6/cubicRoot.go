package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, c, d int
	var m float64

	fmt.Fscan(os.Stdin, &a)
	fmt.Fscan(os.Stdin, &b)
	fmt.Fscan(os.Stdin, &c)
	fmt.Fscan(os.Stdin, &d)

	l := -1000.0
	r := 1000.0

	eps := 0.000001

	for r - l > eps {
		m = (r + l) / 2
		fmt.Println(float64(a) * m * m * m + float64(b) * m * m + float64(c) * m + float64(d))
		if float64(a) * m * m * m + float64(b) * m * m + float64(c) * m + float64(d) > 0 {
			r = m
		} else {
			l = m
		}
	}
	fmt.Println(l)
}
