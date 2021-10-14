// Author: yann
// Date: 2021/10/13
// Desc: field_utils

package field_utils

import "strconv"

//是否大于0
func IsGtZero(number interface{}) bool {
	switch number.(type) {
	case int64:
		return number.(int64) > 0
	case uint64:
		return number.(uint64) > 0
	case int:
		return number.(int) > 0
	case uint:
		return number.(uint) > 0
	case int32:
		return number.(int32) > 0
	case uint32:
		return number.(uint32) > 0
	case int8:
		return number.(int8) > 0
	case uint8:
		return number.(uint8) > 0
	case float32:
		return number.(float32) > 0
	case float64:
		return number.(float64) > 0
	case string:
		num, err := strconv.ParseFloat(number.(string), 64)
		if err != nil {
			return false
		}
		return num > 0
	default:
		return false
	}
}

//是否小于0
func IsLtZero(number interface{}) bool {
	switch number.(type) {
	case int64:
		return number.(int64) < 0
	case uint64:
		return number.(uint64) < 0
	case int:
		return number.(int) < 0
	case uint:
		return number.(uint) < 0
	case int32:
		return number.(int32) < 0
	case uint32:
		return number.(uint32) < 0
	case int8:
		return number.(int8) < 0
	case uint8:
		return number.(uint8) < 0
	case float32:
		return number.(float32) < 0
	case float64:
		return number.(float64) < 0
	case string:
		num, err := strconv.ParseFloat(number.(string), 64)
		if err != nil {
			return false
		}
		return num < 0
	default:
		return false
	}
}
