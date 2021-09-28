package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"sort"
)

func main() {
	var max_num, max_count, num int
	var yes string
	var result []int

	vars := make(map[int]int)
	nums := make(map[int]int)

	Reader := bufio.NewReaderSize(os.Stdin, 300000)

	fmt.Fscan(Reader, &max_num)

	for i := 1; i <= max_num; i++ {
		vars[i] = 0
	}

	for ; ; {
		_, err := fmt.Fscan(Reader, &num)
		if err != nil {
			_, err = fmt.Fscan(Reader, &yes)
			if yes == "YES" {
				for key, _ := range nums {
					vars[key] += 1
					if max_count < vars[key] {
						max_count = vars[key]
					}
					delete(nums, key)
				}
			} else if yes == "NO" {
				for key, _ := range nums {
					_, ok := vars[key]
					if ok == true {
						delete(vars, key)
					}
					delete(nums, key)
				}
			} else {
				for key, value := range vars {
					if value == max_count {
						result = append(result, key)
					}
				}
				sort.Sort(sort.IntSlice(result))
				fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
				break
			}
		} else {
			nums[num] = 1
		}
	}
}
