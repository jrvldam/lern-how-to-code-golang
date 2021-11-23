package word

import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder

// UseCount returns a map counting the coincidences of words in the given string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in the given string
func Count(s string) int {
	ws := strings.Split(s, " ")
	return len(ws)
}
