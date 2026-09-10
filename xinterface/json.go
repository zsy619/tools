package xinterface

import "github.com/zsy619/tools/xjson"

// ObjectToJson 将任意值序列化为 JSON 字符串。
//
// 内部直接复用 xjson.Marshal，行为与 encoding/json 完全一致，但支持
// json-iterator 等更高性能的 JSON 实现。
//
// 参数 value 可以是结构体指针、map、切片等任意可被 JSON 序列化的值；
// 序列化失败时返回错误与空字符串。
func ObjectToJson(value any) (string, error) {
	meta, err := xjson.Marshal(value)
	return string(meta), err
}

// ObjectsToJson 批量将一组对象转换为 JSON 字符串。
//
// 返回值为字符串切片（每个元素对应 values 中对应下标的对象），与
// 原 values 同序。一旦遇到任何元素序列化失败，立即中断并返回错误。
func ObjectsToJson(values []any) ([]any, error) {
	result := make([]any, 0, len(values))
	for _, currValue := range values {
		meta, err := ObjectToJson(currValue)
		if err != nil {
			return nil, err
		}
		result = append(result, meta)
	}
	return result, nil
}
