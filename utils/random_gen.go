package utils

import "math/rand"

func Random_Gen(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
