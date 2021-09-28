package main

import (
	"fmt"
	"os"
)

func main() {
	var stations, first_station, last_station int64

	fmt.Fscan(os.Stdin, &stations, &first_station, &last_station)

	if last_station > first_station {
		num1 := last_station - first_station - 1
		num2 := (stations - last_station) + (first_station - 1)
		if (num1 > num2) {
			fmt.Println(num2)
		} else {
			fmt.Println(num1)
		}
	} else {
		num1 := first_station - last_station - 1
		num2 := (stations - first_station) + (last_station - 1)
		if (num1 > num2) {
			fmt.Println(num2)
		} else {
			fmt.Println(num1)
		}
	}
}
