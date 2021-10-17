package main

import "fmt"

func main() {
	favSport := "baseball"

	switch favSport {
	case "basketball":
		fmt.Println("basketball")
	case "football":
		fmt.Println("football")
	case "baseball":
		fmt.Println("balseball")
	default:
		fmt.Println("vim-go")
	}
}
