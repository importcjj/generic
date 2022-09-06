package iter

import "github.com/importcjj/generic"

func Find[T any](v []T, f func(T) bool) generic.Option[T] {
	for _, a := range v {
		if f(a) {
			return generic.Some(a)
		}
	}

	return generic.None[T]()
}
