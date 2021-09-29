package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func GetData(reader *bufio.Reader) (string, string, bool) {
	var first, second string

	_, ok := fmt.Fscan(reader, &first, &second)
	if ok == nil {
		return first, second, true
	}
	return "", "", false
}

func AddElem(m map[string]string, name_child, name_parent string) {
	if len(m) == 0 {
		m[name_child] = name_parent
		m[name_parent] = ""
	} else {
		m[name_child] = name_parent
	}
}

func FindParent(m map[string]string, parent, child string) (string, bool) {

	curr := m[child]

	if curr == "" {
		return "", false
	} else {
		if curr == parent {
			return parent, true
		} else {
			return FindParent(m, parent, curr)
		}
	}
}

func main() {
	var count int
	var res []int

	reader := bufio.NewReaderSize(os.Stdin, 65536)

	fmt.Fscan(reader, &count)

	m := make(map[string]string)

	for i := 0; ; i++ {
		first, second, status := GetData(reader)
		if status == false {
			break
		}
		if i < count - 1 {
			AddElem(m, first, second)
		} else {
			_, ans := FindParent(m, first, second)
			if ans == true {
				res = append(res, 1)
			} else {
				_, ans := FindParent(m, second, first)
				if ans == true {
					res = append(res, 2)
				} else {
					res = append(res, 0)
				}
			}
		}
	}
	fmt.Println(strings.Trim(fmt.Sprint(res), "[]"))
}
