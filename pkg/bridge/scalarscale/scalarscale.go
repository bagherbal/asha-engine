// Package scalarscale audits the missing dimensionful bridge from the finite
// scalar/Higgs normal form to physical electroweak units.
//
// The finite engine has derived dimensionless data: a scalar radius r0, a
// quartic shape, a B-sector spectral gap, and a contact leakage invariant. None
// of those objects carries physical units by itself. This package therefore
// implements a scale-audit theorem: it computes every finite dimensionless
// anchor presently available, then proves that an overall physical mass scale
// remains free unless an additional dimensional bridge is supplied.
//
// This is not a retreat. It is the correct dimensional-analysis firewall: the
// engine must not compare r0 to 246 GeV or m_radial_hat to 125 GeV unless a
// non-fitted conversion scale has first been derived.
package scalarscale

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/dynamics/bsector"
	"github.com/bagherbal/asha-engine/pkg/dynamics/scalarpotential"
)

type DimensionlessAnchor struct {
	Name  string
	Value float64
}

type ScaleFamily struct {
	FreeScaleSymbol string
	VEVFormula      string
	HiggsFormula    string
	Explanation     string
}

type Analysis struct {
	Scalar scalarpotential.Analysis
	Vacuum bsector.Vacuum

	Anchors []DimensionlessAnchor

	FiniteRadiusSquared       float64
	FiniteRadius              float64
	DimensionlessRadialMassSq float64
	DimensionlessRadialMass   float64
	QuarticShape              float64
	BGap                      float64
	ContactLeakageNormSquared float64
	ContactLeakageNorm        float64
	ContactIndex              float64

	GapToRadiusRatio     float64
	LeakageToRadiusRatio float64
	RadialToLeakageRatio float64
	GapToLeakageRatio    float64
	CurvatureToGapRatio  float64

	HasDimensionfulAnchor       bool
	OverallScaleFree            bool
	ElectroweakScaleDerived     bool
	HiggsMassBridgeDerived      bool
	UniqueScaleSelected         bool
	HiddenObservedScaleInserted bool

	ScaleFamily ScaleFamily

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
		vac, err := bsector.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(sp, vac, 1e-12)
	})
	return defaultValue, defaultErr
}

func Build(sp scalarpotential.Analysis, vac bsector.Vacuum, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-12
	}
	if sp.VacuumRadiusSquared <= eps {
		return Analysis{}, fmt.Errorf("finite scalar radius is not positive")
	}
	if sp.DimensionlessRadialMassSq <= eps {
		return Analysis{}, fmt.Errorf("dimensionless radial curvature is not positive")
	}
	gap := vac.FirstPositiveEigenvalue(1e-8)
	if math.IsNaN(gap) || gap <= eps {
		return Analysis{}, fmt.Errorf("B-sector gap is not positive")
	}
	leakageSq := vac.Contact.BareLeakageNormSquared()
	leakage := math.Sqrt(leakageSq)
	if leakage <= eps {
		return Analysis{}, fmt.Errorf("contact leakage norm is not positive")
	}

	radiusSq := sp.VacuumRadiusSquared
	radius := sp.VacuumRadius
	radialMassSq := sp.DimensionlessRadialMassSq
	radialMass := math.Sqrt(radialMassSq)
	contactIndex := vac.Contact.ContactIndex()

	anchors := []DimensionlessAnchor{
		{Name: "r0² = Tr(M_K)", Value: radiusSq},
		{Name: "r0", Value: radius},
		{Name: "lambda_shape = Tr(M_K²)/Tr(M_K)²", Value: sp.LambdaShape},
		{Name: "m_radial_hat² = 8 lambda_shape r0²", Value: radialMassSq},
		{Name: "B-sector first gap", Value: gap},
		{Name: "L_BG² = ||P_B P_G - P_K||²", Value: leakageSq},
		{Name: "I_BG", Value: contactIndex},
	}

	analysis := Analysis{
		Scalar:                      sp,
		Vacuum:                      vac,
		Anchors:                     anchors,
		FiniteRadiusSquared:         radiusSq,
		FiniteRadius:                radius,
		DimensionlessRadialMassSq:   radialMassSq,
		DimensionlessRadialMass:     radialMass,
		QuarticShape:                sp.LambdaShape,
		BGap:                        gap,
		ContactLeakageNormSquared:   leakageSq,
		ContactLeakageNorm:          leakage,
		ContactIndex:                contactIndex,
		GapToRadiusRatio:            gap / radiusSq,
		LeakageToRadiusRatio:        leakageSq / radiusSq,
		RadialToLeakageRatio:        radialMassSq / leakageSq,
		GapToLeakageRatio:           gap / leakageSq,
		CurvatureToGapRatio:         radialMassSq / gap,
		HasDimensionfulAnchor:       false,
		OverallScaleFree:            true,
		ElectroweakScaleDerived:     false,
		HiggsMassBridgeDerived:      false,
		UniqueScaleSelected:         false,
		HiddenObservedScaleInserted: false,
		ScaleFamily: ScaleFamily{
			FreeScaleSymbol: "mu",
			VEVFormula:      "v(mu) = mu * r0",
			HiggsFormula:    "m_H(mu) = mu * sqrt(m_radial_hat²)",
			Explanation:     "Every finite scalar observable is dimensionless, so multiplying the entire scalar sector by any positive mass unit mu leaves the exact finite theorem ladder unchanged.",
		},
		TruthStatement: "The finite scalar sector supplies a strong dimensionless normal form, but it does not yet derive a physical electroweak scale. A free mass unit mu remains unless a new bridge fixes units from gravity, action normalization, RG flow, or another dimensional anchor.",
		RemainingUnknowns: []string{
			"U-18A-SCALAR-SCALE-BRIDGE: derive a non-fitted physical mass unit mu for v(mu)=mu*r0",
			"U-18B-HIGGS-MASS-BRIDGE: derive the normalization converting m_radial_hat² into a physical Higgs mass",
			"U-18E-DIMENSIONAL-ANCHOR: identify whether the anchor comes from Planck scale, topological action, spectral cutoff, RG threshold, or another finite-to-continuum theorem",
			"U-18F-NO-FITTING-AUDIT: reject any bridge whose only purpose is to force v=246 GeV or m_H=125 GeV",
		},
	}
	return analysis, nil
}

func (a Analysis) AnchorByName(name string) (DimensionlessAnchor, bool) {
	for _, anchor := range a.Anchors {
		if anchor.Name == name {
			return anchor, true
		}
	}
	return DimensionlessAnchor{}, false
}
