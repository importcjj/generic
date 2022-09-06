package iter

func Any[T any](v []T, f func(T) bool) bool {
	for _, a := range v {
		if f(a) {
			return true
		}
	}
	return false
}
