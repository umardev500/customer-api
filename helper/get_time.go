package helper

import (
	"customer-api/variable"
	"time"
)

func GetTime(timeType *string) (res int64) {
	t := time.Now().UTC()

	switch timeType {
	case &variable.UnixNano:
		res = t.UnixNano()
	case &variable.UnixMili:
		res = t.UnixMilli()
	case &variable.UnixMicro:
		res = t.UnixMicro()
	default:
		res = t.Unix()
	}

	return
}
