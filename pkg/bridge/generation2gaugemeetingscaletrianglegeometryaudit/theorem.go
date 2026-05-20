package generation2gaugemeetingscaletrianglegeometryaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2GaugeMeetingScaleTriangleGeometryAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 gauge meeting-scale triangle geometry audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate608 meeting-scale triangle audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate607 meeting-scale triangle", Passed: a.Inherited.Lambda12GeV > 0 && a.Inherited.Lambda13GeV > a.Inherited.Lambda12GeV && a.Inherited.Lambda23GeV > a.Inherited.Lambda13GeV, Detail: FormatInherited(a.Inherited)},
			{Name: "compute pairwise meeting table", Passed: len(a.PairwiseScales) == 3 && containsPair(a.PairwiseScales, "Lambda_12") && containsPair(a.PairwiseScales, "Lambda_13") && containsPair(a.PairwiseScales, "Lambda_23"), Detail: FormatPairwiseTable(a.PairwiseScales)},
			{Name: "compute log triangle geometry", Passed: a.LogGeometry.SpreadDecades > 2.9 && a.LogGeometry.SpreadDecades < 3.0 && a.LogGeometry.Ratio13Over12 > 10 && a.LogGeometry.Ratio23Over13 > 80, Detail: FormatLogGeometry(a.LogGeometry)},
			{Name: "classify boundary choices", Passed: len(a.BoundaryChoices) == 4 && containsBoundaryChoice(a.BoundaryChoices, "Lambda_12") && containsBoundaryChoice(a.BoundaryChoices, "Lambda_geom"), Detail: FormatBoundaryChoices(a.BoundaryChoices)},
			{Name: "define beta deformation diagnostics", Passed: len(a.BetaDeformations) == 3 && containsBetaStrategy(a.BetaDeformations, "deform b3") && containsBetaStrategy(a.BetaDeformations, "minimal ||Delta b||"), Detail: FormatBetaDeformations(a.BetaDeformations)},
			{Name: "classify threshold origin slots", Passed: len(a.ThresholdOriginSlots) >= 6 && containsThresholdSlot(a.ThresholdOriginSlots, "two-loop") && containsThresholdSlot(a.ThresholdOriginSlots, "heavy threshold") && containsThresholdSlot(a.ThresholdOriginSlots, "extra colored"), Detail: FormatThresholdOriginSlots(a.ThresholdOriginSlots)},
			{Name: "audit native ASHA status", Passed: !a.NativeStatus.ProvidesNativeThresholdSpectrum && !a.NativeStatus.ProvidesNativeLambdaUSelection && !a.NativeStatus.ClaimsUnification, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "record scalar relation without mixing closure", Passed: a.ScalarRelation.ZeroCrossingGeV > 0 && a.ScalarRelation.ZeroCrossingGeV < a.ScalarRelation.Lambda12GeV, Detail: FormatScalarRelation(a.ScalarRelation)},
			{Name: "preserve firewalls", Passed: !a.Firewalls.ClaimsUnification && !a.Firewalls.IntroducesNewFields && !a.Firewalls.FitsThresholds && !a.Firewalls.PromotesLambdaU && !a.Firewalls.DerivesEndpoint && !a.Firewalls.UsesScalarToClose, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
