package homework

import (
	"testing"
)

func Test5(t *testing.T) {
	input := map[int]string{2: "a", 0: "b", 1: "c"}
	want := []string{"b", "c", "a"}
	if got := sortMapValues(input); !equal_string_slices(got, want) {
		t.Errorf("average() = %v, want %v", got, want)
	}
}

func Test6(t *testing.T) {
	input := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	want := []string{"bb", "aa", "cc"}
	if got := sortMapValues(input); !equal_string_slices(got, want) {
		t.Errorf("average() = %v, want %v", got, want)
	}
}

func equal_string_slices(a, b []string) bool {
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
