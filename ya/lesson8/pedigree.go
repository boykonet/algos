package main

import (
	"os"
	"fmt"
	"bufio"
)

func GetData(reader *bufio.Reader) (string, string, bool) {
	var name_child, name_parent string

	_, ok := fmt.Fscan(reader, &name_child, &name_parent)
	if ok == nil {
		return name_child, name_parent, true
	}
	return "", "", false
}

func main() {
	var count int

	reader := bufio.NewReaderSize(os.Stdin, 65536)

	fmt.Fscan(reader, &count)

	for i := 0; ; i++ {
		name_child, name_parent, status := GetData(reader)
		if status == false {
			break
		}
		if i < count - 1 {
			fmt.Println("child =", name_child, "parent =", name_parent)
		} else {
			fmt.Println(1)
		}
	}
}
