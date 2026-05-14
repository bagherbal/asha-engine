package combinatorics

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
