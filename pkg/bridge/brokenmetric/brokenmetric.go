// Package brokenmetric implements Gate 92: broken-image metric / kinetic
// normalization audit.
//
// Gate 91 showed that, after quotienting the arbitrary protected O(3) frame,
// the quotient-safe protected-to-broken comparison retains a real broken-image
// metric.  That metric is anisotropic: the neutral broken image is four times
// the charged broken images in the current diagnostic normalization.
//
// Gate 92 asks whether that anisotropy is physical, a gauge/generator
// normalization artifact, or unresolved kinetic data.  The key observation is
// that the anisotropy is exactly removed by normalizing the neutral broken
// generator by 1/2.  This does not derive physical couplings; it shows that the
// current anisotropy is not yet an intrinsic physical prediction.  It belongs
// to the still-open scalar/gauge kinetic-normalization layer.
package brokenmetric

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/gaugeeating"
	"github.com/bagherbal/asha-engine/pkg/bridge/quotientedcorrespondence"
)

type Analysis struct {
	QuotientedCorrespondence quotientedcorrespondence.Analysis
	GaugeEating              gaugeeating.Analysis

	ChargedEigenvalue      float64
	NeutralEigenvalue      float64
	RawCondition           float64
	RawAnisotropic         bool
	NeutralToChargedRatio  float64
	NeutralNormFactor      float64
	NormalizedNeutralEigen float64
	IsotropizedEigenvalues []float64
	IsotropizedCondition   float64
	IsotropizationPossible bool
	IsotropizationExact    bool

	PhysicalAnisotropyDerived          bool
	GaugeNormalizationArtifactPossible bool
	ScalarKineticNormalizationSelected bool
	GaugeKineticNormalizationSelected  bool
	BrokenMetricPhysicalPrediction     bool
	GaugeEatingTheoremCompleted        bool

	QuotientSafeConclusion string
	RejectedClaims         []string
	RemainingUnknowns      []string
	TruthStatement         string
	RecommendedNextGate    string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		qc, err := quotientedcorrespondence.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		ge, err := gaugeeating.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(qc, ge, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(qc quotientedcorrespondence.Analysis, ge gaugeeating.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if qc.BrokenRank != 3 || ge.BrokenImageRank != 3 {
		return Analysis{}, fmt.Errorf("broken metric audit expects broken rank 3; got quotient=%d gaugeEating=%d", qc.BrokenRank, ge.BrokenImageRank)
	}
	charged := ge.BrokenImageMinEigen
	neutral := ge.BrokenImageMaxEigen
	if charged <= eps || neutral <= eps {
		return Analysis{}, fmt.Errorf("broken metric audit requires positive charged/neutral eigenvalues; got %.12g %.12g", charged, neutral)
	}
	ratio := neutral / charged
	factor := math.Sqrt(charged / neutral)
	normalizedNeutral := neutral * factor * factor
	iso := []float64{charged, charged, normalizedNeutral}
	minIso, maxIso := minMax(iso)
	condIso := maxIso / minIso
	exact := math.Abs(condIso-1) < 1e-8
	possible := math.Abs(factor-0.5) < 1e-8 || exact

	truth := "The broken-generator image metric is anisotropic in the raw diagnostic normalization, with neutral-to-charged ratio 4. This anisotropy is exactly removable by normalizing the neutral broken generator by 1/2, producing an isotropic quotient-safe broken metric. Therefore the current anisotropy is not yet a physical prediction; it is unresolved scalar/gauge kinetic and generator-normalization data. The finite gauge-eating bridge remains structurally correct at the count/image level but not yet action-normalized."

	return Analysis{
		QuotientedCorrespondence: qc,
		GaugeEating:              ge,

		ChargedEigenvalue:      charged,
		NeutralEigenvalue:      neutral,
		RawCondition:           ratio,
		RawAnisotropic:         math.Abs(ratio-1) > eps,
		NeutralToChargedRatio:  ratio,
		NeutralNormFactor:      factor,
		NormalizedNeutralEigen: normalizedNeutral,
		IsotropizedEigenvalues: iso,
		IsotropizedCondition:   condIso,
		IsotropizationPossible: possible,
		IsotropizationExact:    exact,

		PhysicalAnisotropyDerived:          false,
		GaugeNormalizationArtifactPossible: true,
		ScalarKineticNormalizationSelected: false,
		GaugeKineticNormalizationSelected:  false,
		BrokenMetricPhysicalPrediction:     false,
		GaugeEatingTheoremCompleted:        false,

		QuotientSafeConclusion: "only eigenvalue ratios/conditions are quotient-safe; component-wise broken/protected frame matching remains rejected",
		RejectedClaims: []string{
			"raw broken-image anisotropy is a physical W/Z mass prediction",
			"the neutral generator normalization is derived by the current finite action",
			"the protected-to-broken isometry follows from dimension matching",
			"the gauge-eating theorem is complete before scalar and gauge kinetic actions are selected",
		},
		RemainingUnknowns: []string{
			"U-18C7C3-BROKEN-METRIC-NORMALIZATION: derive the neutral generator normalization from a finite gauge kinetic Hessian",
			"U-18C6-SCALAR-KINETIC-ACTION: derive the scalar/contact kinetic metric from finite action data",
			"U-18C8-GAUGE-HESSIAN-COUPLINGS: derive gauge-field kinetic Hessian and couplings",
			"U-18C7C2-QUOTIENTED-INTERTWINER: derive a quotient-safe protected-to-broken map after normalization",
		},
		TruthStatement:      truth,
		RecommendedNextGate: "Gate 93 — Normalized Broken-Generator Basis / Gauge-Kinetic Candidate Search",
	}, nil
}

func minMax(xs []float64) (float64, float64) {
	mn := math.Inf(1)
	mx := math.Inf(-1)
	for _, x := range xs {
		if x < mn {
			mn = x
		}
		if x > mx {
			mx = x
		}
	}
	return mn, mx
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
