package main

import "fmt"

func main() {
	start := 33
	end := 122

	for i := start; i <= end; i += 1 {
		fmt.Printf("%v -- %#x -- %#U\n", i, i, i)
	}
}
