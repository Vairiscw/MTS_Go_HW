package sevices

import (
	"math/rand"
)

func GetRandomInt(maxNumber int) int {
	return rand.Intn(maxNumber)
}
