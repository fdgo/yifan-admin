package define

import (
	gonanoid "github.com/matoous/go-nanoid"
	"math/rand"
	"strconv"
	"time"
)

func GetRandIpId() uint {
	idStr, _ := gonanoid.Generate("123456789", 9)
	id, _ := strconv.Atoi(idStr)
	return uint(id)
}
func GetRandSeriesId() uint {
	idStr, _ := gonanoid.Generate("123456789", 9)
	id, _ := strconv.Atoi(idStr)
	return uint(id)
}

func GetRandFanId() uint {
	idStr, _ := gonanoid.Generate("123456789", 9)
	id, _ := strconv.Atoi(idStr)
	return uint(id)
}
func GetRandGoodId() uint {
	idStr, _ := gonanoid.Generate("123456789", 9)
	id, _ := strconv.Atoi(idStr)
	return uint(id)
}
func GetRandBoxId() uint {
	idStr, _ := gonanoid.Generate("123456789", 9)
	id, _ := strconv.Atoi(idStr)
	return uint(id)
}
func GetRandPrizeId() uint {
	idStr, _ := gonanoid.Generate("123456789", 9)
	id, _ := strconv.Atoi(idStr)
	return uint(id)
}
func GetRandUserId() uint {
	idStr, _ := gonanoid.Generate("123456789", 9)
	id, _ := strconv.Atoi(idStr)
	return uint(id)
}

func RandPrizeIndex(size int64) int64 {
	seed := time.Now().UnixNano()
	n := rand.New(rand.NewSource(seed)).Intn(int(size))
	return int64(n)
}
func RangeRand(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(max-min) + min
	return r
}
func GetRandRums(min, max, many int) []int {
	numSlice := make([]int, 0)
	for n := 0; n < many; n++ {
		numSlice = append(numSlice, RangeRand(min, max+1))
	}
	isExist := false
	for i := 0; i < len(numSlice); i++ {
		for j := 0; j < i; j++ {
			if numSlice[i] == numSlice[j] {
				isExist = true
			}
		}
	}
	if isExist {
		return GetRandRums(min, max, many)
	} else {
		return numSlice
	}
}
