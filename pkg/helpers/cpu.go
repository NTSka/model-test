package helpers

import "crypto/sha256"

func CPUBoundTask() {
	st := ""
	for i := 0; i < 100; i++ {
		st += GenerateString(10)
	}

	v := sha256.New()
	v.Write([]byte(st))
}
