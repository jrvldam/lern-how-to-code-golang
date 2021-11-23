package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	got := Years(10)

	if got != 70 {
		t.Error("Want 70. Got", got)
	}
}

func TestYearsTwo(t *testing.T) {
	got := YearsTwo(10)

	if got != 70 {
		t.Error("Want 70. Got", got)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output: 70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output: 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		YearsTwo(10)
	}
}
