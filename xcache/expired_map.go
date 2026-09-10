package xcache

import (
	"sync"
	"time"
)

// val 存储 ExpiredMap 中键对应的实际数据与绝对过期时间（Unix 秒）。
type val struct {
	data        interface{}
	expiredTime int64
}

// Data 返回 val 中存储的实际数据。
func (v *val) Data() interface{} {
	return v.data
}

// ExpiredTime 返回 val 中数据的绝对过期时间（Unix 秒）。
func (v *val) ExpiredTime() int64 {
	return v.expiredTime
}

// 过期时间
const (
	expiredTime = 10 * time.Minute
)

// delChannelMap 删除消息通道的容量上限，用于平衡后台 goroutine 的删除速率。
const delChannelCap = 100

// ExpiredMap 一个带过期时间的线程安全 map，底层使用双 map（按 key 索引 + 按过期时间索引）。
//
// 字段说明：
//   - m: 主存储，键 -> val。
//   - timeMap: 辅助索引，过期时间（Unix 秒）-> 在该时刻过期的 key 列表。
//   - lck: 全局互斥锁，保证所有操作的并发安全。
//   - stop: 用于通知后台清理 goroutine 停止的信号通道。
type ExpiredMap struct {
	m       map[interface{}]*val
	timeMap map[int64][]interface{}
	lck     *sync.Mutex
	stop    chan struct{}
}

// NewExpiredMap 创建并返回一个初始化完成的 ExpiredMap，同时启动后台过期清理 goroutine。
//
// 返回值：新创建的 ExpiredMap 指针。
func NewExpiredMap() *ExpiredMap {
	e := ExpiredMap{
		m:       make(map[interface{}]*val),
		lck:     new(sync.Mutex),
		timeMap: make(map[int64][]interface{}),
		stop:    make(chan struct{}),
	}
	go e.run(time.Now().Unix())
	return &e
}

// delMsg 后台删除消息：包含在某个时间点过期的所有 key 和该过期时间戳。
type delMsg struct {
	keys []interface{}
	t    int64
}

// List 直接返回内部主存储 map（key -> val）。注意：返回值是内部 map 的引用，调用方不应直接修改。
func (e *ExpiredMap) List() map[interface{}]*val {
	return e.m
}

// background goroutine 主动删除过期的key
// 数据实际删除时间比应该删除的时间稍晚一些，这个误差会在查询的时候被解决。
func (e *ExpiredMap) run(now int64) {
	t := time.NewTicker(time.Second * 1)
	defer t.Stop()
	delCh := make(chan *delMsg, delChannelCap)
	go func() {
		for v := range delCh {
			e.multiDelete(v.keys, v.t)
		}
	}()
	for {
		select {
		case <-t.C:
			now++ // 这里用now++的形式，直接用time.Now().Unix()可能会导致时间跳过1s，导致key未删除。
			e.lck.Lock()
			if keys, found := e.timeMap[now]; found {
				e.lck.Unlock()
				delCh <- &delMsg{keys: keys, t: now}
			} else {
				e.lck.Unlock()
			}
		case <-e.stop:
			close(delCh)
			return
		}
	}
}

// GetFirst 遍历 map，返回第一个尚未过期的值。
//
// 返回值：
//   - found: 是否找到未过期的元素。
//   - value: 找到的第一个未过期元素的数据（任意类型）。
func (e *ExpiredMap) GetFirst() (found bool, value interface{}) {
	e.lck.Lock()
	defer e.lck.Unlock()
	for _, v := range e.m {
		if v.expiredTime > time.Now().Unix() {
			found = true
			value = v.data
			break
		}
	}
	return
}

// Set 在 map 中写入一个键值对，并设置过期时间。
//
// 参数：
//   - key: 任意类型的键。
//   - value: 任意类型的值。
//   - expireSeconds: 相对当前时间的过期秒数，必须 > 0，否则直接返回 false，不写入。
//
// 返回值：写入是否成功。
func (e *ExpiredMap) Set(key, value interface{}, expireSeconds int64) bool {
	if expireSeconds <= 0 {
		return false
	}
	e.lck.Lock()
	defer e.lck.Unlock()
	expiredTime := time.Now().Unix() + expireSeconds
	e.m[key] = &val{
		data:        value,
		expiredTime: expiredTime,
	}
	e.timeMap[expiredTime] = append(e.timeMap[expiredTime], key) // 过期时间作为key，放在map中
	return true
}

// Get 根据键读取未过期的值。
//
// 参数：
//   - key: 要查找的键。
//
// 返回值：
//   - found: 是否存在且未过期。
//   - value: 命中时返回存储的数据，否则为 nil。
func (e *ExpiredMap) Get(key interface{}) (found bool, value interface{}) {
	e.lck.Lock()
	defer e.lck.Unlock()
	if found = e.checkDeleteKey(key); !found {
		return
	}
	value = e.m[key].data
	return
}

// Delete 删除指定键的缓存项；不会清理 timeMap 中的索引（保留索引以便后台删除时校验）。
func (e *ExpiredMap) Delete(key interface{}) {
	e.lck.Lock()
	delete(e.m, key)
	e.lck.Unlock()
}

// Remove 是 Delete 的别名。
func (e *ExpiredMap) Remove(key interface{}) {
	e.Delete(key)
}

// multiDelete 由后台删除 goroutine 调用，批量删除指定过期时刻 t 的所有 key，并清理 timeMap 索引。
//
// 参数：
//   - keys: 在时刻 t 应当过期的键列表。
//   - t: 过期时刻（Unix 秒）。
func (e *ExpiredMap) multiDelete(keys []interface{}, t int64) {
	e.lck.Lock()
	defer e.lck.Unlock()
	delete(e.timeMap, t)
	for _, key := range keys {
		delete(e.m, key)
	}
}

// Length 返回 map 中元素数量。注意：结果可能不准确，因为存在尚未被后台清理的过期 key。
func (e *ExpiredMap) Length() int { // 结果是不准确的，因为有未删除的key
	e.lck.Lock()
	defer e.lck.Unlock()
	return len(e.m)
}

// Size 是 Length 的别名。
func (e *ExpiredMap) Size() int {
	return e.Length()
}

// 返回key的剩余生存时间 key不存在返回负数
// TTL 返回指定键的剩余生存时间（秒）。
//
// 参数：
//   - key: 要查询的键。
//
// 返回值：剩余秒数；键不存在或已过期时返回 -1。
func (e *ExpiredMap) TTL(key interface{}) int64 {
	e.lck.Lock()
	defer e.lck.Unlock()
	if !e.checkDeleteKey(key) {
		return -1
	}
	return e.m[key].expiredTime - time.Now().Unix()
}

// Clear 清空整个 map 的所有数据，但不会停止后台清理 goroutine。
func (e *ExpiredMap) Clear() {
	e.lck.Lock()
	defer e.lck.Unlock()
	e.m = make(map[interface{}]*val)
	e.timeMap = make(map[int64][]interface{})
}

// Close 通知后台清理 goroutine 退出。调用后 stop 通道将不再可写。
// 注意：当前实现 Close 后再次操作可能会被永久阻塞（todo）。
func (e *ExpiredMap) Close() { // todo 关闭后在使用怎么处理
	e.lck.Lock()
	defer e.lck.Unlock()
	e.stop <- struct{}{}
	// e.m = nil
	// e.timeMap = nil
}

// Stop 是 Close 的别名。
func (e *ExpiredMap) Stop() {
	e.Close()
}

// DoForEach 遍历所有未过期的元素，对每个元素调用 handler(key, val)。
//
// 参数：
//   - handler: 遍历回调函数，参数为 (key, val)。
func (e *ExpiredMap) DoForEach(handler func(interface{}, interface{})) {
	e.lck.Lock()
	defer e.lck.Unlock()
	for k, v := range e.m {
		if !e.checkDeleteKey(k) {
			continue
		}
		handler(k, v)
	}
}

// DoForEachWithBreak 与 DoForEach 类似，但 handler 返回 true 时会中断遍历。
//
// 参数：
//   - handler: 遍历回调函数；返回 true 时立即停止遍历。
func (e *ExpiredMap) DoForEachWithBreak(handler func(interface{}, interface{}) bool) {
	e.lck.Lock()
	defer e.lck.Unlock()
	for k, v := range e.m {
		if !e.checkDeleteKey(k) {
			continue
		}
		if handler(k, v) {
			break
		}
	}
}

// checkDeleteKey 在读取前检查 key 是否过期；如果已过期则顺带删除并返回 false。
//
// 返回值：true 表示 key 存在且未过期；false 表示不存在或已过期。
func (e *ExpiredMap) checkDeleteKey(key interface{}) bool {
	if val, found := e.m[key]; found {
		if val.expiredTime <= time.Now().Unix() {
			delete(e.m, key)
			// delete(e.timeMap, val.expiredTime)
			return false
		}
		return true
	}
	return false
}
