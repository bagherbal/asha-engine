package fullthresholdrgtransport

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FullThresholdRGTransportConditionalHiggsMassPredictionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-FULL-THRESHOLD-RG-TRANSPORT-CONDITIONAL-HIGGS-MASS-PREDICTION-AUDIT"
	const name = "Full Threshold RG Transport / Conditional Higgs Mass Prediction Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 322 full threshold RG transport audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "two-stage flattened-top RG protocol is formalized", Passed: a.Protocol.Formalized && a.Protocol.LaneName == "gauge_only_zero_top_lower_envelope" && a.Protocol.TopYukawaUV == 0 && a.Protocol.OneLoopOnly && a.Protocol.LambdaUV > 0.258 && a.Protocol.LambdaUV < 0.259, Detail: FormatProtocol(a.Protocol)},
			{Name: "Gate-321 derived threshold jump is inserted with the correct negative sign", Passed: a.Insertion.Inserted && a.Insertion.DeltaLambda < -0.097 && a.Insertion.DeltaLambda > -0.098 && a.Insertion.LowersQuartic && !a.Insertion.DerivedAsFullPotential, Detail: FormatInsertion(a.Insertion)},
			{Name: "transport computes a stable positive running quartic near the comparison target", Passed: a.Transport.Computed && a.Transport.Perturbative && a.Transport.VacuumStableAtEndpoint && a.Transport.FinalLambdaAtV > 0.128 && a.Transport.FinalLambdaAtV < 0.130 && a.Transport.RunningMassGeV > 124.0 && a.Transport.RunningMassGeV < 126.0 && a.Transport.NearObservedWithinOnePct, Detail: FormatTransport(a.Transport)},
			{Name: "baseline floor and threshold-shifted mass are both recorded", Passed: a.Transport.BaselineMassGeV > 158.0 && a.Transport.BaselineMassGeV < 159.0 && a.Transport.LambdaAtThresholdPlus > 0.226 && a.Transport.LambdaAtThresholdPlus < 0.227 && a.Transport.LambdaAtThresholdMinus > 0.128 && a.Transport.LambdaAtThresholdMinus < 0.129, Detail: FormatTransport(a.Transport)},
			{Name: "precision gap sieve keeps running mass separate from pole mass", Passed: a.Precision.Formalized && a.Precision.WithinOnePercent && a.Precision.RunningMassNotPoleMass && a.Precision.TwoLoopRequired && a.Precision.PoleMatchingRequired && a.Precision.ExactThresholdScaleRequired, Detail: FormatPrecision(a.Precision)},
			{Name: "firewalls preserve no final collider mass claim", Passed: a.Firewalls.NoPoleMassClaimed && a.Firewalls.NoTwoLoopClaimed && a.Firewalls.NoExactThresholdScaleClaimed && a.Firewalls.NoPhysicalTopSectorClaimed && a.Firewalls.NoFullSigmaPotentialClaimed && a.Firewalls.NoFinalColliderClaimed && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "summary records conditional transport success without overclaim", Passed: a.Summary.TwoStageExecuted && a.Summary.DerivedJumpInserted && a.Summary.RunningMassComputed && a.Summary.NearObserved && !a.Summary.FinalMassClaimed && a.Summary.FirewallsPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 322 uses the Gate-321 derived Delta lambda, not the fitted Gate-314 target jump.", "The output is a conditional one-loop running-mass proxy near 125 GeV; two-loop and pole matching remain mandatory before any collider-mass claim."}}
	}}
}
