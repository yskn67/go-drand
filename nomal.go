package drand

import (
	"math/rand"
)

func Nomal(mu, sigma float64, seed int64) float64 {
	r := rand.New(rand.NewSource(seed))
	return stdLibNomal(r, mu, sigma)
}

func stdLibNomal(r *rand.Rand, mu, sigma float64) float64 {
	return mu + sigma*r.NormFloat64()
}
