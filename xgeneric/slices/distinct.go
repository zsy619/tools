package slices

// Distinct remove the same items
func Distinct[T any](slice []T, condition func(item1, item2 T) bool) []T {
	var newSlice []T
	ll := len(slice)
	for i := 0; i < ll; i++ {
		for j := 0; j <= i; j++ {
			if j == i {
				newSlice = append(newSlice, slice[i])
				break
			}
			if condition(slice[j], slice[i]) {
				break
			}
		}
	}
	return newSlice
}
