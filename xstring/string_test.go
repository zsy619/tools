// Package time_test contains tests for the string strings

package xstring

import (
	"reflect"
	"testing"
	"unicode/utf8"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// addressesString    = []string{"0x77e5aaBddb760FBa989A1C4B2CDd4aA8Fa3d311d", "0xDFe273082089bB7f70Ee36Eebcde64832FE97E55"}
	addressesCommon    = []common.Address{common.HexToAddress("0x77e5aaBddb760FBa989A1C4B2CDd4aA8Fa3d311d"), common.HexToAddress("0xDFe273082089bB7f70Ee36Eebcde64832FE97E55")}
	addressesOneString = "0x77e5aaBddb760FBa989A1C4B2CDd4aA8Fa3d311d,0xDFe273082089bB7f70Ee36Eebcde64832FE97E55"
)

func TestJoinInts(t *testing.T) {
	// test empty slice
	is := []int64{}
	s := JoinInts(is)
	if s != "" {
		t.Errorf("input:%v,output:%s,result is incorrect", is, s)
	} else {
		t.Logf("input:%v,output:%s", is, s)
	}
	// test len(slice)==1
	is = []int64{1}
	s = JoinInts(is)
	if s != "1" {
		t.Errorf("input:%v,output:%s,result is incorrect", is, s)
	} else {
		t.Logf("input:%v,output:%s", is, s)
	}
	// test len(slice)>1
	is = []int64{1, 2, 3}
	s = JoinInts(is)
	if s != "1,2,3" {
		t.Errorf("input:%v,output:%s,result is incorrect", is, s)
	} else {
		t.Logf("input:%v,output:%s", is, s)
	}
}

func TestSplitInts(t *testing.T) {
	// test empty slice
	s := ""
	is, err := SplitInts(s)
	if err != nil || len(is) != 0 {
		t.Error(err)
	}
	// test split int64
	s = "1,2,3"
	is, err = SplitInts(s)
	if err != nil || len(is) != 3 {
		t.Error(err)
	}
}

func BenchmarkJoinInts(b *testing.B) {
	is := make([]int64, 10000, 10000)
	for i := int64(0); i < 10000; i++ {
		is[i] = i
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			JoinInts(is)
		}
	})
}

func TestStrOrEmptyStr(t *testing.T) {
	res := StrOrEmptyStr(nil)
	if res != "" {
		t.Errorf("Should have returned an empty string")
	}

	testStr := "thisisatest"
	res = StrOrEmptyStr(&testStr)
	if res == "" {
		t.Errorf("Should not have returned an empty string")
	}
	if res != testStr {
		t.Errorf("Should have returned the test string")
	}

	testStr = ""
	res = StrOrEmptyStr(&testStr)
	if res != "" {
		t.Errorf("Should have returned an empty string")
	}
}

func TestStrToPtr(t *testing.T) {
	testStr := "thisisatest"
	strPtr := StrToPtr(testStr)
	if *strPtr != testStr {
		t.Errorf("Should have returned the test string")
	}
}

func TestIsValidEthAPIURL(t *testing.T) {
	if IsValidEthAPIURL("thisisnotavalidurl") {
		t.Error("Should have failed on an invalid eth API url")
	}
	if IsValidEthAPIURL("http//thisisnotavalidurl.com") {
		t.Error("Should have failed on an invalid eth API url")
	}
	if IsValidEthAPIURL("http/thisisnotavalidurl.com") {
		t.Error("Should have failed on an invalid eth API url")
	}
	if !IsValidEthAPIURL("http://thisisvalid.co") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !IsValidEthAPIURL("http://thisisvalid.com") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !IsValidEthAPIURL("https://thisisvalid.com") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !IsValidEthAPIURL("https://thisisvalid.longtld") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !IsValidEthAPIURL("ws://thisisvalid.ether/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !IsValidEthAPIURL("wss://thisisvalid.com/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !IsValidEthAPIURL("wss://localhost/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !IsValidEthAPIURL("wss://localhost:8545/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !IsValidEthAPIURL("wss://127.0.0.1/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
	if !IsValidEthAPIURL("wss://127.0.0.1:8545/ws") {
		t.Error("Should have not failed on an valid eth API url")
	}
}

func TestIsValidContractAddress(t *testing.T) {
	if IsValidContractAddress("") {
		t.Error("Should have failed on an empty contract address")
	}
	if IsValidContractAddress("thisisnotavalidaddress") {
		t.Error("Should have failed on an invalid contract address")
	}
	if IsValidContractAddress("0xdfe273082089bb7f70ee36eebcde64832fe97e55f") {
		t.Error("Should have failed on an invalid contract address")
	}
	if !IsValidContractAddress("0xdfe273082089bb7f70ee36eebcde64832fe97e55") {
		t.Error("Should have not have failed on an valid contract address")
	}
}

func TestRandomHex(t *testing.T) {
	s, err := RandomHexStr(32)
	if err != nil {
		t.Errorf("Should not have failed on call to random hex str: err: %v", err)
	}
	if len(s) != 64 {
		t.Errorf("Should have been a 64 char hex string: %v", len(s))
	}

	s, err = RandomHexStr(10)
	if err != nil {
		t.Errorf("Should not have failed on call to random hex str: err: %v", err)
	}
	if len(s) != 20 {
		t.Errorf("Should have been a 20 char hex string: %v", len(s))
	}

	s, err = RandomHexStr(0)
	if err != nil {
		t.Errorf("Should not have failed on call to random hex str: err: %v", err)
	}
	if len(s) != 0 {
		t.Errorf("Should have been a 0 char hex string: %v", len(s))
	}
}

func TestListCommonAddressesToString(t *testing.T) {
	stringConverted := ListCommonAddressesToString(addressesCommon)
	if stringConverted != addressesOneString {
		t.Errorf("string is not what it should be, %v", stringConverted)
	}
}

func TestStringToCommonAddressesList(t *testing.T) {
	commonAddressConverted := StringToCommonAddressesList(addressesOneString)
	if !reflect.DeepEqual(commonAddressConverted, addressesCommon) {
		t.Errorf("common.Address slice is not what it should be, %v", commonAddressConverted)
	}
}

func TestSubString(t *testing.T) {
	type args struct {
		source string
		start  int
		end    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubString(tt.args.source, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("SubString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToUint8(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint8(tt.args.input); got != tt.want {
				t.Errorf("StrToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomHexStr(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandomHexStr(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomHexStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RandomHexStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListCommonAddressToListString(t *testing.T) {
	type args struct {
		addresses []common.Address
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListCommonAddressToListString(tt.args.addresses); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListCommonAddressToListString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListStringToListCommonAddress(t *testing.T) {
	type args struct {
		addresses []string
	}
	tests := []struct {
		name string
		args args
		want []common.Address
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListStringToListCommonAddress(tt.args.addresses); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListStringToListCommonAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListIntToListString(t *testing.T) {
	type args struct {
		listInt []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListIntToListString(tt.args.listInt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListIntToListString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithout(t *testing.T) {
	type args struct {
		arr    []string
		remove string
	}
	tests := []struct {
		name        string
		args        args
		wantChopped []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotChopped := Without(tt.args.arr, tt.args.remove); !reflect.DeepEqual(gotChopped, tt.wantChopped) {
				t.Errorf("Without() = %v, want %v", gotChopped, tt.wantChopped)
			}
		})
	}
}

func TestGenerateSlug(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name     string
		args     args
		wantSlug string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSlug := GenerateSlug(tt.args.str); gotSlug != tt.wantSlug {
				t.Errorf("GenerateSlug() = %v, want %v", gotSlug, tt.wantSlug)
			}
		})
	}
}

func TestInChain(t *testing.T) {
	type args struct {
		needle   string
		haystack []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InChain(tt.args.needle, tt.args.haystack); got != tt.want {
				t.Errorf("InChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimHtml(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimHtml(tt.args.src); got != tt.want {
				t.Errorf("TrimHtml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCutstrHtml(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CutstrHtml(tt.args.src); got != tt.want {
				t.Errorf("CutstrHtml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrVal(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrVal(tt.args.value); got != tt.want {
				t.Errorf("StrVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev, err := Reverse(tc.in)
		if rev != tc.want || err != nil {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345", "中国"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, _ := Reverse(orig)
		doubleRev, _ := Reverse(rev)
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
