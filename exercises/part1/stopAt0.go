package main

import (
	"fmt"
	"strconv"
)

func main() {

	for {
		fmt.Println("Input a number to continue. 0 to quit.")
		var answer string
		fmt.Scanln(&answer)
		input, err := strconv.Atoi(answer)
		if err == nil {
			if input == 0 {
				fmt.Println("Got 0, bye now.")
				break
			} else {
				fmt.Printf("Got %d, here we go again \n", input)
			}
		} else {
			fmt.Println("That's not very nice. Numbers only here! Try again.")
		}

	}
}
