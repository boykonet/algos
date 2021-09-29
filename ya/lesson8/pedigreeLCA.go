package main

import (
	"os"
	"fmt"
	"bufio"
)

type Parent struct {
	NameParent	string
	Depth		int
}

func GetData(reader *bufio.Reader) (string, string, bool) {
	var first, second string

	_, ok := fmt.Fscan(reader, &first, &second)
	if ok == nil {
		return first, second, true
	}
	return "", "", false
}

func AddElem(m map[string]Parent, name_child, name_parent string) {
	if len(m) == 0 {
		m[name_child] = Parent{ NameParent:name_parent, Depth:1 }
		m[name_parent] = Parent{ NameParent:"", Depth:0 }
	} else {
		m[name_child] = Parent{ name_parent, m[name_parent].Depth + 1 }
	}
}

func FindParentToChild(m map[string]Parent, parent, child string) (string, bool) {

	curr := m[child]

	if curr.NameParent == "" {
		return "", false
	} else {
		if curr.NameParent == parent {
			return parent, true
		} else {
			return FindParentToChild(m, parent, curr.NameParent)
		}
	}
}

func FindParentToTwoChild(m map[string]Parent, first, second string) string {
	firstParent := m[first]
	secondParent := m[second]

	if firstParent.NameParent == "" && secondParent.NameParent == "" {
		return first
	}

	if firstParent.Depth == secondParent.Depth {
		if firstParent.NameParent == secondParent.NameParent {
			return firstParent.NameParent
		} else {
			return FindParentToTwoChild(m, firstParent.NameParent, secondParent.NameParent)
		}
	} else if firstParent.Depth > secondParent.Depth {
		return FindParentToTwoChild(m, firstParent.NameParent, second)
	} else {
		return FindParentToTwoChild(m, first, secondParent.NameParent)
	}
}

func main() {
	var count int
	var res []string

	reader := bufio.NewReaderSize(os.Stdin, 65536)

	fmt.Fscan(reader, &count)

	m := make(map[string]Parent)

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
					res = append(res, FindParentToTwoChild(m, first, second))
				}
			}
		}
	}
	fmt.Println(m)
	for _, v := range res {
		fmt.Println(v)
	}
}
