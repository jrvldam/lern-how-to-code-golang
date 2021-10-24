package main

import "fmt"

func main() {
	array := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	chunk1 := array[:5]
	fmt.Println(chunk1)

	chunk2 := array[5:]
	fmt.Println(chunk2)

	chunk3 := array[2:7]
	fmt.Println(chunk3)

	chunk4 := array[1:6]
	fmt.Println(chunk4)
}
