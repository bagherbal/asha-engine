// Package gauge contains finite Lie-algebra tests for the Asha engine.
//
// This file implements the contact-copy centralizer theorem inside g₂. It is
// deliberately standard in its naming: octonion derivations generate g₂, an
// involution R acts on Im(O), and the centralizer g₂ᴿ is recovered as the
// kernel of the commutator map X ↦ [X,R].
package gauge

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/octonion"
)

type Centralizer struct {
	RawDerivations          []linear.Matrix
	G2Frame                 linear.Matrix // 49 x 14, vectorized 7x7 matrices.
	G2Rank                  int
	Involution              linear.Matrix // 7 x 7.
	CentralizerFrame        linear.Matrix // 49 x d, vectorized centralizer matrices.
	CentralizerCoefficients linear.Matrix // 14 x d, coordinates inside G2Frame.
	CentralizerDimension    int
	CenterDimension         int
	DerivedDimension        int
	G2GramEigenvalues       []float64
	CentralizerEigenvalues  []float64
}

func BuildContactCentralizer() (Centralizer, error) {
	raw, err := OctonionDerivations()
	if err != nil {
		return Centralizer{}, err
	}
	g2Frame, g2Values, err := orthonormalSpan(raw, 1e-9)
	if err != nil {
		return Centralizer{}, err
	}
	r := ContactCopyInvolution()

	constraint := linear.NewMatrix(49, g2Frame.Cols())
	for col := 0; col < g2Frame.Cols(); col++ {
		x := unvectorizeOperator(column(g2Frame, col), 7)
		comm, err := commutator(x, r)
		if err != nil {
			return Centralizer{}, err
		}
		vec := vectorizeOperator(comm)
		for row, v := range vec {
			constraint.Set(row, col, v)
		}
	}
	gram, err := constraint.Transpose().Mul(constraint)
	if err != nil {
		return Centralizer{}, err
	}
	eig, err := linear.SymmetricEigenJacobi(gram, 1e-13, 0)
	if err != nil {
		return Centralizer{}, err
	}
	values, vectors, err := sortEigenAscending(eig.Values, eig.Vectors)
	if err != nil {
		return Centralizer{}, err
	}
	zero := make([]int, 0)
	for i, value := range values {
		if math.Abs(value) < 1e-8 {
			zero = append(zero, i)
		}
	}
	if len(zero) == 0 {
		return Centralizer{}, fmt.Errorf("centralizer kernel is empty")
	}
	coeff := linear.NewMatrix(g2Frame.Cols(), len(zero))
	for newCol, oldCol := range zero {
		for row := 0; row < vectors.Rows(); row++ {
			coeff.Set(row, newCol, vectors.At(row, oldCol))
		}
	}
	centralFrame, err := g2Frame.Mul(coeff)
	if err != nil {
		return Centralizer{}, err
	}
	centerDim, err := centerDimension(centralFrame, 1e-8)
	if err != nil {
		return Centralizer{}, err
	}
	derivedDim, err := derivedDimension(centralFrame, 1e-8)
	if err != nil {
		return Centralizer{}, err
	}
	return Centralizer{
		RawDerivations:          raw,
		G2Frame:                 g2Frame,
		G2Rank:                  g2Frame.Cols(),
		Involution:              r,
		CentralizerFrame:        centralFrame,
		CentralizerCoefficients: coeff,
		CentralizerDimension:    centralFrame.Cols(),
		CenterDimension:         centerDim,
		DerivedDimension:        derivedDim,
		G2GramEigenvalues:       g2Values,
		CentralizerEigenvalues:  values,
	}, nil
}

// OctonionDerivations returns the 21 standard derivations D_ab on Im(O).
// Their span has dimension 14 and is the compact Lie algebra g₂.
func OctonionDerivations() ([]linear.Matrix, error) {
	left := make([]linear.Matrix, 8)
	right := make([]linear.Matrix, 8)
	for i := 0; i < 8; i++ {
		l, err := leftMultiplication(i)
		if err != nil {
			return nil, err
		}
		r, err := rightMultiplication(i)
		if err != nil {
			return nil, err
		}
		left[i] = l
		right[i] = r
	}
	out := make([]linear.Matrix, 0, 21)
	for a := 1; a <= 7; a++ {
		for b := a + 1; b <= 7; b++ {
			ll, err := commutator(left[a], left[b])
			if err != nil {
				return nil, err
			}
			lr, err := commutator(left[a], right[b])
			if err != nil {
				return nil, err
			}
			rr, err := commutator(right[a], right[b])
			if err != nil {
				return nil, err
			}
			sum, err := ll.Add(lr)
			if err != nil {
				return nil, err
			}
			sum, err = sum.Add(rr)
			if err != nil {
				return nil, err
			}
			out = append(out, imaginaryBlock(sum))
		}
	}
	return out, nil
}

// ContactCopyInvolution is the R used by the finite contact-copy centralizer.
// It acts on Im(O) in the ordered basis e₁,...,e₇.
func ContactCopyInvolution() linear.Matrix {
	return linear.Diagonal([]float64{-1, +1, -1, +1, +1, +1, +1})
}

func (c Centralizer) G2SkewResidual() float64 {
	max := 0.0
	for i := 0; i < c.G2Frame.Cols(); i++ {
		x := unvectorizeOperator(column(c.G2Frame, i), 7)
		diff, _ := x.Add(x.Transpose())
		if n := diff.FrobeniusNorm(); n > max {
			max = n
		}
	}
	return max
}

func (c Centralizer) InvolutionResidual() float64 {
	r2, _ := c.Involution.Mul(c.Involution)
	diff, _ := r2.Sub(linear.Identity(c.Involution.Rows()))
	return diff.FrobeniusNorm()
}

func (c Centralizer) CentralizerResidual() float64 {
	max := 0.0
	for i := 0; i < c.CentralizerFrame.Cols(); i++ {
		x := unvectorizeOperator(column(c.CentralizerFrame, i), 7)
		comm, _ := commutator(x, c.Involution)
		if n := comm.FrobeniusNorm(); n > max {
			max = n
		}
	}
	return max
}

func (c Centralizer) ClosureResidual() float64 {
	max := 0.0
	p, _ := c.CentralizerFrame.Mul(c.CentralizerFrame.Transpose())
	for i := 0; i < c.CentralizerFrame.Cols(); i++ {
		x := unvectorizeOperator(column(c.CentralizerFrame, i), 7)
		for j := i + 1; j < c.CentralizerFrame.Cols(); j++ {
			y := unvectorizeOperator(column(c.CentralizerFrame, j), 7)
			bracket, _ := commutator(x, y)
			vec := vectorColumn(vectorizeOperator(bracket))
			proj, _ := p.Mul(vec)
			diff, _ := vec.Sub(proj)
			if n := diff.FrobeniusNorm(); n > max {
				max = n
			}
		}
	}
	return max
}

func (c Centralizer) FrameIsometryResidual() float64 {
	qtq, _ := c.CentralizerFrame.Transpose().Mul(c.CentralizerFrame)
	diff, _ := qtq.Sub(linear.Identity(c.CentralizerFrame.Cols()))
	return diff.FrobeniusNorm()
}

func leftMultiplication(idx int) (linear.Matrix, error) {
	m := linear.NewMatrix(8, 8)
	for j := 0; j < 8; j++ {
		sign, k, err := octonion.BasisProduct(idx, j)
		if err != nil {
			return linear.Matrix{}, err
		}
		m.Set(k, j, float64(sign))
	}
	return m, nil
}

func rightMultiplication(idx int) (linear.Matrix, error) {
	m := linear.NewMatrix(8, 8)
	for j := 0; j < 8; j++ {
		sign, k, err := octonion.BasisProduct(j, idx)
		if err != nil {
			return linear.Matrix{}, err
		}
		m.Set(k, j, float64(sign))
	}
	return m, nil
}

func commutator(a, b linear.Matrix) (linear.Matrix, error) {
	ab, err := a.Mul(b)
	if err != nil {
		return linear.Matrix{}, err
	}
	ba, err := b.Mul(a)
	if err != nil {
		return linear.Matrix{}, err
	}
	return ab.Sub(ba)
}

func imaginaryBlock(m linear.Matrix) linear.Matrix {
	out := linear.NewMatrix(7, 7)
	for r := 0; r < 7; r++ {
		for c := 0; c < 7; c++ {
			out.Set(r, c, m.At(r+1, c+1))
		}
	}
	return out
}

func orthonormalSpan(matrices []linear.Matrix, eps float64) (linear.Matrix, []float64, error) {
	if len(matrices) == 0 {
		return linear.Matrix{}, nil, fmt.Errorf("empty operator list")
	}
	rows := matrices[0].Rows()
	cols := matrices[0].Cols()
	raw := linear.NewMatrix(rows*cols, len(matrices))
	for c, op := range matrices {
		if op.Rows() != rows || op.Cols() != cols {
			return linear.Matrix{}, nil, fmt.Errorf("operator shape mismatch")
		}
		vec := vectorizeOperator(op)
		for r, v := range vec {
			raw.Set(r, c, v)
		}
	}
	gram, err := raw.Transpose().Mul(raw)
	if err != nil {
		return linear.Matrix{}, nil, err
	}
	eig, err := linear.SymmetricEigenJacobi(gram, 1e-13, 0)
	if err != nil {
		return linear.Matrix{}, nil, err
	}
	values, vectors, err := linear.SortEigenDescending(eig.Values, eig.Vectors)
	if err != nil {
		return linear.Matrix{}, nil, err
	}
	positive := make([]int, 0)
	for i, value := range values {
		if value > eps {
			positive = append(positive, i)
		}
	}
	if len(positive) == 0 {
		return linear.Matrix{}, nil, fmt.Errorf("operator span rank is zero")
	}
	vpos := linear.NewMatrix(vectors.Rows(), len(positive))
	inv := linear.NewMatrix(len(positive), len(positive))
	selected := make([]float64, len(positive))
	for newCol, oldCol := range positive {
		selected[newCol] = values[oldCol]
		inv.Set(newCol, newCol, 1/math.Sqrt(values[oldCol]))
		for row := 0; row < vectors.Rows(); row++ {
			vpos.Set(row, newCol, vectors.At(row, oldCol))
		}
	}
	rv, err := raw.Mul(vpos)
	if err != nil {
		return linear.Matrix{}, nil, err
	}
	frame, err := rv.Mul(inv)
	if err != nil {
		return linear.Matrix{}, nil, err
	}
	return frame, selected, nil
}

func centerDimension(frame linear.Matrix, eps float64) (int, error) {
	dim := frame.Cols()
	constraint := linear.NewMatrix(dim*49, dim)
	row := 0
	for i := 0; i < dim; i++ {
		x := unvectorizeOperator(column(frame, i), 7)
		for j := 0; j < dim; j++ {
			y := unvectorizeOperator(column(frame, j), 7)
			bracket, err := commutator(y, x)
			if err != nil {
				return 0, err
			}
			vec := vectorizeOperator(bracket)
			for k, v := range vec {
				constraint.Set(row+k, j, v)
			}
		}
		row += 49
	}
	gram, err := constraint.Transpose().Mul(constraint)
	if err != nil {
		return 0, err
	}
	eig, err := linear.SymmetricEigenJacobi(gram, 1e-13, 0)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, value := range eig.Values {
		if math.Abs(value) < eps {
			count++
		}
	}
	return count, nil
}

func derivedDimension(frame linear.Matrix, eps float64) (int, error) {
	dim := frame.Cols()
	brackets := make([]linear.Matrix, 0)
	for i := 0; i < dim; i++ {
		x := unvectorizeOperator(column(frame, i), 7)
		for j := i + 1; j < dim; j++ {
			y := unvectorizeOperator(column(frame, j), 7)
			bracket, err := commutator(x, y)
			if err != nil {
				return 0, err
			}
			if bracket.FrobeniusNorm() > eps {
				brackets = append(brackets, bracket)
			}
		}
	}
	if len(brackets) == 0 {
		return 0, nil
	}
	frameBrackets, _, err := orthonormalSpan(brackets, eps)
	if err != nil {
		return 0, err
	}
	return frameBrackets.Cols(), nil
}

func vectorizeOperator(m linear.Matrix) []float64 {
	out := make([]float64, 0, m.Rows()*m.Cols())
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			out = append(out, m.At(r, c))
		}
	}
	return out
}

func unvectorizeOperator(values []float64, n int) linear.Matrix {
	m := linear.NewMatrix(n, n)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			m.Set(r, c, values[r*n+c])
		}
	}
	return m
}

func column(m linear.Matrix, c int) []float64 {
	out := make([]float64, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		out[r] = m.At(r, c)
	}
	return out
}

func vectorColumn(values []float64) linear.Matrix {
	m := linear.NewMatrix(len(values), 1)
	for i, v := range values {
		m.Set(i, 0, v)
	}
	return m
}

func sortEigenAscending(values []float64, vectors linear.Matrix) ([]float64, linear.Matrix, error) {
	if len(values) != vectors.Cols() || vectors.Rows() != vectors.Cols() {
		return nil, linear.Matrix{}, fmt.Errorf("eigenvectors must be square with one column per eigenvalue")
	}
	idx := make([]int, len(values))
	for i := range idx {
		idx[i] = i
	}
	for i := 0; i < len(idx); i++ {
		min := i
		for j := i + 1; j < len(idx); j++ {
			if values[idx[j]] < values[idx[min]] {
				min = j
			}
		}
		idx[i], idx[min] = idx[min], idx[i]
	}
	sortedValues := make([]float64, len(values))
	sortedVectors := linear.NewMatrix(vectors.Rows(), vectors.Cols())
	for newCol, oldCol := range idx {
		sortedValues[newCol] = values[oldCol]
		for row := 0; row < vectors.Rows(); row++ {
			sortedVectors.Set(row, newCol, vectors.At(row, oldCol))
		}
	}
	return sortedValues, sortedVectors, nil
}
