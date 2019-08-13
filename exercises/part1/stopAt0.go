package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	min := math.MaxInt32
	max := math.MinInt32
	for {
		fmt.Println("Input a number to continue. 0 to quit.")
		var answer string
		fmt.Scanln(&answer)
		input, err := strconv.Atoi(answer)
		if err == nil {
			if input == 0 {
				fmt.Println("Got 0, bye now.")
				fmt.Printf("Min: %d, Max: %d", min, max)
				break
			} else {
				if input >= max {
					max = input
					fmt.Println("new max")
					fmt.Println(max)
				}

				if input <= min {
					min = input
					fmt.Println("new min")
				}

				fmt.Printf("Got %d, here we go again \n", input)
			}
		} else {
			fmt.Println("That's not very nice. Numbers only here! Try again.")
		}

	}
}
