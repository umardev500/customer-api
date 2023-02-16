package helper

import "strconv"

func ToInt(str string) int64 {
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return int64(n)
}
