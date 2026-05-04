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
func Sm3Hash(data string, b string) string {
	hash := sm3.New()
	hash.Write([]byte(data))
	hashed := hash.Sum([]byte(b))
	println("sm3 hash = ", hex.EncodeToString(hashed))
	return fmt.Sprintf("%x", hashed)
}
