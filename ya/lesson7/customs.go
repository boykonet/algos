package main

import (
	"os"
	"fmt"
	"bufio"
	"sort"
)

type A struct {
	Time		int
	Event		int
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ReadData(reader *bufio.Reader, count int) []A {
	var l, r int
	var a []A
	var mKeys []int

	m := make(map[int][]int)

	for i := 0; i < count; i++ {
		fmt.Fscan(reader, &l)
		fmt.Fscan(reader, &r)
		m[l] = append(m[l], 1)
		m[r + l] = append(m[r + l], -1)
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

func CountDevices(a []A) int {
	var max int
	var countEvents int

	for indx, _ := range a {
		if a[indx].Event == 1 {
			countEvents++
		} else if a[indx].Event == -1 {
			countEvents--
		}
		max = Max(max, countEvents)
	}
	return max
}

func main() {
	var countNums int

	reader := bufio.NewReaderSize(os.Stdin, 32768)

	fmt.Fscan(reader, &countNums)

	a := ReadData(reader, countNums)

	count := CountDevices(a)

	fmt.Println(count)
}

