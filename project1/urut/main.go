package main

import (
	"fmt"
	"sync"
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

	var mutex sync.Mutex
	for i := 1; i <= 4; i++ {
		mutex.Lock()
		go func(i int) {
			defer mutex.Unlock()
			bisa.printData(i)
		}(i)
		mutex.Lock()
		go func(i int) {
			defer mutex.Unlock()
			coba.printData(i)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
