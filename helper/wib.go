package helper

import "time"

func WIB() (loc *time.Location, err error) {
	loc, err = time.LoadLocation("Asia/Jakarta")
	return
}
