package main

import (
	"os"
	"fmt"
	"bufio"
	"sort"
)

func main() {
	var count, key, value, indx int64
	var keys []int64

	Reader := bufio.NewReaderSize(os.Stdin, 1000000)

	m := make(map[int64]int64)

	fmt.Fscan(Reader, &count)

	for ; indx < count; indx++ {
		fmt.Fscan(Reader, &key, &value)
		_, ok := m[key]
		if ok == false {
			m[key] = 0
		}
		m[key] += value
	}
	for k, _ := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
