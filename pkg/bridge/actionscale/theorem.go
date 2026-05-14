package actionscale

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ActionNormalizationScaleBridgeTheorem() theorem.Theorem {
	const id = "BRIDGE-ACTION-NORMALIZATION-SCALE"
	const name = "gravity/action-normalization scale bridge audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct action-normalization scale audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "unit contact index anchor", Passed: a.UnitIndexAvailable, Detail: fmt.Sprintf("I_BG=%.10f", a.ContactIndex)},
			{Name: "dimensionless topological action seal", Passed: a.DimensionlessActionDerived && a.TopologicalActionSeal > 0, Detail: fmt.Sprintf("S_top=8π²I_BG=%.10f; exp(-S_top)=%.3e", a.TopologicalActionSeal, a.ActionWeight)},
			{Name: "finite scalar anchors remain available", Passed: a.RadiusSquared > 0 && a.RadialMassSq > 0, Detail: fmt.Sprintf("r0²=%.10f, m_radial_hat²=%.10f", a.RadiusSquared, a.RadialMassSq)},
			{Name: "action-to-finite ratios", Passed: a.ActionToRadiusRatio > 0 && a.ActionToGapRatio > 0, Detail: fmt.Sprintf("S_top/r0²=%.10f, S_top/gap=%.10f, S_top/L_BG²=%.10f", a.ActionToRadiusRatio, a.ActionToGapRatio, a.ActionToLeakageRatio)},
			{Name: "continuum index bridge", Passed: a.ContinuumIndexBridgeDerived, Detail: "not yet derived; finite I_BG=1 is not automatically a continuum instanton theorem"},
			{Name: "coupling normalization", Passed: !a.CouplingNormalizationOpen, Detail: "open; S_top may normalize a coupling/action weight, but the coupling bridge is not yet fixed"},
			{Name: "dimensionful mass unit", Passed: a.DimensionfulUnitDerived, Detail: "not derived; a dimensionless action cannot by itself choose mu in v(mu)=mu*r0"},
			{Name: "gravity scale derived", Passed: a.GravityScaleDerived, Detail: "not derived; no Planck-scale comparison is allowed without a mass/length unit bridge"},
			{Name: "scalar scale fixed", Passed: a.ScalarScaleFixed, Detail: "not fixed; electroweak vev and Higgs mass remain forbidden comparisons"},
			{Name: "hidden observed scale insertion", Passed: !a.HiddenObservedScaleInserted, Detail: "no observed physical scale was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("remaining unknowns: %v", a.RemainingUnknowns),
		}}
	}}
}
