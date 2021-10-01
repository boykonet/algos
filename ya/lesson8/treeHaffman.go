package main

import (
	"fmt"
	"os"
	"bufio"
	// "strings"
)

type Node struct {
	Left	*Node
	Right	*Node
	Up		*Node
	Type	string
}

func BuildTree(s string) *Node {
	head := &Node{ Left: nil, Right: nil, Up: nil, Type: "head" }

	currnode := head

	for _, r := range s {
		if r == 'D' {
			currnode.Left = &Node{ Left: nil, Right: nil, Up: currnode, Type: "left" }
			currnode = currnode.Left
		} else if r == 'U' {
			for currnode.Type == "right" {
				currnode = currnode.Up
			}
			currnode = currnode.Up
			currnode.Right = &Node{ Left: nil, Right: nil, Up: currnode, Type: "right" }
			currnode = currnode.Right
		}
	}
	return head
}

func Decoding(head *Node, prefix string, ans []string) []string {

	if head.Left == nil && head.Right == nil {
		ans = append(ans, prefix)
		return ans
	}

	prefix += "0"

	Decoding(head.Left, prefix, ans)

	prefix = prefix[:len(prefix) - 1]
	prefix += "1"

	Decoding(head.Right, prefix, ans)
	prefix = prefix[:len(prefix) - 1]
	return ans
}

func main() {
	var num int
	var str string
	var codes []string

	reader := bufio.NewReaderSize(os.Stdin, 150000)

	fmt.Fscan(reader, &num)

	for i := 0; i < num; i++ {
		fmt.Fscan(reader, &str)
		codes = append(codes, str)
	}
	fmt.Println(codes)
	for _, s := range codes {
		var ans []string
		var prefix string

		head := BuildTree(s)
		ans = Decoding(head, prefix, ans)

		fmt.Println("len =", len(ans))
		for _, a := range ans {
			fmt.Println(a)
		}
	}
}
