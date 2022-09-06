package slice

import (
	"reflect"
	"testing"
)

func TestRetainRemoveHead(t *testing.T) {
	var v = []int{1, 2, 3, 4}
	v = Retain(v, func(elm int) bool {
		return elm > 3
	})

	if !reflect.DeepEqual(v, []int{4}) {
		t.Errorf("%v expected, got %v", v, []int{4})
	}
}

func TestRetainRemoveTail(t *testing.T) {
	var v = []int{1, 2, 3, 4}
	v = Retain(v, func(elm int) bool {
		return elm < 3
	})

	if !reflect.DeepEqual(v, []int{1, 2}) {
		t.Errorf("%v expected, got %v", v, []int{1, 2})
	}
}

func TestRetainRemoveMiddle(t *testing.T) {
	var v = []int{1, 2, 3, 4}
	v = Retain(v, func(elm int) bool {
		return elm%2 == 0
	})

	if !reflect.DeepEqual(v, []int{2, 4}) {
		t.Errorf("%v expected, got %v", v, []int{2, 4})
	}
}
