// Package lift tests whether tangent-level gauge generators survive the finite
// Boolean support. The central question is not whether g₂ᴿ closes inside g₂;
// Gate 7 already verifies that. The question here is stricter:
//
//  1. lift the contact-preserving generators to Λ⁴R⁸;
//  2. compress them through the Boolean incidence isometry W;
//  3. restrict them to the Boolean contact complement Kᵀ;
//  4. test whether their commutators still close after this finite projection.
//
// This is an intentionally harsh gate. A large residual is not hidden; it is a
// no-go diagnostic for the naive finite gauge lift.
package lift

import (
	"sync"

	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/combinatorics"
	"github.com/bagherbal/asha-engine/pkg/gauge"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type Compression struct {
	Contact              contact.Space
	Centralizer          gauge.Centralizer
	ExteriorGenerators   []linear.Matrix // 70x70 operators on Λ⁴R⁸.
	BooleanGenerators    []linear.Matrix // Wᵀρ(X)W on Boolean coordinates.
	RestrictedGenerators []linear.Matrix // P_C Wᵀρ(X)W P_C.

	BooleanContactFrame        linear.Matrix // 56x7 frame WᵀQ_K.
	BooleanContactProjector    linear.Matrix // 56x56 projector onto K in Boolean coordinates.
	BooleanComplementProjector linear.Matrix // 56x56 projector onto K⊥ inside Im(P_B).

	CompressedFrameRank int
	CompressedFrame     linear.Matrix // vectorized orthonormal frame of restricted generators.

	MaxSkewResidual         float64
	MaxBoundaryLeakage      float64
	ClosureResidual         float64
	ClosureRelativeResidual float64
}

var (
	liftDefaultOnce  sync.Once
	liftDefaultValue Compression
	liftDefaultErr   error
)

func BuildDefault() (Compression, error) {
	liftDefaultOnce.Do(func() {
		liftDefaultValue, liftDefaultErr = buildLiftDefaultUncached()
	})
	return liftDefaultValue, liftDefaultErr
}

func buildLiftDefaultUncached() (Compression, error) {
	k, err := contact.BuildDefault()
	if err != nil {
		return Compression{}, err
	}
	c, err := gauge.BuildContactCentralizer()
	if err != nil {
		return Compression{}, err
	}
	return Build(k, c, 1e-9)
}

func Build(k contact.Space, c gauge.Centralizer, eps float64) (Compression, error) {
	if eps <= 0 {
		eps = 1e-9
	}
	w := k.BooleanSupport.Normalized
	wt := w.Transpose()

	kBoolean, err := wt.Mul(k.ContactFrame)
	if err != nil {
		return Compression{}, err
	}
	pkb, err := kBoolean.Mul(kBoolean.Transpose())
	if err != nil {
		return Compression{}, err
	}
	pc, err := linear.Identity(w.Cols()).Sub(pkb)
	if err != nil {
		return Compression{}, err
	}

	exterior := make([]linear.Matrix, 0, c.CentralizerDimension)
	booleanOps := make([]linear.Matrix, 0, c.CentralizerDimension)
	restricted := make([]linear.Matrix, 0, c.CentralizerDimension)
	maxSkew := 0.0
	maxLeak := 0.0

	for col := 0; col < c.CentralizerFrame.Cols(); col++ {
		x7 := unvectorize7(column(c.CentralizerFrame, col))
		rho, err := ExteriorLiftToLambda4R8(x7, k.G2Support.Basis)
		if err != nil {
			return Compression{}, err
		}
		exterior = append(exterior, rho)

		bw, err := rho.Mul(w)
		if err != nil {
			return Compression{}, err
		}
		bgen, err := wt.Mul(bw)
		if err != nil {
			return Compression{}, err
		}
		booleanOps = append(booleanOps, bgen)

		leftLeak, err := pkb.Mul(bgen)
		if err != nil {
			return Compression{}, err
		}
		leftLeak, err = leftLeak.Mul(pc)
		if err != nil {
			return Compression{}, err
		}
		rightLeak, err := pc.Mul(bgen)
		if err != nil {
			return Compression{}, err
		}
		rightLeak, err = rightLeak.Mul(pkb)
		if err != nil {
			return Compression{}, err
		}
		if n := leftLeak.FrobeniusNorm(); n > maxLeak {
			maxLeak = n
		}
		if n := rightLeak.FrobeniusNorm(); n > maxLeak {
			maxLeak = n
		}

		tmp, err := pc.Mul(bgen)
		if err != nil {
			return Compression{}, err
		}
		rgen, err := tmp.Mul(pc)
		if err != nil {
			return Compression{}, err
		}
		restricted = append(restricted, rgen)
		skew, err := rgen.Add(rgen.Transpose())
		if err != nil {
			return Compression{}, err
		}
		if n := skew.FrobeniusNorm(); n > maxSkew {
			maxSkew = n
		}
	}

	frame, _, err := orthonormalOperatorSpan(restricted, eps)
	if err != nil {
		return Compression{}, err
	}
	closure, rel, err := closureResidual(restricted, frame)
	if err != nil {
		return Compression{}, err
	}

	return Compression{
		Contact:                    k,
		Centralizer:                c,
		ExteriorGenerators:         exterior,
		BooleanGenerators:          booleanOps,
		RestrictedGenerators:       restricted,
		BooleanContactFrame:        kBoolean,
		BooleanContactProjector:    pkb,
		BooleanComplementProjector: pc,
		CompressedFrameRank:        frame.Cols(),
		CompressedFrame:            frame,
		MaxSkewResidual:            maxSkew,
		MaxBoundaryLeakage:         maxLeak,
		ClosureResidual:            closure,
		ClosureRelativeResidual:    rel,
	}, nil
}

// ExteriorLiftToLambda4R8 lifts a 7x7 operator on Im(O) to an operator on
// Λ⁴R⁸. The real axis e₀ is fixed, and the 7 imaginary directions e₁,...,e₇
// are acted on by x7. Matrix columns represent input basis vectors.
func ExteriorLiftToLambda4R8(x7 linear.Matrix, basis []combinatorics.Subset) (linear.Matrix, error) {
	if x7.Rows() != 7 || x7.Cols() != 7 {
		return linear.Matrix{}, fmt.Errorf("expected 7x7 imaginary-octonion operator, got %dx%d", x7.Rows(), x7.Cols())
	}
	index := combinatorics.IndexByKey(basis)
	out := linear.NewMatrix(len(basis), len(basis))

	for col, blade := range basis {
		if len(blade) != 4 {
			return linear.Matrix{}, fmt.Errorf("Λ⁴ basis expected blades of length 4, got %d", len(blade))
		}
		for pos, oldAxis := range blade {
			// e₀ is fixed by the lifted g₂ action.
			if oldAxis == 0 {
				continue
			}
			oldImag := oldAxis - 1
			for newImag := 0; newImag < 7; newImag++ {
				coeff := x7.At(newImag, oldImag)
				if math.Abs(coeff) == 0 {
					continue
				}
				newAxis := newImag + 1
				seq := append([]int{}, blade...)
				seq[pos] = newAxis
				if hasDuplicate(seq) {
					continue
				}
				sorted := sortedCopy(seq)
				row, ok := index[combinatorics.Subset(sorted).Key()]
				if !ok {
					return linear.Matrix{}, fmt.Errorf("missing lifted Λ⁴ basis row for %v", sorted)
				}
				out.Set(row, col, out.At(row, col)+float64(paritySign(seq))*coeff)
			}
		}
	}
	return out, nil
}

func closureResidual(ops []linear.Matrix, frame linear.Matrix) (float64, float64, error) {
	if len(ops) == 0 {
		return 0, 0, fmt.Errorf("empty operator list")
	}
	projector, err := frame.Mul(frame.Transpose())
	if err != nil {
		return 0, 0, err
	}
	maxAbs := 0.0
	maxRel := 0.0
	for i := 0; i < len(ops); i++ {
		for j := i + 1; j < len(ops); j++ {
			bracket, err := commutator(ops[i], ops[j])
			if err != nil {
				return 0, 0, err
			}
			v := vectorColumn(vectorizeOperator(bracket))
			pv, err := projector.Mul(v)
			if err != nil {
				return 0, 0, err
			}
			diff, err := v.Sub(pv)
			if err != nil {
				return 0, 0, err
			}
			abs := diff.FrobeniusNorm()
			norm := v.FrobeniusNorm()
			rel := 0.0
			if norm > 0 {
				rel = abs / norm
			}
			if abs > maxAbs {
				maxAbs = abs
			}
			if rel > maxRel {
				maxRel = rel
			}
		}
	}
	return maxAbs, maxRel, nil
}

func orthonormalOperatorSpan(matrices []linear.Matrix, eps float64) (linear.Matrix, []float64, error) {
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

func unvectorize7(values []float64) linear.Matrix {
	m := linear.NewMatrix(7, 7)
	for r := 0; r < 7; r++ {
		for c := 0; c < 7; c++ {
			m.Set(r, c, values[r*7+c])
		}
	}
	return m
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

func vectorColumn(values []float64) linear.Matrix {
	m := linear.NewMatrix(len(values), 1)
	for i, v := range values {
		m.Set(i, 0, v)
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

func hasDuplicate(values []int) bool {
	seen := map[int]bool{}
	for _, v := range values {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}

func sortedCopy(values []int) []int {
	out := append([]int{}, values...)
	for i := 1; i < len(out); i++ {
		key := out[i]
		j := i - 1
		for j >= 0 && out[j] > key {
			out[j+1] = out[j]
			j--
		}
		out[j+1] = key
	}
	return out
}

// paritySign returns the sign of the permutation that sorts the current axis
// sequence into increasing order.
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
