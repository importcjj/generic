package slice

func retain[T any](v []T, f func(T) bool) []T {
	return retain_mut(v, func(elm T, _ int) bool { return f(elm) })
}

func retain_mut[T any](v []T, f func(T, int) bool) []T {
	deleted := 0
	processed := 0
	originLen := len(v)

	for processed < originLen {
		elm := v[processed]
		if !f(elm, processed) {
			deleted += 1
			processed += 1
			break
		}
		processed += 1
	}

	for processed < originLen {
		elm := v[processed]
		if !f(elm, processed) {
			deleted += 1
			processed += 1
			continue
		}

		copy(v[processed-deleted:], v[processed:processed+1])

		processed += 1
	}

	return v[:processed-deleted]
}
