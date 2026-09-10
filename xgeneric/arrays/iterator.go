package arrays

import "github.com/zsy619/tools/xgeneric/utils"

type T any

// ArrayIterator 实现了 utils.RandomAccessIterator 接口（编译期断言）。
var _ utils.RandomAccessIterator[T] = (*ArrayIterator[T])(nil)

// ArrayIterator 是 Array 的迭代器实现，支持随机访问。
// 内含数组引用与当前位置 position，可通过 Next/Prev 等方法移动。
type ArrayIterator[T any] struct {
	array    *Array[T]
	position int
}

// IsValid 判断迭代器是否指向数组中的有效位置。
// 返回：当 position 在 [0, array.Size()) 范围内时返回 true，否则返回 false（包含负数与越界位置）。
func (iter *ArrayIterator[T]) IsValid() bool {
	if iter.position >= 0 && iter.position < iter.array.Size() {
		return true
	}
	return false
}

// Value 返回迭代器当前位置对应的数组元素。
// 返回：array.At(position)。
// 当 position 越界时，At 会 panic("index out off range")。
func (iter *ArrayIterator[T]) Value() T {
	return iter.array.At(iter.position)
}

// SetValue 将数组在迭代器当前位置的值设置为 val。
// 参数：val 为要写入的新值。
// 副作用：若 position 越界则静默忽略（依赖 Array.Set 的越界保护）。
func (iter *ArrayIterator[T]) SetValue(val T) {
	iter.array.Set(iter.position, val)
}

// Next 将迭代器位置向后移动一位。
// 返回：自身（满足迭代器接口的链式调用）。
// 副作用：当 position < array.Size() 时递增；越界后再次调用不再变化。
func (iter *ArrayIterator[T]) Next() utils.ConstIterator[T] {
	if iter.position < iter.array.Size() {
		iter.position++
	}
	return iter
}

// Prev 将迭代器位置向前移动一位。
// 返回：自身（满足双向迭代器接口的链式调用）。
// 副作用：当 position >= 0 时递减；越界后再次调用不再变化（最小为 -1）。
func (iter *ArrayIterator[T]) Prev() utils.ConstBidIterator[T] {
	if iter.position >= 0 {
		iter.position--
	}
	return iter
}

// Clone 复制当前迭代器，生成一个状态相同的新迭代器。
// 返回：与当前迭代器引用同一数组且 position 相同的新 *ArrayIterator[T]，对新迭代器的修改不会影响原迭代器。
func (iter *ArrayIterator[T]) Clone() utils.ConstIterator[T] {
	return &ArrayIterator[T]{array: iter.array, position: iter.position}
}

// IteratorAt 在保持数组引用不变的前提下，返回指向新位置 pos 的迭代器。
// 参数：pos 为新迭代器的初始位置。
// 返回：position 为 pos 的新 *ArrayIterator[T]；调用方需自行保证 pos 在 [-1, array.Size()] 范围内。
func (iter *ArrayIterator[T]) IteratorAt(pos int) utils.RandomAccessIterator[T] {
	return &ArrayIterator[T]{array: iter.array, position: pos}
}

// Position 返回迭代器当前指向的下标。
// 返回：position 字段值，范围可能为 [-1, array.Size()]。
func (iter *ArrayIterator[T]) Position() int {
	return iter.position
}

// Equal 判断两个迭代器是否指向同一位置且引用同一数组。
// 参数：other 为待比较的迭代器（必须能转换为 *ArrayIterator[T]）。
// 返回：当 other 指向同一数组的同一位置时返回 true，否则返回 false；若 other 不是 ArrayIterator 类型同样返回 false。
func (iter *ArrayIterator[T]) Equal(other utils.ConstIterator[T]) bool {
	otherIter, ok := other.(*ArrayIterator[T])
	if !ok {
		return false
	}
	if otherIter.array == iter.array && otherIter.position == iter.position {
		return true
	}
	return false
}
