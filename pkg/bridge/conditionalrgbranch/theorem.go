package conditionalrgbranch

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ConditionalRGBoundaryScaleSolvabilityAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-CONDITIONAL-RG-U1-BOUNDARY-SCALE-SOLVABILITY"
	const name = "conditional RG boundary-scale solvability under quarantined u=1 branch"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build conditional RG branch audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "quarantined u=1 branch is inherited", Passed: a.Input.Gate175ConditionalBranchAvailable && !a.Input.Gate175StrictAbsoluteUDerived && a.Input.RelativeGaugeRatioClosed && a.Input.WeakAngleSeedClosed && a.Input.ConditionalUInverseGStar == 1 && !a.Input.UsesObservedInputForDerivation, Detail: FormatInput(a.Input)},
			{Name: "observed ledger is comparison-only", Passed: a.Observed.QuarantinedComparison && !a.Observed.UsedForDerivation && a.Input.ObservedComparisonQuarantined, Detail: FormatObserved(a.Observed)},
			{Name: "conditional boundary point reproduces the finite seed", Passed: a.BoundaryPoint.LogIntervalL == 0 && a.BoundaryPoint.PositiveKinetic && a.BoundaryPoint.AlphaEMInverse > 0 && a.BoundaryPoint.Sin2 > 0, Detail: FormatPoint(a.BoundaryPoint)},
			{Name: "single-observable fits do not give a simultaneous viable M_Z point", Passed: len(a.Fits) == 4 && !a.Firewall.AnyObservedFitViable && a.Firewall.ConditionalBranchRejectedByMZ, Detail: FormatFits(a.Fits)},
			{Name: "ratio-only running check is independent of u but fails comparison", Passed: a.Ratio.IndependentOfU && a.Ratio.UsesGUTNormalizedAlpha1 && a.Ratio.ObservedComparisonOnly && !a.Ratio.LIntervalsAgree && !a.Firewall.RatioCheckPasses, Detail: FormatRatio(a.Ratio)},
			{Name: "threshold and normalization firewalls remain closed", Passed: !a.Firewall.ThresholdCorrectionsIncluded && a.Firewall.StrictUStillOpen && !a.Firewall.BoundaryScaleDerivedStrict && !a.Firewall.BoundaryScaleDerivedConditional, Detail: FormatFirewall(a.Firewall)},
			{Name: "strict nullity and physical constants remain sealed", Passed: a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3 && a.Firewall.ConditionalNullityBefore == 2 && a.Firewall.ConditionalNullityAfter == 2 && !a.Firewall.PhysicalConstantsDerived && !a.Firewall.HiddenObservedInputUsedForDerivation, Detail: FormatFirewall(a.Firewall)},
		}, Notes: []string{
			a.TruthStatement,
			"The conditional u=1 branch is not promoted to a physical coupling theorem; it is tested and rejected as an unthresholded M_Z phenomenological branch.",
			"The ratio-only audit uses GUT-normalized α1 differences and is independent of the absolute intercept u, but it also fails without thresholds or new matching data.",
		}}
	}}
}
