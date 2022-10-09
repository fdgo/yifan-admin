package define

import (
	"math/rand"
	"time"
	"yifan/pkg/randId"
)

func GetRandIpId() uint {
	return randId.RandID()
}
func GetRandSeriesId() uint {
	return randId.RandID()
}

func GetRandFanId() uint {
	return randId.RandID()
}
func GetRandGoodId() uint {
	return randId.RandID()
}
func GetRandBoxId() uint {
	return randId.RandID()
}
func GetRandPrizeId() uint {
	return randId.RandID()
}
func GetRandUserId() uint {
	return randId.RandID()
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
