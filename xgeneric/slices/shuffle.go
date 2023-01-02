package slices

import "math/rand"

// Shuffle slice
func Shuffle[T any](slice []T) {
	ll := len(slice)
	for i := 0; i < ll; i++ {
		rand.Shuffle(ll, func(i, j int) {
			slice[i], slice[j] = slice[j], slice[i]
		})
	}
}
