package xbyte

import (
	"bytes"
	"reflect"
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

func TestHexStringToByte32(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    [32]byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexStringToByte32(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexStringToByte32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexStringToByte32() = %v, want %v", got, tt.want)
			}
		})
	}
}
