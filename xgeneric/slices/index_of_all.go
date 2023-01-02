package slices

// IndexOfAll slice[i] == item
func IndexOfAll[T comparable](slice []T, item T) []int {
	var indexes []int
	ll := len(slice)
	for i := 0; i < ll; i++ {
		if slice[i] == item {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
