package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	var v, count int

	Reader := bufio.NewReaderSize(os.Stdin, 200000)
	m := make(map[int]int)


	for ; ; {
		_, err := fmt.Fscan(Reader, &v)
		if err != nil {
			break
		}
		_, ok := m[v]
		if ok == true {
			m[v] += 1
			count++
		} else {
			m[v] = 1
		}
	}
	fmt.Println(count)
}
