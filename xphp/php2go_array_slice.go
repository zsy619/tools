package xphp

import "reflect"

// array_unique — Removes duplicate values from an array
func ArrayUnique(arr []string) []string {
	size := len(arr)
	result := make([]string, 0, size)
	temp := map[string]struct{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; !ok {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}

// ArraySlice array_slice()
func ArraySlice(s []interface{}, offset, length int) []interface{} {
	if offset > len(s) {
		panic("offset: the offset is less than the length of s")
	}
	end := offset + length
	if end < len(s) {
		return s[offset:end]
	}
	return s[offset:]
}

func ArraySliceString(s []string, offset, length int) []string {
	if offset > len(s) {
		panic("offset: the offset is less than the length of s")
	}
	end := offset + length
	if end < len(s) {
		return s[offset:end]
	}
	return s[offset:]
}

// array_diff — Computes the difference of arrays
// Compares array1 against one or more other arrays and returns the values in array1 that are not present in any of the other arrays.
func ArrayDiff(array1 []string, arrayOthers ...[]string) []string {
	c := make(map[string]bool)
	for i := 0; i < len(array1); i++ {
		if _, hasKey := c[array1[i]]; hasKey {
			c[array1[i]] = true
		} else {
			c[array1[i]] = false
		}
	}
	for i := 0; i < len(arrayOthers); i++ {
		for j := 0; j < len(arrayOthers[i]); j++ {
			if _, hasKey := c[arrayOthers[i][j]]; hasKey {
				c[arrayOthers[i][j]] = true
			} else {
				c[arrayOthers[i][j]] = false
			}
		}
	}
	result := make([]string, 0)
	for k, v := range c {
		if !v {
			result = append(result, k)
		}
	}
	return result
}

// array_intersect — Computes the intersection of arrays
func ArrayIntersect(array1 []string, arrayOthers ...[]string) []string {
	c := make(map[string]bool)
	for i := 0; i < len(array1); i++ {
		if _, hasKey := c[array1[i]]; hasKey {
			c[array1[i]] = true
		} else {
			c[array1[i]] = false
		}
	}
	for i := 0; i < len(arrayOthers); i++ {
		for j := 0; j < len(arrayOthers[i]); j++ {
			if _, hasKey := c[arrayOthers[i][j]]; hasKey {
				c[arrayOthers[i][j]] = true
			} else {
				c[arrayOthers[i][j]] = false
			}
		}
	}
	result := make([]string, 0)
	for k, v := range c {
		if v {
			result = append(result, k)
		}
	}
	return result
}

// array_search — Searches the array for a given value and returns the first corresponding key if successful
func ArraySearch(needle interface{}, hystack interface{}) (index int) {
	index = -1

	switch reflect.TypeOf(hystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(hystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) {
				index = i
				return
			}
		}
	}
	return
}
