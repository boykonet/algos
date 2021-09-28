package main

import (
	"os"
	"fmt"
	"bufio"
	"sort"
)

type A struct {
	Coordinate	int
	Event		int
}


func ReadData(reader *bufio.Reader, count int) []A {
	var l, r int
	var a []A
	var mKeys []int

	m := make(map[int][]int)

	for i := 0; i < count; i++ {
		fmt.Fscan(reader, &l)
		fmt.Fscan(reader, &r)
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

func SegmentCalculation(a []A) int64 {
	var lenSegment int64
	var countEvents int

	for indx, _ := range a {
		if countEvents > 0 {
			lenSegment += int64(a[indx].Coordinate) - int64(a[indx - 1].Coordinate)
		}
		if a[indx].Event == -1 {
			countEvents++
		} else if a[indx].Event == 1 {
			countEvents--
		}
	}
	return lenSegment
}

func main() {
	var count int

	reader := bufio.NewReaderSize(os.Stdin, 32768)

	fmt.Fscan(reader, &count)

	a := ReadData(reader, count)

	lenSegment := SegmentCalculation(a)

	fmt.Println(lenSegment)
}
