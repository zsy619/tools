package xinterface

import "haedu.gov.cn/tools/xjson"

// ObjectToJson
/**
 * @description: any转json
 * @param {any} value
 * @return {string error}
 */
func ObjectToJson(value any) (string, error) {
	meta, err := xjson.Marshal(value)
	return string(meta), err
}

// ObjectsToJson
/**
 * @description: any转json
 * @param {[]any} values
 * @return {*}
 */
func ObjectsToJson(values []any) ([]any, error) {
	result := [](any){}

	for _, currValue := range values {
		meta, err := ObjectToJson(currValue)

		if err != nil {
			return nil, err
		}
		result = append(result, meta)
	}

	return result, nil
}
