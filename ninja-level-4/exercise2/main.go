package main

import "fmt"

func main() {
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", slice)
}
