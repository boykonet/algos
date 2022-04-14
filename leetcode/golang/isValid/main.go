package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
}

func isValid(s string) bool {
	stack := make([]string, 0)

	for _, v := range s {
		value := string(v)

		if len(stack) > 0 {
			if (value == ")" && stack[len(stack) - 1] == "(") ||
			(value == "]" && stack[len(stack) - 1] == "[") ||
			(value == "}" && stack[len(stack) - 1] == "{") {
				stack = stack[:len(stack) - 1]
				continue
			}
		}
		stack = append(stack, value)
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
