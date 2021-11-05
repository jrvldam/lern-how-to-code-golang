package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var counter int

	for i := 0; i < 100; i += 1 {
		wg.Add(1)
		go func() {
			a := counter
			runtime.Gosched()
			a += 1
			counter = a
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
