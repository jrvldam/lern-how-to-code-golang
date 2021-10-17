package main

import "fmt"

func main() {
	for i := 0; i < 10; i += 1 {
		if md := i % 2; md == 0 {
			fmt.Println(i)
		}
	}
}
