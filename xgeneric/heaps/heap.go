package heaps

import (
	"container/heap"
	"fmt"

	"github.com/zsy619/tools"
)

// heapST 是对标准库 container/heap 中 heap.Interface 的实现包装。
// 通过持有元素切片与自定义比较函数 cmp，支持任意类型的堆操作。
type heapST[E any] struct {
	data []E
	cmp  tools.CMP[E]
}

// Len 返回堆中元素的数量（满足 heap.Interface）。
func (h *heapST[E]) Len() int { return len(h.data) }

// Less 比较下标 i 与 j 处元素的大小，使用 cmp 函数（满足 heap.Interface）。
// 返回：cmp(data[i], data[j]) < 0 的结果。
func (h *heapST[E]) Less(i, j int) bool {
	v := h.cmp(h.data[i], h.data[j])
	return v < 0
}

// Swap 交换下标 i 与 j 处的元素（满足 heap.Interface）。
func (h *heapST[E]) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

// Push 将元素 x 添加到堆末尾（满足 heap.Interface）。
// 使用指针接收者是因为该方法会修改切片的长度，而非仅修改其内容。
func (h *heapST[E]) Push(x any) {
	v := append(h.data, x.(E))
	h.data = v
}

// Pop 移除并返回堆末尾的元素（满足 heap.Interface）。
// 注意：标准库的 heap.Pop 会先调用此方法，再执行上浮/下沉调整。
func (h *heapST[E]) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

// Heap 是基于泛型的通用堆数据结构。
// 内部通过 heapST 持有元素与比较函数，对外提供类型安全的 API。
type Heap[E any] struct {
	data *heapST[E]
}

// Push 将元素 v 压入堆中。
// 时间复杂度：O(log n)，其中 n = h.Len()。
// 副作用：会就地修改 h 内部的堆结构。
func (h *Heap[E]) Push(v E) {
	heap.Push(h.data, v)
}

// Pop 移除并返回堆顶元素（依据 Less 定义的最值）。
// 时间复杂度：O(log n)，其中 n = h.Len()。
// 等价于 Remove(h, 0)。
func (h *Heap[E]) Pop() E {
	return heap.Pop(h.data).(E)
}

// Element 返回堆中下标为 index 的元素值。
// 参数：index 为目标下标。
// 返回：对应元素及 nil 错误；当 index 越界（<0 或 >= Len()）时返回零值与 error（信息为 "out of index"）。
func (h *Heap[E]) Element(index int) (e E, err error) {
	if index < 0 || index >= h.data.Len() {
		return e, fmt.Errorf("out of index")
	}
	return h.data.data[index], nil
}

// Remove 移除并返回堆中下标为 index 的元素。
// 时间复杂度：O(log n)，其中 n = h.Len()。
// 注意：调用方需保证 index 在 [0, Len()) 范围内；越界访问会引发运行时 panic。
func (h *Heap[E]) Remove(index int) E {
	return heap.Remove(h.data, index).(E)
}

// Len 返回堆中元素的数量。
func (h *Heap[E]) Len() int {
	return len(h.data.data)
}

// Copy 复制当前堆，生成一个状态完全一致的新堆。
// 返回：内部数据已复制并完成堆初始化的新 *Heap[E]；与原堆共享比较函数但拥有独立的数据切片。
func (h *Heap[E]) Copy() *Heap[E] {
	ret := heapST[E]{cmp: h.data.cmp}
	ret.data = make([]E, len(h.data.data))
	copy(ret.data, h.data.data)
	heap.Init(&ret)
	return &Heap[E]{&ret}
}

// NewHeap 创建一个堆并使用 heap.Init 完成初始化。
// 参数：t 为初始元素切片（顺序任意，初始化时会调整为堆结构），cmp 为比较函数（与 Less 一致）。
// 返回：堆化完毕的 *Heap[E]，初始大小与 t 相同。
func NewHeap[E any](t []E, cmp tools.CMP[E]) *Heap[E] {
	ret := heapST[E]{data: t, cmp: cmp}
	heap.Init(&ret)
	return &Heap[E]{&ret}
}
