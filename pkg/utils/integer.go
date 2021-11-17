package utils

import (
	"encoding/json"
	"strconv"
)

func Interface2Int64(inter interface{}) int64 {
	var temp int64
	switch inter.(type) {
	case string:
		temp, _ = strconv.ParseInt(inter.(string), 10, 64)
		break
	case int64:
		temp = inter.(int64)
		break
	case int:
		s := strconv.Itoa(inter.(int))
		temp, _ = strconv.ParseInt(s, 10, 64)
		break
	case int32:
		temp = int64(inter.(int32))
		break
	case float64:
		tempStr := strconv.FormatFloat(inter.(float64), 'f', -1, 64)
		temp, _ = strconv.ParseInt(tempStr, 10, 64)
		break
	case json.Number:
		temp, _ = inter.(json.Number).Int64()
	}
	return temp
}
