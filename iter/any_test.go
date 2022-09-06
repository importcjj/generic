package iter

import "testing"

func TestAnylanguage(t *testing.T) {
	var (
		input = []string{"Rust", "Go", "Python"}
		f     = func(s string) bool { return s == "Rust" }
		exist = true
	)

	if ans := Any(input, f); ans != exist {
		t.Errorf("%v expected, got %v", exist, ans)
	}
}

func TestAnyOddNumber(t *testing.T) {
	var (
		input = []int{2, 4, 6, 8, 10}
		f     = func(s int) bool { return s%2 == 1 }
		exist = false
	)

	if ans := Any(input, f); ans != exist {
		t.Errorf("%v expected, got %v", exist, ans)
	}
}
