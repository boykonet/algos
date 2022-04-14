package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

var m = map[string]int{
		"I": 1, "IV": 4, "V": 5, "IX": 9,
		"X": 10, "XL": 40, "L": 50, "XC": 90,
		"C": 100, "CD": 400, "D": 500, "CM": 900,
		"M": 1000,
}

func romanToInt(s string) int {
	res := 0
	for len(s) > 0 {
		if len(s) > 1 {
			twoSymbols := s[0:2]
			value, ok := m[twoSymbols]
			if ok == true {
				res += value
				s = s[2:]
				continue
			}
		}
		oneSymbol := s[0:1]
		value, ok := m[oneSymbol]
		if ok == true {
			res += value
			s = s[1:]
		}
	}
	return res
}
