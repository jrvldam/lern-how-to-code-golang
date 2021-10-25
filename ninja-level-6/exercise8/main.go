package main

import "fmt"

func main() {
	count := makeCounter()
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
}

func makeCounter() func() int {
	counter := 0
	return func() int {
		counter += 1
		return counter
	}
}
