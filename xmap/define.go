package xmap

import "errors"

type H map[string]interface{}

func (h H) Get(key string) interface{} {
	return h[key]
}

func (h H) GetString(key string) string {
	return h[key].(string)
}

func (h H) GetInt(key string) (int, error) {
	out := h[key]
	outInt, ok := out.(int)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not int")
}

func (h H) GetInt32(key string) (int32, error) {
	out := h[key]
	outInt, ok := out.(int32)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not int32")
}

func (h H) GetInt16(key string) (int16, error) {
	out := h[key]
	outInt, ok := out.(int16)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not int16")
}

func (h H) GetInt8(key string) (int8, error) {
	out := h[key]
	outInt, ok := out.(int8)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not int8")
}

func (h H) GetUint(key string) (uint, error) {
	out := h[key]
	outInt, ok := out.(uint)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not uint")
}

func (h H) GetUint32(key string) (uint32, error) {
	out := h[key]
	outInt, ok := out.(uint32)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not uint32")
}

func (h H) GetUint16(key string) (uint16, error) {
	out := h[key]
	outInt, ok := out.(uint16)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not uint16")
}

func (h H) GetUint8(key string) (uint8, error) {
	out := h[key]
	outInt, ok := out.(uint8)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not uint8")
}

func (h H) GetInt64(key string) (int64, error) {
	out := h[key]
	outInt, ok := out.(int64)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not int64")
}

func (h H) GetFloat64(key string) (float64, error) {
	out := h[key]
	outInt, ok := out.(float64)
	if ok {
		return outInt, nil
	}
	return 0, errors.New("not float64")
}

func (h H) GetBool(key string) (bool, error) {
	out := h[key]
	outInt, ok := out.(bool)
	if ok {
		return outInt, nil
	}
	return false, errors.New("not bool")
}

func (h H) GetMap(key string) (H, error) {
	out := h[key]
	outInt, ok := out.(H)
	if ok {
		return outInt, nil
	}
	return nil, errors.New("not H")
}

func (h H) GetSlice(key string) ([]interface{}, error) {
	out := h[key]
	outInt, ok := out.([]interface{})
	if ok {
		return outInt, nil
	}
	return nil, errors.New("not []interface{}")
}

func (h H) Set(key string, value interface{}) {
	h[key] = value
}

func (h H) Keys() []string {
	keys := make([]string, 0, len(h))
	for key := range h {
		keys = append(keys, key)
	}

	return keys
}

func (h H) Values() []interface{} {
	values := make([]interface{}, 0, len(h))
	for _, value := range h {
		values = append(values, value)
	}

	return values
}

func (h H) Len() int {
	return len(h)
}

func (h H) DeepCopy() H {
	newMap := make(H)
	for k, v := range h {
		newMap[k] = DeepCopy(v)
	}

	return newMap
}
