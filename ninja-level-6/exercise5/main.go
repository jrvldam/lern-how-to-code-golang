package main

import "fmt"

func main() {
	sqr := square{
		side: 4,
	}

	crl := circle{
		radius: 6.5,
	}

	info(sqr)
	info(crl)
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return 3.1416 * c.radius * 2
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}
