package iter

func Partition[T any](v []T, f func(T) bool) (left, right []T) {
	for _, a := range v {
		if f(a) {
			left = append(left, a)
		} else {
			right = append(right, a)
		}
	}

	return
}

// The relative order of partitioned items is not matained.
func PartitionInPlace[T any](v []T, f func(T) bool) ([]T, []T) {
	i := 0
	j := len(v)

	trueCount := 0

	for i < j {
		if f(v[i]) {
			trueCount++
			i++
			continue
		}

		for j-1 > i {
			j--
			if f(v[j]) {
				trueCount++
				v[i], v[j] = v[j], v[i]
			}
		}
		i++
	}

	return v[:trueCount], v[trueCount:]
}
