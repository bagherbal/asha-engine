// Package gaugeeating implements Gate 85: finite scalar kinetic normalization /
// gauge-eating theorem search.
//
// Gate 84 constructed an abstract scalar covariant derivative and found the
// dimensionless W/Z/photon mass-matrix signature.  Gate 85 asks the stronger
// question: does this diagnostic become a finite gauge-eating theorem?  The
// gate therefore separates three levels:
//
//  1. exact representation diagnostics: the broken generators map the scalar
//     vacuum into a three-dimensional tangent/Goldstone image while Q_em
//     annihilates it;
//  2. kinetic diagnostics: the active scalar frame admits the canonical
//     Euclidean metric used by the mass-matrix computation;
//  3. bridge failures: no finite action has yet selected the scalar kinetic
//     normalization, gauge-field Hessian, physical couplings, or vacuum
//     orientation.
//
// The result is intentionally not labelled as a physical W/Z mass theorem.  It
// is a strengthened gauge-eating diagnostic: the correct null and broken-image
// structure appears, but the action-level normalizations are still open.
package gaugeeating

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/scalarcovariant"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type Analysis struct {
	ScalarCovariant scalarcovariant.Analysis

	ActiveRealDimension    int
	GaugeGeneratorCount    int
	BrokenGeneratorCount   int
	UnbrokenGeneratorCount int

	ScalarKineticMetric          linear.Matrix
	KineticTrace                 float64
	KineticDeterminant           float64
	KineticMetricPositive        bool
	KineticNormalizationSelected bool

	RadialDirection         []float64
	RadialNormResidual      float64
	BrokenImageGram         linear.Matrix
	BrokenImageRank         int
	BrokenImageMinEigen     float64
	BrokenImageMaxEigen     float64
	BrokenImageCondition    float64
	BrokenImagesIndependent bool

	EMNullNorm       float64
	PhotonNullVector []float64

	GoldstoneImageTheoremDiagnostic bool
	GaugeEatingCountDiagnostic      bool
	FiniteGaugeEatingTheoremDerived bool
	GaugeBosonMassMatrixDerived     bool
	PhysicalMassesDerived           bool

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
		sc, err := scalarcovariant.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(sc, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(sc scalarcovariant.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if sc.ActiveRealDimension != 4 {
		return Analysis{}, fmt.Errorf("gauge-eating audit expects a 4-real scalar frame, got %d", sc.ActiveRealDimension)
	}
	if len(sc.VacuumVector) != 4 || sc.VacuumRadius <= eps {
		return Analysis{}, fmt.Errorf("gauge-eating audit requires a nonzero 4D vacuum vector")
	}

	kin := linear.Identity(4)
	ktr, _ := kin.Trace()
	kdet := 1.0 // determinant of I_4

	radial := make([]float64, 4)
	for i, v := range sc.VacuumVector {
		radial[i] = v / sc.VacuumRadius
	}
	radialResidual := math.Abs(norm(radial) - 1)

	// In the equal-coupling diagnostic, the broken neutral generator is the
	// orthogonal complement to Q_em=T3+Yphi in the neutral span: Z=T3-Yphi.
	z, err := sc.T3.Sub(sc.YPhi)
	if err != nil {
		return Analysis{}, err
	}
	broken := []linear.Matrix{sc.T1, sc.T2, z}
	images := make([][]float64, len(broken))
	for i, g := range broken {
		images[i] = matVec(g, sc.VacuumVector)
	}
	gram := gram(images)
	eig, err := linear.SymmetricEigenJacobi(gram, eps, 200)
	if err != nil {
		return Analysis{}, err
	}
	rank := linear.RankFromEigenvalues(eig.Values, eps)
	minPositive := math.Inf(1)
	maxEigen := 0.0
	for _, v := range eig.Values {
		if v > eps && v < minPositive {
			minPositive = v
		}
		if v > maxEigen {
			maxEigen = v
		}
	}
	if math.IsInf(minPositive, 1) {
		minPositive = 0
	}
	cond := math.Inf(1)
	if minPositive > eps {
		cond = maxEigen / minPositive
	}

	emNorm := norm(matVec(sc.QEM, sc.VacuumVector))
	independent := rank == 3 && minPositive > eps
	countDiag := sc.ActiveRealDimension == 4 && rank == 3 && emNorm < eps
	truth := "Gate 85 strengthens the Gate 84 diagnostic: the chosen scalar vacuum has one radial direction, three independent broken-generator images, and an electromagnetic null generator. This is the finite Goldstone/gauge-eating image signature. It is still not a completed gauge-eating theorem because the Euclidean scalar kinetic metric, vacuum orientation, gauge-field Hessian, and physical couplings remain bridge-level rather than action-selected."

	return Analysis{
		ScalarCovariant: sc,

		ActiveRealDimension:    4,
		GaugeGeneratorCount:    4,
		BrokenGeneratorCount:   3,
		UnbrokenGeneratorCount: 1,

		ScalarKineticMetric:          kin,
		KineticTrace:                 ktr,
		KineticDeterminant:           kdet,
		KineticMetricPositive:        true,
		KineticNormalizationSelected: false,

		RadialDirection:         radial,
		RadialNormResidual:      radialResidual,
		BrokenImageGram:         gram,
		BrokenImageRank:         rank,
		BrokenImageMinEigen:     minPositive,
		BrokenImageMaxEigen:     maxEigen,
		BrokenImageCondition:    cond,
		BrokenImagesIndependent: independent,
		EMNullNorm:              emNorm,
		PhotonNullVector:        []float64{0, 0, 1, 1},

		GoldstoneImageTheoremDiagnostic: independent && emNorm < eps,
		GaugeEatingCountDiagnostic:      countDiag,
		FiniteGaugeEatingTheoremDerived: false,
		GaugeBosonMassMatrixDerived:     false,
		PhysicalMassesDerived:           false,

		TruthStatement:      truth,
		RecommendedNextGate: "Gate 86 — Scalar Vacuum Orientation / Finite Minimizer Search",
		RemainingUnknowns: []string{
			"U-18C6-SCALAR-KINETIC-ACTION: derive the scalar kinetic metric from finite contact/BF action instead of using the Euclidean frame metric diagnostically",
			"U-18C7-GAUGE-EATING-INTERTWINER: derive a canonical isometry from protected contact directions to broken generator images",
			"U-18C8-GAUGE-HESSIAN-COUPLINGS: derive gauge-field kinetic Hessian and couplings before physical W/Z masses",
			"U-19A1-VACUUM-ORIENTATION: derive phi0 as a finite minimizer rather than a unitary-gauge diagnostic choice",
		},
	}, nil
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

func gram(vs [][]float64) linear.Matrix {
	m := linear.NewMatrix(len(vs), len(vs))
	for i := range vs {
		for j := range vs {
			m.Set(i, j, dot(vs[i], vs[j]))
		}
	}
	return m
}

func dot(a, b []float64) float64 {
	s := 0.0
	for i := range a {
		s += a[i] * b[i]
	}
	return s
}

func norm(v []float64) float64 { return math.Sqrt(dot(v, v)) }

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

func FormatGram(m linear.Matrix) string {
	return fmt.Sprintf("[[%.10f, %.10f, %.10f], [%.10f, %.10f, %.10f], [%.10f, %.10f, %.10f]]",
		m.At(0, 0), m.At(0, 1), m.At(0, 2),
		m.At(1, 0), m.At(1, 1), m.At(1, 2),
		m.At(2, 0), m.At(2, 1), m.At(2, 2))
}
