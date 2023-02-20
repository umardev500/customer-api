package helper

import "reflect"

func IsZero(i any) (res bool) {
	return reflect.ValueOf(i).IsZero()
}
