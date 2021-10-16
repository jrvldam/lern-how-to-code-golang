package main

import (
	"fmt"
)

type myOwnType int

var x myOwnType

func main() {
	fmt.Println("zero value:", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
