package main

import (
	"fmt"
	"os"
	"bufio"
	"sort"
	"strings"
)

func GetSortArrayAndIndices(reader *bufio.Reader, countRead int) ([]int, map[int][]int) {
	var num int
	var arr []int

	m := make(map[int][]int)

	for indx := 0; indx < countRead; indx++ {
		fmt.Fscan(reader, &num)
		arr = append(arr, num)
		m[num] = append(m[num], indx)
	}
	sort.Ints(arr)
	return arr, m
}

func GetIndexFromMultiply(m map[int][]int, key int) int {
	indices_for_key := m[key]
	first_indx := indices_for_key[0]
	m[key] = indices_for_key[1:]

	return first_indx
}

func main() {
	var n, m, count int

	reader := bufio.NewReaderSize(os.Stdin, 8192)

	fmt.Fscan(reader, &n, &m)

	n_arr, n_map := GetSortArrayAndIndices(reader, n)
	m_arr, m_map := GetSortArrayAndIndices(reader, m)

	res := make([]int, n)
	count_n := 0
	for count_m, _ := range m_arr {
		if m_arr[count_m] >= n_arr[count_n] + 1 {

			indx_n := GetIndexFromMultiply(n_map, n_arr[count_n])
			indx_m := GetIndexFromMultiply(m_map, m_arr[count_m])

			res[indx_n] = indx_m + 1
			count++
			count_n++
		}
	}
	fmt.Println(count)
	fmt.Println(strings.Trim(fmt.Sprint(res), "[]"))
}
