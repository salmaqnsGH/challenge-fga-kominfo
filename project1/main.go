package main

import (
	"fmt"
	"time"
)

type data interface {
	printData(i int)
}

type Bisa struct {
	text []string
}

func (b Bisa) printData(i int) {
	fmt.Println(b.text, i)
}

type Coba struct {
	text []string
}

func (c Coba) printData(i int) {
	fmt.Println(c.text, i)

}

func main() {
	var bisa data = Bisa{text: []string{"bisa1", "bisa2", "bisa3"}}
	var coba data = Coba{text: []string{"coba1", "coba2", "coba3"}}

	for i := 0; i < 4; i++ {
		go bisa.printData(i)
		go coba.printData(i)
	}

	time.Sleep(1 * time.Second)
}
