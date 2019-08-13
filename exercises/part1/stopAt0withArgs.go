package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	min := math.MaxInt32
	max := math.MinInt32

	for i := 0; i < len(arguments); i++ {
		num, err := strconv.Atoi(arguments[i])
		if err == nil {
			if num == 0 {
				// fmt.Println("Found a 0. Stopping.")
				break
			} else {
				if num > max {
					max = num
				}
				if num < min {
					min = num
				}
			}
		}
	}
	if len(arguments) <= 1 {
		min, max = 0, 0
	}
	fmt.Printf("Min: %d, Max: %d", min, max)
}
