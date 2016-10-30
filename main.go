package main

import (
	"fmt"
	"sync"
)

var initGlobal sync.Once

func ExampleOnce() {
	cb := func(wg *sync.WaitGroup) {
		defer wg.Done()

		initGlobal.Do(func() {
			fmt.Println("do once!")
		})
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go cb(&wg)
	}

	wg.Wait()
}

func main() {
	ExampleOnce()
}
