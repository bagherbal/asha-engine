package topologicalnormalization

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func SpectralActionTopologicalNormalizationTheorem() theorem.Theorem {
	const id = "BRIDGE-SPECTRAL-ACTION-TOPOLOGICAL-NORMALIZATION"
	const name = "spectral-action normalization from topological action seal"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build topological normalization audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "closed gauge ratio and sealed mass route are inherited", Passed: a.Input.GaugeRatioClosed && a.Input.WeakAngleSeedClosed && a.Input.MassGenerationSealed && !a.Input.UsesObservedInput, Detail: FormatInput(a.Input)},
			{Name: "topological action seal is available", Passed: a.Input.TopologicalSealAvailable && a.Input.ContactIndex > 0 && a.Input.TopologicalActionSeal > 0, Detail: fmt.Sprintf("I_BG=%.12g, S_top=%.12g", a.Input.ContactIndex, a.Input.TopologicalActionSeal)},
			{Name: "conditional instanton matching computes an absolute boundary coupling", Passed: a.Matching.ConditionalMatchingAvailable && close(a.Matching.ConditionalUInverseGStar, 1, 1e-10) && close(a.Matching.ConditionalGStarSquared, 1, 1e-10), Detail: FormatMatching(a.Matching)},
			{Name: "strict instanton matching bridge is not yet derived", Passed: !a.Matching.CanonicalStrictMatchingDerived && a.Matching.RequiresContinuumIndexBridge && a.Matching.RequiresTraceKineticBridge && !a.Matching.ContinuumIndexBridgeDerived && !a.Matching.TraceKineticBridgeDerived && !a.Matching.TopologicalSealAloneSufficient, Detail: FormatMatching(a.Matching)},
			{Name: "spectral-action prefactor is convention-dependent while boundary ratio is stable", Passed: len(a.Conventions) == 2 && a.Conventions[0].SameBoundaryPhysics && a.Conventions[1].SameBoundaryPhysics && a.Conventions[0].ConditionalF0 != a.Conventions[1].ConditionalF0, Detail: FormatConventions(a.Conventions)},
			{Name: "strict nullity does not reduce", Passed: !a.Firewall.StrictAbsoluteUDerived && !a.Firewall.StrictF0Derived && a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3, Detail: FormatFirewall(a.Firewall)},
			{Name: "conditional branch would reduce nullity to two", Passed: a.Firewall.ConditionalAbsoluteUAvailable && a.Firewall.ConditionalNullityAfter == 2 && len(a.Firewall.ConditionalRemainingUnknowns) == 2, Detail: fmt.Sprintf("conditional remaining unknowns: %v", a.Firewall.ConditionalRemainingUnknowns)},
			{Name: "physical constants remain sealed", Passed: !a.Firewall.BoundaryScaleDerived && !a.Firewall.ThresholdCorrectionsDerived && !a.Firewall.PhysicalCouplingsDerived && !a.Firewall.PhysicalFineStructureDerived && !a.Firewall.PhysicalMassesDerived && !a.Firewall.HiddenObservedInputUsed, Detail: FormatFirewall(a.Firewall)},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("strict remaining unknowns: %v", a.Firewall.RemainingStrictUnknowns),
			"Next gate should derive or reject the finite-to-continuum instanton trace-normalization bridge rather than insert an observed coupling.",
		}}
	}}
}
