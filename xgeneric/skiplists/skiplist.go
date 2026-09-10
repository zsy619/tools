package skiplists

import (
	"errors"
	"math/rand"
	gosync "sync"
	"time"

	"github.com/zsy619/tools/xgeneric/utils"
)

var (
	defaultMaxLevel = 10
	defaultLocker   utils.FakeLocker
)
var ErrorNotFound = errors.New("not found")

// Options 持有 Skiplist 的配置项（内部使用，外部通过 Option 函数设置）。
type Options struct {
	maxLevel int
	locker   utils.Locker
}

// Option 是用于配置 Options 的函数（Functional Options 模式）。
type Option func(option *Options)

// WithGoroutineSafe 启用 Skiplist 的并发安全特性（使用 sync.RWMutex 保护内部数据）。
func WithGoroutineSafe() Option {
	return func(option *Options) {
		option.locker = &gosync.RWMutex{}
	}
}

// WithMaxLevel 设置 Skiplist 的最大层数。
// 参数：maxLevel 为期望的最大层数（必须 >= 1）。
func WithMaxLevel(maxLevel int) Option {
	return func(option *Options) {
		option.maxLevel = maxLevel
	}
}

// Node 表示跳表中的一个层级节点，仅持有指向各层下一个 Element 的指针数组。
type Node[K, V any] struct {
	next []*Element[K, V]
}

// Element 是跳表中存储键值对数据的节点，内嵌 Node 并额外持有 key 与 value。
type Element[K, V any] struct {
	Node[K, V]
	key   K
	value V
}

// Skiplist 是通过以空间换时间实现快速查找的数据结构（跳表）。
// 支持自定义键比较器与可选的并发安全模式。
type Skiplist[K, V any] struct {
	locker         utils.Locker
	head           Node[K, V]
	maxLevel       int
	keyCmp         utils.Comparator[K]
	len            int
	prevNodesCache []*Node[K, V]
	rander         *rand.Rand
}

// New 创建一个跳表。
// 参数：cmp 为键比较器（返回 <0/<0/>0 分别表示 a<b/a==b/a>b）；opts 为可选的配置项（如 WithGoroutineSafe、WithMaxLevel）。
// 返回：已分配好头节点与缓存的 *Skiplist[K, V]。
func New[K, V any](cmp utils.Comparator[K], opts ...Option) *Skiplist[K, V] {
	option := Options{
		maxLevel: defaultMaxLevel,
		locker:   defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	l := &Skiplist[K, V]{
		locker:   option.locker,
		maxLevel: option.maxLevel,
		keyCmp:   cmp,
		rander:   rand.New(rand.NewSource(time.Now().Unix())),
	}
	l.head.next = make([]*Element[K, V], l.maxLevel)
	l.prevNodesCache = make([]*Node[K, V], l.maxLevel)
	return l
}

// Insert 将键值对插入跳表。
// 参数：key 为键，value 为值。
// 当 key 已存在时直接更新其 value。
// 副作用：会就地修改跳表结构（持有锁时执行，goroutine-safe 模式下安全）。
func (sl *Skiplist[K, V]) Insert(key K, value V) {
	sl.locker.Lock()
	defer sl.locker.Unlock()
	prevs := sl.findPrevNodes(key)

	if prevs[0].next[0] != nil && sl.keyCmp(prevs[0].next[0].key, key) == 0 {
		// same key, update value
		prevs[0].next[0].value = value
		return
	}

	level := sl.randomLevel()

	e := &Element[K, V]{
		key:   key,
		value: value,
		Node: Node[K, V]{
			next: make([]*Element[K, V], level),
		},
	}

	for i := range e.next {
		e.next[i] = prevs[i].next[i]
		prevs[i].next[i] = e
	}

	sl.len++
}

// Get 根据 key 查找对应的值。
// 参数：key 为要查找的键。
// 返回：找到时返回对应的 value 与 nil；未找到时返回 V 的零值与 ErrorNotFound 错误。
func (sl *Skiplist[K, V]) Get(key K) (V, error) {
	sl.locker.RLock()
	defer sl.locker.RUnlock()

	pre := &sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		cur := pre.next[i]
		for ; cur != nil; cur = cur.next[i] {
			cmpRet := sl.keyCmp(cur.key, key)
			if cmpRet == 0 {
				return cur.value, nil
			}
			if cmpRet > 0 {
				break
			}
			pre = &cur.Node
		}
	}
	return *new(V), ErrorNotFound
}

// Remove 从跳表中移除指定 key 对应的键值对。
// 参数：key 为要移除的键。
// 返回：成功移除返回 true；若 key 不存在则返回 false（无副作用）。
func (sl *Skiplist[K, V]) Remove(key K) bool {
	sl.locker.Lock()
	defer sl.locker.Unlock()

	prevs := sl.findPrevNodes(key)
	element := prevs[0].next[0]
	if element == nil {
		return false
	}
	if element != nil && sl.keyCmp(element.key, key) != 0 {
		return false
	}

	for i, v := range element.next {
		prevs[i].next[i] = v
	}
	sl.len--
	return true
}

// Len 返回跳表中键值对的数量。
func (sl *Skiplist[K, V]) Len() int {
	sl.locker.RLock()
	defer sl.locker.RUnlock()
	return sl.len
}

// randomLevel 使用幂律分布随机生成新节点的层数（在 [1, maxLevel] 之间）。
func (sl *Skiplist[K, V]) randomLevel() int {
	total := uint64(1)<<uint64(sl.maxLevel) - 1 // 2^n-1
	k := sl.rander.Uint64() % total
	levelN := uint64(1) << (uint64(sl.maxLevel) - 1)

	level := 1
	for total -= levelN; total > k; level++ {
		levelN >>= 1
		total -= levelN
	}
	return level
}

// findPrevNodes 查找各层中目标 key 的前驱节点，缓存到 prevNodesCache 中以复用内存。
func (sl *Skiplist[K, V]) findPrevNodes(key K) []*Node[K, V] {
	prevs := sl.prevNodesCache
	prev := &sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		if sl.head.next[i] != nil {
			for next := prev.next[i]; next != nil; next = next.next[i] {
				if sl.keyCmp(next.key, key) >= 0 {
					break
				}
				prev = &next.Node
			}
		}
		prevs[i] = prev
	}
	return prevs
}

// Traversal 按 key 升序遍历跳表中的所有键值对。
// 每访问一对 (key, value) 时调用 visitor；当 visitor 返回 false 时停止遍历。
func (sl *Skiplist[K, V]) Traversal(visitor utils.KvVisitor[K, V]) {
	sl.locker.RLock()
	defer sl.locker.RUnlock()

	for e := sl.head.next[0]; e != nil; e = e.next[0] {
		if !visitor(e.key, e.value) {
			return
		}
	}
}

// Keys 返回跳表中所有键组成的切片（按 key 升序）。
// 返回：包含所有键的 []K；空跳表时返回 nil。
func (sl *Skiplist[K, V]) Keys() []K {
	var keys []K
	sl.Traversal(func(key K, value V) bool {
		keys = append(keys, key)
		return true
	})
	return keys
}
