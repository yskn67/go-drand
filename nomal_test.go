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

func TestStdLibNomal(t *testing.T) {
	rnd := rand.New(rand.NewSource(0))

	// nomal random sample range test
	var r float64
	rnd = rand.New(rand.NewSource(0))
	outliers := 0
	for i := 0; i < 100; i++ {
		r = stdLibNomal(rnd, 0, 1)
		if r > 3 || r < -3 {
			outliers++
		}
	}
	if outliers > 1 {
		t.Errorf("outlier number is %d in seed 0", outliers)
	}

	rnd = rand.New(rand.NewSource(11))
	outliers = 0
	for i := 0; i < 100; i++ {
		r = stdLibNomal(rnd, 0, 1)
		if r > 3 || r < -3 {
			outliers++
		}
	}
	if outliers > 1 {
		t.Errorf("outlier number is %d in seed 11", outliers)
	}

	rnd = rand.New(rand.NewSource(65))
	outliers = 0
	for i := 0; i < 100; i++ {
		r = stdLibNomal(rnd, 0, 1)
		if r > 3 || r < -3 {
			outliers++
		}
	}
	if outliers > 1 {
		t.Errorf("outlier number is %d in seed 65", outliers)
	}
}

func BenchmarkStdLibNomal(b *testing.B) {
	rnd := rand.New(rand.NewSource(0))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = stdLibNomal(rnd, 0, 1)
	}
}
