package main

import "fmt"

func main() {
	const born = 1972
	const current = 2021
	year := born

	for year <= current {
		fmt.Println(year)
		year += 1
	}
}
