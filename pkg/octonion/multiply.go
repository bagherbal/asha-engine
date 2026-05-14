package octonion

import "fmt"

// BasisProduct returns the product of two octonion basis units.
// Basis index 0 is the scalar unit 1. Indices 1..7 are imaginary units e₁..e₇.
// The returned pair is (sign, basisIndex).
func BasisProduct(i, j int) (int, int, error) {
	if i < 0 || i > 7 || j < 0 || j > 7 {
		return 0, 0, fmt.Errorf("octonion basis index out of range: %d,%d", i, j)
	}
	if i == 0 {
		return 1, j, nil
	}
	if j == 0 {
		return 1, i, nil
	}
	a := i - 1
	b := j - 1
	if a == b {
		return -1, 0, nil
	}
	for _, term := range StandardFanoTerms() {
		tri := []int{term.I, term.J, term.K}
		if contains(tri, a) && contains(tri, b) {
			c := -1
			for _, v := range tri {
				if v != a && v != b {
					c = v
					break
				}
			}
			if c < 0 {
				return 0, 0, fmt.Errorf("malformed Fano triple")
			}
			return term.Sign * paritySign([]int{a, b, c}), c + 1, nil
		}
	}
	return 0, 0, fmt.Errorf("no Fano line contains imaginary pair %d,%d", a, b)
}

func contains(values []int, needle int) bool {
	for _, v := range values {
		if v == needle {
			return true
		}
	}
	return false
}
