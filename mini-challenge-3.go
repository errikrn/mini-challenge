package main

import (
	"fmt"
	"sync"
)

func interface1(data interface{}) {
	fmt.Println("[coba1 coba2 coba3]", data)
	wg.Done()
}

func interface2(data interface{}) {
	fmt.Println("[bisa1 bisa2 bisa3]", data)
	wg.Done()
}

var wg sync.WaitGroup

func main() {

	wg.Add(8)
	for i := 1; i < 5; i++ {
		go interface1(i)
	}

	for i := 1; i < 5; i++ {
		go interface2(i)
	}
	wg.Wait()
}
