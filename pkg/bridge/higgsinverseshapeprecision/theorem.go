package higgsinverseshapeprecision

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ExactInverseHiggsShapeDeviationPrecisionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-EXACT-INVERSE-HIGGS-SHAPE-DEVIATION-PRECISION-AUDIT"
	const name = "Exact Inverse Higgs Shape Deviation / Full-Precision Diagnostic Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 336 inverse precision audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 335 exact branch inherited", Passed: a.Inputs.HighestInheritedGate == inheritedHighestGate && a.Inputs.NativeShape.RatString() == "1197/4624" && a.Inputs.LambdaH.RatString() == "1197/9248", Detail: FormatInputs(a.Inputs)},
			{Name: "observed inverse shape computed exactly", Passed: a.Inverse.ObservedShape.RatString() == "39125025/151560721" && a.Inverse.ObservedLambda.RatString() == "39125025/303121442", Detail: FormatInverse(a.Inverse)},
			{Name: "contact-shape and lambda deviations computed exactly", Passed: a.Deviation.ShapeDelta.RatString() == "504067437/700816773904" && a.Deviation.LambdaDelta.RatString() == "504067437/1401633547808" && nearlyFloat(a.Deviation.RelativeShapeError, 0.278622502908778, 1e-14), Detail: FormatDeviation(a.Deviation)},
			{Name: "self-energy equivalence target preserved", Passed: a.SelfEnergy.RequiredRePiGeV2.RatString() == "504067437/11560000" && nearlyRat(a.SelfEnergy.RequiredRePiGeV2, 43.60444956747405, 1e-12), Detail: FormatSelfEnergy(a.SelfEnergy)},
			{Name: "required VEV counterfactual computed with high precision", Passed: nearlyFloat(a.RequiredVEV.VRequiredForTargetGeV, 245.87770295825946, 1e-12) && nearlyFloat(a.RequiredVEV.VShiftGeV, -0.34229704174053516, 1e-12), Detail: FormatRequiredVEV(a.RequiredVEV)},
			{Name: "firewalls preserve no-fit and no-pole claims", Passed: a.Firewalls.NoPoleCorrection && a.Firewalls.NoFittingShape && a.Firewalls.NoColliderClaim && !a.Ledger.UsesFloat64Core, Detail: FormatFirewalls(a.Firewalls) + " " + FormatLedger(a.Ledger)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary), FormatStatuses(Statuses(a))}}
	}}
}
