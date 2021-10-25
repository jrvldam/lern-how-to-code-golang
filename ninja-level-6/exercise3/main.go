package main

import "fmt"

func main() {
	defer foo()
	baz()
}

func foo() {
	fmt.Println("Hello from foo")
}

func baz() {
	fmt.Println("Hello from baz")
}
