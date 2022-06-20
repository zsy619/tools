package xjson

import (
	"fmt"
)

// Convert json string to map
func JsonToMap[K comparable, V any](jsonStr string) (map[K]V, error) {
	m := make(map[K]V)
	err := Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}

	// for k, v := range m {
	// 	fmt.Printf("%v: %v\n", k, v)
	// }

	return m, nil
}

// Convert map json string
func MapToJson[K comparable, V any](m map[K]V) (string, error) {
	jsonByte, err := Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", err
	}

	return string(jsonByte), nil
}
