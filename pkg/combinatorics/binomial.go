package combinatorics

import "fmt"

// Binomial returns n choose k using integer arithmetic.
func Binomial(n, k int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("n must be non-negative: %d", n)
	}
	if k < 0 || k > n {
		return 0, fmt.Errorf("k must satisfy 0 <= k <= n: n=%d k=%d", n, k)
	}
	if k > n-k {
		k = n - k
	}
	result := 1
	for i := 1; i <= k; i++ {
		result = result * (n - k + i) / i
	}
	return result, nil
}

// GradeDimensions returns dim Λ^k R^n for k=0..n.
func GradeDimensions(n int) ([]int, error) {
	if n < 0 {
		return nil, fmt.Errorf("dimension must be non-negative: %d", n)
	}
	dims := make([]int, n+1)
	for k := 0; k <= n; k++ {
		v, err := Binomial(n, k)
		if err != nil {
			return nil, err
		}
		dims[k] = v
	}
	return dims, nil
}

func Sum(values []int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
