package iter

func Map[T, U any](v []T, f func(T) U) []U {
	var ans = make([]U, 0, len(v))
	for _, a := range v {
		ans = append(ans, f(a))
	}

	return ans
}

func FlatMap[T, U any](v []T, f func(T) []U) []U {
	var ans = make([]U, 0, len(v))
	for _, a := range v {
		ans = append(ans, f(a)...)
	}

	return ans
}
