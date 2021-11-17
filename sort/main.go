package main

import "fmt"

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func BubbleSort(array []int) []int {
	lenArray := len(array)

	for i := 0; i + 1 < lenArray; i++ {
		for j := 0; j + 1 < lenArray - i; j++ {
			if array[j + 1] < array[j] {
				Swap(&array[j], &array[j + 1])
			}
		}
	}
	return array
}

func InsertionSort(array []int) []int {
	lenArray := len(array)

	for i := 1; i < lenArray; i++ {
		curr := array[i]
		j := i
		for ; j > 0; j-- {
			if array[j - 1] > curr {
				array[j] = array[j-1]
			} else {
				break
			}
		}
		array[j] = curr
	}
	return array
}

func Partition(array []int, l, r int) []int {
	x := array[r]
	less := l

	for i := l; i < r; i++ {
		if array[i] <= x {
			
		}
	}

	return array
}

func main() {
	//array := make([]int, 0)
	//array = append(array, 6, 5)
	//
	//fmt.Println("before:", array)
	//Swap(&array[0], &array[1])
	//fmt.Println("After:", array)

	array := make([]int, 0)

	array = append(array, 8, 4, 9, 6, 7, 1, 2, 10, 65, 33, 77, 22, 45)
	//array = append(array, 9, 8, 7, 6, 5, 4, 3, 2, 1)


	//fmt.Println(BubbleSort(array))

	fmt.Println(InsertionSort(array))
}