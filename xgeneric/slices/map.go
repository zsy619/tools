package slices

import (
	"bytes"
	"encoding/gob"
	"encoding/json"

	"haedu.gov.cn/tools/xgeneric"
)

// Map used to apply a function to each element in slice
// return a new slice contains the transformed elements
func Map[S, D any](s []S, f func(S) D) []D {
	r := make([]D, 0, len(s))
	for _, e := range s {
		r = append(r, f(e))
	}
	return r
}

// Keys creates an array of the map keys.
func Keys[K comparable, V any](in map[K]V) []K {
	result := make([]K, 0, len(in))

	for k := range in {
		result = append(result, k)
	}

	return result
}

// Values creates an array of the map values.
func Values[K comparable, V any](in map[K]V) []V {
	result := make([]V, 0, len(in))

	for _, v := range in {
		result = append(result, v)
	}

	return result
}

// PickBy returns same map type filtered by given predicate.
func PickBy[K comparable, V any](in map[K]V, predicate func(K, V) bool) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if predicate(k, v) {
			r[k] = v
		}
	}
	return r
}

// PickByKeys returns same map type filtered by given keys.
func PickByKeys[K comparable, V any](in map[K]V, keys []K) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if xgeneric.Contains(keys, k) {
			r[k] = v
		}
	}
	return r
}

// PickByValues returns same map type filtered by given values.
func PickByValues[K comparable, V comparable](in map[K]V, values []V) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if xgeneric.Contains(values, v) {
			r[k] = v
		}
	}
	return r
}

// PickBy returns same map type filtered by given predicate.
func OmitBy[K comparable, V any](in map[K]V, predicate func(K, V) bool) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if !predicate(k, v) {
			r[k] = v
		}
	}
	return r
}

// OmitByKeys returns same map type filtered by given keys.
func OmitByKeys[K comparable, V any](in map[K]V, keys []K) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if !xgeneric.Contains(keys, k) {
			r[k] = v
		}
	}
	return r
}

// OmitByValues returns same map type filtered by given values.
func OmitByValues[K comparable, V comparable](in map[K]V, values []V) map[K]V {
	r := map[K]V{}
	for k, v := range in {
		if !xgeneric.Contains(values, v) {
			r[k] = v
		}
	}
	return r
}

// Entries transforms a map into array of key/value pairs.
func Entries[K comparable, V any](in map[K]V) []xgeneric.Entry[K, V] {
	entries := make([]xgeneric.Entry[K, V], 0, len(in))

	for k, v := range in {
		entries = append(entries, xgeneric.Entry[K, V]{
			Key:   k,
			Value: v,
		})
	}

	return entries
}

// FromEntries transforms an array of key/value pairs into a map.
func FromEntries[K comparable, V any](entries []xgeneric.Entry[K, V]) map[K]V {
	out := map[K]V{}

	for _, v := range entries {
		out[v.Key] = v.Value
	}

	return out
}

// Invert creates a map composed of the inverted keys and values. If map
// contains duplicate values, subsequent values overwrite property assignments
// of previous values.
func Invert[K comparable, V comparable](in map[K]V) map[V]K {
	out := map[V]K{}

	for k, v := range in {
		out[v] = k
	}

	return out
}

// Assign merges multiple maps from left to right.
func Assign[K comparable, V any](maps ...map[K]V) map[K]V {
	out := map[K]V{}

	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
	}

	return out
}

// MapKeys manipulates a map keys and transforms it to a map of another type.
func MapKeys[K comparable, V any, R comparable](in map[K]V, iteratee func(V, K) R) map[R]V {
	result := map[R]V{}

	for k, v := range in {
		result[iteratee(v, k)] = v
	}

	return result
}

// MapValues manipulates a map values and transforms it to a map of another type.
func MapValues[K comparable, V any, R any](in map[K]V, iteratee func(V, K) R) map[K]R {
	result := map[K]R{}

	for k, v := range in {
		result[k] = iteratee(v, k)
	}

	return result
}

func DeepCopy[K comparable, V any](value map[K]V) map[K]V {
	newMap := make(map[K]V)
	for k, v := range value {
		newMap[k] = v
	}

	return newMap
}

func DeepCopyByGob[K comparable, V any](value map[K]V) (dst map[K]V, err error) {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(value); err != nil {
		return nil, err
	}

	err = gob.NewDecoder(&buffer).Decode(dst)
	return
}

func DeepCopyByJson[K comparable, V any](value map[K]V) (dst map[K]V, err error) {
	b, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, dst)
	return dst, err
}
