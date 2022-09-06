package iter

import (
	"testing"
)

func TestFindlanguage(t *testing.T) {
	var (
		input = []string{"Rust", "Go", "Python"}
		f     = func(s string) bool { return s == "Rust" }
	)

	ans := Find(input, f)

	if ans.IsNone() {
		t.Errorf("Some expected, got None")
	}

	l := ans.Unwrap()

	if l != "Rust" {
		t.Errorf("Rust expected, got %s", l)
	}
}

func TestFindOddNumber(t *testing.T) {
	var (
		input = []int{2, 4, 6, 8, 10}
		f     = func(s int) bool { return s%2 == 1 }
	)

	ans := Find(input, f)

	if ans.IsSome() {
		t.Errorf("None expected, got Some")
	}
}
