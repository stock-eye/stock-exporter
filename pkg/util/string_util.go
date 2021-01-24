package util

import "strconv"

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func Split(r rune) bool {
	return r == ',' || r == '，' || r == ' ' || r == '+'
}
