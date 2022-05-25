package internal

import (
	"math/rand"
	"time"
)

func RandomMessage() string {
	rand.Seed(time.Now().UnixMicro())
	switch rand.Intn(3) {
	case 0:
		return "Hello World!"
	case 1:
		return "Hola Mundo!"
	default:
		return ""
	}
}
