package iter

import "github.com/importcjj/generic"

func Reduce[T any](v []T, f func(T, T) T) generic.Option[T] {
	if len(v) == 0 {
		return generic.None[T]()
	}

	a := v[0]
	for _, b := range v[1:] {
		a = f(a, b)
	}

	return generic.Some(a)
}
