package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
    "strings"
)

type Houses struct {
	Index, MinDistance int
}

type Shops struct {
	Index int
}

func numbers(s string) []int {
    var n []int
    for _, f := range strings.Fields(s) {
        i, err := strconv.Atoi(f)
        if err == nil {
            n = append(n, i)
        }
    }
    return n
}

func main() {
	var count_houses, count_shops int
	var nums []int

	houses := make(map[int]Houses)
	shops := make(map[int]Shops)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums = numbers(scanner.Text())

	len_nums := len(nums)

	for i := 0; i < len_nums; i++ {
		if nums[i] == 1 {
			curr_house := Houses{ i, len_nums }
			if count_shops != 0 {
				last_shop := shops[count_shops]
				curr_house.MinDistance = curr_house.Index - last_shop.Index
				// fmt.Println("add new house:", "i =", i, "min_distance =", curr_house.MinDistance)
			}
			houses[count_houses + 1] = curr_house
			count_houses++
		} else if nums[i] == 2 {
			curr_shop := Shops{ i }
			if count_houses != 0 {
				for j := 1; j <= count_houses; j++ {
					curr_house := houses[j]
					if curr_house.MinDistance > curr_shop.Index - curr_house.Index {
						// fmt.Println("curr_house.MinDistance =", curr_house.MinDistance, "curr_shop.Index - curr_house.Index", curr_shop.Index - curr_house.Index)
						curr_house.MinDistance = (curr_shop.Index - curr_house.Index)
						houses[j] = curr_house
						// fmt.Println("add new shop: index shop =", curr_shop.Index, "Index =", curr_house.Index, "min_distance =", curr_house.MinDistance)
					}
				}
			}
			shops[count_shops + 1] = curr_shop
			count_shops++
		}
	}
	max_distanse := houses[1].MinDistance
	for i := 2; i <= count_houses; i++ {
		if max_distanse < houses[i].MinDistance {
			max_distanse = houses[i].MinDistance
		}
	}
	fmt.Println(max_distanse)
}
