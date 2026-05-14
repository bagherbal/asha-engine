package conditionalhiggsmassrgtransport

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ConditionalHiggsMassFromQuarticRGTransportAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-CONDITIONAL-HIGGS-MASS-QUARTIC-RG-TRANSPORT-AUDIT"
	const name = "Conditional Higgs Mass from Quartic RG Transport Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 309 RG transport audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 308 quartic boundary and gstar²=1 seal are inherited", Passed: a.Boundary.BoundaryInherited && a.Boundary.GStarSquaredSealed == 1.0 && a.Boundary.LambdaUV > 0.258 && a.Boundary.LambdaUV < 0.259 && !a.Boundary.AbsoluteFinalClaim, Detail: FormatBoundary(a.Boundary)},
			{Name: "one-loop gauge, top-Yukawa, and Higgs-quartic RG system is formalized", Passed: a.Beta.OneLoopOnly && a.Beta.UsesStandardContinuumQFT && !a.Beta.DerivedAsFiniteCoreTheorem, Detail: FormatBeta(a.Beta)},
			{Name: "PeV threshold and top-Yukawa lanes are present", Passed: len(a.Thresholds) == 3 && len(a.TopLanes) == 2 && a.Thresholds[1].ConditionalOnPeVSeal && a.Thresholds[2].ConditionalOnPeVSeal && a.TopLanes[1].YtUV > 1.28 && a.TopLanes[1].YtUV < 1.29, Detail: FormatThreshold(a.Thresholds[1]) + " | " + FormatThreshold(a.Thresholds[2]) + " | " + FormatTopLane(a.TopLanes[1])},
			{Name: "pure SM high-scale lane is rejected before electroweak extraction", Passed: a.Prediction.PureSMRunInvalid, Detail: FormatPrediction(a.Prediction)},
			{Name: "conditional r_plus PeV-threshold Higgs mass is computed as a diagnostic number", Passed: a.Prediction.RPlusTopLaneComputed && a.Prediction.PrimaryConditionalMassGeV > 320 && a.Prediction.PrimaryConditionalMassGeV < 340 && a.Prediction.PrimaryLambdaAtV > 0.85 && a.Prediction.PrimaryLambdaAtV < 0.95, Detail: FormatPrediction(a.Prediction)},
			{Name: "firewalls prevent final collider-scale mass claim", Passed: a.Firewalls.NoObservedHiggsMassUsedForDerivation && a.Firewalls.NoObservedTopMassUsedForDerivation && a.Firewalls.NoTwoLoopTermsInserted && a.Firewalls.NoThresholdMatchingInserted && a.Firewalls.NoPoleMassMatchingInserted && a.Firewalls.PeVThresholdsRemainSealed && a.Firewalls.TopYukawaOriginRemainsSealed && a.Firewalls.PureSMPathologyRecorded && !a.Firewalls.FinalColliderPredictionClaimed && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "summary reports conditional number and preserves final-prediction firewall", Passed: a.Summary.Gate308Inherited && a.Summary.GStarSealActivated && a.Summary.RGSystemFormalized && a.Summary.ConditionalTransportRun && a.Summary.PrimaryMassComputed && a.Summary.PureSMPathologyFound && !a.Summary.FinalMassClaimed && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 309 intentionally computes conditional GeV diagnostics, but the values are not final pole-mass predictions.", "The r_+ one-loop lane is a tension diagnostic: the next legal step is two-loop RG plus threshold/pole matching."}}
	}}
}
