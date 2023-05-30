package utils

import (
	"math/rand"
	"strings"
	"time"
)

var numberRunes = []rune("0123456789")

func RandomNumericRunes(length int) string {
	rand.Seed(time.Now().UnixNano())
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(numberRunes[rand.Intn(len(numberRunes))])
	}
	str := b.String()
	return str
}
