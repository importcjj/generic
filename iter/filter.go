package iter

import "github.com/importcjj/generic"

func Filter[T any](v []T, f func(T) bool) []T {
	var r []T

	for _, item := range v {
		if f(item) {
			r = append(r, item)
		}
	}

	return r
}

func FilterMap[T, U any](v []T, f func(T) generic.Option[U]) []U {
	var r []U

	for _, item := range v {
		if b := f(item); b.IsSome() {
			r = append(r, b.UnwrapUnchecked())
		}
	}

	return r
}
