package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`x`] = []string{`Physics`, `Books`, `Baseball`}

	for k, v := range m {
		fmt.Println(k, ":")
		for i, xs := range v {
			fmt.Println("\t", i, xs)
		}
	}
}
