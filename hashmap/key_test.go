package hashmap

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	scores := map[string]int64{
		"Alice": 100,
		"Bob":   99,
		"Cindy": 60,
	}

	keys := Keys(scores)
	expected := []string{"Alice", "Bob", "Cindy"}
	assert.ElementsMatch(t, keys, expected)
}

func TestMapKeys(t *testing.T) {
	scores := map[string]int64{
		"Alice": 100,
		"Bob":   99,
		"Cindy": 60,
	}

	keys := MapKeys(scores, strings.ToUpper)
	expected := []string{"ALICE", "BOB", "CINDY"}
	assert.ElementsMatch(t, keys, expected)
}
