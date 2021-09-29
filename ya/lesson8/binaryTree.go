package main


import (
	"fmt"
	"os"
	"bufio"
)

type List struct {
	LeftChild	*List
	Data		int
	RightChild	*List
}

type tree *List

func AddElem(head *List, elem int) *List {
	if head == nil {
		head = &List{ LeftChild: nil, Data: elem, RightChild: nil }
		fmt.Println("DONE")
		return head
	}

	curr := head.Data

	if curr == elem {
		fmt.Println("ALREADY")
		return head
	} else if curr < elem {
		left := head.LeftChild
		if left == nil {
			head.LeftChild = &List{ LeftChild: nil, Data: elem, RightChild: nil }
			fmt.Println("DONE")
			return head.LeftChild
		} else {
			AddElem(left, elem)
		}
	} else if curr > elem {
		right := head.RightChild
		if right == nil {
			head.RightChild = &List{ LeftChild: nil, Data: elem, RightChild: nil }
			fmt.Println("DONE")
			return head.RightChild
		} else {
			AddElem(right, elem)
		}
	}
	return head
}

func FindElem(head *List, elem int) *List {
	if head == nil {
		fmt.Println("NO")
		return nil
	}

	curr := head.Data

	if curr == elem {
		fmt.Println("YES")
		return head
	} else if curr < elem {
		left := head.LeftChild
		if left == nil {
			fmt.Println("NO")
			return nil
		} else {
			FindElem(left, elem)
		}
	} else if curr > elem {
		right := head.RightChild
		if right == nil {
			fmt.Println("NO")
			return nil
		} else {
			FindElem(right, elem)
		}
	}
	return nil
}

func ParseData(reader *bufio.Reader) (string, int, bool) {
	var command string
	var num int

	_, ok := fmt.Fscan(reader, &command)
	if ok != nil {
		return "", 0, false
	}
	if command == "PRINTTREE" {
		return command, 0, true
	}
	fmt.Fscan(reader, &num)
	return command, num, true
}

func PrintTree(head *List, depth string) {
	if head == nil {
		return
	}

	if head.RightChild != nil {
		PrintTree(head.RightChild, depth + ".")
	}

	fmt.Printf("%v%v\n", depth, head.Data)

	if head.LeftChild != nil {
		PrintTree(head.LeftChild, depth + ".")
	}
}

func main() {
	var flag int
	var head *List

	reader := bufio.NewReaderSize(os.Stdin, 100000)

	for ; ; {
		command, num, ok := ParseData(reader)
		if ok == false {
			break
		}
		if command == "SEARCH" {
			FindElem(head, num)
		} else if command == "ADD" {
			if flag == 0 {
				head = AddElem(head, num)
				flag = 1
			} else {
				AddElem(head, num)
			}
		} else if command == "PRINTTREE" {
			PrintTree(head, "")
		}
	}
}
