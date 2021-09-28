package main

import (
	"os"
	"fmt"
)

type AutoNums struct {
	num		string
	count	int
}

func parseNumsWitnesses(countWitnesses int) map[int]map[rune]struct{} {
	var v string

	w := make(map[int]map[rune]struct{})

	for i := 0 ; i < countWitnesses; i++ {
		fmt.Fscan(os.Stdin, &v)
		w[i] = make(map[rune]struct{})
		for _, char := range v {
			w[i][char] = struct{}{}
		}
	}
	return w
}

func parseAutoNums(w map[int]map[rune]struct{}, countAutoNums, countWitnesses int) (int, []AutoNums) {
	var v string
	var result []AutoNums
	var max int

	for i := 0; i < countAutoNums; i++ {
		add := 0
		a := make(map[rune]struct{})
		result = append(result, AutoNums{ "", 0 })
		fmt.Fscan(os.Stdin, &v)
		for _, char := range v {
			a[char] = struct{}{}
		}
		for j := 0; j < countWitnesses; j++ {
			count := 0
			for char, _ := range w[j] {
				_, found := a[char]
				if found == true {
					count++
				}
			}
			if count == len(w[j]) {
				add += 1
			}
			if max < add {
				max = add
			}
			result[i] = AutoNums{ v, add }
		}
	}
	return max, result
}

func main() {
	var countWitnesses, countAutoNums int
	var result []AutoNums

	fmt.Fscan(os.Stdin, &countWitnesses)

	w := parseNumsWitnesses(countWitnesses)

	fmt.Fscan(os.Stdin, &countAutoNums)

	max, result := parseAutoNums(w, countAutoNums, countWitnesses)

	for _, val := range result {
		if (val.count == max) {
			fmt.Println(val.num)
		}
	}
}
