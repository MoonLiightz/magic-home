package util

import "strconv"

// ParseStringToUint8 ...
func ParseStringToUint8(v string) uint8 {
	res, _ := strconv.ParseUint(v, 10, 8)
	return uint8(res)
}
