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
