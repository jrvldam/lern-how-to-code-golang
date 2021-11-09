package main

import "fmt"

func main() {
	c := make(chan int)

	for i := 0; i < 10; i += 1 {
		go func() {
			for j := 0; j < 10; j += 1 {
				c <- j
			}
		}()
	}

	for k := 0; k < 100; k += 1 {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")
}
