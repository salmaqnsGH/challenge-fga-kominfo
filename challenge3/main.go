package main

import (
	"fmt"
)

func main() {
	input := "selamat malam"

	myMap := map[string]int{}

	for _, value := range input {
		fmt.Println(string(value))

		myMap[string(value)] = 0
	}

	for _, value := range input {
		if isExist(myMap, string(value)) {
			myMap[string(value)] += 1
		}
	}

	fmt.Println(myMap)
}

func isExist(myMap map[string]int, str string) bool {
	for key := range myMap {
		if key == str {
			return true
		}
	}
	return false
}
