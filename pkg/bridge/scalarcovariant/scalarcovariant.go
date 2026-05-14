// Package scalarcovariant implements Gate 84: finite scalar covariant
// derivative / gauge-boson mass matrix search.
//
// Gates 48-50 established that the active scalar/contact sector has the right
// 4-real-dimensional Higgs-doublet size and can host an abstract SU(2) doublet
// representation, but the full scalar SU(2) action and U(1) kinetic Hessian are
// still bridge-level data.  Gate 83 then showed that physical gauge couplings
// cannot be derived until a finite gauge-field action is selected.
//
// This gate therefore builds the most conservative object now available: an
// abstract finite covariant-derivative template on the 4D scalar/contact frame,
// with generators {T1,T2,T3,Yphi}.  It computes the dimensionless mass matrix
// produced by a chosen vacuum orientation.  The result has the electroweak
// signature — two degenerate charged modes, one neutral massive mode, and one
// photon null direction — but it is still only a bridge diagnostic because the
// vacuum orientation, scalar kinetic normalization, gauge couplings, and finite
// gauge action Hessian are not yet derived.
package scalarcovariant

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/gaugehessian"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarsu2"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type Analysis struct {
	ScalarSU2    scalarsu2.Analysis
	GaugeHessian gaugehessian.Analysis

	ActiveRealDimension int
	VacuumRadiusSquared float64
	VacuumRadius        float64

	VacuumOrientationChosen bool
	VacuumOrientationName   string
	VacuumVector            []float64

	T1   linear.Matrix
	T2   linear.Matrix
	T3   linear.Matrix
	YPhi linear.Matrix
	QEM  linear.Matrix

	AbstractCovariantDerivativeTemplate bool
	GeneratorSkewResidual               float64
	EMAnnihilatesVacuumNorm             float64

	MassMatrix                linear.Matrix
	ChargedMassSquaredHat     float64
	ChargedDegeneracyResidual float64
	NeutralMassiveSquaredHat  float64
	PhotonMassSquaredHat      float64
	PhotonNullResidual        float64
	MassMatrixRank            int

	DimensionlessWZPhotonSignature          bool
	FiniteScalarKineticNormalizationDerived bool
	GaugeCouplingsDerived                   bool
	GaugeActionHessianDerived               bool
	PhysicalMassesDerived                   bool

	TruthStatement      string
	RecommendedNextGate string
	RemainingUnknowns   []string
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
		gh, err := gaugehessian.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(su2, gh, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(su2 scalarsu2.Analysis, gh gaugehessian.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if su2.ActiveRealDimension != 4 || !su2.AbstractDoubletRepresentation {
		return Analysis{}, fmt.Errorf("Gate 84 requires a four-real-dimensional abstract scalar doublet representation")
	}
	r2 := su2.Scalar.VacuumRadiusSquared
	if r2 <= 0 {
		return Analysis{}, fmt.Errorf("scalar vacuum radius must be positive")
	}
	r := math.Sqrt(r2)

	// Real basis: (Re z1, Im z1, Re z2, Im z2).  The diagnostic vacuum is the
	// lower-component real direction.  This is the usual unitary-gauge choice,
	// but Gate 84 records it as a convention because a canonical finite vacuum
	// orientation has not yet been derived.
	phi0 := []float64{0, 0, r, 0}

	yphi := scalarHyperchargeGenerator()
	qem, err := su2.T3.Add(yphi)
	if err != nil {
		return Analysis{}, err
	}

	gens := []linear.Matrix{su2.T1, su2.T2, su2.T3, yphi}
	mass, err := gaugeMassMatrix(gens, phi0)
	if err != nil {
		return Analysis{}, err
	}

	emNorm := vectorNorm(matVec(qem, phi0))
	skew := maxSkew(su2.T1, su2.T2, su2.T3, yphi)

	charged1 := mass.At(0, 0)
	charged2 := mass.At(1, 1)
	chargedResidual := math.Abs(charged1 - charged2)
	// Neutral block [[a,-a],[-a,a]] with a=r²/4 in the equal-coupling diagnostic.
	neutralMassive := mass.At(2, 2) + mass.At(3, 3)
	photonResidual := math.Abs(mass.At(2, 2)+mass.At(2, 3)) + math.Abs(mass.At(3, 2)+mass.At(3, 3))
	photonMass := 0.0
	rank := 0
	for _, x := range []float64{charged1, charged2, neutralMassive, photonMass} {
		if math.Abs(x) > eps {
			rank++
		}
	}
	signature := charged1 > eps && charged2 > eps && chargedResidual < eps && neutralMassive > eps && photonResidual < eps && rank == 3

	truth := "Gate 84 constructs an abstract finite scalar covariant-derivative template and its dimensionless gauge-boson mass-matrix diagnostic. The W/Z/photon signature appears in the bridge representation: two degenerate charged modes, one neutral massive mode, and one electromagnetic null direction. This is not yet a physical mass theorem because the scalar vacuum orientation, scalar kinetic normalization, gauge couplings, and finite gauge-field Hessian remain unselected by the finite action."

	return Analysis{
		ScalarSU2:               su2,
		GaugeHessian:            gh,
		ActiveRealDimension:     su2.ActiveRealDimension,
		VacuumRadiusSquared:     r2,
		VacuumRadius:            r,
		VacuumOrientationChosen: false,
		VacuumOrientationName:   "lower-component real unitary-gauge diagnostic",
		VacuumVector:            phi0,
		T1:                      su2.T1, T2: su2.T2, T3: su2.T3, YPhi: yphi, QEM: qem,
		AbstractCovariantDerivativeTemplate:     true,
		GeneratorSkewResidual:                   skew,
		EMAnnihilatesVacuumNorm:                 emNorm,
		MassMatrix:                              mass,
		ChargedMassSquaredHat:                   charged1,
		ChargedDegeneracyResidual:               chargedResidual,
		NeutralMassiveSquaredHat:                neutralMassive,
		PhotonMassSquaredHat:                    photonMass,
		PhotonNullResidual:                      photonResidual,
		MassMatrixRank:                          rank,
		DimensionlessWZPhotonSignature:          signature,
		FiniteScalarKineticNormalizationDerived: false,
		GaugeCouplingsDerived:                   false,
		GaugeActionHessianDerived:               gh.HessianSelected,
		PhysicalMassesDerived:                   false,
		TruthStatement:                          truth,
		RecommendedNextGate:                     "Gate 85 — Finite Scalar Kinetic Normalization / Gauge-Eating Theorem Search",
		RemainingUnknowns: []string{
			"U-18C3-SCALAR-COVARIANT-DERIVATIVE: derive DΦ from finite contact geometry rather than an abstract representation template",
			"U-18C4-GAUGE-BOSON-MASS-MATRIX: derive scalar kinetic normalization and gauge couplings before interpreting W/Z masses",
			"U-18C5-PHOTON-NULL-THEOREM: connect Q=T3+Y to a finite gauge-field Hessian and unbroken U(1)_em field",
			"U-19A1-VACUUM-ORIENTATION: derive the scalar vacuum direction from finite contact/BF dynamics instead of unitary-gauge convention",
		},
	}, nil
}

// scalarHyperchargeGenerator is the real anti-Hermitian U(1) action with
// scalar hypercharge +1/2 on both complex components.  In the chosen real basis
// it is +1/2 times the same phase rotation on both complex planes.
func scalarHyperchargeGenerator() linear.Matrix {
	y := linear.NewMatrix(4, 4)
	y.Set(0, 1, 0.5)
	y.Set(1, 0, -0.5)
	y.Set(2, 3, 0.5)
	y.Set(3, 2, -0.5)
	return y
}

func gaugeMassMatrix(gens []linear.Matrix, phi []float64) (linear.Matrix, error) {
	n := len(gens)
	out := linear.NewMatrix(n, n)
	images := make([][]float64, n)
	for i, g := range gens {
		images[i] = matVec(g, phi)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			out.Set(i, j, dot(images[i], images[j]))
		}
	}
	return out, nil
}

func matVec(m linear.Matrix, v []float64) []float64 {
	out := make([]float64, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		sum := 0.0
		for c := 0; c < m.Cols(); c++ {
			sum += m.At(r, c) * v[c]
		}
		out[r] = sum
	}
	return out
}
func dot(a, b []float64) float64 {
	s := 0.0
	for i := range a {
		s += a[i] * b[i]
	}
	return s
}
func vectorNorm(v []float64) float64 { return math.Sqrt(dot(v, v)) }
func maxSkew(ms ...linear.Matrix) float64 {
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

func FormatVector(v []float64) string {
	out := "["
	for i, x := range v {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%.10f", x)
	}
	return out + "]"
}

func FormatMassMatrix(m linear.Matrix) string {
	return fmt.Sprintf("[[%.10f, %.10f, %.10f, %.10f], [%.10f, %.10f, %.10f, %.10f], [%.10f, %.10f, %.10f, %.10f], [%.10f, %.10f, %.10f, %.10f]]",
		m.At(0, 0), m.At(0, 1), m.At(0, 2), m.At(0, 3),
		m.At(1, 0), m.At(1, 1), m.At(1, 2), m.At(1, 3),
		m.At(2, 0), m.At(2, 1), m.At(2, 2), m.At(2, 3),
		m.At(3, 0), m.At(3, 1), m.At(3, 2), m.At(3, 3))
}
