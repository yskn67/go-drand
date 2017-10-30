package drand

import (
	"math/rand"
	"testing"
)

func TestNomal(t *testing.T) {
	r := Nomal(0, 1, 0)
	t.Logf("random sample is %f, mu: 0, sigma: 1", r)
	r = Nomal(5, 1, 0)
	t.Logf("random sample is %f, mu: 5, sigma: 1", r)
	r = Nomal(0, 3, 0)
	t.Logf("random sample is %f, mu: 0, sigma: 3", r)
}

func TestBoxMuller(t *testing.T) {
	rnd := rand.New(rand.NewSource(0))

	// boxMullerFlag test
	if boxMullerFlag == 0 {
		_ = boxMuller(rnd, 0, 1)
		if boxMullerFlag != 1 {
			t.Error("boxMullerFlag is not toggled")
		}
	} else {
		_ = boxMuller(rnd, 0, 1)
		if boxMullerFlag != 0 {
			t.Error("boxMullerFlag is not toggled")
		}
	}

	// nomal random sample range test
	var r float64
	rnd = rand.New(rand.NewSource(0))
	outliers := 0
	for i := 0; i < 100; i++ {
		r = boxMuller(rnd, 0, 1)
		if r > 3 || r < -3 {
			outliers++
		}
	}
	if outliers > 1 {
		t.Errorf("outlier number is %d in seed 0", outliers)
	}

	rnd = rand.New(rand.NewSource(71))
	outliers = 0
	for i := 0; i < 100; i++ {
		r = boxMuller(rnd, 0, 1)
		if r > 3 || r < -3 {
			outliers++
		}
	}
	if outliers > 1 {
		t.Errorf("outlier number is %d in seed 71", outliers)
	}

	rnd = rand.New(rand.NewSource(84))
	outliers = 0
	for i := 0; i < 100; i++ {
		r = boxMuller(rnd, 0, 1)
		if r > 3 || r < -3 {
			outliers++
		}
	}
	if outliers > 1 {
		t.Errorf("outlier number is %d in seed 84", outliers)
	}
}
