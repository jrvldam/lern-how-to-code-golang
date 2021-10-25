package main

import "fmt"

func main() {
	ns := []int{2, 3, 4, 5, 6, 7, 8, 9}

	x := filter(ns, func(i int) bool { return i%2 != 0 })

	fmt.Println(x)
}

func filter(is []int, fn func(i int) bool) []int {
	var result []int

	for _, v := range is {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}
