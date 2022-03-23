package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var cost int

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	word := scan.Text()

	i := 0
	j := len(word) - 1
	for ; i < j; {
		if word[i] != word[j] {
			cost++
		}
		i++
		j--
	}
	fmt.Println(cost)
}
