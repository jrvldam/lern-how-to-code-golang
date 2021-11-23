package mymath

import "testing"

func TestCenteredAvg(t *testing.T) {
	s := [][]int{
		[]int{1, 2, 3, 4},
		[]int{4, 3, 2, 1},
	}

	for _, v := range s {
		got := CenteredAvg(v)

		if got != 2.5 {
			t.Error("Went '2.5'. Got", got)
		}
	}
}
