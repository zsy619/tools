package xmap

import (
	"sync"

	"github.com/zsy619/tools"
	"github.com/zsy619/tools/xreflect"
)

// Map 是带互斥锁的泛型映射，类似于 Go 内置的 map[interface{}]interface{}，但提供了更多实用的操作方法。
type Map[K comparable, V any] struct {
	mp         map[K]V
	empty      V
	sync.Mutex // 创建互斥锁
}

// NewMap 创建一个新的 Map。
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{mp: make(map[K]V)}
}

// Put 向 map 中写入 key 和 value，并返回该 key 原先对应的值。
func (m *Map[K, V]) Put(key K, value V) V {
	m.Lock()
	defer m.Unlock()
	v := m.mp[key]
	m.mp[key] = value
	return v
}

// Get 获取 key 对应的 value，并返回该 key 是否存在于 map 中。
func (m *Map[K, V]) Get(key K) (V, bool) {
	m.Lock()
	defer m.Unlock()
	v, ok := m.mp[key]
	return v, ok
}

// IsEmpty 当 map 中没有任何 key 时返回 true。
func (m *Map[K, V]) IsEmpty() (empty bool) {
	m.Lock()
	defer m.Unlock()
	return m.mp == nil || len(m.mp) == 0
}

// Size 返回 map 中键值对的数量。
func (m *Map[K, V]) Size() int {
	m.Lock()
	defer m.Unlock()
	if m.mp == nil {
		return 0
	}
	return len(m.mp)
}

// ToMap 将 map 中的键值复制并转换为原始的 map 结构后返回。
func (m *Map[K, V]) ToMap() map[K]V {
	m.Lock()
	defer m.Unlock()
	if m.mp == nil {
		return nil
	}

	return Clone(m.mp)
}

// Range 依次对 map 中每个键值对调用函数 f；当 f 返回 false 时提前终止遍历。
func (m *Map[K, V]) Range(f tools.BiFunc[bool, K, V]) {
	m.Lock()
	defer m.Unlock()
	if m.mp == nil {
		return
	}

	for k, v := range m.mp {
		ok := f(k, v)
		if !ok {
			break
		}
	}
}

// Values 返回 map 中所有 value 组成的切片。
func (m *Map[K, V]) Values() []V {
	m.Lock()
	defer m.Unlock()
	ret := make([]V, 0)
	m.Range(func(key K, value V) bool {
		ret = append(ret, value)
		return true
	})
	return ret
}

// Keys 返回 map 中所有 key 组成的切片。
func (m *Map[K, V]) Keys() []K {
	m.Lock()
	defer m.Unlock()
	ret := make([]K, 0)
	m.Range(func(key K, value V) bool {
		ret = append(ret, key)
		return true
	})
	return ret
}

// Clear 移除 map 中所有的键值。
func (m *Map[K, V]) Clear() {
	m.Lock()
	defer m.Unlock()
	Clear(m.mp)
}

// Copy 将当前 map 的所有键值复制到一个新的 Map 中。
func (m *Map[K, V]) Copy() *Map[K, V] {
	m.Lock()
	defer m.Unlock()
	ret := NewMap[K, V]()
	m.Range(func(key K, value V) bool {
		ret.Put(key, value)
		return true
	})

	return ret
}

// Exist 当指定的 key 存在于 map 中时返回 true。
func (m *Map[K, V]) Exist(key K) bool {
	m.Lock()
	defer m.Unlock()
	_, ok := m.Get(key)
	return ok
}

// ExistValue 当 map 中存在与 value 深度相等的值时返回 true，并返回其对应的 key。
func (m *Map[K, V]) ExistValue(value V) (k K, exist bool) {
	m.Lock()
	defer m.Unlock()
	de := xreflect.NewDeepEquals(value)
	m.Range(func(key K, val V) bool {
		if de.Matches(val) {
			exist = true
			k = key
			return false
		}
		return true
	})
	return
}

// ExistValueWithComparator 当 map 中存在与 value 按比较器 equal 判定相等的值时返回 true，并返回其对应的 key。
func (m *Map[K, V]) ExistValueWithComparator(value V, equal tools.EQL[V]) (k K, exist bool) {
	m.Lock()
	defer m.Unlock()
	m.Range(func(key K, val V) bool {
		if equal(value, val) {
			exist = true
			k = key
			return false
		}
		return true
	})
	return
}

// Remove 从 map 中移除指定的 key，并返回该 key 是否原本存在。
func (m *Map[K, V]) Remove(key K) bool {
	m.Lock()
	defer m.Unlock()
	_, ok := m.mp[key]
	if ok {
		delete(m.mp, key)
	}
	return ok
}

// MinValue 依据比较器 compare 返回 map 中值最小的键值对。
func (m *Map[K, V]) MinValue(compare tools.CMP[V]) (key K, v V) {
	m.Lock()
	defer m.Unlock()
	return selectByCompareValue(m, func(o1, o2 V) int {
		return compare(o1, o2)
	})
}

// MaxValue 依据比较器 compare 返回 map 中值最大的键值对。
func (m *Map[K, V]) MaxValue(compare tools.CMP[V]) (key K, v V) {
	return selectByCompareValue(m, func(o1, o2 V) int {
		return compare(o2, o1)
	})
}

func selectByCompareValue[K comparable, V any](mp *Map[K, V], compare tools.CMP[V]) (key K, v V) {
	var ret V
	i := 0
	mp.Range(func(k K, v V) bool {
		if i == 0 {
			ret = v
			key = k
		} else {
			if compare(ret, v) > 0 {
				ret = v
				key = k
			}
		}
		i++
		return true
	})
	return key, ret
}

// MinKey 依据比较器 compare 返回 map 中键最小的键值对。
func (m *Map[K, V]) MinKey(compare tools.CMP[K]) (key K, v V) {
	m.Lock()
	defer m.Unlock()
	return selectByCompareKey(m, func(o1, o2 K) int {
		return compare(o1, o2)
	})
}

// MaxKey 依据比较器 compare 返回 map 中键最大的键值对。
func (m *Map[K, V]) MaxKey(compare tools.CMP[K]) (key K, v V) {
	m.Lock()
	defer m.Unlock()
	return selectByCompareKey(m, func(o1, o2 K) int {
		return compare(o2, o1)
	})
}

func selectByCompareKey[K comparable, V any](mp *Map[K, V], compare tools.CMP[K]) (key K, value V) {
	var ret K
	i := 0
	mp.Range(func(k K, v V) bool {
		if i == 0 {
			ret = k
			value = v
		} else {
			if compare(ret, k) > 0 {
				ret = k
				value = v
			}
		}
		i++
		return true
	})
	return ret, value
}

// Equals 按比较器 eql 判断两个 map 的所有键值是否完全相同。
func (m *Map[K, V]) Equals(mp *Map[K, V], eql tools.EQL[V]) bool {
	m.Lock()
	defer m.Unlock()
	if m == mp {
		return true
	}

	if m.Size() != mp.Size() {
		return false
	}

	neq := true
	m.Range(func(k K, v V) bool {
		nv, exist := mp.Get(k)
		if !exist {
			neq = false
			return false
		}
		if !eql(v, nv) {
			neq = false
			return false
		}
		return true
	})
	return neq
}
