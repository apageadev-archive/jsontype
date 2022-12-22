package jsontype

import (
	"github.com/goccy/go-reflect"
)

func IsString(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.String
}

func IsNumber(value interface{}) bool {
	numberTypes := []string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64"}
	isNumberType := false
	for _, t := range numberTypes {
		if reflect.TypeOf(value).Kind().String() == t {
			isNumberType = true
		}
	}
	return isNumberType
}

func IsBool(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Bool
}

func IsObject(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Map
}

// NOTE: items in an array must be of the same type, otherwise it is a list
func IsArray(value interface{}) bool {
	if reflect.TypeOf(value).Kind() != reflect.Slice {
		return false
	}

	// check if all items in the array are of the same type
	if len(value.([]interface{})) > 0 {
		firstItem := value.([]interface{})[0]
		for _, item := range value.([]interface{}) {
			if reflect.TypeOf(item) != reflect.TypeOf(firstItem) {
				return false
			}
		}
	}
	return true
}

func IsList(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Slice
}

func IsNull(value interface{}) bool {
	return value == nil
}

func IsType(value interface{}, typeToValidate string) bool {
	switch typeToValidate {
	case "string":
		return IsString(value)
	case "number":
		return IsNumber(value)
	case "bool":
		return IsBool(value)
	case "object":
		return IsObject(value)
	case "array":
		return IsArray(value)
	case "list":
		return IsList(value)
	case "null":
		return IsNull(value)
	}
	return false
}
