package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

type Topic struct {
	Theme	string
	Message	string
}

func GetNumber(s string) (int, bool) {
	num, err := strconv.Atoi(s)

	if err != nil {
		return num, false
	}
	return num, true
}

func main() {
	var count, num, max int
	var theme, message string
	var keys []string

	m := make(map[int]Topic)
	themes := make(map[string]int)

	reader := bufio.NewReaderSize(os.Stdin, 8192)

	s, _ := reader.ReadString('\n')
	s = s[:len(s) - 1]
	count, _ = GetNumber(s)

	for i := 0; i < count; i++ {
		s, _ = reader.ReadString('\n')
		s = s[:len(s) - 1]
		num, _ = GetNumber(s)
		if num == 0 {
			theme, _ = reader.ReadString('\n')
			theme = theme[:len(theme) - 1]

			message, _ = reader.ReadString('\n')
			message = message[:len(message) - 1]
			m[i + 1] = Topic{ theme, message }

			themes[theme] = 0
			keys = append(keys, theme)
		} else {
			message, _ = reader.ReadString('\n')
			message = message[:len(message) - 1]
			theme = m[num].Theme
			m[i + 1] = Topic{ theme, message }
		}
		themes[theme] += 1
	}

	for _, value := range keys {
		if max < themes[value] {
			max = themes[value]
			theme = value
		}
	}
	fmt.Println(theme)
}
