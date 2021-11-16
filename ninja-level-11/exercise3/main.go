package main

import "fmt"

type customError struct {
	info string
}

func (ce customError) Error() string {
	return fmt.Sprintf("new error %v", ce.info)
}

func main() {
	err := customError{
		info: "my cutom error",
	}

	foo(err)
}

func foo(e error) {
	fmt.Println("recieve error", e)
}
