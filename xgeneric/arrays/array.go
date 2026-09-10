package arrays

import "fmt"

// Array 是一个固定大小的切片容器。
// 内部使用切片 values 存储元素，构造后其长度不可变。
type Array[T any] struct {
	values []T
}

// New 创建一个长度为 size 的新 Array。
// 参数：size 为数组的固定长度。
// 返回：内部 values 已分配好容量的 *Array[T]。当 size < 0 时会 panic。
func New[T any](size int) *Array[T] {
	return &Array[T]{values: make([]T, size, size)}
}

// NewFromArray 通过复制另一个 Array 的所有元素创建一个新 Array。
// 参数：other 为源 Array。
// 返回：长度与 other 相同且元素已复制的新 *Array[T]。
func NewFromArray[T any](other *Array[T]) *Array[T] {
	a := &Array[T]{values: make([]T, other.Size(), other.Size())}
	for i := range other.values {
		a.values[i] = other.values[i]
	}
	return a
}

// Fill 将数组中的所有元素设置为 val。
// 参数：val 为要填充的值。
// 副作用：会就地修改 a.values 中的全部元素。
func (a *Array[T]) Fill(val T) {
	for i := range a.values {
		a.values[i] = val
	}
}

// Set 将下标为 pos 的元素设置为 val。
// 参数：pos 为目标下标，val 为新值。
// 副作用：当 pos 越界（<0 或 >= len）时静默忽略，不会修改任何元素。
func (a *Array[T]) Set(pos int, val T) {
	if pos < 0 || pos >= len(a.values) {
		return
	}
	a.values[pos] = val
}

// At 返回数组中下标为 pos 的元素值。
// 参数：pos 为目标下标。
// 返回：该位置上的元素值。
// 当 pos 越界（<0 或 >= len）时 panic("index out off range")。
func (a *Array[T]) At(pos int) T {
	if pos < 0 || pos >= len(a.values) {
		panic("index out off range")
	}
	return a.values[pos]
}

// Front 返回数组的第一个元素。
// 返回：values[0]；当数组为空时会因 At(0) 越界而 panic。
func (a *Array[T]) Front() T {
	return a.At(0)
}

// Back 返回数组的最后一个元素。
// 返回：values[len-1]；当数组为空时会因 At(len-1) 越界而 panic。
func (a *Array[T]) Back() T {
	return a.At(len(a.values) - 1)
}

// Size 返回数组中元素的数量。
// 返回：内部切片的长度（即数组的固定长度）。
func (a *Array[T]) Size() int {
	return len(a.values)
}

// Empty 判断数组是否为空。
// 返回：当 len(values) == 0 时返回 true，否则返回 false。
func (a *Array[T]) Empty() bool {
	return len(a.values) == 0
}

// SwapArray 交换两个数组的内部数据。
// 参数：other 为要交换的另一个 Array。
// 副作用：当 a 与 other 长度不同时直接返回，不会发生交换。
func (a *Array[T]) SwapArray(other *Array[T]) {
	if a.Size() != other.Size() {
		return
	}
	a.values, other.values = other.values, a.values
}

// Data 返回数组内部的元素切片。
// 返回：直接暴露内部 values 引用，对其修改会影响数组内容。
func (a *Array[T]) Data() []T {
	return a.values
}

// Begin 返回指向数组起始位置的迭代器。
// 返回：position 为 0 的 *ArrayIterator[T]（相当于 First()）。
func (a *Array[T]) Begin() *ArrayIterator[T] {
	return a.First()
}

// End 返回指向数组末尾之后位置的迭代器（哨兵位置）。
// 返回：position 等于 a.Size() 的 *ArrayIterator[T]。
func (a *Array[T]) End() *ArrayIterator[T] {
	return a.IterAt(a.Size())
}

// First 返回指向数组起始位置（下标 0）的迭代器。
// 返回：position 为 0 的 *ArrayIterator[T]。
func (a *Array[T]) First() *ArrayIterator[T] {
	return a.IterAt(0)
}

// Last 返回指向数组最后一个有效位置的迭代器。
// 返回：position 为 a.Size()-1 的 *ArrayIterator[T]；空数组将得到 position=-1 的迭代器。
func (a *Array[T]) Last() *ArrayIterator[T] {
	return a.IterAt(a.Size() - 1)
}

// IterAt 返回指向指定位置的迭代器。
// 参数：pos 为迭代器指向的下标。
// 返回：position 为 pos 的 *ArrayIterator[T]；不检查越界，调用方需自行保证 pos 在 [-1, a.Size()] 范围内。
func (a *Array[T]) IterAt(pos int) *ArrayIterator[T] {
	return &ArrayIterator[T]{array: a, position: pos}
}

// String 返回数组的字符串表示。
// 返回：使用 fmt.Sprintf("%v", values) 格式化得到的字符串，等价于 values 的默认表示。
func (a *Array[T]) String() string {
	return fmt.Sprintf("%v", a.values)
}
