package main

import (
	"fmt"
)

type huminta int

var x huminta
var y int

func main() {
	fmt.Println("zero value:", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	// convertion
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
