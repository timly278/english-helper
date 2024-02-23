package service

import "math/rand"

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomIntSlice(n, min, max int) []int {
	if n > max-min+1 {
		return nil // Not enough unique elements in the range
	}

	perm := rand.Perm(max - min + 1)

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = min + perm[i]
	}

	return slice
}