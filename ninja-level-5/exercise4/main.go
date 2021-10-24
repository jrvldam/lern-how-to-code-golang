package main

import "fmt"

func main() {
	a := struct {
		legs  int
		color string
		name  string
	}{
		legs:  4,
		color: "black",
		name:  "Minino",
	}

	fmt.Println(a)
}
