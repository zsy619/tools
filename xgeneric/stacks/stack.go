package stacks

import "sync"

var default_init_size = 16

// Stack 是基于动态数组实现的并发安全栈结构（后进先出）。
// 通过 sync.Mutex 保护内部状态，支持自动扩容。
type Stack[E any] struct {
	data []E
	pos  int

	lock sync.Mutex

	empty E
}

// NewStack 创建一个使用默认初始容量（16）的空栈。
func NewStack[E any]() *Stack[E] {
	return NewStackSize[E](default_init_size)
}

// NewStackSize 创建一个指定初始容量的空栈。
// 参数：initSize 为期望的初始容量；<= 0 时使用默认值 16。
func NewStackSize[E any](initSize int) *Stack[E] {
	if initSize <= 0 {
		initSize = default_init_size
	}
	s := &Stack[E]{data: make([]E, initSize), pos: -1}
	return s
}

// resize 将内部 data 切片容量翻倍以支持更多元素（仅由 Push 在容量耗尽时调用）。
func (s *Stack[E]) resize() {
	l := len(s.data)
	newData := make([]E, l*2)
	copy(newData, s.data)
	s.data = newData
}

// Push 将元素 e 压入栈顶。
// 当内部容量不足时会自动调用 resize 扩容。
// 副作用：在锁保护下修改 data 与 pos。
func (s *Stack[E]) Push(e E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.Cap() <= 0 {
		s.resize()
	}
	s.pos++
	s.data[s.pos] = e
}

// Pop 从栈顶弹出一个元素并将其从栈中移除。
// 返回：栈顶元素；当栈为空时返回 T 的零值（无错误，调用方应先用 IsEmpty 判断）。
// 副作用：在锁保护下修改 data 与 pos，并清空原位置的值以避免内存泄漏。
func (s *Stack[E]) Pop() (e E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.pos >= 0 {
		e = s.data[s.pos]
		s.data[s.pos] = s.empty
		s.pos--
	}
	return
}

// Cap 返回栈的剩余可用容量。
// 返回：len(data) - pos - 1。
func (s *Stack[E]) Cap() int {
	return len(s.data) - s.pos - 1
}

// IsEmpty 判断栈是否为空。
// 返回：当 pos < 0 时返回 true，否则返回 false。
func (s *Stack[E]) IsEmpty() bool {
	return s.pos < 0
}

// Size 返回栈中当前的元素数量。
func (s *Stack[E]) Size() int {
	return s.pos + 1
}

// Copy 复制当前栈，生成一个数据完全相同的新栈。
// 返回：内部 data 已复制且 pos 与原栈相同的新 *Stack[E]；与原栈共享锁（不会并发安全）。
func (s *Stack[E]) Copy() *Stack[E] {
	data := make([]E, len(s.data))
	copy(data, s.data)
	r := &Stack[E]{data: data, pos: s.pos}
	return r
}
