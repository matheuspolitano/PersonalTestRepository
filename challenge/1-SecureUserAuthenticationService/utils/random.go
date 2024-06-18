package utils

import (
	"math/rand"
	"strings"
	"time"
)

var mainRand *rand.Rand

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rsource := rand.NewSource(time.Now().UnixNano())
	mainRand = rand.New(rsource)

}

func RandString(size int) string {
	bAlphabet := []byte(alphabet)
	lenAlphabet := len(alphabet)

	var stringBuilder strings.Builder
	for i := 0; i < size; i++ {
		stringBuilder.WriteByte(bAlphabet[mainRand.Intn(lenAlphabet)])
	}
	return stringBuilder.String()

}
