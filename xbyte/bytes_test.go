package xbyte

import (
	"bytes"
	"testing"

	"haedu.gov.cn/tools/xstring"
)

// TestByte32Utils tests encoding, then decoding from byte[32] to string
func TestByte32Utils(t *testing.T) {
	hexStr, _ := xstring.RandomHexStr(32)
	testBytes := []byte(hexStr)
	testFixed := [32]byte{}
	copy(testFixed[:], testBytes)

	result := Byte32ToHexString(testFixed)
	if result == "" {
		t.Error("Empty result for string")
	}

	bys, err := HexStringToByte32(result)
	if err != nil {
		t.Errorf("Should not have returned an error: err: %v", err)
	}

	if !bytes.Equal(testFixed[:], bys[:]) {
		t.Error("[32]bytes are not equal")
	}
}
