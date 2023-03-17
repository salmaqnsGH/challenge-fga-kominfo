package main

import "fmt"

func main() {
	data1 := [][][]int{
		{
			{1, 2, 3}, // 6
			{1, 2, 3}, // 6
		},
		{
			{1, 2, 3}, // 6
			{1, 2},    // 3
			{1},       // 1
		},
	} // 22

	deepSum(data1)
	fmt.Println(deepSum(data1))

	data2 := [][][]int{
		{
			{1, 1, 1}, // 3
			{1},       // 1
		},
		{
			{1, 1, 2}, // 4
			{3, 2, 1}, // 6
		},
		{{3}}, // 3
	} // 17

	deepSum(data2)
	fmt.Println(deepSum(data2))

}

func deepSum(input [][][]int) int {
	// logic to sum all the int from multidimensional array
	x := 0
	for i := range input {
		for j := range input[i] {
			for k := range input[i][j] {
				x += input[i][j][k]
			}
		}
	}
	return x
}
