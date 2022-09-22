package hashmap

func FromSlice[T any, K comparable, V any](v []T, f func(T) (K, V)) map[K]V {
	var m = make(map[K]V)
	for _, a := range v {
		k, v := f(a)
		m[k] = v
	}

	return m
}
