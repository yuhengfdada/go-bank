package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var dict string = "abcdefghijklmnopqrstuvwxyz"

func GetRandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
func GetRandomInt64(min, max int) int64 {
	return rand.Int63n(int64(max)-int64(min)+1) + int64(min)
}
func GetRandomString(l int) string {
	var builder strings.Builder
	for ; l > 0; l -= 1 {
		builder.WriteByte(dict[GetRandomInt(0, len(dict)-1)])
	}
	return builder.String()
}
