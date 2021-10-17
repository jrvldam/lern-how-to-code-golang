package main

import "fmt"

func main() {
	A := 65
	Z := 90

	for i := A; i <= Z; i += 1 {
		fmt.Printf("%v\n", i)
		for j := 0; j < 3; j += 1 {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
