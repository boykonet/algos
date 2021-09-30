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

func CountLen(m map[int][]int, head, elem, count int) int {
	s := m[head]

	if elem == head {
		return count
	} else {
		if s == nil {
			return count
		}
		for i := 0; i < len(s); i++ {
			return CountLen(m, s[i], elem, count + 1)
		}
	}
	return count



	// e := m[head]

	// fmt.Printf("count = %v, elem = %v, head = %v\n", count, elem, head)
	// if head == elem {
	// 	return count
	// } else {
	// 	count++
	// 	for i := 0; i < len(e); i++ {
	// 		if head == elem {
	// 			return CountLen(m, head, elem, count)
	// 		} else {
	// 			return CountLen(m, e[i], elem, count)
	// 		}
	// 	}
	// }
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindDiameter(m map[int][]int, nilElements []int, head int) int {
	var lens []int

	if len(nilElements) == 0 {
		return 0
	}
	for _, v := range nilElements {
		a := CountLen(m, head, v, 0)
		fmt.Println("a =", a)
		lens = append(lens, a)
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
	fmt.Println(m)
	m[head] = m[head][1:]
	// fmt.Println(head)

	e := FindNilElements(m)
	fmt.Println(FindDiameter(m, e, head))
	// fmt.Println(CountLen(m, head, e[0], 0))
	// fmt.Println(CountLen(m, head, e[1], 0))
	// fmt.Println(CountLen(m, head, e[2], 0))
}
