package main

import "fmt"

const (
	x = 2021
	a = x + iota
	b = x + iota
	c = x + iota
	d = x + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
