package main

import (
	"fmt"
	"os"
)

func main() {
	var code, verdict_interactor, verdict_checker int64

	fmt.Fscan(os.Stdin, &code, &verdict_interactor, &verdict_checker)

	if verdict_interactor == 0 {
		if code != 0 {
			fmt.Println(3)
		} else {
			fmt.Println(verdict_checker)
		}
	} else if verdict_interactor == 1 {
		fmt.Println(verdict_checker)
	} else if verdict_interactor == 4 {
		if code != 0 {
			fmt.Println(3)
		} else {
			fmt.Println(4)
		}
	} else if verdict_interactor == 6 {
		fmt.Println(0)
	} else if verdict_interactor == 7 {
		fmt.Println(1)
	} else {
		fmt.Println(verdict_interactor)
	}
}
