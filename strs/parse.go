package strs

import (
	"fmt"
	"reflect"
	"strconv"
)

func Parse(value interface{}) string {
	if value == nil {
		return ""
	}

	switch r := value.(type) {
	case string:
		return r
	case []byte:
		return string(r)
	}

	r := reflect.ValueOf(value)
	switch r.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(r.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(r.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(r.Float(), 'g', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(r.Float(), 'g', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(r.Bool())
	}

	return fmt.Sprintf("%v", value)
}
