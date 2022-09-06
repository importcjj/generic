package iter

import (
	"reflect"
	"testing"
)

type MapElement[T any] struct {
	Field T
	Set   []T
}

func TestMap(t *testing.T) {
	caseA := []MapElement[int]{
		{Field: 1, Set: []int{1, 2, 3}},
		{Field: 2, Set: []int{4, 5, 6}},
	}

	ansA := Map(caseA, func(i MapElement[int]) int { return i.Field * 2 })
	ansB := Map(caseA, func(i MapElement[int]) int { return i.Field - 10 })

	if !reflect.DeepEqual(ansA, []int{2, 4}) {
		t.Errorf("%v expected, got %v", []int{2, 4}, ansA)
	}

	if !reflect.DeepEqual(ansB, []int{-9, -8}) {
		t.Errorf("%v expected, got %v", []int{-9, -8}, ansA)
	}
}

func TestFlatMap(t *testing.T) {
	caseA := []MapElement[int]{
		{Field: 1, Set: []int{1, 2, 3}},
		{Field: 2, Set: []int{4, 5, 6}},
	}

	ansA := FlatMap(caseA, func(i MapElement[int]) []int { return i.Set })

	if !reflect.DeepEqual(ansA, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("%v expected, got %v", []int{1, 2, 3, 4, 5, 6}, ansA)
	}

}
