package drand

import (
	"math"
	"math/rand"
)

var (
	boxMullerFlag = 0
)

func Nomal(mu, sigma float64, seed int64) float64 {
	rnd := rand.New(rand.NewSource(seed))
	return boxMuller(rnd, mu, sigma)
}

func boxMuller(rnd *rand.Rand, mu, sigma float64) float64 {
	r := math.Sqrt(-2 * math.Log(rnd.Float64()))
	omega := 2 * math.Pi * rnd.Float64()
	if boxMullerFlag == 0 {
		boxMullerFlag = 1
		return mu + sigma*(r*math.Cos(omega))
	} else {
		boxMullerFlag = 0
		return mu + sigma*(r*math.Sin(omega))
	}
}
