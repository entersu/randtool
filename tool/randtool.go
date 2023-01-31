package tool

import "math/rand"

// RandNumber create random number

func RandNumber() int {
	n := 10
	//  returns, as an int, a non-negative pseudo-random number in [0,n)
	return rand.Intn(n)
}