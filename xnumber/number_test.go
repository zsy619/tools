package xnumber

import (
	"math/big"
	"testing"
)

func TestBigIntToFloat64(t *testing.T) {
	floatVal := float64(3)
	bigInt := big.NewInt(3)
	floatNum := BigIntToFloat64(bigInt)
	if floatVal != floatNum {
		t.Errorf("Bigint to Float64 conversion failed")
	}
}

func TestFloat64ToBigInt(t *testing.T) {
	bigIntVal := big.NewInt(34)
	float := float64(34)
	bigInt := Float64ToBigInt(float)
	if bigInt == bigIntVal {
		t.Errorf("Float64 to Bigint conversion failed")
	}
}

func TestIntOrEmptyInt(t *testing.T) {
	res := IntOrEmptyInt(nil)
	if res != 0 {
		t.Errorf("Should have returned an empty int")
	}

	testInt := 100
	res = IntOrEmptyInt(&testInt)
	if res == 0 {
		t.Errorf("Should not have returned an empty int")
	}
	if res != testInt {
		t.Errorf("Should have returned the test string")
	}

	testInt = 0
	res = IntOrEmptyInt(&testInt)
	if res != 0 {
		t.Errorf("Should have returned an empty int")
	}
}

func TestIntToPtr(t *testing.T) {
	testInt := 100
	strPtr := IntToPtr(testInt)
	if *strPtr != testInt {
		t.Errorf("Should have returned the test int")
	}
}
