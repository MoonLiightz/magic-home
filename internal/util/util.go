package util

import "strconv"

// ParseStringToUint8 ...
func ParseStringToUint8(v string) uint8 {
	res, _ := strconv.ParseUint(v, 10, 8)
	return uint8(res)
}

// Min compares the two numbers and returns the smaller number
func Min(x uint8, y uint8) uint8 {
	if x < y {
		return x
	}
	return y
}

// Max compares the two numbers and returns the larger number
func Max(x uint8, y uint8) uint8 {
	if x > y {
		return x
	}
	return y
}

// Clamp "clamps" a value between a pair of boundary values
func Clamp(value uint8, minValue uint8, maxValue uint8) uint8 {
	return Min(maxValue, Max(minValue, value))
}
