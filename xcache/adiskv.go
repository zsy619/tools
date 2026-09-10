package xcache

import (
	"fmt"
	"strconv"

	"github.com/peterbourgon/diskv"
)

// diskKeyValue 基于 diskv 的全局磁盘缓存实例，根目录为当前目录下的 data 文件夹，最大内存缓存 1MB。
var diskKeyValue *diskv.Diskv

// init 初始化 diskv 全局缓存实例，初始化时会打印一行日志。
func init() {
	fmt.Println("init diskv.Diskv")
	// Simplest transform function: put all the data files into the base dir.
	// 使用最简单的 transform 函数：所有数据文件都直接放到基础目录下，不做分层。
	flatTransform := func(s string) []string { return []string{} }
	// Initialize a new diskv store, rooted at "my-data-dir", with a 1MB cache.
	// 创建一个新的 diskv 存储，根目录为当前目录下的 "data"，最大内存缓存为 1MB。
	diskKeyValue = diskv.New(diskv.Options{
		BasePath:     "./data",
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})
}

// SetDisCahce 将字节切片写入磁盘缓存；写入错误被忽略。
//
// 参数：
//   - k: 缓存键。
//   - x: 要写入磁盘的字节内容。
func SetDisCahce(k string, x []byte) {
	_ = diskKeyValue.Write(k, x)
}

// GetDisCahce 从磁盘缓存读取字符串值。
//
// 参数：
//   - k: 缓存键。
//
// 返回值：
//   - 第一个返回值为读到的字符串。
//   - 第二个返回值在键不存在或读取失败时返回非 nil 错误。
func GetDisCahce(k string) (string, error) {
	value, err := diskKeyValue.Read(k)
	if err != nil {
		return "", err
	}
	return string(value), nil
}

// GetDisCahceOk 从磁盘缓存读取字符串值，但不返回错误；命中失败返回 ("", false)。
//
// 参数：
//   - k: 缓存键。
//
// 返回值：
//   - 第一个返回值为读到的字符串（命中失败时为空字符串）。
//   - 第二个返回值表示是否成功读取到内容。
func GetDisCahceOk(k string) (string, bool) {
	value, err := diskKeyValue.Read(k)
	if err != nil {
		return "", false
	}
	return string(value), true
}

// DisCahceExists 判断指定键在磁盘缓存中是否存在。
//
// 参数：
//   - k: 缓存键。
//
// 返回值：键存在时返回 true，否则返回 false。
func DisCahceExists(k string) bool {
	_, err := diskKeyValue.Read(k)
	if err != nil {
		return false
	}
	return true
}

// GetDiskvInt 从磁盘缓存读取一个整数值。
//
// 参数：
//   - key: 缓存键。
//
// 返回值：读取并解析成功的整数；若键不存在、读取失败或值不是合法整数则返回 0。
func GetDiskvInt(key string) int {
	value, err := GetDisCahce(key)
	if err != nil {
		return 0
	}
	md, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return md
}

// SetDiskvInt 将整数值以字符串形式写入磁盘缓存；写入错误被忽略。
//
// 参数：
//   - key: 缓存键。
//   - value: 要写入的整数值。
func SetDiskvInt(key string, value int) {
	x := strconv.Itoa(value)
	_ = diskKeyValue.Write(key, []byte(x))
}

// SetDiskvString 将字符串写入磁盘缓存；写入错误被忽略。
//
// 参数：
//   - key: 缓存键。
//   - value: 要写入的字符串内容。
func SetDiskvString(key string, value string) {
	_ = diskKeyValue.Write(key, []byte(value))
}

// GetDiskvString 从磁盘缓存读取一个字符串值。
//
// 参数：
//   - key: 缓存键。
//
// 返回值：读取到的字符串；若读取失败则返回空字符串。
func GetDiskvString(key string) string {
	value, err := GetDisCahce(key)
	if err != nil {
		return ""
	}
	return value
}
