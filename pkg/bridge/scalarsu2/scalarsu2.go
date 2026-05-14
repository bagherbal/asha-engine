// Package scalarsu2 audits whether the finite scalar/contact active frame
// already carries an SU(2)_L action.
//
// Gate 48 established a count-level resonance: four active real scalar/contact
// directions split into one radial plus three angular directions, matching the
// three electroweak broken directions.  This package tests the next stronger
// claim.  A four-real-dimensional frame can always host the realification of a
// complex SU(2) doublet, but that is not the same as deriving the action from
// the finite scalar spectrum itself.
//
// The active scalar response has pair spectrum (a,a,b,b).  That pair structure
// canonically supports two independent pair rotations, and the standard doublet
// representation can be written on the frame.  However, when a != b, the full
// SU(2) generators do not commute with the finite scalar response; only the
// diagonal T3-like pair rotation does.  Therefore this gate classifies the
// result as a bridge-level representation candidate, not as a completed finite
// scalar-contact SU(2) theorem.
package scalarsu2

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/dynamics/scalarpotential"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type Analysis struct {
	Scalar scalarpotential.Analysis

	ActiveRealDimension     int
	ComplexDoubletDimension int
	ActiveSpectrum          []float64
	PairDegenerate          bool
	PairSplit               float64

	ScalarResponse linear.Matrix
	T1             linear.Matrix
	T2             linear.Matrix
	T3             linear.Matrix

	SkewResidual             float64
	SU2ClosureResidual       float64
	ScalarResponseCommT1Norm float64
	ScalarResponseCommT2Norm float64
	ScalarResponseCommT3Norm float64
	MaxFullSU2CommNorm       float64

	PairRotationCommutantDimension int
	AbstractDoubletRepresentation  bool
	FullSU2SelectedByScalarData    bool
	U1PairRotationSelected         bool
	CanonicalComplexStructure      bool
	CovariantDerivativeDerived     bool
	GaugeEatingTheoremDerived      bool

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
		sp, err := scalarpotential.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(sp, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(sp scalarpotential.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if sp.ActiveRealDimension != 4 {
		return Analysis{}, fmt.Errorf("scalar SU(2) audit expects four active real directions, got %d", sp.ActiveRealDimension)
	}
	if len(sp.ActiveSpectrum) != 4 {
		return Analysis{}, fmt.Errorf("scalar SU(2) audit expects four active spectrum entries, got %d", len(sp.ActiveSpectrum))
	}
	spectrum := append([]float64(nil), sp.ActiveSpectrum...)
	response := linear.Diagonal(spectrum)
	t1, t2, t3 := standardRealDoubletGenerators()

	skewResidual := maxSkewResidual(t1, t2, t3)
	closureResidual, err := su2ClosureResidual(t1, t2, t3)
	if err != nil {
		return Analysis{}, err
	}
	c1, err := linear.Commutator(response, t1)
	if err != nil {
		return Analysis{}, err
	}
	c2, err := linear.Commutator(response, t2)
	if err != nil {
		return Analysis{}, err
	}
	c3, err := linear.Commutator(response, t3)
	if err != nil {
		return Analysis{}, err
	}
	c1n, c2n, c3n := c1.FrobeniusNorm(), c2.FrobeniusNorm(), c3.FrobeniusNorm()
	maxComm := math.Max(c1n, math.Max(c2n, c3n))
	pairSplit := math.Abs(spectrum[0] - spectrum[2])

	fullSelected := maxComm < eps
	u1Selected := c3n < eps && (c1n > eps || c2n > eps)
	abstract := skewResidual < eps && closureResidual < eps

	return Analysis{
		Scalar:                         sp,
		ActiveRealDimension:            sp.ActiveRealDimension,
		ComplexDoubletDimension:        sp.ComplexDoubletDimension,
		ActiveSpectrum:                 spectrum,
		PairDegenerate:                 sp.PairDegenerate,
		PairSplit:                      pairSplit,
		ScalarResponse:                 response,
		T1:                             t1,
		T2:                             t2,
		T3:                             t3,
		SkewResidual:                   skewResidual,
		SU2ClosureResidual:             closureResidual,
		ScalarResponseCommT1Norm:       c1n,
		ScalarResponseCommT2Norm:       c2n,
		ScalarResponseCommT3Norm:       c3n,
		MaxFullSU2CommNorm:             maxComm,
		PairRotationCommutantDimension: 2,
		AbstractDoubletRepresentation:  abstract,
		FullSU2SelectedByScalarData:    fullSelected,
		U1PairRotationSelected:         u1Selected,
		CanonicalComplexStructure:      false,
		CovariantDerivativeDerived:     false,
		GaugeEatingTheoremDerived:      false,
		TruthStatement:                 truth(abstract, fullSelected, u1Selected),
		RemainingUnknowns: []string{
			"U-19A-SCALAR-COMPLEX-STRUCTURE: derive a canonical complex/quaternionic structure on the four active contact directions, not merely an abstract realification",
			"U-19B-SCALAR-SU2-GEOMETRIC-ORIGIN: derive the SU(2)_L scalar action directly from Boolean/contact connection geometry",
			"U-19C-COVARIANT-DERIVATIVE: construct finite D_mu Phi and scalar kinetic normalization",
			"U-18C-GAUGE-EATING-THEOREM: map scalar angular directions to broken gauge generators by a canonical intertwiner",
		},
	}, nil
}

// standardRealDoubletGenerators are the real 4x4 representation of the
// anti-Hermitian su(2) generators -i sigma_a/2 acting on C^2, with real basis
// (Re z1, Im z1, Re z2, Im z2).
func standardRealDoubletGenerators() (linear.Matrix, linear.Matrix, linear.Matrix) {
	t1 := linear.NewMatrix(4, 4)
	t1.Set(0, 3, 0.5)
	t1.Set(1, 2, -0.5)
	t1.Set(2, 1, 0.5)
	t1.Set(3, 0, -0.5)

	t2 := linear.NewMatrix(4, 4)
	t2.Set(0, 2, -0.5)
	t2.Set(1, 3, -0.5)
	t2.Set(2, 0, 0.5)
	t2.Set(3, 1, 0.5)

	t3 := linear.NewMatrix(4, 4)
	t3.Set(0, 1, 0.5)
	t3.Set(1, 0, -0.5)
	t3.Set(2, 3, -0.5)
	t3.Set(3, 2, 0.5)
	return t1, t2, t3
}

func maxSkewResidual(ms ...linear.Matrix) float64 {
	max := 0.0
	for _, m := range ms {
		s, err := m.Add(m.Transpose())
		if err != nil {
			return math.Inf(1)
		}
		if n := s.FrobeniusNorm(); n > max {
			max = n
		}
	}
	return max
}

func su2ClosureResidual(t1, t2, t3 linear.Matrix) (float64, error) {
	r12, err := closureOne(t1, t2, t3)
	if err != nil {
		return 0, err
	}
	r23, err := closureOne(t2, t3, t1)
	if err != nil {
		return 0, err
	}
	r31, err := closureOne(t3, t1, t2)
	if err != nil {
		return 0, err
	}
	return math.Max(r12, math.Max(r23, r31)), nil
}

func closureOne(a, b, expected linear.Matrix) (float64, error) {
	c, err := linear.Commutator(a, b)
	if err != nil {
		return 0, err
	}
	d, err := c.Sub(expected)
	if err != nil {
		return 0, err
	}
	return d.FrobeniusNorm(), nil
}

func truth(abstract, fullSelected, u1Selected bool) string {
	switch {
	case abstract && fullSelected:
		return "The scalar/contact active frame carries a full finite SU(2) action selected by the scalar response."
	case abstract && u1Selected:
		return "The four active scalar/contact directions support an abstract SU(2) doublet, but the anisotropic finite scalar response selects only a commuting T3/U(1) pair-rotation direction; the full scalar SU(2)_L action remains a bridge theorem."
	case abstract:
		return "An abstract SU(2) doublet representation exists on the four active directions, but it is not selected by the current finite scalar data."
	default:
		return "The current scalar/contact active frame does not even support the standard abstract SU(2) doublet representation."
	}
}

func FormatSpectrum(values []float64) string {
	if len(values) == 0 {
		return "[]"
	}
	out := "["
	for i, v := range values {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%.10f", v)
	}
	return out + "]"
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
