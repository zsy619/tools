package lists

import (
	"sort"

	"github.com/zsy619/tools"
)

type sortableList[E any] struct {
	data *List[E]
	cmp  tools.CMP[E]
}

func (s sortableList[E]) Len() int { return s.data.Len() }
func (s sortableList[E]) Swap(i, j int) {
	e1, _ := s.data.get(i)
	e2, _ := s.data.get(j)
	e1.Value, e2.Value = e2.Value, e1.Value
}

func (s sortableList[E]) Less(i, j int) bool {
	e1, _ := s.data.Get(i)
	e2, _ := s.data.Get(j)
	return s.cmp(e1, e2) <= 0
}

// Element 表示双向链表中的一个节点。
type Element[E any] struct {
	// 双向链表中的前驱与后继指针。
	// 为简化实现，链表内部按环状结构组织，&l.root 同时为链表最后一个元素（l.Back()）的后继
	// 与第一个元素（l.Front()）的前驱。
	next, prev *Element[E]

	// 当前节点所属的链表。
	list *List[E]

	// 当前节点存储的值。
	Value E
}

// Next 返回当前节点的后继节点；若不存在则返回 nil。
func (e *Element[E]) Next() *Element[E] {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev 返回当前节点的前驱节点；若不存在则返回 nil。
func (e *Element[E]) Prev() *Element[E] {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// List 表示一个双向链表。
// List 的零值即为可用的空链表。
type List[E any] struct {
	root Element[E] // 哨兵节点，仅使用 &root、root.prev、root.next
	len  int        // 当前链表长度（不含哨兵节点）
}

// Init 初始化或清空链表 l。
// 返回：经过初始化后的 l，便于链式调用。
func (l *List[E]) Init() *List[E] {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New 返回一个已初始化的空链表。
func NewList[E any]() *List[E] { return new(List[E]).Init() }

// NewListFromArray 返回一个已初始化的链表，并使用目标切片的元素填充（顺序追加到尾部）。
// 参数：arr 为输入元素切片。
// 返回：以 arr 中顺序依次 PushBack 元素得到的新 *List[E]。
func NewListFromArray[E any](arr []E) *List[E] {
	l := new(List[E]).Init()
	for _, v := range arr {
		l.PushBack(v)
	}

	return l
}

// NewListOf 返回一个已初始化的链表，并使用可变参数中的元素填充（顺序追加到尾部）。
// 参数：e 为可变数量的元素。
// 返回：以参数顺序依次 PushBack 元素得到的新 *List[E]。
func NewListOf[E any](e ...E) *List[E] {
	l := new(List[E]).Init()
	for _, v := range e {
		l.PushBack(v)
	}
	return l
}

// Len 返回链表中的元素数量。
// 时间复杂度：O(1)。
func (l *List[E]) Len() int { return l.len }

// IsEmpty 判断链表是否为空。
// 返回：当 len == 0 时返回 true，否则返回 false。
func (l *List[E]) IsEmpty() bool { return l.len == 0 }

// Front 返回链表的第一个节点；链表为空时返回 nil。
func (l *List[E]) Front() *Element[E] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// FrontValue 返回链表第一个节点的值。
// 返回：第一个节点存储的值；链表为空时返回 T 的零值。
func (l *List[E]) FrontValue() (v E) {
	if e := l.Front(); e != nil {
		v = e.Value
	}
	return
}

// Back 返回链表的最后一个节点；链表为空时返回 nil。
func (l *List[E]) Back() *Element[E] {
	if l.IsEmpty() {
		return nil
	}
	return l.root.prev
}

// BackValue 返回链表最后一个节点的值。
// 返回：最后一个节点存储的值；链表为空时返回 T 的零值。
func (l *List[E]) BackValue() (v E) {
	if e := l.Back(); e != nil {
		v = e.Value
	}
	return
}

// lazyInit 在第一次使用时对零值 List 进行延迟初始化。
// 当 root.next 还未指向自身时调用 Init 完成初始化。
func (l *List[E]) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert 将 e 插入到 at 之后，增加链表长度并返回 e。
func (l *List[E]) insert(e, at *Element[E]) *Element[E] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue 是 insert(&Element{Value: v}, at) 的便捷包装。
func (l *List[E]) insertValue(v E, at *Element[E]) *Element[E] {
	return l.insert(&Element[E]{Value: v}, at)
}

// remove 将 e 从所属链表中移除并递减链表长度，同时清空 e 的指针以避免内存泄漏。
func (l *List[E]) remove(e *Element[E]) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
}

// move 将节点 e 移动到 at 之后（保持链表环形结构完整）。
// 当 e == at 时直接返回，不会产生变化。
func (l *List[E]) move(e, at *Element[E]) {
	if e == at {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
}

// RemoveElement 将节点 e 从链表 l 中移除，并返回其值 e.Value。
// 当 e 不属于 l 时链表保持不变（但仍然返回 e.Value）。
// 注意：e 不能为 nil。
func (l *List[E]) RemoveElement(e *Element[E]) E {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// PushFront 将值为 v 的新节点插入链表头部，并返回该节点。
func (l *List[E]) PushFront(v E) *Element[E] {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack 将值为 v 的新节点插入链表尾部，并返回该节点。
func (l *List[E]) PushBack(v E) *Element[E] {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore 将值为 v 的新节点插入到 mark 节点之前，并返回该节点。
// 当 mark 不属于链表 l 时链表保持不变并返回 nil。
// 注意：mark 不能为 nil。
func (l *List[E]) InsertBefore(v E, mark *Element[E]) *Element[E] {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark.prev)
}

// InsertAfter 将值为 v 的新节点插入到 mark 节点之后，并返回该节点。
// 当 mark 不属于链表 l 时链表保持不变并返回 nil。
// 注意：mark 不能为 nil。
func (l *List[E]) InsertAfter(v E, mark *Element[E]) *Element[E] {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

// MoveToFront 将节点 e 移动到链表 l 的头部。
// 当 e 不属于链表 l 或 e 已在头部时链表保持不变。
// 注意：e 不能为 nil。
func (l *List[E]) MoveToFront(e *Element[E]) {
	if e.list != l || l.root.next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, &l.root)
}

// MoveToBack 将节点 e 移动到链表 l 的尾部。
// 当 e 不属于链表 l 或 e 已在尾部时链表保持不变。
// 注意：e 不能为 nil。
func (l *List[E]) MoveToBack(e *Element[E]) {
	if e.list != l || l.root.prev == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, l.root.prev)
}

// MoveBefore 将节点 e 移动到 mark 节点之前。
// 当 e 或 mark 不属于链表 l，或 e == mark 时链表保持不变。
// 注意：e 与 mark 均不能为 nil。
func (l *List[E]) MoveBefore(e, mark *Element[E]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter 将节点 e 移动到 mark 节点之后。
// 当 e 或 mark 不属于链表 l，或 e == mark 时链表保持不变。
// 注意：e 与 mark 均不能为 nil。
func (l *List[E]) MoveAfter(e, mark *Element[E]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

// PushBackList 将另一个链表 other 的副本依次插入到链表 l 的尾部。
// l 与 other 可以是同一个链表；两者均不能为 nil。
func (l *List[E]) PushBackList(other *List[E]) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList 将另一个链表 other 的副本依次插入到链表 l 的头部。
// l 与 other 可以是同一个链表；两者均不能为 nil。
func (l *List[E]) PushFrontList(other *List[E]) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

// ToArray 将链表中的所有元素按顺序转换为切片。
// 返回：按链表从前到后顺序排列的 []E；链表为空时返回空切片。
func (l *List[E]) ToArray() []E {
	if l.IsEmpty() {
		return []E{}
	}
	ret := make([]E, l.len)
	e := l.Front()
	i := 0
	for e != nil {
		ret[i] = e.Value
		e = e.Next()
		i++
	}
	return ret
}

// WriteToArray 将链表中的元素按顺序写入传入的切片 v。
// 当 v 的长度小于链表长度时仅写入 v 的容量个元素；当 v 较长时多余位置保持不变。
// 链表为空或 v 为 nil 或长度为 0 时函数直接返回。
func (l *List[E]) WriteToArray(v []E) {
	if l.IsEmpty() || v == nil || len(v) == 0 {
		return
	}

	size := len(v)
	pos := 0
	l.iterate(func(e *Element[E]) bool {
		if pos < size {
			v[pos] = e.Value
		} else {
			return false
		}
		pos++
		return true
	})
}

// Iterate 按从前到后的顺序遍历所有元素。
// 参数：f 为遍历回调函数，对每个元素调用一次，返回 false 时停止遍历。
// 链表为空时直接返回。
func (l *List[E]) Iterate(f tools.Func[bool, E]) {
	if l.IsEmpty() {
		return
	}
	l.iterate(func(e *Element[E]) bool {
		return f(e.Value)
	})
}

// IterateReverse 按从后到前的顺序遍历所有元素。
// 参数：f 为遍历回调函数，对每个元素调用一次，返回 false 时停止遍历。
// 链表为空时直接返回。
func (l *List[E]) IterateReverse(f tools.Func[bool, E]) {
	if l.IsEmpty() {
		return
	}
	l.iterateReverse(func(e *Element[E]) bool {
		return f(e.Value)
	})
}

// iterate 是内部遍历辅助，按从前到后的顺序对节点调用 f，返回 false 时停止。
func (l *List[E]) iterate(f tools.Func[bool, *Element[E]]) {
	e := l.Front()
	for e != nil {
		next := e.Next()
		if !f(e) {
			break
		}
		e = next
	}
}

// iterateReverse 是内部遍历辅助，按从后到前的顺序对节点调用 f，返回 false 时停止。
func (l *List[E]) iterateReverse(f tools.Func[bool, *Element[E]]) {
	e := l.Back()
	for e != nil {
		next := e.Prev()
		if !f(e) {
			break
		}
		e = next
	}
}

// Contains 判断链表中是否存在与 v 相等的元素（由 f 判断相等）。
// 参数：v 为要查找的目标值，f 为相等性比较函数。
// 返回：找到则返回 true，否则返回 false。
func (l *List[E]) Contains(v E, f tools.EQL[E]) (contains bool) {
	return l.Index(v, f) != -1
}

// Index 返回链表中第一个与 v 相等的元素的下标（从 0 开始）。
// 参数：v 为目标值，f 为相等性比较函数。
// 返回：首个匹配元素的下标；未找到时返回 -1。
func (l *List[E]) Index(v E, f tools.EQL[E]) (index int) {
	index = -1
	matched := false
	l.Iterate(func(e E) bool {
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

// LastIndex 返回链表中最后一个与 v 相等的元素的下标（从 0 开始）。
// 参数：v 为目标值，f 为相等性比较函数。
// 返回：最后一个匹配元素的下标；未找到时返回 -1。
func (l *List[E]) LastIndex(v E, f tools.EQL[E]) (index int) {
	index = l.Len()
	matched := false
	l.IterateReverse(func(e E) bool {
		index--
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

// Clear 移除链表中的所有元素，使其变为空链表。
func (l *List[E]) Clear() {
	l.iterate(func(e *Element[E]) bool {
		l.remove(e)
		return true
	})
}

// Remove 删除链表中第一个与 v 相等的元素（由 f 判断相等）。
// 参数：v 为要删除的目标值，f 为相等性比较函数。
// 返回：被删除元素的值与 true；未找到时返回 T 的零值与 false。
func (l *List[E]) Remove(v E, f tools.EQL[E]) (ret E, removed bool) {
	return l.removeMatches(v, f, true)
}

// RemoveAll 删除链表中所有与 v 相等的元素（由 f 判断相等）。
// 参数：v 为要删除的目标值，f 为相等性比较函数。
// 返回：最后一个被删除元素的值与 true；未找到时返回 T 的零值与 false。
func (l *List[E]) RemoveAll(v E, f tools.EQL[E]) (ret E, removed bool) {
	return l.removeMatches(v, f, false)
}

// removeMatches 删除链表中所有（或首个）与 v 相等的元素；first=true 时仅删除第一个匹配项。
func (l *List[E]) removeMatches(v E, f tools.EQL[E], first bool) (ret E, removed bool) {
	l.iterate(func(e *Element[E]) bool {
		if f(v, e.Value) {
			removed = true
			ret = l.RemoveElement(e)
			return !first
		}
		return true
	})
	return
}

// Get 返回下标为 index 的节点的值。
// 参数：index 为目标下标。
// 返回：该位置上的元素值与 true；当 index 越界时返回 T 的零值与 false。
func (l *List[E]) Get(index int) (E, bool) {
	var ret E
	if !l.isElementIndex(index) {
		return ret, false
	}

	e, b := l.get(index)
	return e.Value, b
}

// get 返回下标为 index 的内部节点。
// 参数：index 为目标下标。
// 返回：对应 *Element[E] 与 true；当 index 越界时返回 nil 与 false。
func (l *List[E]) get(index int) (*Element[E], bool) {
	if !l.isElementIndex(index) {
		return nil, false
	}

	return l.node(index), true
}

// Set 将下标为 index 的节点值设置为 v。
// 参数：index 为目标下标，v 为新值。
// 返回：设置成功返回 true；当 index 越界时返回 false 且链表保持不变。
func (l *List[E]) Set(index int, v E) bool {
	if !l.isElementIndex(index) {
		return false
	}
	e := l.node(index)
	e.Value = v
	return true
}

// Add 将值为 v 的新节点插入到指定 index 之后。
// 参数：index 为参考下标（其后插入新节点），v 为新值。
// 返回：插入成功返回 true；当 index 越界时返回 false。空链表时直接 PushFront 并返回 true。
func (l *List[E]) Add(index int, v E) bool {
	if l.IsEmpty() {
		l.PushFront(v)
		return true
	}
	if !l.isElementIndex(index) {
		return false
	}
	e := l.node(index)
	l.InsertAfter(v, e)
	return true
}

// RemoveFront 移除并返回链表头部的元素。
// 返回：被移除节点的值；链表为空时返回 T 的零值。
func (l *List[E]) RemoveFront() E {
	var ret E
	if l.IsEmpty() {
		return ret
	}
	e := l.Front()
	ret = e.Value
	l.remove(e)
	return ret
}

// RemoveBack 移除并返回链表尾部的元素。
// 返回：被移除节点的值；链表为空时返回 T 的零值。
func (l *List[E]) RemoveBack() E {
	var ret E
	if l.IsEmpty() {
		return ret
	}
	e := l.Back()
	ret = e.Value
	l.remove(e)
	return ret
}

func (l *List[E]) node(index int) *Element[E] {
	if index < (l.len >> 1) { // pos value is before middle value
		e := l.Front()
		for i := 0; i < index; i++ {
			e = e.next
		}
		return e
	} else {
		e := l.Back()
		for i := l.len - 1; i > index; i-- {
			e = e.prev
		}
		return e
	}
}

func (l *List[E]) isElementIndex(index int) bool {
	return index >= 0 && index < l.len
}

// Filter 根据谓词函数过滤元素，返回一个仅包含满足条件元素的新链表。
// 参数：test 为判断函数，对每个元素调用，返回 true 时保留。
// 返回：新创建的 *List[E]；当 test 为 nil 时返回空链表。
func (l *List[E]) Filter(test tools.Evaluate[E]) *List[E] {
	ret := NewList[E]()
	if test == nil {
		return ret
	}

	l.iterate(func(e *Element[E]) bool {
		if test(e.Value) {
			ret.PushBack(e.Value)
		}
		return true
	})

	return ret
}

// Min 在链表中查找最小元素。
// 参数：compare 为比较函数（compare(a,b) < 0 表示 a 小于 b）。
// 返回：最小元素的值；链表为空时返回 T 的零值。
func (l *List[E]) Min(compare tools.CMP[E]) (min E) {
	return selectByCompare(l, func(o1, o2 E) int {
		return compare(o1, o2)
	})
}

// Max 在链表中查找最大元素。
// 参数：compare 为比较函数（compare(a,b) < 0 表示 a 小于 b）。
// 返回：最大元素的值；链表为空时返回 T 的零值。
func (l *List[E]) Max(compare tools.CMP[E]) (min E) {
	return selectByCompare(l, func(o1, o2 E) int {
		return compare(o2, o1)
	})
}

func selectByCompare[E any](l *List[E], compare tools.CMP[E]) (v E) {
	i := 0
	l.iterate(func(e *Element[E]) bool {
		if i == 0 {
			v = e.Value
		} else {
			if compare(v, e.Value) > 0 {
				v = e.Value
			}
		}
		i++
		return true
	})
	return
}

// Sort 根据 compare 比较函数对链表中的元素进行原地排序。
// 参数：compare 为比较函数（compare(a,b) < 0 表示 a 应排在 b 之前）。
// 副作用：会就地调整 l 中元素的顺序。
func (l *List[E]) Sort(compare tools.CMP[E]) {
	sortobject := sortableList[E]{data: l, cmp: compare}
	sort.Sort(sortobject)
}

// Copy 复制链表中的所有元素到新链表。
// 返回：包含相同元素的新 *List[E]；修改新链表不会影响原链表。
func (l *List[E]) Copy() *List[E] {
	ret := NewList[E]()
	l.iterate(func(e *Element[E]) bool {
		ret.PushBack(e.Value)
		return true
	})
	return ret
}

// Range 与 Iterate 等价，按从前到后顺序遍历元素。
// 参数：c 为遍历回调函数（返回 false 时停止）。
func (l *List[E]) Range(c tools.Func[bool, E]) {
	l.Iterate(c)
}
