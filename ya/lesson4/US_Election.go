package main

import (
	"os"
	"fmt"
	"bufio"
	"sort"
)

func main() {
	var key string
	var value int
	var keys []string

	m := make(map[string]int)

	Reader := bufio.NewReaderSize(os.Stdin, 500000)

	for ; ; {
		_, err := fmt.Fscan(Reader, &key, &value)
		if err != nil {
			break
		}
		_, ok := m[key]
		if ok == false {
			m[key] = 0
		}
		m[key] += value
	}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
