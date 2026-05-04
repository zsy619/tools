package xcache

import (
	"fmt"
	"strconv"

	"github.com/peterbourgon/diskv"
)

var diskKeyValue *diskv.Diskv

func init() {
	fmt.Println("init diskv.Diskv")
	// Simplest transform function: put all the data files into the base dir.
	flatTransform := func(s string) []string { return []string{} }
	// Initialize a new diskv store, rooted at "my-data-dir", with a 1MB cache.
	diskKeyValue = diskv.New(diskv.Options{
		BasePath:     "./data",
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})
}

func SetDisCahce(k string, x []byte) {
	_ = diskKeyValue.Write(k, x)
}

func GetDisCahce(k string) (string, error) {
	value, err := diskKeyValue.Read(k)
	if err != nil {
		return "", err
	}
	return string(value), nil
}

func GetDisCahceOk(k string) (string, bool) {
	value, err := diskKeyValue.Read(k)
	if err != nil {
		return "", false
	}
	return string(value), true
}

func DisCahceExists(k string) bool {
	_, err := diskKeyValue.Read(k)
	if err != nil {
		return false
	}
	return true
}

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

func SetDiskvInt(key string, value int) {
	x := strconv.Itoa(value)
	_ = diskKeyValue.Write(key, []byte(x))
}

func SetDiskvString(key string, value string) {
	_ = diskKeyValue.Write(key, []byte(value))
}

func GetDiskvString(key string) string {
	value, err := GetDisCahce(key)
	if err != nil {
		return ""
	}
	return value
}
