package xjson

import (
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
