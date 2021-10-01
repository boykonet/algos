package main

import (
	"fmt"
	"os"
	"bufio"
	"sort"
)

func AddElem(m map[int][]int, first, second int) {
	m[first] = append(m[first], second)
	m[second] = append(m[second], first)
}

func FindSheets(m map[int][]int) []int {
	var elements []int

	for key, value := range m {
		if len(value) == 1 {
			elements = append(elements, key)
		}
	}
	return elements
}

func CountLen(m map[int][]int, vertexs map[int]bool, head, elem, count int, max *int) {
	s := m[head]

	if elem == head {
		*max = count
		return
	}
	for i := 0; i < len(s); i++ {
		vertexs[s[i]] = true
		CountLen(m, vertexs, s[i], elem, count + 1, max)
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindDiameter(m map[int][]int, nilElements []int, head int) int {
	var lens []int
	var max int


	if len(nilElements) == 0 {
		return 0
	}
	for _, v := range nilElements {
		CountLen(m, make(map[int]bool), head, v, 0, &max)
		lens = append(lens, max)
	}
	sort.Ints(lens)
	if len(nilElements) == 1 {
		return lens[0] + 1
	}
	return lens[len(lens) - 1] + lens[len(lens) - 2] + 1
}

func main() {
	var countBeads, l, r int

	reader := bufio.NewReaderSize(os.Stdin, 32768)

	fmt.Fscan(reader, &countBeads)

	m := make(map[int][]int)

	for i := 0; i < countBeads - 1; i++ {
		fmt.Fscan(reader, &l, &r)
		AddElem(m, l, r)
	}

	fmt.Println(m)

	e := FindSheets(m)
	sort.Ints(e)
	fmt.Println(e)
	// fmt.Println(FindDiameter(m, e, head))
}
