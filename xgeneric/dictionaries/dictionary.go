package dictionaries

// Dictionary 是 Go 内置 map 类型的别名。
// 用于以更具语义的名称表达键值对集合；与 map[K]V 完全等价。
type Dictionary[K comparable, V any] map[K]V
