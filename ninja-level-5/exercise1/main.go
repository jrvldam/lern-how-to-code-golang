package main

import "fmt"

type person struct {
	first_name                 string
	last_name                  string
	favorite_ice_cream_flavors []string
}

func main() {
	p1 := person{
		first_name:                 "James",
		last_name:                  "Bond",
		favorite_ice_cream_flavors: []string{"lemon", "cookies", "pippermint"},
	}
	p2 := person{
		first_name:                 "Miss",
		last_name:                  "Monypenny",
		favorite_ice_cream_flavors: []string{"chocolate", "strawberry", "vanila"},
	}

	ps := []person{p1, p2}

	for _, p := range ps {
		fmt.Println(p.first_name)
		fmt.Println(p.last_name)
		for _, f := range p.favorite_ice_cream_flavors {
			fmt.Println("\t", f)
		}
	}
}
