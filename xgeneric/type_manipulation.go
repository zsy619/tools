package xgeneric

// FromPtr returns the pointer value or empty.
func FromPtr[T any](x *T) T {
	if x == nil {
		return Empty[T]()
	}

	return *x
}

// FromPtrOr returns the pointer value or the fallback value.
func FromPtrOr[T any](x *T, fallback T) T {
	if x == nil {
		return fallback
	}

	return *x
}

// ToAnySlice returns a slice with all elements mapped to `any` type
func ToAnySlice[T any](collection []T) []any {
	result := make([]any, len(collection))
	for i, item := range collection {
		result[i] = item
	}
	return result
}

// FromAnySlice returns an `any` slice with all elements mapped to a type.
// Returns false in case of type conversion failure.
func FromAnySlice[T any](in []any) (out []T, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			out = []T{}
			ok = false
		}
	}()

	result := make([]T, len(in))
	for i, item := range in {
		result[i] = item.(T)
	}
	return result, true
}

// IsEmpty returns true if argument is a zero value.
func IsEmpty[T comparable](v T) bool {
	var zero T
	return zero == v
}

// IsNotEmpty returns true if argument is not a zero value.
func IsNotEmpty[T comparable](v T) bool {
	var zero T
	return zero != v
}
