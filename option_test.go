package generic

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionConstructSome(t *testing.T) {

	type Foo struct {
	}

	type Bar interface{}

	Some(1)
	Some("Hello")

	Some(Foo{})
	Some(&Foo{})
	Some([]int{1, 2, 3})
	Some([]string{"1", "2", "3"})
	Some([]*Foo{{}, {}, {}})

	var i Bar = Foo{}
	Some(i)
	Some(&i)
}

func TestOptionIsSome(t *testing.T) {
	type Foo struct {
	}

	type Bar interface{}

	assert.True(t, Some(1).IsSome())
	assert.True(t, Some("Hello").IsSome())

	assert.True(t, Some(Foo{}).IsSome())
	assert.True(t, Some(&Foo{}).IsSome())
	assert.True(t, Some([]int{1, 2, 3}).IsSome())
	assert.True(t, Some([]string{"1", "2", "3"}).IsSome())
	assert.True(t, Some([]*Foo{{}, {}, {}}).IsSome())

	var i Bar = Foo{}
	assert.True(t, Some(i).IsSome())
	assert.True(t, Some(&i).IsSome())
}

func TestOptionIsSomeAnd(t *testing.T) {
	type Foo struct {
		Age int
	}

	type Bar interface{}

	assert.True(t, Some(1).IsSomeAnd(func(i int) bool { return i == 1 }))
	assert.True(t, Some("Hello").IsSomeAnd(func(s string) bool { return strings.ContainsAny(s, "ee") }))

	assert.True(t, Some(Foo{Age: 20}).IsSomeAnd(func(f Foo) bool { return f.Age > 10 }))
	assert.True(t, Some(&Foo{Age: 20}).IsSomeAnd(func(f *Foo) bool { return f.Age%2 == 0 }))
	assert.True(t, Some([]int{1, 2, 3}).IsSomeAnd(func(i []int) bool { return len(i) == 3 }))
	assert.True(t, Some([]string{"1", "2", "3"}).IsSomeAnd(func(i []string) bool { return len(i) == 3 }))
	assert.True(t, Some([]*Foo{{}, {}, {}, {}}).IsSomeAnd(func(f []*Foo) bool { return len(f) == 4 }))

	var i Bar = Foo{Age: 100}
	assert.True(t, Some(i).IsSomeAnd(func(b Bar) bool { return true }))
	assert.True(t, Some(&i).IsSomeAnd(func(b *Bar) bool { return true }))
}

func TestOptionConstructNone(t *testing.T) {
	type Foo struct {
	}

	type Bar interface{}

	None[int]()
	None[string]()
	None[Foo]()
	None[*Foo]()
	None[**Foo]()
	None[Bar]()
	None[*Bar]()
	None[[]int]()
	None[[]*Foo]()
}

func TestOptionIsNone(t *testing.T) {
	type Foo struct {
	}

	type Bar interface{}

	assert.True(t, None[int]().IsNone())
	assert.True(t, None[string]().IsNone())
	assert.True(t, None[Foo]().IsNone())
	assert.True(t, None[*Foo]().IsNone())
	assert.True(t, None[**Foo]().IsNone())
	assert.True(t, None[Bar]().IsNone())
	assert.True(t, None[*Bar]().IsNone())
	assert.True(t, None[[]int]().IsNone())
	assert.True(t, None[[]*Foo]().IsNone())
}

func TestOptionIsSomeAnd2(t *testing.T) {
	type Foo struct {
	}

	type Bar interface{}

	assert.False(t, None[int]().IsSomeAnd(func(i int) bool { return true }))
	assert.False(t, None[string]().IsSomeAnd(func(s string) bool { return true }))
	assert.False(t, None[Foo]().IsSomeAnd(func(f Foo) bool { return true }))
	assert.False(t, None[*Foo]().IsSomeAnd(func(f *Foo) bool { return true }))
	assert.False(t, None[**Foo]().IsSomeAnd(func(f **Foo) bool { return true }))
	assert.False(t, None[Bar]().IsSomeAnd(func(b Bar) bool { return true }))
	assert.False(t, None[*Bar]().IsSomeAnd(func(b *Bar) bool { return true }))
	assert.False(t, None[[]int]().IsSomeAnd(func(i []int) bool { return true }))
	assert.False(t, None[[]*Foo]().IsSomeAnd(func(f []*Foo) bool { return true }))
}
