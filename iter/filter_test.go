package iter

import (
	"strconv"
	"testing"

	"github.com/importcjj/generic"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	isPositive := func(n int) bool {
		return n > 0
	}

	inputA := []int{0, 1, 2, 3}
	inputB := []int{-1, -2, -3}
	inputC := []int{100, 100, 100}

	assert.Equal(t, Filter(inputA, isPositive), []int{1, 2, 3})
	assert.Equal(t, Filter(inputB, isPositive), ([]int)(nil))
	assert.Equal(t, Filter(inputC, isPositive), []int{100, 100, 100})
}

func TestFilterMap(t *testing.T) {
	parseNumeric := func(s string) generic.Option[int64] {
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return generic.None[int64]()
		}

		return generic.Some(n)
	}

	inputA := []string{"1", "two", "NaN", "four"}
	inputB := []string{"12zx", "two", "NaN", "four"}
	inputC := []string{"100", "100", "100"}

	assert.Equal(t, FilterMap(inputA, parseNumeric), []int64{1})
	assert.Equal(t, FilterMap(inputB, parseNumeric), ([]int64)(nil))
	assert.Equal(t, FilterMap(inputC, parseNumeric), []int64{100, 100, 100})
}
