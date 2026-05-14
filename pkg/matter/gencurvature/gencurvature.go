// Package gencurvature restricts the finite projected-connection curvature to
// the three protected contact directions exposed by the Higgs/contact sector.
//
// Gate 29 found a diagonal generation-breaking spurion from Higgs/contact
// anisotropy, but it did not find a mixing operator. This package tests the
// next natural source: the second-fundamental curvature of the block connection.
//
// The key point is categorical. The projection-curvature identity has two
// sides. On the contact-complement side it explains why PAP does not close as a
// Lie representation. On the contact side it produces the mirror curvature
//
//	R^K_AB = P_K A P_C B P_K - P_K B P_C A P_K,
//
// which is a real skew operator on K. Restricting R^K_AB to the three unmixed
// protected contact directions tests whether the finite geometry supplies a
// non-diagonal generation-rotation carrier. This is not yet a Yukawa mass
// matrix; it is the first honest curvature source for generation mixing.
package gencurvature

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/dynamics/higgspotential"
	"github.com/bagherbal/asha-engine/pkg/gauge/connection"
	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/generationbreak"
)

type Operator struct {
	Pair             string
	Matrix           linear.Matrix // 3x3 restriction to protected generation carrier.
	Norm             float64
	OffDiagonalNorm  float64
	SkewResidual     float64
	DiagonalResidual float64
}

type Analysis struct {
	GenerationBreak generationbreak.Analysis
	Potential       higgspotential.Analysis
	Connection      connection.Analysis

	CarrierFrame        linear.Matrix // 56x3 protected contact carrier in Boolean coordinates.
	CarrierDimension    int
	CarrierGramResidual float64
	ContactKernelCount  int

	ActiveFrame        linear.Matrix // 56x4 active Higgs/contact carrier in Boolean coordinates.
	ActiveDimension    int
	ActiveGramResidual float64

	Operators            []Operator // restricted to protected carrier.
	NonzeroOperators     int
	MaxCurvatureNorm     float64
	MaxOffDiagonalNorm   float64
	MaxSkewResidual      float64
	MaxDiagonalResidual  float64
	OperatorSpanRank     int
	ClosureResidual      float64
	ClosureRelative      float64
	NonCommutingPairNorm float64

	ActiveOperators        []Operator // restricted to active Higgs/contact carrier.
	ActiveNonzeroOperators int
	ActiveMaxCurvatureNorm float64
	ActiveOperatorSpanRank int

	FullSO3CarrierFound       bool
	NonDiagonalMixingFound    bool
	SymmetricMassTextureFound bool
	CKMDerived                bool
	PMNSDerived               bool

	TruthStatement    string
	RemainingUnknowns []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		defaultValue, defaultErr = buildDefaultUncached()
	})
	return defaultValue, defaultErr
}

func buildDefaultUncached() (Analysis, error) {
	gb, err := generationbreak.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	hp, err := higgspotential.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	cn, err := connection.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(gb, hp, cn, 1e-8)
}

func Build(gb generationbreak.Analysis, hp higgspotential.Analysis, cn connection.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-8
	}
	carrier, active, kernelCount, gramResidual, activeGramResidual, err := contactCarrierFrames(hp, cn, eps)
	if err != nil {
		return Analysis{}, err
	}

	p := cn.Compression.BooleanComplementProjector
	q := cn.Compression.BooleanContactProjector
	ops := make([]Operator, 0)
	maxNorm := 0.0
	maxOffDiag := 0.0
	maxSkew := 0.0
	maxDiag := 0.0
	nonzero := 0

	generators := cn.Compression.BooleanGenerators
	for i := 0; i < len(generators); i++ {
		for j := i + 1; j < len(generators); j++ {
			curv, err := contactSideSecondFundamental(q, p, generators[i], generators[j])
			if err != nil {
				return Analysis{}, err
			}
			restricted, err := restrict(carrier, curv)
			if err != nil {
				return Analysis{}, err
			}
			op, err := summarizeOperator(fmt.Sprintf("[%d,%d]", i, j), restricted)
			if err != nil {
				return Analysis{}, err
			}
			ops = append(ops, op)
			if op.Norm > eps {
				nonzero++
			}
			if op.Norm > maxNorm {
				maxNorm = op.Norm
			}
			if op.OffDiagonalNorm > maxOffDiag {
				maxOffDiag = op.OffDiagonalNorm
			}
			if op.SkewResidual > maxSkew {
				maxSkew = op.SkewResidual
			}
			if op.DiagonalResidual > maxDiag {
				maxDiag = op.DiagonalResidual
			}
		}
	}

	activeOps := make([]Operator, 0)
	activeNonzero := 0
	activeMaxNorm := 0.0
	for i := 0; i < len(generators); i++ {
		for j := i + 1; j < len(generators); j++ {
			curv, err := contactSideSecondFundamental(q, p, generators[i], generators[j])
			if err != nil {
				return Analysis{}, err
			}
			restricted, err := restrict(active, curv)
			if err != nil {
				return Analysis{}, err
			}
			op, err := summarizeOperator(fmt.Sprintf("[%d,%d]", i, j), restricted)
			if err != nil {
				return Analysis{}, err
			}
			activeOps = append(activeOps, op)
			if op.Norm > eps {
				activeNonzero++
			}
			if op.Norm > activeMaxNorm {
				activeMaxNorm = op.Norm
			}
		}
	}
	activeSpanRank, _, err := operatorSpan(activeOps, eps)
	if err != nil {
		return Analysis{}, err
	}

	spanRank, frame, err := operatorSpan(ops, eps)
	if err != nil {
		return Analysis{}, err
	}
	closureAbs, closureRel, noncomm, err := closureDiagnostics(ops, frame, eps)
	if err != nil {
		return Analysis{}, err
	}

	fullSO3 := carrier.Cols() == 3 && spanRank == 3 && closureRel < 1e-7 && maxSkew < 1e-7 && maxNorm > eps
	nonDiag := maxOffDiag > eps && nonzero > 0

	return Analysis{
		GenerationBreak:           gb,
		Potential:                 hp,
		Connection:                cn,
		CarrierFrame:              carrier,
		CarrierDimension:          carrier.Cols(),
		CarrierGramResidual:       gramResidual,
		ContactKernelCount:        kernelCount,
		ActiveFrame:               active,
		ActiveDimension:           active.Cols(),
		ActiveGramResidual:        activeGramResidual,
		Operators:                 ops,
		NonzeroOperators:          nonzero,
		MaxCurvatureNorm:          maxNorm,
		MaxOffDiagonalNorm:        maxOffDiag,
		MaxSkewResidual:           maxSkew,
		MaxDiagonalResidual:       maxDiag,
		OperatorSpanRank:          spanRank,
		ClosureResidual:           closureAbs,
		ClosureRelative:           closureRel,
		NonCommutingPairNorm:      noncomm,
		ActiveOperators:           activeOps,
		ActiveNonzeroOperators:    activeNonzero,
		ActiveMaxCurvatureNorm:    activeMaxNorm,
		ActiveOperatorSpanRank:    activeSpanRank,
		FullSO3CarrierFound:       fullSO3,
		NonDiagonalMixingFound:    nonDiag,
		SymmetricMassTextureFound: false,
		CKMDerived:                false,
		PMNSDerived:               false,
		TruthStatement:            truthStatement(fullSO3, nonDiag),
		RemainingUnknowns: []string{
			"U-16D-YUKAWA-SCALE-BRIDGE: convert finite generation operators into physical Yukawa coupling strengths without fitting masses",
			"U-16E-SYMMETRIC-TEXTURE-FORMATION: build Hermitian/symmetric mass textures from the skew curvature generators and diagonal spurions",
			"U-16F-CKM-PMNS-COMMUTATOR: derive at least two non-commuting symmetric texture pairs before claiming CKM/PMNS",
			"U-17-BF-CURVATURE: implement the finite BF curvature operator and compare it with the second-fundamental generation curvature",
		},
	}, nil
}

func contactCarrierFrames(hp higgspotential.Analysis, cn connection.Analysis, eps float64) (protected linear.Matrix, active linear.Matrix, kernelCount int, protectedGramResidual float64, activeGramResidual float64, err error) {
	kFrame := cn.Compression.BooleanContactFrame // 56x7, orthonormal contact frame.
	kt := kFrame.Transpose()
	tmp, err := hp.Mixing.VacuumMixingOperator.Mul(kFrame)
	if err != nil {
		return linear.Matrix{}, linear.Matrix{}, 0, 0, 0, err
	}
	kOp, err := kt.Mul(tmp) // 7x7 contact-side mixing operator.
	if err != nil {
		return linear.Matrix{}, linear.Matrix{}, 0, 0, 0, err
	}
	eig, err := linear.SymmetricEigenJacobi(kOp, 1e-12, 0)
	if err != nil {
		return linear.Matrix{}, linear.Matrix{}, 0, 0, 0, err
	}
	zeroCols := make([]int, 0)
	activeCols := make([]int, 0)
	for i, v := range eig.Values {
		if math.Abs(v) <= eps {
			zeroCols = append(zeroCols, i)
		} else if v > eps {
			activeCols = append(activeCols, i)
		}
	}
	if len(zeroCols) != 3 {
		return linear.Matrix{}, linear.Matrix{}, len(zeroCols), 0, 0, fmt.Errorf("expected 3 protected contact eigenvectors, got %d", len(zeroCols))
	}
	if len(activeCols) != 4 {
		return linear.Matrix{}, linear.Matrix{}, len(zeroCols), 0, 0, fmt.Errorf("expected 4 active contact eigenvectors, got %d", len(activeCols))
	}
	protectedZ := eigenSubframe(eig.Vectors, zeroCols)
	activeZ := eigenSubframe(eig.Vectors, activeCols)
	protected, err = kFrame.Mul(protectedZ)
	if err != nil {
		return linear.Matrix{}, linear.Matrix{}, 0, 0, 0, err
	}
	active, err = kFrame.Mul(activeZ)
	if err != nil {
		return linear.Matrix{}, linear.Matrix{}, 0, 0, 0, err
	}
	protectedGramResidual, err = gramResidual(protected)
	if err != nil {
		return linear.Matrix{}, linear.Matrix{}, 0, 0, 0, err
	}
	activeGramResidual, err = gramResidual(active)
	if err != nil {
		return linear.Matrix{}, linear.Matrix{}, 0, 0, 0, err
	}
	return protected, active, len(zeroCols), protectedGramResidual, activeGramResidual, nil
}

func eigenSubframe(vectors linear.Matrix, cols []int) linear.Matrix {
	z := linear.NewMatrix(vectors.Rows(), len(cols))
	for newCol, oldCol := range cols {
		for r := 0; r < vectors.Rows(); r++ {
			z.Set(r, newCol, vectors.At(r, oldCol))
		}
	}
	return z
}

func gramResidual(frame linear.Matrix) (float64, error) {
	gram, err := frame.Transpose().Mul(frame)
	if err != nil {
		return 0, err
	}
	diff, err := gram.Sub(linear.Identity(frame.Cols()))
	if err != nil {
		return 0, err
	}
	return diff.FrobeniusNorm(), nil
}

func contactSideSecondFundamental(q, p, a, b linear.Matrix) (linear.Matrix, error) {
	aqb, err := triple(a, p, b)
	if err != nil {
		return linear.Matrix{}, err
	}
	qapbq, err := sandwich(q, aqb, q)
	if err != nil {
		return linear.Matrix{}, err
	}
	bpa, err := triple(b, p, a)
	if err != nil {
		return linear.Matrix{}, err
	}
	qbpaq, err := sandwich(q, bpa, q)
	if err != nil {
		return linear.Matrix{}, err
	}
	return qapbq.Sub(qbpaq)
}

func restrict(frame, op linear.Matrix) (linear.Matrix, error) {
	tmp, err := op.Mul(frame)
	if err != nil {
		return linear.Matrix{}, err
	}
	return frame.Transpose().Mul(tmp)
}

func summarizeOperator(pair string, m linear.Matrix) (Operator, error) {
	skew, err := m.Add(m.Transpose())
	if err != nil {
		return Operator{}, err
	}
	diag := 0.0
	off := 0.0
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			v := m.At(r, c)
			if r == c {
				diag += v * v
			} else {
				off += v * v
			}
		}
	}
	return Operator{
		Pair:             pair,
		Matrix:           m,
		Norm:             m.FrobeniusNorm(),
		OffDiagonalNorm:  math.Sqrt(off),
		SkewResidual:     skew.FrobeniusNorm(),
		DiagonalResidual: math.Sqrt(diag),
	}, nil
}

func operatorSpan(ops []Operator, eps float64) (int, linear.Matrix, error) {
	if len(ops) == 0 {
		return 0, linear.NewMatrix(0, 0), nil
	}
	rows := ops[0].Matrix.Rows() * ops[0].Matrix.Cols()
	cols := len(ops)
	v := linear.NewMatrix(rows, cols)
	for col, op := range ops {
		idx := 0
		for r := 0; r < op.Matrix.Rows(); r++ {
			for c := 0; c < op.Matrix.Cols(); c++ {
				v.Set(idx, col, op.Matrix.At(r, c))
				idx++
			}
		}
	}
	gram, err := v.Transpose().Mul(v)
	if err != nil {
		return 0, linear.Matrix{}, err
	}
	eig, err := linear.SymmetricEigenJacobi(gram, 1e-12, 0)
	if err != nil {
		return 0, linear.Matrix{}, err
	}
	rank := linear.RankFromEigenvalues(eig.Values, eps)
	frame := linear.NewMatrix(rows, rank)
	outCol := 0
	for i, value := range eig.Values {
		if math.Abs(value) <= eps {
			continue
		}
		norm := math.Sqrt(math.Abs(value))
		for r := 0; r < rows; r++ {
			sum := 0.0
			for c := 0; c < cols; c++ {
				sum += v.At(r, c) * eig.Vectors.At(c, i)
			}
			frame.Set(r, outCol, sum/norm)
		}
		outCol++
	}
	return rank, frame, nil
}

func closureDiagnostics(ops []Operator, frame linear.Matrix, eps float64) (float64, float64, float64, error) {
	if frame.Cols() == 0 {
		return 0, 0, 0, nil
	}
	projector, err := frame.Mul(frame.Transpose())
	if err != nil {
		return 0, 0, 0, err
	}
	maxAbs := 0.0
	maxRel := 0.0
	maxCom := 0.0
	for i := 0; i < len(ops); i++ {
		for j := i + 1; j < len(ops); j++ {
			br, err := commutator(ops[i].Matrix, ops[j].Matrix)
			if err != nil {
				return 0, 0, 0, err
			}
			n := br.FrobeniusNorm()
			if n > maxCom {
				maxCom = n
			}
			v := vectorize(br)
			pv, err := projector.Mul(v)
			if err != nil {
				return 0, 0, 0, err
			}
			diff, err := v.Sub(pv)
			if err != nil {
				return 0, 0, 0, err
			}
			abs := diff.FrobeniusNorm()
			rel := 0.0
			if n > eps {
				rel = abs / n
			}
			if abs > maxAbs {
				maxAbs = abs
			}
			if rel > maxRel {
				maxRel = rel
			}
		}
	}
	return maxAbs, maxRel, maxCom, nil
}

func vectorize(m linear.Matrix) linear.Matrix {
	v := linear.NewMatrix(m.Rows()*m.Cols(), 1)
	idx := 0
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			v.Set(idx, 0, m.At(r, c))
			idx++
		}
	}
	return v
}

func sandwich(left, middle, right linear.Matrix) (linear.Matrix, error) {
	lm, err := left.Mul(middle)
	if err != nil {
		return linear.Matrix{}, err
	}
	return lm.Mul(right)
}

func triple(a, b, c linear.Matrix) (linear.Matrix, error) {
	ab, err := a.Mul(b)
	if err != nil {
		return linear.Matrix{}, err
	}
	return ab.Mul(c)
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

func truthStatement(fullSO3, nonDiag bool) string {
	switch {
	case fullSO3:
		return "The contact-side second-fundamental curvature restricts to a closed three-dimensional skew algebra on the protected contact carrier. This supplies a finite generation-rotation source, but not yet symmetric Yukawa masses."
	case nonDiag:
		return "The contact-side second-fundamental curvature is nonzero and non-diagonal on the protected generation carrier, but it does not yet form a closed full so(3) carrier."
	default:
		return "The second-fundamental curvature is flat on the protected generation carrier in the current implementation. It acts on the active Higgs/contact directions instead, so generation mixing still requires an additional projection or BF curvature bridge."
	}
}

func FormatOperators(ops []Operator) string {
	xs := make([]string, 0, len(ops))
	for _, op := range ops {
		if op.Norm <= 1e-10 {
			continue
		}
		xs = append(xs, fmt.Sprintf("%s:norm=%.6g,off=%.6g", op.Pair, op.Norm, op.OffDiagonalNorm))
	}
	if len(xs) == 0 {
		return "[]"
	}
	return "[" + strings.Join(xs, "; ") + "]"
}
