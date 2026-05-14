package combinatorics

import (
	"fmt"
	"strings"
)

type Subset []int

// Subsets returns all k-subsets of {0,...,n-1} in lexicographic order.
func Subsets(n, k int) ([]Subset, error) {
	if _, err := Binomial(n, k); err != nil {
		return nil, err
	}
	out := make([]Subset, 0)
	current := make([]int, 0, k)

	var walk func(start int)
	walk = func(start int) {
		if len(current) == k {
			s := make([]int, k)
			copy(s, current)
			out = append(out, s)
			return
		}
		remaining := k - len(current)
		for i := start; i <= n-remaining; i++ {
			current = append(current, i)
			walk(i + 1)
			current = current[:len(current)-1]
		}
	}

	walk(0)
	return out, nil
}

func (s Subset) Key() string {
	parts := make([]string, len(s))
	for i, v := range s {
		parts[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(parts, ",")
}

func (s Subset) Contains(value int) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

func (s Subset) ContainsAll(other Subset) bool {
	for _, v := range other {
		if !s.Contains(v) {
			return false
		}
	}
	return true
}

func IndexByKey(subsets []Subset) map[string]int {
	index := make(map[string]int, len(subsets))
	for i, subset := range subsets {
		index[subset.Key()] = i
	}
	return index
}
