package utils

// ConstIterator 是常量（只读）迭代器的接口。
// 实现需提供 IsValid、Next、Value、Clone、Equal 等方法。
type ConstIterator[T any] interface {
	IsValid() bool
	Next() ConstIterator[T]
	Value() T
	Clone() ConstIterator[T]
	Equal(other ConstIterator[T]) bool
}

// Iterator 是可变迭代器的接口（在 ConstIterator 基础上增加 SetValue）。
type Iterator[T any] interface {
	ConstIterator[T]
	SetValue(value T)
}

// ConstKvIterator 是常量键值对迭代器的接口（ConstIterator[V] 加 Key() 方法）。
type ConstKvIterator[K, V any] interface {
	ConstIterator[V]
	Key() K
}

// KvIterator 是可变键值对迭代器的接口（在 ConstKvIterator 基础上增加 SetValue）。
type KvIterator[K, V any] interface {
	ConstKvIterator[K, V]
	SetValue(value V)
}

// ConstBidIterator 是常量双向迭代器的接口（ConstIterator 加 Prev() 方法）。
type ConstBidIterator[T any] interface {
	ConstIterator[T]
	Prev() ConstBidIterator[T]
}

// BidIterator 是可变双向迭代器的接口（在 ConstBidIterator 基础上增加 SetValue）。
type BidIterator[T any] interface {
	ConstBidIterator[T]
	SetValue(value T)
}

// ConstKvBidIterator 是常量键值对双向迭代器的接口。
type ConstKvBidIterator[K, V any] interface {
	ConstKvIterator[K, V]
	BidIterator[V]
}

// KvBidIterator 是可变键值对双向迭代器的接口。
type KvBidIterator[K, V any] interface {
	KvIterator[K, V]
	BidIterator[V]
}

// RandomAccessIterator 是可变随机访问迭代器的接口（BidIterator 加 IteratorAt 与 Position）。
type RandomAccessIterator[T any] interface {
	BidIterator[T]
	// IteratorAt 返回指向指定 position 的新迭代器。
	IteratorAt(position int) RandomAccessIterator[T]
	// Position 返回当前迭代器对应的下标。
	Position() int
}
