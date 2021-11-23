package word

import (
	"fmt"
	"testing"
)

var s = "one two three. four four"

func TestCount(t *testing.T) {
	got := Count(s)

	if got != 5 {
		t.Error("Went 4. Got", got)
	}
}

func TestUseCount(t *testing.T) {
	got := UseCount(s)

	went := map[string]int{
		"one":    1,
		"two":    1,
		"three.": 1,
		"four":   2,
	}

	for w := range went {
		if got[w] != went[w] {
			t.Error("Went", went[w], ". Got", got[w])
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		Count(s)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		UseCount(s)
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three. four four"))
	// Output: 5
}
