package homework

import (
	"testing"
)

func Test3(t *testing.T) {
	input := []int64{1, 5, 3, 7, 8, 4, 3, 6, 7, 8, 3, 5, 7, 3, 1}
	want := []int64{1, 3, 7, 5, 3, 8, 7, 6, 3, 4, 8, 7, 3, 5, 1}
	if got := reverse(input); !equal(got, want) {
		t.Errorf("reverse() = %v, want %v", got, want)
	}
}

func Test4(t *testing.T) {
	input := []int64{1, 2, 5, 15}
	want := []int64{15, 5, 2, 1}
	if got := reverse(input); !equal(got, want) {
		t.Errorf("reverse() = %v, want %v", got, want)
	}
}

func equal(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
