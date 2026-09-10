package utils

// Visitor 是用于遍历数据结构的访问者函数类型（返回 false 表示停止遍历）。
type Visitor[V any] func(value V) bool

// KvVisitor 是用于遍历键值对数据结构的访问者函数类型（返回 false 表示停止遍历）。
type KvVisitor[K, V any] func(key K, value V) bool
