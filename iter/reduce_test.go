package iter

import (
	"testing"

	"github.com/importcjj/generic"
	"github.com/stretchr/testify/assert"
)

func TestReduceToFindMax(t *testing.T) {
	a := []int{10, 20, 5, -23, 0}
	b := []int{}

	maxInA := Reduce(a, func(accum, item int) int {
		if accum > item {
			return accum
		}

		return item
	})

	assert.EqualValues(t, maxInA, generic.Some(20))

	maxInB := Reduce(b, func(accum, item int) int {
		if accum > item {
			return accum
		}

		return item
	})

	assert.EqualValues(t, maxInB, generic.None[int]())
}
