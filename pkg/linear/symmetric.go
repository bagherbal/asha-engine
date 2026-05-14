package linear

import (
	"fmt"
	"math"
	"sort"
)

type Eigenpair struct {
	Value float64
	Index int
}

type SymmetricEigenResult struct {
	Values  []float64
	Vectors Matrix // columns are normalized eigenvectors
	Sweeps  int
}

// SymmetricEigenJacobi computes the eigen-decomposition of a real symmetric matrix
// using the classical Jacobi rotation algorithm. It is intentionally small and
// auditable: good enough for the finite 56x56 and 70x70 theorem gates.
func SymmetricEigenJacobi(input Matrix, eps float64, maxSweeps int) (SymmetricEigenResult, error) {
	if input.Rows() != input.Cols() {
		return SymmetricEigenResult{}, fmt.Errorf("symmetric eigensolver requires a square matrix: %dx%d", input.Rows(), input.Cols())
	}
	if !input.IsSymmetric(100 * eps) {
		return SymmetricEigenResult{}, fmt.Errorf("symmetric eigensolver received a non-symmetric matrix")
	}
	n := input.Rows()
	a := input.Clone()
	v := Identity(n)
	if maxSweeps <= 0 {
		maxSweeps = 10 * n * n
	}

	sweep := 0
	for ; sweep < maxSweeps; sweep++ {
		p, q, max := largestOffDiagonal(a)
		if max < eps {
			break
		}

		app := a.At(p, p)
		aqq := a.At(q, q)
		apq := a.At(p, q)
		angle := 0.5 * math.Atan2(2*apq, aqq-app)
		c := math.Cos(angle)
		s := math.Sin(angle)

		for k := 0; k < n; k++ {
			if k == p || k == q {
				continue
			}
			akp := a.At(k, p)
			akq := a.At(k, q)
			newKP := c*akp - s*akq
			newKQ := s*akp + c*akq
			a.Set(k, p, newKP)
			a.Set(p, k, newKP)
			a.Set(k, q, newKQ)
			a.Set(q, k, newKQ)
		}

		newPP := c*c*app - 2*s*c*apq + s*s*aqq
		newQQ := s*s*app + 2*s*c*apq + c*c*aqq
		a.Set(p, p, newPP)
		a.Set(q, q, newQQ)
		a.Set(p, q, 0)
		a.Set(q, p, 0)

		for k := 0; k < n; k++ {
			vkp := v.At(k, p)
			vkq := v.At(k, q)
			v.Set(k, p, c*vkp-s*vkq)
			v.Set(k, q, s*vkp+c*vkq)
		}
	}

	values := make([]float64, n)
	for i := 0; i < n; i++ {
		values[i] = a.At(i, i)
	}
	return SymmetricEigenResult{Values: values, Vectors: v, Sweeps: sweep}, nil
}

func largestOffDiagonal(a Matrix) (int, int, float64) {
	n := a.Rows()
	p, q := 0, 1
	max := 0.0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			abs := math.Abs(a.At(i, j))
			if abs > max {
				max = abs
				p = i
				q = j
			}
		}
	}
	return p, q, max
}

func SortEigenDescending(values []float64, vectors Matrix) ([]float64, Matrix, error) {
	if len(values) != vectors.Cols() || vectors.Rows() != vectors.Cols() {
		return nil, Matrix{}, fmt.Errorf("eigenvectors must be square with one column per eigenvalue")
	}
	pairs := make([]Eigenpair, len(values))
	for i, value := range values {
		pairs[i] = Eigenpair{Value: value, Index: i}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Value > pairs[j].Value })

	sortedValues := make([]float64, len(values))
	sortedVectors := NewMatrix(vectors.Rows(), vectors.Cols())
	for newCol, pair := range pairs {
		sortedValues[newCol] = pair.Value
		for row := 0; row < vectors.Rows(); row++ {
			sortedVectors.Set(row, newCol, vectors.At(row, pair.Index))
		}
	}
	return sortedValues, sortedVectors, nil
}

func RankFromEigenvalues(values []float64, eps float64) int {
	rank := 0
	for _, v := range values {
		if math.Abs(v) > eps {
			rank++
		}
	}
	return rank
}
