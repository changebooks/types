package objs

import (
	"fmt"
	"strconv"
)

func ParseBool(data []byte) (bool, error) {
	if len(data) == 0 {
		return false, nil
	}

	if data[0] == 0x00 {
		return false, nil
	}

	if data[0] == 0x01 {
		return true, nil
	}

	return strconv.ParseBool(string(data))
}

// float32 - 3.14 => float64 - 3.140000104904175
func ParseFloat(value interface{}) (float64, error) {
	switch value.(type) {
	case float64:
		return value.(float64), nil
	case float32:
		return float64(value.(float32)), nil
	case int:
		return float64(value.(int)), nil
	case uint:
		return float64(value.(uint)), nil
	case int8:
		return float64(value.(int8)), nil
	case uint8:
		return float64(value.(uint8)), nil
	case int16:
		return float64(value.(int16)), nil
	case uint16:
		return float64(value.(uint16)), nil
	case int32:
		return float64(value.(int32)), nil
	case uint32:
		return float64(value.(uint32)), nil
	case int64:
		return float64(value.(int64)), nil
	case uint64:
		return float64(value.(uint64)), nil
	case string:
		return strconv.ParseFloat(value.(string), 64)
	case []byte:
		return strconv.ParseFloat(string(value.([]byte)), 64)
	default:
		return 0, fmt.Errorf("unsupported type %v", value)
	}
}

func ParseUint(value interface{}) (uint64, error) {
	if value == nil {
		return 0, nil
	}

	switch value.(type) {
	case uint64:
		return value.(uint64), nil
	case int:
		return uint64(value.(int)), nil
	case uint:
		return uint64(value.(uint)), nil
	case int8:
		return uint64(value.(int8)), nil
	case uint8:
		return uint64(value.(uint8)), nil
	case int16:
		return uint64(value.(int16)), nil
	case uint16:
		return uint64(value.(uint16)), nil
	case int32:
		return uint64(value.(int32)), nil
	case uint32:
		return uint64(value.(uint32)), nil
	case int64:
		return uint64(value.(int64)), nil
	case float32:
		return uint64(value.(float32)), nil
	case float64:
		return uint64(value.(float64)), nil
	case string:
		return strconv.ParseUint(value.(string), 10, 64)
	case []byte:
		return strconv.ParseUint(string(value.([]byte)), 10, 64)
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
