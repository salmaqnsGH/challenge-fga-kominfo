package main

import "fmt"

func main() {
	arr := []rune{'С', 'А', 'Ш', 'А', 'Р', 'В', 'О'}

	for j := 0; j <= 10; j++ {
		if j == 0 {
			for i := 0; i <= 4; i++ {
				fmt.Println("Nilai i =", i)
			}
		} else if j == 5 {
			for i, v := range arr {
				fmt.Printf("character %U '%c' starts at byte position %d \n", v, v, i*2)
			}
			continue
		}
		fmt.Println("Nilai j =", j)
	}
}
