package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Individual struct {
	Name string
	Age  int8
}

func TestFromSlice(t *testing.T) {
	var inputs = []*Individual{
		{Name: "Alice", Age: 10},
		{Name: "Bob", Age: 20},
		{Name: "Cindy", Age: 30},
	}

	var expected = make(map[string]int8)
	for _, individual := range inputs {
		expected[individual.Name] = individual.Age
	}

	m := FromSlice(inputs, func(i *Individual) (string, int8) { return i.Name, i.Age })
	assert.EqualValues(t, m, expected)
}

func TestFromSlice2(t *testing.T) {
	var inputs = []*Individual{
		{Name: "Alice", Age: 10},
		{Name: "Bob", Age: 20},
		{Name: "Cindy", Age: 30},
	}

	var expected = make(map[string]struct{})
	for _, individual := range inputs {
		expected[individual.Name] = struct{}{}
	}

	m := FromSlice(inputs, func(i *Individual) (string, struct{}) { return i.Name, struct{}{} })
	assert.EqualValues(t, m, expected)
}
