package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func() {
		fmt.Println("foo")
		wg.Done()
	}()

	go func() {
		fmt.Println("bar")
		wg.Done()
	}()

	wg.Wait()
}
