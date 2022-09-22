package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	var a = []int{1, 2, 3}
	var isEven = func(n int) bool {
		return n%2 == 0
	}

	even, old := Partition(a, isEven)
	assert.Equal(t, []int{1, 2, 3}, a)
	assert.Equal(t, []int{2}, even)
	assert.Equal(t, []int{1, 3}, old)
}

func TestPartitionInPlace(t *testing.T) {
	var a = []int{1, 2, 3}
	var isEven = func(n int) bool {
		return n%2 == 0
	}

	even, old := PartitionInPlace(a, isEven)
	assert.Equal(t, []int{2, 1, 3}, a)
	assert.Equal(t, []int{2}, even)
	assert.Equal(t, []int{1, 3}, old)
}

func TestPartitionInPlace2(t *testing.T) {
	var a = []int{1, 1, 1}
	var isEven = func(n int) bool {
		return n%2 == 0
	}

	even, old := PartitionInPlace(a, isEven)
	assert.Equal(t, []int{1, 1, 1}, a)
	assert.Equal(t, []int{}, even)
	assert.Equal(t, []int{1, 1, 1}, old)
}

func TestPartitionInPlace3(t *testing.T) {
	var a = []int{2, 2, 2}
	var isEven = func(n int) bool {
		return n%2 == 0
	}

	even, old := PartitionInPlace(a, isEven)
	assert.Equal(t, []int{2, 2, 2}, a)
	assert.Equal(t, []int{2, 2, 2}, even)
	assert.Equal(t, []int{}, old)
}
