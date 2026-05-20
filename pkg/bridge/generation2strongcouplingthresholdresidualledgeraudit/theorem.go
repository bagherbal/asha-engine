package generation2strongcouplingthresholdresidualledgeraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2StrongCouplingThresholdResidualLedgerAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 strong-coupling threshold residual ledger audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate607 strong threshold residual ledger", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate606 gauge spine", Passed: a.Inherited.GaugeSpinePresent && a.Inherited.Delta3Runtime < 0 && a.Inherited.R3 > 1, Detail: FormatInherited(a.Inherited)},
			{Name: "convert strong residual in multiple schemes", Passed: len(a.ResidualConversions) >= 8 && containsResidual(a.ResidualConversions, "Delta g3") && containsResidual(a.ResidualConversions, "required Delta alpha3^-1"), Detail: FormatResidualTable(a.ResidualConversions)},
			{Name: "define strong threshold correction slot", Passed: len(a.ThresholdSlots) >= 5 && containsSlot(a.ThresholdSlots, "delta_3^threshold") && a.ThresholdSlots[0].RequiredValue > 0, Detail: FormatThresholdSlots(a.ThresholdSlots)},
			{Name: "compute beta coefficient deformation size", Passed: a.BetaDeformation.DeltaB3Required < 0 && a.BetaDeformation.FractionOfAbsSMb3 > 0.1 && a.BetaDeformation.FractionOfAbsSMb3 < 0.2, Detail: FormatBetaDeformation(a.BetaDeformation)},
			{Name: "compute meeting scale triangle", Passed: len(a.MeetingScales) == 3 && containsMeeting(a.MeetingScales, "Lambda_12") && containsMeeting(a.MeetingScales, "Lambda_13") && containsMeeting(a.MeetingScales, "Lambda_23") && a.MeetingScales[1].ScaleGeV > a.MeetingScales[0].ScaleGeV, Detail: FormatMeetingScales(a.MeetingScales)},
			{Name: "classify possible correction sources without asserting them", Passed: len(a.SourceClassifications) >= 6 && containsSource(a.SourceClassifications, "two-loop") && containsSource(a.SourceClassifications, "heavy threshold"), Detail: FormatSources(a.SourceClassifications)},
			{Name: "audit native ASHA status", Passed: !a.NativeStatus.ProvidesNativeStrongThreshold && !a.NativeStatus.ProvidesExtraColoredContent && !a.NativeStatus.ClaimsFullUnification, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "separate scalar sector", Passed: !a.ScalarRelation.MixedIntoStrongLedger && a.ScalarRelation.LambdaLambda12 < 0, Detail: FormatScalarRelation(a.ScalarRelation)},
			{Name: "preserve firewalls", Passed: !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsNewPhysics && !a.Firewalls.DerivesEndpoint && !a.Firewalls.DerivesWZPhoton && !a.Firewalls.ThresholdAssertedReal && !a.Firewalls.ScalarMixedIntoStrong, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
