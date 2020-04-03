package ints

import (
	"fmt"
	"strconv"
)

func Parse(value interface{}) (int64, error) {
	if value == nil {
		return 0, nil
	}

	switch value.(type) {
	case int:
		return int64(value.(int)), nil
	case uint:
		return int64(value.(uint)), nil
	case int8:
		return int64(value.(int8)), nil
	case uint8:
		return int64(value.(uint8)), nil
	case int16:
		return int64(value.(int16)), nil
	case uint16:
		return int64(value.(uint16)), nil
	case int32:
		return int64(value.(int32)), nil
	case uint32:
		return int64(value.(uint32)), nil
	case int64:
		return value.(int64), nil
	case uint64:
		return int64(value.(uint64)), nil
	case float32:
		return int64(value.(float32)), nil
	case float64:
		return int64(value.(float64)), nil
	case string:
		return strconv.ParseInt(value.(string), 10, 64)
	case []byte:
		return strconv.ParseInt(string(value.([]byte)), 10, 64)
	case bool:
		if value.(bool) {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, fmt.Errorf("unsupported type %v", value)
	}
}
