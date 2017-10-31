package drand

import (
	"math"
	"math/rand"
)

var (
	boxMullerFlag = 0
	polarFlag     = 0
)

func Nomal(mu, sigma float64, seed int64) float64 {
	rnd := rand.New(rand.NewSource(seed))
	return mu + sigma*rnd.NormFloat64()
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

func polar(rnd *rand.Rand, mu, sigma float64) float64 {
	var v, v1, v2 float64
	for {
		v1 = 2*rnd.Float64() - 1
		v2 = 2*rnd.Float64() - 1
		v = math.Pow(v1, 2) + math.Pow(v2, 2)
		if v > 0 && v < 1 {
			break
		}
	}
	w := math.Sqrt((-2 * math.Log(v)) / v)
	if polarFlag == 0 {
		polarFlag = 1
		return mu + sigma*v1*w
	} else {
		polarFlag = 0
		return mu + sigma*v2*w
	}
}
