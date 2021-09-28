package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func GetNumber(s string) (int, bool) {
	num, err := strconv.Atoi(s)

	if err != nil {
		return num, false
	}
	return num, true
}

func main() {
	var sum, v, num int
	var a1, a2, lens []int

	reader := bufio.NewReaderSize(os.Stdin, 1000000)

	a3 := make(map[int]int)

	fmt.Fscan(reader, &sum)
	for i := 0; i < 3; i++ {
		fmt.Fscan(reader, &v)
		for j := 0; j < v; j++ {
			fmt.Fscan(reader, &num)
			if i == 0 {
				a1 = append(a1, num)
			} else if i == 1 {
				a2 = append(a2, num)
			} else if i == 2 {
				_, ok := a3[num]
				if ok == false {
					a3[num] = j
				}
			}
		}
		lens = append(lens, v)
	}

	for i := 0; i < lens[0]; i++ {
		if sum - a1[i] >= 2 {
			for j := 0; j < lens[1]; j++ {
				if a1[i] + a2[j] < sum {
					_, ok := a3[sum - a1[i] - a2[j]]
					if ok == true {
						fmt.Println(i, j, a3[sum - a1[i] - a2[j]])
						return
					}
				}
			}
		}
	}
	fmt.Println(-1)
}
