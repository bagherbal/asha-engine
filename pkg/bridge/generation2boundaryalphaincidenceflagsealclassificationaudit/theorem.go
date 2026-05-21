package generation2boundaryalphaincidenceflagsealclassificationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_BOUNDARY_ALPHA_INCIDENCE_FLAG_SEAL_CLASSIFICATION_AUDIT"
	theoremName = "Gate 880 — BoundaryAlpha IncidenceFlag Seal Classification Audit"
)

func Generation2BoundaryAlphaIncidenceFlagSealClassificationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "freeze official ledger and classify impact", Passed: a.Ledger.OfficialFrozen && !a.Ledger.CanUpdate && !a.Impact.CanUpdateNEff && !a.Impact.CanPromoteToR3 && a.Impact.SealName == SealName, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "classify BoundaryAlpha incidence-flag seal", Passed: a.Seal.Name == SealName && a.Seal.FullName == FullSealName && a.Seal.ReducedExteriorShape && a.Seal.DegreeIndexedFlagSelector && !a.Seal.NativeFunctor && near(a.Seal.Alpha, AlphaB) && containsAll(a.Seal.Supports, []string{SupportBoundaryAlphaIncidenceFlagSeal, SupportReducedExteriorShape, SupportDegreeIndexSelectsQuotients, SupportAlphaReconstructedFromRanks}) && containsAll(a.Seal.Failures, []string{FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion, FailureAlphaStillSealed}), Detail: FormatSeal(a.Seal)},
			{Name: "reassemble conditional trace proxy chain", Passed: a.Chain.ReducedExteriorToAlpha && a.Chain.AlphaToSocketMagnitudes && a.Chain.SocketMagnitudesToYDagY && a.Chain.YDagYToHAgg && a.Chain.HAggToNEff && a.Chain.CoherentGivenSeal && containsAll(a.Chain.Supports, []string{SupportFullConditionalTraceProxyCoherent, SupportConditionalYDagYReadout, SupportOperatorNEffDiagnostic}) && containsAll(a.Chain.Failures, []string{FailureAlphaStillSealed, FailureConditionalProxyNotR3}), Detail: FormatChain(a.Chain)},
			{Name: "file exact native theorem still missing", Passed: a.MissingTheorem.Name == "BoundaryExteriorIncidenceFlagFunctor" && a.MissingTheorem.RequiredForR3 && !a.MissingTheorem.Native && containsAll(a.MissingTheorem.Supports, []string{SupportIncidenceFunctorIsExactMissingObject}) && containsAll(a.MissingTheorem.Failures, []string{FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion}), Detail: FormatMissing(a.MissingTheorem)},
			{Name: "block native R3/R4 eligibility", Passed: a.Eligibility.ConditionalTraceProxyMature && a.Eligibility.BoundaryAlphaSeal && !a.Eligibility.NativeIncidenceFlagFunctor && !a.Eligibility.NativeCrossLaneExclusion && !a.Eligibility.NativeSocketMagnitudeSource && !a.Eligibility.EligibleForR3 && !a.Eligibility.EligibleForR4 && containsAll(a.Eligibility.Failures, []string{FailureConditionalProxyNotR3, FailureNoNativeR3, FailureNoNativeSocketMagnitude, FailureNoNativeSectorTraceMagnitude, FailureNoNativeYukawaOperator}), Detail: FormatEligibility(a.Eligibility)},
			{Name: "preserve Gate 880 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatSeal(a.Seal), FormatChain(a.Chain), FormatMissing(a.MissingTheorem), FormatEligibility(a.Eligibility), FormatImpact(a.Impact), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
