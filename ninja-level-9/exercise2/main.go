package main

import "fmt"

func main() {
	p := person{
		first: "James",
		last:  "Bond",
		says:  "Mixed not stirred",
	}

	saySomething(&p)
	// this doesn't compile
	// saySomething(p)
}

type person struct {
	first string
	last  string
	says  string
}

func (p *person) speak() {
	fmt.Println(p.says)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
