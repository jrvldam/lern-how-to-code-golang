package main

import "fmt"

func main() {
	for i := 10; i <= 100; i += 1 {
		fmt.Println(i, i%4)
	}
}
