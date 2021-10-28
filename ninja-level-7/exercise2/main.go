package main

import "fmt"

func main() {
	p := person{
		name: "Jane",
	}
	fmt.Println(p)

	changeMe(&p)
	fmt.Println(p)
}

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "Joe"
}
