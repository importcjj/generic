package generic

type Option[T any] struct {
	value *T
}

func Some[T any](t T) Option[T] {
	return Option[T]{value: &t}
}

func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

func (option Option[T]) IsSome() bool {
	return !option.IsNone()
}

func (option Option[T]) IsNone() bool {
	return option.value == nil
}

func (option Option[T]) Unwrap() T {
	if option.IsNone() {
		panic("called `Option::unwrap()` on a `None` value")
	}

	return *option.value
}

func (option Option[T]) Expect(msg string) T {
	if option.IsNone() {
		panic(msg)
	}

	return *option.value
}
