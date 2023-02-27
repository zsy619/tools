package xarray

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// 数组旋转
func ArrayRotateInt(matrix [][]int) {
	m := len(matrix)
	tmp := 0

	for i := 0; i < m/2; i++ {
		for j := i; j < m-1-i; j++ {
			tmp = matrix[i][j]
			matrix[i][j] = matrix[m-1-j][i]
			matrix[m-1-j][i] = matrix[m-1-i][m-1-j]
			matrix[m-1-i][m-1-j] = matrix[j][m-1-i]
			matrix[j][m-1-i] = tmp
		}
	}
}

func ArrayRotate[T any](matrix [][]T) {
	m := len(matrix)

	for i := 0; i < m/2; i++ {
		for j := i; j < m-1-i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[m-1-j][i]
			matrix[m-1-j][i] = matrix[m-1-i][m-1-j]
			matrix[m-1-i][m-1-j] = matrix[j][m-1-i]
			matrix[j][m-1-i] = tmp
		}
	}
}

// 任意类型的数组倒转
func ArrayReverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// 随机打乱数组/Slice
func ArrayShuffleAny(slice []interface{}) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

// ArrayShuffle 随机打乱数组/slice
func ArrayShuffle[T any](slice []T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := len(slice) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// ArrayMerge 数组合并
func ArrayMerge[T any](ary ...[]T) ([]T, error) {
	if ary == nil || len(ary) <= 0 {
		return nil, errors.New("参数错误")
	}
	var b []T
	for _, item := range ary {
		b = append(b, item...)
	}
	return b, nil
}

// ArrayUnique 数组去重
func ArrayUnique[T comparable](slice []T) []T {
	m := make(map[T]bool)
	var result []T
	for _, v := range slice {
		if !m[v] {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}

// ArrayChar 创建一个26元素的数组，分别放置'A'-'Z'
func ArrayChar() []string {
	myChars := [26]string{}
	for i := 0; i < 26; i++ {
		myChars[i] = fmt.Sprintf("%c", 'A'+byte(i))
	}
	return myChars[:]
}
