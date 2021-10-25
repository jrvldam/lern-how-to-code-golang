package main

import "fmt"

func main() {
	is := []int{2, 3, 4, 5, 6}
	n := foo(is...)
	fmt.Println(n)

	n2 := bar(is)
	fmt.Println(n2)
}

func foo(ns ...int) int {
	total := 0

	for _, v := range ns {
		total += v
	}

	return total
}

func bar(ns []int) int {
	total := 0

	for _, v := range ns {
		total += v
	}

	return total
}
