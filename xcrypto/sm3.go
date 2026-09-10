package xcrypto

import (
	"encoding/hex"
	"fmt"

	"github.com/tjfoc/gmsm/sm3"
)

/**
 * @description: 哈希
 * @param {string} data
 * @return {*}
 */
// Sm3Hash 使用国密 SM3 算法对 data 进行哈希计算。
//
// 参数：
//   - data: 待哈希的字符串。
//   - b: 追加到 data 后参与 Sum 的字节串（通常传空字符串）。
//
// 返回值：SM3 哈希值的十六进制字符串；同时会向标准输出打印一行调试信息。
func Sm3Hash(data string, b string) string {
	hash := sm3.New()
	hash.Write([]byte(data))
	hashed := hash.Sum([]byte(b))
	println("sm3 hash = ", hex.EncodeToString(hashed))
	return fmt.Sprintf("%x", hashed)
}
