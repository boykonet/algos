package main

import (
	"os"
	"fmt"
	"bufio"
	"sort"
	"strings"
)

type KeyValue struct {
	Key		string
	Value	int
}

func Min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

func main() {
	var key string
	var sortedKeyValue []KeyValue

	Reader := bufio.NewReaderSize(os.Stdin, 500000)

	m := make(map[string]int)

	for ;; {
		_, err := fmt.Fscan(Reader, &key)
		if err != nil {
			break
		}
		_, ok := m[key]
		if ok == false {
			m[key] = 0
		}
		m[key] += 1
	}

	for k, v := range m {
		sortedKeyValue = append(sortedKeyValue, KeyValue{ k, v })
	}

	sort.Slice(sortedKeyValue, func(i, j int) bool {
		return sortedKeyValue[i].Value > sortedKeyValue[j].Value ||
		(sortedKeyValue[i].Value == sortedKeyValue[j].Value &&
		strings.Compare(sortedKeyValue[i].Key, sortedKeyValue[j].Key) == -1)
	})

	for _, k_v := range sortedKeyValue {
		fmt.Println(k_v.Key)
	}
}
