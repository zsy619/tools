package xarray

import (
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
