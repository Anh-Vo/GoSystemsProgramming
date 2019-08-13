package main

import "fmt"

func main() {
	myArray := [4]int{1, 2, 4, -4}
	twoD := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	threeD := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	for i, number := range myArray {
		fmt.Printf("index [%d] contains: %d \n", i, number)
	}

	for _, arr := range twoD {
		for _, num := range arr {
			fmt.Printf("%d ", num)
		}
		fmt.Printf("\n")
	}

	for _, arr := range threeD {
		for _, arr2 := range arr {
			for _, num := range arr2 {
				fmt.Printf("%d ", num)
			}
			fmt.Printf("\n")
		}
	}
}
