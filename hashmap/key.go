package hashmap

func Keys[K comparable, V any](m map[K]V) []K {
	var a = make([]K, 0, len(m))
	for k := range m {
		a = append(a, k)
	}

	return a
}

func MapKeys[K comparable, V any, Q any](m map[K]V, f func(K) Q) []Q {
	var a = make([]Q, 0, len(m))
	for k := range m {
		a = append(a, f(k))
	}

	return a
}
