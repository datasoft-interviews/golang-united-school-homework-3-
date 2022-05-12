package homework

import (
	"math"
	"testing"
)

func Test1(t *testing.T) {
	input := [15]float32{1, 5, 3, 7, 8, 4, 3, 6, 7, 8, 3.3, 5, 7, 3, 1}
	const threshold = 1e-6
	var want float32 = 4.753334
	if got := average(input); math.Abs(float64(got-want)) > threshold {
		t.Errorf("average() = %f, want %f", got, want)
	}
}

func Test2(t *testing.T) {
	input := [15]float32{1, 2, 3, 4, 5, 6}
	const threshold = 1e-6
	var want float32 = 1.4
	if got := average(input); math.Abs(float64(got-want)) > threshold {
		t.Errorf("average() = %f, want %f", got, want)
	}
}
