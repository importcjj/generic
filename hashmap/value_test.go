package hashmap

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	scores := map[string]int64{
		"Alice": 100,
		"Bob":   99,
		"Cindy": 60,
	}

	values := Values(scores)
	expected := []int64{100, 99, 60}
	assert.ElementsMatch(t, values, expected)
}

func TestMapValues(t *testing.T) {
	scores := map[string]int64{
		"Alice": 100,
		"Bob":   99,
		"Cindy": 60,
	}

	values := MapValues(scores, func(v int64) string {
		return strconv.FormatInt(v, 10)
	})
	expected := []string{"100", "99", "60"}
	assert.ElementsMatch(t, values, expected)
}

func TestMapValues2(t *testing.T) {
	scores := map[string]int64{
		"Alice": 100,
		"Bob":   99,
		"Cindy": 60,
	}

	values := MapValues(scores, func(v int64) string {
		return strconv.FormatInt(v, 2)
	})
	expected := []string{"1100100", "1100011", "111100"}
	assert.ElementsMatch(t, values, expected)
}
