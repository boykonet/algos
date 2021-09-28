package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	var v int

	m := make(map[int]bool)

	Reader := bufio.NewReaderSize(os.Stdin, 100000)

	for ;; {
		_, err := fmt.Fscan(Reader, &v)
		if err != nil {
			break
		}
		_, value := m[v]
		if value == true {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
			m[v] = true
		}
	}
}
