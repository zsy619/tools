package rings

import (
	"sort"

	"github.com/zsy619/tools"
)

type sortableRing[E any] struct {
	data *Ring[E]
	cmp  tools.CMP[E]
}

func (s sortableRing[E]) Len() int { return s.data.Len() }
func (s sortableRing[E]) Swap(i, j int) {
	e1 := s.data.get(i)
	e2 := s.data.get(j)
	e1.Value, e2.Value = e2.Value, e1.Value
}

func (s sortableRing[E]) Less(i, j int) bool {
	e1 := s.data.get(i)
	e2 := s.data.get(j)
	return s.cmp(e1.Value, e2.Value) <= 0
}

// Ring 表示一个环形链表中的节点。
// 环形链表没有明确的起点或终点，指向任意节点的指针都可代表整个环。
// 空环用 nil 指针表示；Ring 的零值是一个 Value 为 T 零值的单元素环。
type Ring[E any] struct {
	next, prev *Ring[E]
	Value      E // 供调用方使用；本库不会修改该字段
}

// init 将零值环初始化为单元素环（自己指向自己），并返回自身。
func (r *Ring[E]) init() *Ring[E] {
	r.next = r
	r.prev = r
	return r
}

// Next 返回环中的下一个节点；r 不能为空环。
// 注意：当 r 为零值（未初始化）时会自动 init。
func (r *Ring[E]) Next() *Ring[E] {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// Prev 返回环中的上一个节点；r 不能为空环。
// 注意：当 r 为零值（未初始化）时会自动 init。
func (r *Ring[E]) Prev() *Ring[E] {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

// Move 在环中向前（n >= 0）或向后（n < 0）移动 n % r.Len() 个位置，并返回指向最终位置的节点。
// r 不能为空环；n 为 0 时返回自身。
func (r *Ring[E]) Move(n int) *Ring[E] {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// Get 返回环中向前（n >= 0）或向后（n < 0）移动 n % r.Len() 个位置后那个节点的值。
// r 不能为空环；等价于 r.Move(n).Value。
func (r *Ring[E]) Get(n int) (ret E) {
	return r.get(n).Value
}

// get 是 Get 的内部实现，返回目标位置的节点指针。
func (r *Ring[E]) get(n int) (ret *Ring[E]) {
	if r.next == nil {
		return r.init()
	}
	ret = r
	switch {
	case n < 0:
		for ; n < 0; n++ {
			ret = ret.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			ret = ret.next
		}
	}
	return ret
}

// NewRing 创建包含 n 个节点的环形链表，每个节点的 Value 均为 T 的零值。
// 参数：n 为期望节点数。
// 返回：n <= 0 时返回 nil，否则返回包含 n 个节点的 *Ring[E]。
func NewRing[E any](n int) *Ring[E] {
	if n <= 0 {
		return nil
	}
	r := new(Ring[E])
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring[E]{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

// NewRingOf 使用可变参数中的元素依次创建一个环形链表。
// 参数：e 为元素列表。
// 返回：e 为空时返回 nil，否则返回包含 e 所有元素的 *Ring[E]。
func NewRingOf[E any](e ...E) *Ring[E] {
	if len(e) == 0 {
		return nil
	}
	r := &Ring[E]{Value: e[0]}
	p := r

	for i := 1; i < len(e); i++ {
		p.next = &Ring[E]{prev: p, Value: e[i]}
		p = p.next
	}

	p.next = r
	r.prev = p
	return r
}

// Link 将环 r 与环 s 连接，使 r.Next() 指向 s，并返回 r.Next() 原本指向的节点。
// r 不能为空环。
//
// 当 r 与 s 指向同一个环时，Link 会移除 r 与 s 之间的所有元素；
// 被移除的元素形成一个子环作为返回值（若未移除任何元素，则返回原 r.Next()，不为 nil）。
//
// 当 r 与 s 指向不同的环时，Link 会把 s 中的元素插入到 r 之后，合并为单一环；
// 返回值指向 s 最后一个元素被插入之后的那个节点。
func (r *Ring[E]) Link(s *Ring[E]) *Ring[E] {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		// Note: Cannot use multiple assignment because
		// evaluation order of LHS is not specified.
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// LinkValue 在 r 之后插入一个值为 e 的新节点（便捷包装 Link）。
func (r *Ring[E]) LinkValue(e E) {
	nr := &Ring[E]{Value: e}
	r.Link(nr)
}

// Unlink 从环 r 中移除从 r.Next() 开始的 n % r.Len() 个节点。
// 当 n % r.Len() == 0 时 r 保持不变。
// 返回：被移除节点形成的子环；r 不能为空环。n <= 0 时返回 nil。
func (r *Ring[E]) Unlink(n int) *Ring[E] {
	if n <= 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

// Len 计算环 r 中的节点数量。
// 时间复杂度与节点数成正比。
// r 为 nil 时返回 0。
func (r *Ring[E]) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}

// Do 按正向顺序对环中每个节点调用一次 f。
// 注意：若 f 修改了 *r，则 Do 的行为是未定义的。
// r 为 nil 时不进行任何调用。
func (r *Ring[E]) Do(f func(E)) {
	if r != nil {
		f(r.Value)
		for p := r.Next(); p != r; p = p.next {
			f(p.Value)
		}
	}
}

// Iterate 按正向顺序遍历所有元素，回调返回 false 时停止遍历。
// r 为 nil 时不进行任何调用。
func (r *Ring[E]) Iterate(f func(E) bool) {
	if r != nil {
		if !f(r.Value) {
			return
		}
		for p := r.Next(); p != r; p = p.next {
			if !f(p.Value) {
				return
			}
		}
	}
}

// Range 与 Iterate 等价，按正向顺序遍历所有元素。
func (r *Ring[E]) Range(f func(E) bool) {
	r.Iterate(f)
}

// Min 在环中查找最小元素。
// 参数：compare 为比较函数（compare(a,b) < 0 表示 a 小于 b）。
// 返回：最小元素的值；环为空时返回 T 的零值。
func (r *Ring[E]) Min(compare tools.CMP[E]) (min E) {
	return selectByCompareRing(r, func(o1, o2 E) int {
		return compare(o1, o2)
	})
}

// Max 在环中查找最大元素。
// 参数：compare 为比较函数（compare(a,b) < 0 表示 a 小于 b）。
// 返回：最大元素的值；环为空时返回 T 的零值。
func (r *Ring[E]) Max(compare tools.CMP[E]) (min E) {
	return selectByCompareRing(r, func(o1, o2 E) int {
		return compare(o2, o1)
	})
}

// selectByCompareRing 是 Min/Max 共用的内部辅助函数，使用 compare 选出最值。
func selectByCompareRing[E any](r *Ring[E], compare tools.CMP[E]) (v E) {
	i := 0
	r.Do(func(e E) {
		if i == 0 {
			v = e
		} else {
			if compare(v, e) > 0 {
				v = e
			}
		}
		i++
	})
	return
}

// Sort 根据 compare 比较函数对环中的元素进行原地排序（按正向索引顺序重排值）。
// 参数：compare 为比较函数（compare(a,b) < 0 表示 a 应排在 b 之前）。
// 副作用：会就地修改环中节点 Value 的顺序。
func (r *Ring[E]) Sort(compare tools.CMP[E]) {
	sortobject := sortableRing[E]{data: r, cmp: compare}
	sort.Sort(sortobject)
}

// Index 返回环中第一个与 v 相等的元素的下标（从 0 开始）。
// 参数：v 为目标值，f 为相等性比较函数。
// 返回：首个匹配元素的下标；未找到时返回 -1。
func (r *Ring[E]) Index(v E, f tools.EQL[E]) (index int) {
	index = -1
	matched := false
	r.Iterate(func(e E) bool {
		index++
		if f(v, e) {
			matched = true
			return false
		}
		return true
	})
	if !matched {
		index = -1
	}
	return
}

// Contains 判断环中是否包含与 v 相等的元素（由 f 判断相等）。
// 参数：v 为目标值，f 为相等性比较函数。
// 返回：找到则返回 true，否则返回 false。
func (r *Ring[E]) Contains(v E, f tools.EQL[E]) bool {
	return r.Index(v, f) != -1
}

// ToArray 将环中所有元素按正向顺序转换为切片。
// 返回：包含所有元素 Value 的 []E；环为空时返回空切片。
func (r *Ring[E]) ToArray() []E {
	ret := make([]E, 0)
	r.Do(func(e E) {
		ret = append(ret, e)
	})

	return ret
}

// WriteToArray 将环中元素按正向顺序写入传入的切片 v。
// 当 v 长度小于环长度时仅写入 v 的容量个元素；v 为 nil 或长度 0 时不进行任何写入。
func (r *Ring[E]) WriteToArray(v []E) {
	size := len(v)
	pos := 0

	r.Iterate(func(e E) bool {
		if pos < size {
			v[pos] = e
		} else {
			return false
		}
		pos++
		return true
	})
}

// Copy 复制当前环中的所有元素到新环。
// 返回：包含相同元素的新 *Ring[E]；修改新环不会影响原环。
func (r *Ring[E]) Copy() *Ring[E] {
	ret := NewRing[E](1)
	i := 0
	r.Do(func(e E) {
		if i == 0 {
			ret.Value = e
		} else {
			ret.LinkValue(e)
		}
		i++
	})
	return ret
}
