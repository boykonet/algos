package main

import (
	"fmt"
	"os"
	"bufio"
	"errors"
)

type stack struct {
	s	[]rune
}

func NewStack() *stack {
	return &stack{ make([]rune, 0), }
}

func (s *stack) Push(v rune) {
	s.s = append(s.s, v)
}

func (s *stack) Pop() (rune, error) {
	l := len(s.s)

	if l == 0 {
		return 0, errors.New("stack is empty")
	}

	num := s.s[l - 1]
	s.s = s.s[:l - 1]
	return num, nil
}

func (s *stack) Top() (rune, error) {
	l := len(s.s)

	if l == 0 {
		return 0, errors.New("stack is empty")
	}

	return s.s[l - 1], nil
}

func (s *stack) Empty() bool {
	if len(s.s) == 0 {
		return true
	}
	return false
}

func main() {
	var value string

	reader := bufio.NewReaderSize(os.Stdin, 11000)

	s := NewStack()

	fmt.Fscan(reader, &value)

	for _, char := range value {
		if char == '(' {
			s.Push(char)
		} else if char == ')' {
			if s.Empty() == true {
				fmt.Println("NO")
				return
			}
			curr, _ := s.Top()
			if curr == '(' {
				s.Pop()
			}
		}
	}
	if s.Empty() == true {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
