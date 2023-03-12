package main

import "fmt"

func main() {
	i := 21
	fmt.Printf("%d \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	j := true
	fmt.Printf("%t \n", j)

	fmt.Printf("\n")

	fmt.Printf("%b \n", i)
	s := '\u042F'
	fmt.Printf("%c\n", s)
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	l := 15
	fmt.Printf("%x \n", l)
	fmt.Printf("%X \n", l)
	fmt.Printf("%U \n", 'Я')

	fmt.Printf("\n")

	var k float64 = 123.456
	fmt.Printf("%f \n", k)
	fmt.Printf("%E \n", k)
}
