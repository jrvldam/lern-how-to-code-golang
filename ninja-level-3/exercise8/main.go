package main

import "fmt"

func main() {
	switch {
	case 1 == 2:
		fmt.Println("1 < 2 ")
	case 1 > 2:
		fmt.Println("1 > 2")
	case 1 < 2:
		fmt.Println("1 == 2")
	default:
		fmt.Println("vim-go")
	}
}
