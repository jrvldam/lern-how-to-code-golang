package main

import "fmt"

func main() {
	array := [5]int{6, 66, 666, 6666, 66666}

	for i, v := range array {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", array)
}
