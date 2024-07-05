package helpers

import (
	"fmt"
	"math/rand"
)

func GenerateAsset() string {
	a := rand.Intn(256)
	b := rand.Intn(256)
	c := rand.Intn(256)
	d := rand.Intn(256)

	return fmt.Sprintf("%d.%d.%d.%d", a, b, c, d)
}

func GenerateString(length int) string {
	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}
