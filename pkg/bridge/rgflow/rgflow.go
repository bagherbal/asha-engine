// Package rgflow audits the renormalization-group bridge required after the
// finite electroweak projection has produced a boundary-candidate normalization.
//
// The previous electroweak projection gate derives the representation identity
// Q=T3_L+Y, the hypercharge normalization k_Y=5/3, and the equal-normalized-
// coupling boundary candidate sin^2(theta)=3/8. None of that is yet a
// low-energy prediction. To reach a physical weak angle or alpha_em one must
// also derive the gauge kinetic normalization, a boundary scale, beta
// coefficients, and threshold/matching data.
//
// This package exposes the formal flow variables and rejects any hidden use of
// observed low-energy couplings.
package rgflow

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/ewprojection"
)

type Analysis struct {
	Electroweak ewprojection.Analysis

	HyperchargeNormalizationKY float64
	BoundarySin2Candidate      float64

	// NormalizedBoundaryRelation is the finite representation-level relation,
	// not a physical low-energy statement.
	NormalizedBoundaryRelation string

	// Formal one-loop placeholders. The engine does not yet derive b1, b2, the
	// boundary coupling gStar, or the log interval L=ln(M*/mu). Therefore they
	// remain symbolic bridge variables.
	FormalFlowEquation string
	RequiredVariables  []string

	BoundaryCouplingFree     bool
	LogScaleIntervalFree     bool
	BetaCoefficientsDerived  bool
	ThresholdSpectrumDerived bool
	GaugeKineticDerived      bool
	BoundaryScaleDerived     bool
	RGEvolutionDetermined    bool
	PhysicalWeakAngleDerived bool
	FineStructureDerived     bool
	HiddenObservedInputUsed  bool

	// Diagnostic only: if one additionally assumes no running and unit normalized
	// boundary coupling g_*^2=1, then e^2=g_*^2/(1+k_Y). This is deliberately
	// labelled non-physical because both assumptions are not derived.
	UnitNoRunningElectromagneticCouplingSq float64
	UnitNoRunningAlpha                     float64
	UnitNoRunningInverseAlpha              float64
	UnitNoRunningDiagnosticPhysical        bool

	MinimumMissingData []string
	TruthStatement     string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		ew, err := ewprojection.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(ew, 1e-12)
	})
	return defaultValue, defaultErr
}

func Build(ew ewprojection.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-12
	}
	if !ew.EqualNormalizedCouplingBoundaryCandidate || ew.HyperchargeNormalizationKY <= eps {
		return Analysis{}, fmt.Errorf("electroweak projection boundary candidate must be available")
	}

	kY := ew.HyperchargeNormalizationKY
	sin2 := ew.EqualNormalizedCouplingBoundarySin2

	// Pure diagnostic under two additional non-derived assumptions:
	//  1. no RG running from the boundary to the evaluation scale;
	//  2. unit normalized boundary coupling g_*^2=1.
	// Then g_Y^2=g_*^2/kY and e^2=g_2^2 g_Y^2/(g_2^2+g_Y^2)=1/(1+kY).
	e2Diag := 1 / (1 + kY)
	alphaDiag := e2Diag / (4 * math.Pi)
	invAlphaDiag := 1 / alphaDiag

	return Analysis{
		Electroweak:                            ew,
		HyperchargeNormalizationKY:             kY,
		BoundarySin2Candidate:                  sin2,
		NormalizedBoundaryRelation:             "finite charge table gives k_Y=5/3; equal normalized couplings give sin²θ*=1/(1+k_Y)=3/8",
		FormalFlowEquation:                     "1/g_i²(μ)=1/g_i²(M*)+B_i·ln(M*/μ), where B_i encodes beta coefficients, thresholds, and finite-to-continuum kinetic normalization",
		RequiredVariables:                      []string{"g_*²", "M*/μ", "B_1", "B_2", "threshold matching", "finite kinetic normalization"},
		BoundaryCouplingFree:                   true,
		LogScaleIntervalFree:                   true,
		BetaCoefficientsDerived:                false,
		ThresholdSpectrumDerived:               false,
		GaugeKineticDerived:                    ew.GaugeKineticNormalizationDerived,
		BoundaryScaleDerived:                   ew.RGBoundaryScaleDerived,
		RGEvolutionDetermined:                  false,
		PhysicalWeakAngleDerived:               false,
		FineStructureDerived:                   false,
		HiddenObservedInputUsed:                false,
		UnitNoRunningElectromagneticCouplingSq: e2Diag,
		UnitNoRunningAlpha:                     alphaDiag,
		UnitNoRunningInverseAlpha:              invAlphaDiag,
		UnitNoRunningDiagnosticPhysical:        false,
		MinimumMissingData: []string{
			"derive finite gauge kinetic terms for SU(2)_L and U(1)_Y, not only charge-trace normalization",
			"derive or justify the RG boundary scale M* from the finite geometry or continuum bridge",
			"derive beta coefficients from the full finite particle/threshold spectrum rather than importing a Standard Model table",
			"derive threshold matching between finite heavy modes, Higgs/contact modes, and low-energy fields",
			"derive the normalized boundary coupling g_*² or an equivalent action normalization",
		},
		TruthStatement: "The finite electroweak projection gives a strong boundary-candidate normalization, k_Y=5/3 and sin²θ*=3/8 under equal normalized couplings. RG flow to a physical weak angle or alpha_em is not determined because the boundary coupling, boundary scale, beta coefficients, thresholds, and finite kinetic normalization are still independent bridge data.",
	}, nil
}
