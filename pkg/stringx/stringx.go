package stringx

import (
	"math/rand"
	"strconv"
	"time"
	"unicode/utf8"
)

func Uint64ToString(i64 uint64) string {
	return strconv.FormatUint(i64, 10)
}
func StringToUint64(str string) uint64 {
	intNum, _ := strconv.Atoi(str)
	return uint64(intNum)
}
//获取随机code
var defaultLetters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
func RandString(size int) string {
	if size <= 0 || size > len(defaultLetters) {
		size = len(defaultLetters)
	}
	result := make([]byte, size)
	for i := 0; i < size; i++ {
		result[i] = defaultLetters[rand.Intn(len(defaultLetters))]
	}
	return string(result)
}
func Rand6NumString() string {
	min := 100000
	max := 999999
	return strconv.Itoa(GetRandNum(min, max))
}
func GetRandNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + int(rand.Int63n(int64(max)-int64(min)+1))
}
func Length(str string) int {
	return utf8.RuneCountInString(str)
}
func SubString(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}