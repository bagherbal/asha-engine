package octonion

import "fmt"

type AssociativeForm struct {
	dim  int
	data []float64
}

type CoassociativeForm struct {
	dim  int
	data []float64
}

func StandardAssociativeForm() AssociativeForm {
	const dim = 7
	f := AssociativeForm{dim: dim, data: make([]float64, dim*dim*dim)}
	for _, term := range StandardFanoTerms() {
		triple := []int{term.I, term.J, term.K}
		for _, p := range permutations(triple) {
			f.set(p[0], p[1], p[2], float64(term.Sign*permutationSign(triple, p)))
		}
	}
	return f
}

func StandardCoassociativeForm(phi AssociativeForm) CoassociativeForm {
	const dim = 7
	f := CoassociativeForm{dim: dim, data: make([]float64, dim*dim*dim*dim)}
	for a := 0; a < dim; a++ {
		for b := a + 1; b < dim; b++ {
			for c := b + 1; c < dim; c++ {
				for d := c + 1; d < dim; d++ {
					quad := []int{a, b, c, d}
					comp := complement(dim, quad)
					full := append(append([]int{}, quad...), comp...)
					val := float64(paritySign(full)) * phi.Value(comp[0], comp[1], comp[2])
					if val == 0 {
						continue
					}
					for _, p := range permutations(quad) {
						f.set(p[0], p[1], p[2], p[3], float64(permutationSign(quad, p))*val)
					}
				}
			}
		}
	}
	return f
}

func (f AssociativeForm) Dimension() int { return f.dim }

func (f AssociativeForm) Value(i, j, k int) float64 {
	if i < 0 || i >= f.dim || j < 0 || j >= f.dim || k < 0 || k >= f.dim {
		panic(fmt.Sprintf("associative form index out of range: %d,%d,%d", i, j, k))
	}
	return f.data[(i*f.dim+j)*f.dim+k]
}

func (f AssociativeForm) NonZeroCanonicalTerms() int {
	count := 0
	for i := 0; i < f.dim; i++ {
		for j := i + 1; j < f.dim; j++ {
			for k := j + 1; k < f.dim; k++ {
				if f.Value(i, j, k) != 0 {
					count++
				}
			}
		}
	}
	return count
}

func (f *AssociativeForm) set(i, j, k int, value float64) {
	f.data[(i*f.dim+j)*f.dim+k] = value
}

func (f CoassociativeForm) Dimension() int { return f.dim }

func (f CoassociativeForm) Value(i, j, k, l int) float64 {
	if i < 0 || i >= f.dim || j < 0 || j >= f.dim || k < 0 || k >= f.dim || l < 0 || l >= f.dim {
		panic(fmt.Sprintf("coassociative form index out of range: %d,%d,%d,%d", i, j, k, l))
	}
	return f.data[((i*f.dim+j)*f.dim+k)*f.dim+l]
}

func (f CoassociativeForm) NonZeroCanonicalTerms() int {
	count := 0
	for i := 0; i < f.dim; i++ {
		for j := i + 1; j < f.dim; j++ {
			for k := j + 1; k < f.dim; k++ {
				for l := k + 1; l < f.dim; l++ {
					if f.Value(i, j, k, l) != 0 {
						count++
					}
				}
			}
		}
	}
	return count
}

func (f *CoassociativeForm) set(i, j, k, l int, value float64) {
	f.data[((i*f.dim+j)*f.dim+k)*f.dim+l] = value
}

func complement(dim int, selected []int) []int {
	out := make([]int, 0, dim-len(selected))
	for i := 0; i < dim; i++ {
		found := false
		for _, v := range selected {
			if i == v {
				found = true
				break
			}
		}
		if !found {
			out = append(out, i)
		}
	}
	return out
}

func permutations(values []int) [][]int {
	out := make([][]int, 0)
	current := append([]int{}, values...)
	var walk func(int)
	walk = func(pos int) {
		if pos == len(current) {
			p := append([]int{}, current...)
			out = append(out, p)
			return
		}
		for i := pos; i < len(current); i++ {
			current[pos], current[i] = current[i], current[pos]
			walk(pos + 1)
			current[pos], current[i] = current[i], current[pos]
		}
	}
	walk(0)
	return out
}

// permutationSign returns the sign of the permutation that maps source to target.
func permutationSign(source, target []int) int {
	if len(source) != len(target) {
		panic("permutationSign length mismatch")
	}
	position := make(map[int]int, len(source))
	for i, v := range source {
		position[v] = i
	}
	perm := make([]int, len(target))
	for i, v := range target {
		perm[i] = position[v]
	}
	return paritySign(perm)
}

func paritySign(values []int) int {
	inv := 0
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				inv++
			}
		}
	}
	if inv%2 == 0 {
		return 1
	}
	return -1
}
