package scalarscale

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ScaleBridgeSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-SCALAR-SCALE-SEARCH"
	const name = "scalar finite-to-physical scale bridge search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct scalar scale bridge audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "finite scalar anchors available", Passed: len(a.Anchors) >= 7, Detail: formatAnchors(a.Anchors)},
			{Name: "finite radius anchor", Passed: a.FiniteRadiusSquared > 0, Detail: fmt.Sprintf("r0²=%.10f, r0=%.10f", a.FiniteRadiusSquared, a.FiniteRadius)},
			{Name: "dimensionless radial curvature", Passed: a.DimensionlessRadialMassSq > 0, Detail: fmt.Sprintf("m_radial_hat²=%.10f, m_radial_hat=%.10f", a.DimensionlessRadialMassSq, a.DimensionlessRadialMass)},
			{Name: "B-sector gap anchor", Passed: a.BGap > 0, Detail: fmt.Sprintf("gap_B=%.10f; dimensionless spectral gap only", a.BGap)},
			{Name: "contact leakage anchor", Passed: a.ContactLeakageNormSquared > 0, Detail: fmt.Sprintf("L_BG²=%.10f, L_BG=%.10f; bare contact frustration, not Λ", a.ContactLeakageNormSquared, a.ContactLeakageNorm)},
			{Name: "scale-free ratios", Passed: a.GapToRadiusRatio > 0 && a.RadialToLeakageRatio > 0, Detail: fmt.Sprintf("gap/r0²=%.10f, L_BG²/r0²=%.10f, m_radial_hat²/L_BG²=%.10f, m_radial_hat²/gap=%.10f", a.GapToRadiusRatio, a.LeakageToRadiusRatio, a.RadialToLeakageRatio, a.CurvatureToGapRatio)},
			{Name: "dimensional-analysis firewall", Passed: !a.HasDimensionfulAnchor && a.OverallScaleFree, Detail: "all available quantities are dimensionless; an overall physical mass unit mu remains free"},
			{Name: "scale family exposed", Passed: a.ScaleFamily.FreeScaleSymbol == "mu", Detail: fmt.Sprintf("%s; %s; %s", a.ScaleFamily.VEVFormula, a.ScaleFamily.HiggsFormula, a.ScaleFamily.Explanation)},
			{Name: "electroweak scale derived", Passed: a.ElectroweakScaleDerived, Detail: "not derived; no comparison to v=246 GeV is allowed without a non-fitted unit bridge"},
			{Name: "Higgs mass bridge derived", Passed: a.HiggsMassBridgeDerived, Detail: "not derived; m_H cannot be compared to 125 GeV while mu is free"},
			{Name: "hidden observed scale insertion", Passed: !a.HiddenObservedScaleInserted, Detail: "no observed physical constants are used inside this bridge search"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("remaining unknowns: %v", a.RemainingUnknowns),
		}}
	}}
}

func formatAnchors(anchors []DimensionlessAnchor) string {
	out := ""
	for i, anchor := range anchors {
		if i > 0 {
			out += "; "
		}
		out += fmt.Sprintf("%s=%.10f", anchor.Name, anchor.Value)
	}
	return out
}
