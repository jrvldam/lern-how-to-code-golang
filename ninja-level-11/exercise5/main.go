package main

import "fmt"

func main() {
	a := 2
	b := 40

	r := Sum(a, b)

	fmt.Println(r)
}

func Sum(a, b int) int {
	return a + b
}
