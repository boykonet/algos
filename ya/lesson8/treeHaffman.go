package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
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

var ans []string

func Decoding(head *Node, prefix []string) []string {
	if head.Left == nil && head.Right == nil {
		return prefix
	}

	prefix = append(prefix, "0")

	a := Decoding(head.Left, prefix)
	if a != nil {
		ans = append(ans, strings.Join(a, ""))
	}

	prefix = prefix[:len(prefix) - 1]
	prefix = append(prefix, "1")

	b := Decoding(head.Right, prefix)
	if b != nil {
		ans = append(ans, strings.Join(b, ""))
	}
	prefix = prefix[:len(prefix) - 1]
	return nil
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

	for _, s := range codes {
		var prefix []string
		ans = []string{}

		head := BuildTree(s)
		Decoding(head, prefix)

		fmt.Println(len(ans))
		for _, a := range ans {
			fmt.Println(a)
		}
	}
}
