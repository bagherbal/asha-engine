package heavypropagatorselfquartic

import "github.com/bagherbal/asha-engine/pkg/theorem"

func HeavyPropagatorSelfQuarticSieveThresholdNormalizationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-HEAVY-PROPAGATOR-SELF-QUARTIC-THRESHOLD-NORMALIZATION-AUDIT"
	const name = "Heavy Propagator & Self-Quartic Sieve / Threshold Normalization Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 321 threshold normalization audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "heavy self-quartic lanes are formalized and raw B-gap quartic is rejected as unnormalized EFT coupling", Passed: a.Quartic.Formalized && a.Quartic.RawSigmaQuartic > 0.010 && a.Quartic.RawSigmaQuartic < 0.011 && a.Quartic.CanonicalRankOneQuartic == 1 && !a.Quartic.RawLanePhysical && a.Quartic.CanonicalLanePhysical, Detail: FormatQuartic(a.Quartic)},
			{Name: "rank-one heavy propagator normalization is formalized on the seesaw support", Passed: a.Prop.Formalized && a.Prop.OverlapIndex == 1 && a.Prop.HeavySupportRank == 1 && a.Prop.HeavyMetric == 1 && a.Prop.PropagatorAtThreshold == 1 && a.Prop.CanonicalNormalization && a.Prop.RawTraceRequiresRescaling, Detail: FormatPropagator(a.Prop)},
			{Name: "threshold synthesis compares raw and canonical lanes", Passed: a.Threshold.Formalized && a.Threshold.CPortal > 0.391 && a.Threshold.CPortal < 0.392 && a.Threshold.PortalWithinOnePercent && !a.Threshold.RawSigmaLane.Viable && a.Threshold.RawSigmaLane.DeltaLambda < -9 && a.Threshold.CanonicalRankOneLane.Viable && a.Threshold.CanonicalRankOneLane.DeltaLambda > -0.098 && a.Threshold.CanonicalRankOneLane.DeltaLambda < -0.097, Detail: FormatThreshold(a.Threshold)},
			{Name: "canonical lane aligns with Gate-314 jump target within one percent", Passed: a.Alignment.Compared && a.Alignment.ResolvesGate314Target && a.Alignment.WithinOnePercent && a.Alignment.CanonicalDeltaLambda < 0 && a.Alignment.StillConditional, Detail: FormatAlignment(a.Alignment)},
			{Name: "firewalls preserve no final Higgs mass, pole mass, or full sigma-potential claim", Passed: a.Firewalls.NoFinalMassClaimed && a.Firewalls.NoPoleMassClaimed && a.Firewalls.NoFullSigmaPotentialClaim && a.Firewalls.NoHeavyMassClaimed && a.Firewalls.NoIndependentLambdaMixClaim && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "summary records conditional threshold-normalization success without overclaim", Passed: a.Summary.HeavyQuarticFormalized && a.Summary.PropagatorFormalized && a.Summary.ThresholdNormalized && a.Summary.RawLaneRejected && a.Summary.CanonicalLaneMatchesTarget && !a.Summary.FinalMassClaimed && a.Summary.FirewallsPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 321 converts the Gate-320 overlap witness into a conditional rank-one EFT threshold-normalization lane.", "The raw B_gap^2 self-quartic lane is rejected; the canonical rank-one lane gives Delta lambda=-0.097846792207, matching Gate 314 within 0.3%, but final Higgs-mass transport remains a later integration gate."}}
	}}
}
