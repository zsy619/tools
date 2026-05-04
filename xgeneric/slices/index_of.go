package slices

// IndexOf slice[i] == item
func IndexOf[T comparable](slice []T, item T) int {
	ll := len(slice)
	for i := 0; i < ll; i++ {
		if slice[i] == item {
			return i
		}
	}
	return -1
}
