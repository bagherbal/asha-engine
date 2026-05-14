// Package scalarcomplex audits whether the active scalar/contact frame carries
// a canonical complex or quaternionic structure.
//
// Gate 49 showed that the four active scalar/contact directions can host an
// abstract SU(2) doublet, but the finite scalar response S_Phi=(a,a,b,b)
// preserves only a pair-rotation subgroup.  This package asks the sharper
// question: does the pair structure itself select the complex/quaternionic data
// needed to turn the bridge representation into a finite theorem?
//
// The answer is deliberately split.  The pair spectrum supports a commuting
// complex-structure candidate J on the two active planes: J^2=-I, J^T=-J, and
// [S_Phi,J]=0.  But the signs/orientations of the two planes are not fixed by
// the scalar response alone, and a full quaternionic triple is not selected
// because the other two quaternionic generators mix the unequal scalar pairs.
package scalarcomplex

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/scalarsu2"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type Analysis struct {
	ScalarSU2 scalarsu2.Analysis

	ActiveRealDimension    int
	PairDegenerate         bool
	PairSplit              float64
	PairPlaneCount         int
	PairOrientationChoices int

	ScalarResponse linear.Matrix
	J              linear.Matrix
	I              linear.Matrix
	K              linear.Matrix

	ComplexSquareResidual     float64
	ComplexSkewResidual       float64
	ComplexCommutesWithScalar float64

	QuaternionicSquareResidual  float64
	QuaternionicClosureResidual float64
	QuaternionicCommIWithScalar float64
	QuaternionicCommJWithScalar float64
	QuaternionicCommKWithScalar float64
	MaxQuaternionicScalarComm   float64

	PairCompatibleComplexAvailable bool
	CanonicalComplexDerived        bool
	QuaternionicTripleAvailable    bool
	QuaternionicTripleSelected     bool
	FullScalarSU2Recovered         bool

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
		su2, err := scalarsu2.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(su2, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(su2 scalarsu2.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if su2.ActiveRealDimension != 4 {
		return Analysis{}, fmt.Errorf("scalar complex audit expects four active real directions, got %d", su2.ActiveRealDimension)
	}
	if !su2.PairDegenerate {
		return Analysis{}, fmt.Errorf("scalar complex audit requires a pair-degenerate active spectrum")
	}

	// The real SU(2) generators in Gate 49 are normalized as T_a^2=-1/4.
	// Multiplying by 2 gives honest complex/quaternionic candidates with square -I.
	I := su2.T1.Scale(2)
	J := su2.T3.Scale(2) // the pair-rotation generator that commutes with S_Phi
	K := su2.T2.Scale(2)
	id := linear.Identity(4)

	complexSquare := squareMinusNegativeIdentity(J, id)
	complexSkew := skewResidual(J)
	complexComm := commNorm(su2.ScalarResponse, J)

	qSquare := max3(squareMinusNegativeIdentity(I, id), squareMinusNegativeIdentity(J, id), squareMinusNegativeIdentity(K, id))
	qClosure := quaternionicClosureResidual(I, J, K)
	cI := commNorm(su2.ScalarResponse, I)
	cJ := commNorm(su2.ScalarResponse, J)
	cK := commNorm(su2.ScalarResponse, K)
	maxQComm := max3(cI, cJ, cK)

	pairComplex := complexSquare < eps && complexSkew < eps && complexComm < eps
	qAvailable := qSquare < eps && qClosure < eps
	qSelected := qAvailable && maxQComm < eps
	canonicalComplex := pairComplex && false // orientation/sign of pair planes is still not fixed by finite scalar data alone.

	return Analysis{
		ScalarSU2:                      su2,
		ActiveRealDimension:            su2.ActiveRealDimension,
		PairDegenerate:                 su2.PairDegenerate,
		PairSplit:                      su2.PairSplit,
		PairPlaneCount:                 2,
		PairOrientationChoices:         4,
		ScalarResponse:                 su2.ScalarResponse,
		J:                              J,
		I:                              I,
		K:                              K,
		ComplexSquareResidual:          complexSquare,
		ComplexSkewResidual:            complexSkew,
		ComplexCommutesWithScalar:      complexComm,
		QuaternionicSquareResidual:     qSquare,
		QuaternionicClosureResidual:    qClosure,
		QuaternionicCommIWithScalar:    cI,
		QuaternionicCommJWithScalar:    cJ,
		QuaternionicCommKWithScalar:    cK,
		MaxQuaternionicScalarComm:      maxQComm,
		PairCompatibleComplexAvailable: pairComplex,
		CanonicalComplexDerived:        canonicalComplex,
		QuaternionicTripleAvailable:    qAvailable,
		QuaternionicTripleSelected:     qSelected,
		FullScalarSU2Recovered:         qSelected,
		TruthStatement:                 truth(pairComplex, canonicalComplex, qAvailable, qSelected),
		RemainingUnknowns: []string{
			"U-19A1-PAIR-ORIENTATION: derive the signs/orientations of the two active scalar planes from finite contact geometry, not from basis convention",
			"U-19A2-QUATERNIONIC-SELECTION: derive the two scalar generators that mix the unequal active pairs, or prove the anisotropic scalar response intentionally breaks them",
			"U-19B-SCALAR-SU2-GEOMETRIC-ORIGIN: connect the scalar complex/quaternionic data back to Boolean/contact connection geometry",
			"U-19C-COVARIANT-DERIVATIVE: construct finite D_mu Phi and scalar kinetic normalization",
		},
	}, nil
}

func squareMinusNegativeIdentity(m, id linear.Matrix) float64 {
	mm, err := m.Mul(m)
	if err != nil {
		return math.Inf(1)
	}
	sum, err := mm.Add(id)
	if err != nil {
		return math.Inf(1)
	}
	return sum.FrobeniusNorm()
}

func skewResidual(m linear.Matrix) float64 {
	s, err := m.Add(m.Transpose())
	if err != nil {
		return math.Inf(1)
	}
	return s.FrobeniusNorm()
}

func commNorm(a, b linear.Matrix) float64 {
	c, err := linear.Commutator(a, b)
	if err != nil {
		return math.Inf(1)
	}
	return c.FrobeniusNorm()
}

func quaternionicClosureResidual(I, J, K linear.Matrix) float64 {
	// With this convention the triple satisfies [I,J] = -2K, [J,K] = -2I, [K,I] = -2J.
	r1 := closureResidual(I, J, K.Scale(-2))
	r2 := closureResidual(J, K, I.Scale(-2))
	r3 := closureResidual(K, I, J.Scale(-2))
	return max3(r1, r2, r3)
}

func closureResidual(a, b, expected linear.Matrix) float64 {
	c, err := linear.Commutator(a, b)
	if err != nil {
		return math.Inf(1)
	}
	d, err := c.Sub(expected)
	if err != nil {
		return math.Inf(1)
	}
	return d.FrobeniusNorm()
}

func max3(a, b, c float64) float64 { return math.Max(a, math.Max(b, c)) }

func truth(pairComplex, canonicalComplex, qAvailable, qSelected bool) string {
	switch {
	case qSelected:
		return "The active scalar/contact frame carries a full quaternionic structure selected by the finite scalar response."
	case pairComplex && qAvailable:
		return "The active scalar/contact frame supports a pair-compatible complex structure and an abstract quaternionic triple, but the anisotropic scalar response selects only the commuting complex direction; the full quaternionic/SU(2) scalar theorem remains open."
	case pairComplex:
		return "The active scalar/contact pair spectrum supports a commuting complex-structure candidate, but its orientation is not canonical and no full quaternionic structure is selected."
	case canonicalComplex:
		return "A canonical complex structure is derived, but its quaternionic completion remains open."
	default:
		return "The current active scalar/contact data does not select a complex or quaternionic structure."
	}
}

func FormatUnknowns(values []string) string {
	if len(values) == 0 {
		return "none"
	}
	out := ""
	for i, v := range values {
		if i > 0 {
			out += "; "
		}
		out += v
	}
	return out
}
