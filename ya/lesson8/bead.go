package main

import (
	"fmt"
	"os"
	"bufio"
	"sort"
)

func AddElem(m map[int][]int, first, second int) {
	_, ok := m[first]
	if ok == false {
		m[first] = append(m[first], 0)
	}
	m[first] = append(m[first], second)
	_, ok = m[second]
	if ok == false {
		m[second] = nil
	}
}

func FindHead(m map[int][]int) int {
	for key, value := range m {
		if len(value) > 0 && value[0] == 0 {
			return key
		}
	}
	return 0
}

func FindNilElements(m map[int][]int) []int {
	var elements []int

	for key, value := range m {
		if len(value) == 0 {
			elements = append(elements, key)
		}
	}
	return elements
}

func CountLen(m map[int][]int, head, elem, count int, max *int) {
	s := m[head]

	if elem == head {
		*max = count
		return
	}
	for i := 0; i < len(s); i++ {
		CountLen(m, s[i], elem, count + 1, max)
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
		CountLen(m, head, v, 0, &max)
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
	head := FindHead(m)
	m[head] = m[head][1:]

	e := FindNilElements(m)
	fmt.Println(FindDiameter(m, e, head))
}
