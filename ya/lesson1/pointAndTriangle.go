package main

import (
	"math"
	"os"
	"fmt"
)

func main() {
	var b, x0, y0 int

	inTriangle := 0
	pointA := 1
	pointB := 2
	pointC := 3

	fmt.Fscan(os.Stdin, &b, &x0, &y0)
	xa := 0
	ya := 0

	xb := b
	yb := 0

	xc := 0
	yc := b

	resA := (xa - x0) * (yb - ya) - (xb - xa) * (ya - y0)
	resB := (xb - x0) * (yc - yb) - (xc - xb) * (yb - y0)
	resC := (xc - x0) * (ya - yc) - (xa - xc) * (yc - y0)

	if (resA <= 0 && resB <= 0 && resC <= 0) ||
	(resA >= 0 && resB >= 0 && resC >= 0) {
		fmt.Println(inTriangle)
	} else {
		vecA := math.Pow(math.Pow(float64(xa - x0), 2.0) + math.Pow(float64(ya - y0), 2.0), 0.5)
		vecB := math.Pow(math.Pow(float64(xb - x0), 2.0) + math.Pow(float64(yb - y0), 2.0), 0.5)
		vecC := math.Pow(math.Pow(float64(xc - x0), 2.0) + math.Pow(float64(yc - y0), 2.0), 0.5)
		if vecA <= vecB && vecA <= vecC {
			fmt.Println(pointA)
		} else if vecB <= vecC && vecB < vecA {
			fmt.Println(pointB)
		} else {
			fmt.Println(pointC)
		}
	}
}
