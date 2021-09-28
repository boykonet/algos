package main

import (
	"os"
	"fmt"
	"bufio"
	"sort"
)

type A struct {
	Point		int
	Event		int
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ReadData(reader *bufio.Reader) []A {
	var l, r int
	var a []A
	var mKeys []int

	m := make(map[int][]int)

	for ; ; {
		fmt.Fscan(reader, &l)
		fmt.Fscan(reader, &r)
		if l == 0 && r == 0 {
			break
		}
		m[l] = append(m[l], -1)
		m[r] = append(m[r], 1)
	}

	for key, _ := range m {
		mKeys = append(mKeys, key)
		sort.Ints(m[key])
	}
	sort.Ints(mKeys[:])

	for _, key := range mKeys {
		events := m[key]
		for _, value := range events {
			a = append(a, A{ key, value })
		}
	}
	return a
}

// func CountDevices(a []A) int {
// 	var max int
// 	var countEvents int

// 	for indx, _ := range a {
// 		if a[indx].Event == 1 {
// 			countEvents++
// 		} else if a[indx].Event == -1 {
// 			countEvents--
// 		}
// 		max = Max(max, countEvents)
// 	}
// 	return max
// }

func main() {
	var lenSegment int

	reader := bufio.NewReaderSize(os.Stdin, 32768)

	fmt.Fscan(reader, &lenSegment)

	a := ReadData(reader)

	fmt.Println(lenSegment, a)

	// count := CountDevices(a)

	// fmt.Println(count)
}


