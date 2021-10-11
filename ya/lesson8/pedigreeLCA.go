package main

import (
	"os"
	"fmt"
	"bufio"
)

func FindDepth(m map[string]string, name string, currDepth int) int {
	if name == "" {
		return currDepth + 1
	} else {
		return FindDepth(m, m[name], currDepth + 1)
	}
}

func GetData(reader *bufio.Reader) (string, string, bool) {
	var first, second string

	_, ok := fmt.Fscan(reader, &first, &second)
	if ok == nil {
		return first, second, true
	}
	return "", "", false
}

func AddElem(m map[string]string, name_child, name_parent string) {
	_, ok := m[name_parent]
	if ok == false {
		m[name_parent] = ""
	}

	m[name_child] = name_parent
}

func FindParentToChild(m map[string]string, parent, child string) (string, bool) {

	curr := m[child]

	if curr == "" {
		return "", false
	} else {
		if curr == parent {
			return parent, true
		} else {
			return FindParentToChild(m, parent, curr)
		}
	}
}

func FindParentToTwoChild(m map[string]string, first, second string, firstDepth, secondDepth int) string {
	if first == second {
		return first
	}

	firstParent := m[first]
	secondParent := m[second]

	if firstDepth == secondDepth {
		return FindParentToTwoChild(m, firstParent, secondParent, firstDepth - 1, secondDepth - 1)
	} else if firstDepth > secondDepth {
		return FindParentToTwoChild(m, firstParent, second, firstDepth - 1, secondDepth)
	} else {
		return FindParentToTwoChild(m, first, secondParent, firstDepth, secondDepth - 1)
	}
}

func main() {
	var count int
	var res []string

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

			p, ans := FindParentToChild(m, first, second)

			if ans == true {
				res = append(res, p)
			} else {
				p, ans := FindParentToChild(m, second, first)
				if ans == true {
					res = append(res, p)
				} else {
					res = append(res, FindParentToTwoChild(m, first, second, FindDepth(m, first, 0), FindDepth(m, second, 0)))
				}
			}
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}
}
