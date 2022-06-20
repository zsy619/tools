package xinterface

import (
	"reflect"

	"haedu.gov.cn/tools/xreflect"
)

// Similar to "extend" in JS, only updates fields that are specified and not empty in newData
//
// Both newData and mainObj must be pointers to struct objects
func Update(mainObj interface{}, newData interface{}) bool {
	newDataVal, mainObjVal := reflect.ValueOf(newData).Elem(), reflect.ValueOf(mainObj).Elem()
	fieldCount := newDataVal.NumField()
	changed := false
	for i := 0; i < fieldCount; i++ {
		newField := newDataVal.Field(i)
		// They passed in a value for this field, update our DB user
		if newField.IsValid() && !xreflect.IsEmpty(newField) {
			dbField := mainObjVal.Field(i)
			dbField.Set(newField)
			changed = true
		}
	}
	return changed
}
