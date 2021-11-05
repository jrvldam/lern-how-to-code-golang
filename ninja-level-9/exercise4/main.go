package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mtx sync.Mutex

func main() {
	var counter int

	for i := 0; i < 100; i += 1 {
		wg.Add(1)
		go func() {
			mtx.Lock()
			a := counter
			a += 1
			counter = a
			mtx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
