package hashmap

func Values[K comparable, V any](m map[K]V) []V {
	var a = make([]V, 0, len(m))
	for _, v := range m {
		a = append(a, v)
	}

	return a
}

func MapValues[K comparable, V any, Q any](m map[K]V, f func(V) Q) []Q {
	var a = make([]Q, 0, len(m))
	for _, v := range m {
		a = append(a, f(v))
	}

	return a
}
