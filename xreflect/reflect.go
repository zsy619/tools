package xreflect

import "reflect"

func GetMapResult(collections map[string]interface{}, key string) bool {
	return collections[key].(bool)
}

func TypeOf(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// IsEmpty checks to see if a field has a set value
//
// Goes beyond usual reflect.IsZero check to handle numbers, strings, and slices
// For structs, iterates over all accessible properties and returns true only if all nested fields
// are also empty.
func IsEmpty(val reflect.Value) bool {
	valType := val.Kind()
	switch valType {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.String:
		return val.String() == ""
	case reflect.Interface, reflect.Slice, reflect.Ptr, reflect.Map, reflect.Chan, reflect.Func:
		// Check for empty slices and props
		if val.IsNil() {
			return true
		} else if valType == reflect.Slice || valType == reflect.Map {
			return val.Len() == 0
		}
	case reflect.Struct:
		fieldCount := val.NumField()
		for i := 0; i < fieldCount; i++ {
			field := val.Field(i)
			if field.IsValid() && !IsEmpty(field) {
				return false
			}
		}
		return true
	default:
		return false
	}
	return false
}
