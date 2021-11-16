package main

import "testing"

func TestSum(t *testing.T) {
	got := Sum(1, 2)

	if got != 3 {
		t.Errorf("expected 3, recieved %v", got)
	}
}
