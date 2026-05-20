package generation2coefficientjacobianrankoneboundarystressaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2CoefficientJacobianRankOneBoundaryStressAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 spectral-action coefficient Jacobian and rank-one boundary stress audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate616 coefficient Jacobian audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate615 coefficient grammar audit", Passed: a.Inherited.XiBoundary > 0 && a.Inherited.Verdict == StatusGate615Inherited, Detail: FormatInherited(a.Inherited)},
			{Name: "build coefficient dependency graph", Passed: len(a.DependencyGraph) >= 8 && hasDependency(a.DependencyGraph, "C_3") && hasDependency(a.DependencyGraph, "lambda") && hasDependency(a.DependencyGraph, "q_boundary"), Detail: FormatDependencyGraph(a.DependencyGraph)},
			{Name: "define normalized shadow map", Passed: !a.ShadowMap.RawPairTypeSafe && a.ShadowMap.PreferredTypeSafe && a.ShadowMap.ColorShadow > 0 && a.ShadowMap.ScalarShadow < 0, Detail: FormatShadowMap(a.ShadowMap)},
			{Name: "audit symbolic Jacobian", Passed: len(a.Jacobian) >= 8 && hasJacobian(a.Jacobian, "C_3", "+", "0") && hasJacobian(a.Jacobian, "lambda", "0", "+") && hasJacobian(a.Jacobian, "q_boundary stress", "+", "-"), Detail: FormatJacobian(a.Jacobian)},
			{Name: "test rank-one source candidates", Passed: len(a.RankOneCandidates) >= 8 && hasBridgeRankOne(a.RankOneCandidates) && !hasNativeRankOne(a.RankOneCandidates), Detail: FormatRankOneCandidates(a.RankOneCandidates)},
			{Name: "classify rank", Passed: !a.RankClassification.NativeRankOneFound && a.RankClassification.BridgeRankOneDefinable && a.RankClassification.RankTwoIndependentSlots, Detail: FormatRankClassification(a.RankClassification)},
			{Name: "audit anti-alignment forcing", Passed: !a.AntiAlignment.CanForceAntiAlignment && !a.AntiAlignment.Native && a.AntiAlignment.ResidualOverXi < 0.03, Detail: FormatAntiAlignment(a.AntiAlignment)},
			{Name: "audit scalar canonical normalization", Passed: !a.CanonicalNormalization.KPhiKnown && !a.CanonicalNormalization.CanonicalScalarLedgerKnown && !a.CanonicalNormalization.CanAuditLambdaBeforeAfterK, Detail: FormatCanonicalNormalization(a.CanonicalNormalization)},
			{Name: "record native status", Passed: !a.NativeStatus.SectorSplitF0 && !a.NativeStatus.NativeQStress && !a.NativeStatus.C3LambdaRelation && !a.NativeStatus.ScalarNormalization && !a.NativeStatus.ThresholdMatching && !a.NativeStatus.NativeXi, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve firewalls", Passed: !a.Firewalls.ClaimsXiNative && !a.Firewalls.ClaimsLambdaZero && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsHiggsStability && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsThresholdExistence && !a.Firewalls.ClaimsNativeCorrection, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func hasDependency(rows []CoefficientDependency, name string) bool {
	for _, r := range rows {
		if r.Source == name {
			return true
		}
	}
	return false
}

func hasJacobian(rows []JacobianEntry, source, dColor, dScalar string) bool {
	for _, r := range rows {
		if r.Source == source && r.DColor == dColor && r.DScalar == dScalar {
			return true
		}
	}
	return false
}

func hasBridgeRankOne(rows []RankOneSourceCandidate) bool {
	for _, r := range rows {
		if r.RankOneBridgeDefinable {
			return true
		}
	}
	return false
}

func hasNativeRankOne(rows []RankOneSourceCandidate) bool {
	for _, r := range rows {
		if r.Native && r.ProducesColorPositive && r.ProducesScalarNegative {
			return true
		}
	}
	return false
}
