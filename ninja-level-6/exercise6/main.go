package main

import "fmt"

func main() {
	foo()()
}

func foo() func() {
	return func() {
		fmt.Println("Hello from anonymus func")
	}
}
