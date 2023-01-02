package slices

// Reverse slice
func Reverse[T any](slice []T) []T {
	ll := len(slice)
	reversed := make([]T, 0, ll)
	for i := ll - 1; i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}
	return reversed
}
