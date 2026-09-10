package xcache

// CacheItemModel 泛型缓存项模型：包装一个带过期时间的内部 ExpiredMap，并对外暴露类型化的存取接口。
//
// 字段说明：
//   - Key: 缓存键（描述用，不参与存储定位）。
//   - Len: 缓存中元素数量。
//   - Note: 缓存说明/备注。
//   - Expired: 默认过期秒数，调用 Set 不传过期参数时使用。
//   - cache: 底层 ExpiredMap，不参与 JSON 序列化。
type CacheItemModel[T any] struct {
	Key     string `json:"key"`     // Key 缓存键
	Len     int    `json:"len"`     // Len 缓存长度
	Note    string `json:"note"`    // Note 缓存说明
	Expired int64  `json:"expired"` // Expired 过期时间

	cache *ExpiredMap `json:"-"` // cache 缓存对象
}

// NewCacheItemModel 创建缓存项
//
// 参数：
//   - key: 缓存标识键（描述用）。
//   - note: 缓存说明/备注。
//   - expired: 默认过期秒数；后续调用 Set 不传 expired 时使用该值。
//
// 返回值：初始化完成的 CacheItemModel 指针。
func NewCacheItemModel[T any](key string, note string, expired int64) *CacheItemModel[T] {
	return &CacheItemModel[T]{
		Key:     key,
		Note:    note,
		Expired: expired,
		cache:   NewExpiredMap(),
	}
}

// Set 设置缓存
//
// 参数：
//   - cacheKey: 实际存储使用的 key。
//   - data: 要缓存的数据，类型为泛型 T。
//   - expired: 可变参数，第一个元素作为本次过期秒数；省略时使用模型默认 Expired。
//
// 返回值：底层 ExpiredMap.Set 的返回值；expired <= 0 时返回 false。
func (cache *CacheItemModel[T]) Set(cacheKey interface{}, data T, expired ...int64) bool {
	var _expired int64
	if len(expired) > 0 {
		_expired = expired[0]
	} else {
		_expired = cache.Expired
	}
	return cache.cache.Set(cacheKey, data, _expired)
}

// Get 获取缓存
//
// 参数：
//   - cacheKey: 缓存键。
//
// 返回值：
//   - ok: 是否成功取出且类型匹配。
//   - v: 取出的值（类型不匹配时为零值）。
//
// 注意：底层存储类型若不是 T，则会返回 (false, 零值)。
func (cache *CacheItemModel[T]) Get(cacheKey interface{}) (ok bool, v T) {
	find, value := cache.cache.Get(cacheKey)
	if !find {
		return find, v
	}
	if v, ok := value.(T); ok {
		return ok, v
	}
	return false, v
}

// GetFirst 获取第一个缓存
//
// 返回值：
//   - ok: 是否找到未过期的元素。
//   - v: 第一个未过期元素的数据（命中失败时为零值）。
func (cache *CacheItemModel[T]) GetFirst() (ok bool, v T) {
	ok, t := cache.cache.GetFirst()
	if ok {
		return ok, t.(T)
	}
	return
}

// GetList 获取列表
//
// 返回值：
//   - ok: 始终返回 true（保留字段）。
//   - v: 内部存储的所有数据组成的切片；若内部为空返回空切片。
func (cache *CacheItemModel[T]) GetList() (ok bool, v []T) {
	list := cache.cache.List()
	out := []T{}
	if len(list) > 0 {
		for _, item := range list {
			out = append(out, item.Data().(T))
		}
	}
	return true, out
}

// GetString 获取字符串
//
// 参数：
//   - cacheKey: 缓存键。
//
// 返回值：
//   - bool: 命中且存储类型为 string 时为 true。
//   - string: 取出的字符串；类型不匹配或未命中时返回 ""。
func (cache *CacheItemModel[T]) GetString(cacheKey interface{}) (bool, string) {
	_, value := cache.cache.Get(cacheKey)
	if str, ok := value.(string); ok {
		return ok, str
	}
	return false, ""
}

// Remove 删除缓存
func (cache *CacheItemModel[T]) Remove(cacheKey interface{}) {
	cache.cache.Remove(cacheKey)
}

// Length 获取缓存长度
func (cache *CacheItemModel[T]) Length() int {
	return cache.cache.Length()
}

// Reset 重置缓存
// 清空当前缓存、关闭底层 ExpiredMap 并重新创建一个新的 ExpiredMap。
func (cache *CacheItemModel[T]) Reset() {
	cache.cache.Clear()
	cache.cache.Close()
	cache.cache = NewExpiredMap()
}
