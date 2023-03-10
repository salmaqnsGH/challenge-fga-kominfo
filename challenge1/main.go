package main

import "fmt"

func main() {
	i := 21
	fmt.Printf("%d \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	j := true
	fmt.Printf("%v \n", j)

	fmt.Printf("\n")

	fmt.Printf("%b \n", i)
	c := 'Я'
	fmt.Printf("%c \n", c)
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%v \n", "f")
	fmt.Printf("%v \n", "F")
	fmt.Printf("%U \n", c)

	fmt.Printf("\n")

	var k float64 = 123.456
	fmt.Printf("%f \n", k)
	fmt.Printf("%E \n", k)
}
