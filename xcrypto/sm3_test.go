package xcrypto

import "testing"

func TestSm3Hash(t *testing.T) {
	t.Log(Sm3Hash("Hello, World!", ""))
	t.Log(Sm3Hash("Hello, World!", "xxx"))
	t.Log(Sm3Hash("sm3是我国国产的哈希算法", ""))
}
