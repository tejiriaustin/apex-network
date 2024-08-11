package utils

import (
	"crypto/rand"
	"math/big"
)

type RandFunc func() int

func genRandomNumber() int {
	number, _ := rand.Int(rand.Reader, big.NewInt(10))
	return int(number.Int64() + 2)
}

func GetRandFunc() RandFunc {
	return genRandomNumber
}
