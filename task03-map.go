package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0)
	for key, _ := range input {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	values := make([]string, 0)
	for i := 0; i < len(keys); i++ {
		values = append(values, input[keys[i]])
	}
	return values
}
