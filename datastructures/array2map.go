package main

import (
	"fmt"
	"strconv"
)

func main() {
	anArray := [4]int{1, 2, 3, 4}
	aMap := make(map[string]int)

	length := len(anArray)
	for i := 0; i < length; i++ {
		aMap[strconv.Itoa(i)] = anArray[i]
	}

	for key, value := range aMap {
		fmt.Printf("key %s, has value %d\n", key, value)
	}
}
