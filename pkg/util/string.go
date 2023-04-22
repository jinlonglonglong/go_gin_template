package util

import "encoding/hex"

func GetStringByLimit(foo string, limit int) string {
	if len(foo) > limit {
		return foo[:limit]
	}
	return foo
}

func DecodeHexString(raw string) string {
	str, _ := hex.DecodeString(raw)
	return string(str)
}
