package xjson

import (
	"fmt"
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	data := struct {
		A time.Time `json:"a,omitempty"`
	}{
		A: time.Time{},
	}
	out, err := Marshal(data)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(string(out))
	}
}

func TestPhp(t *testing.T) {
	type ProductInfo struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}

	str := `{"name":"AppleWatchS8","price":"3199"}`
	data := ProductInfo{}
	if err := Unmarshal([]byte(str), &data); err != nil {
		fmt.Println("error: " + err.Error())
	} else {
		fmt.Println(data)
	}
}
