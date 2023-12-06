package xcache

type CacheItemModel[T any] struct {
	Key     string `json:"key"`     // Key 缓存键
	Len     int    `json:"len"`     // Len 缓存长度
	Note    string `json:"note"`    // Note 缓存说明
	Expired int64  `json:"expired"` // Expired 过期时间

	cache *ExpiredMap `json:"-"` // cache 缓存对象
}

// NewCacheItemModel 创建缓存项
func NewCacheItemModel[T any](key string, note string, expired int64) *CacheItemModel[T] {
	return &CacheItemModel[T]{
		Key:     key,
		Note:    note,
		Expired: expired,
		cache:   NewExpiredMap(),
	}
}

// Set 设置缓存
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
func (cache *CacheItemModel[T]) GetFirst() (ok bool, v T) {
	ok, t := cache.cache.GetFirst()
	if ok {
		return ok, t.(T)
	}
	return
}

// GetList 获取列表
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
func (cache *CacheItemModel[T]) Reset() {
	cache.cache.Clear()
	cache.cache.Close()
	cache.cache = NewExpiredMap()
}
