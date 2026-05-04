package xjson

import (
	"fmt"
	"testing"
	"time"
)

func TestJsonToMap(t *testing.T) {
	jsonStr := `{"ip": "127.0.0.1", "device": "ABESSF0023", "device": "ABESSF0xx023"}`

	// test json string to map
	m, err := JsonToMap[string, string](jsonStr)
	if err != nil {
		fmt.Printf("Convert json to map failed with error: %+v\n", err)
	}

	fmt.Printf("Converted to map result: %+v\n", m)
}

func FuzzJsonToMap(f *testing.F) {
	f.Add(`{"ip": "127.0.0.1", "device": "ABESSF0023", "device": "ABESSF0xx023"}`)
	// f.Add(`{"ip": "127.0.0.1", "device": "ABESSF0023", "device": "}`)
	// f.Add(`{"ip": "127.0.0.1", "device": "ABESSF0023", "device": }`)
	// f.Add(`{"ip": "127.0.0.1", "device": "ABESSF0023", "device"}`)
	f.Add(`{"ip": "127.0.0.1", "device": "ABESSF0023", "a":"120"}`)
	f.Add(`"ip": "127.0.0.1", "device": "ABESSF0023", "a":"120"`)
	f.Fuzz(func(t *testing.T, jsonStr string) {
		m, err := JsonToMap[string, string](jsonStr)
		if err != nil {
			t.Errorf("%q, %v", m, err)
		}

		fmt.Printf("Converted to map result: %+v\n", m)
	})
}

func TestMapToJson(t *testing.T) {
	m := make(map[string]interface{}, 4)
	m["a"] = "aa"
	m["b"] = 12
	m["t"] = time.Now()
	jsonRes, err := MapToJson(m)
	if err != nil {
		fmt.Printf("Convert json to map failed with error: %+v\n", err)
	}

	fmt.Printf("Convert to json string result: %+v\n", jsonRes)
}
