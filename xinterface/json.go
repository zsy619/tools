package xinterface

import "haedu.gov.cn/tools/xjson"

func ObjectToJson(value interface{}) (string, error) {
	meta, err := xjson.Marshal(value)
	return string(meta), err
}

func ObjectsToJson(values []interface{}) ([]interface{}, error) {
	result := [](interface{}){}

	for _, currValue := range values {
		meta, err := ObjectToJson(currValue)

		if err != nil {
			return nil, err
		} else {
			result = append(result, meta)
		}
	}

	return result, nil
}
