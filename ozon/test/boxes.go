package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var countBoxes, b, countApples int
	var boxes []int

	fmt.Fscan(os.Stdin, &countBoxes)
	for i := 0; i < countBoxes; i++ {
		fmt.Fscan(os.Stdin, &b)
		boxes = append(boxes, b)
	}

	sort.Ints(boxes[:])

	for i := 0; i < countBoxes / 2; i++ {
		a := boxes[1] - boxes[0]
		countApples += a
		boxes = boxes[2:]
	}
	fmt.Println(countApples)
}
