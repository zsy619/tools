package xcache

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

// cacheAdapter 全局内存缓存实例，基于 patrickmn/go-cache 实现。
var cacheAdapter *cache.Cache

// init 在包加载时初始化全局缓存适配器。
func init() {
	// 创建一个默认过期时间为5分钟的缓存适配器
	// 每60秒清除一次过期的项目
	cacheAdapter = cache.New(5*time.Minute, 60*time.Second)
}

// SetCahce 将键值对写入缓存，使用指定的过期时间 d。
//
// 参数：
//   - k: 缓存键，字符串类型。
//   - x: 任意类型的缓存值。
//   - d: 过期时间，传 0 表示使用默认过期时间（5分钟）。
//
// 该函数为全局缓存的统一写入入口。
func SetCahce(k string, x interface{}, d time.Duration) {
	cacheAdapter.Set(k, x, d)
}

// GetCache 根据键从全局缓存中读取值。
//
// 参数：
//   - k: 缓存键。
//
// 返回值：
//   - 第一个返回值是缓存中存储的 interface{} 值。
//   - 第二个返回值表示是否命中；未命中或已过期时为 false。
func GetCache(k string) (interface{}, bool) {
	return cacheAdapter.Get(k)
}

// 设置cache 无时间参数
// SetDefaultCahce 将键值对写入缓存，使用默认过期时间（底层 SetDefault）。
func SetDefaultCahce(k string, x interface{}) {
	cacheAdapter.SetDefault(k, x)
}

// 删除 cache
// DeleteCache 删除指定键的缓存项；键不存在时不会产生错误。
func DeleteCache(k string) {
	cacheAdapter.Delete(k)
}

// Add() 加入缓存
// AddCache 仅在键不存在时写入缓存；存在则忽略（错误被忽略）。
//
// 参数：
//   - k: 缓存键。
//   - x: 任意类型的缓存值。
//   - d: 过期时间，0 表示使用默认过期时间。
func AddCache(k string, x interface{}, d time.Duration) {
	_ = cacheAdapter.Add(k, x, d)
}

// IncrementInt() 对已存在的key 值自增n
// IncrementIntCahce 将键对应的整数值自增 n（PKCS 底层使用原子操作）。
//
// 参数：
//   - k: 缓存键，键必须存在且值为 int 类型。
//   - n: 要增加的整数值，可为负。
//
// 返回值：
//   - num: 自增之后的整数值。
//   - err: 当键不存在或存储类型不是 int 时返回错误。
func IncrementIntCahce(k string, n int) (num int, err error) {
	return cacheAdapter.IncrementInt(k, n)
}
