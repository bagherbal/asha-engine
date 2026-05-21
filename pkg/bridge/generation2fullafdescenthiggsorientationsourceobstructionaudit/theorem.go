package generation2fullafdescenthiggsorientationsourceobstructionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_FULL_AF_DESCENT_HIGGS_ORIENTATION_SOURCE_OBSTRUCTION_AUDIT"
	theoremName = "Gate 891 — Full A_F Descent and HiggsOrientation Source Obstruction Audit"
)

func Generation2FullAFDescentHiggsOrientationSourceObstructionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 890 dual-seal/J-mirror descent-blocked status", Passed: len(a.Projectors) == 3 && projectorsOK(a.Projectors) && a.Descent.AFOrientStable && !a.Descent.FullAFStable && containsAll(a.Descent.Supports, []string{SupportAFOrientStable, SupportR3CandidateRequiresOrientationSeal}), Detail: FormatDescent(a.Descent)},
			{Name: "full A_F stability fails because full H mixes h_+ and h_-", Passed: a.Descent.GenericHMixesWeakSockets && !a.Descent.WeakFrameFullHInvariant && containsAll(a.Descent.Failures, []string{FailureFullHMixesWeakSockets, FailureSocketProjectorsNotStableFullH, FailureSocketProjectorsNotStableFullAF}), Detail: FormatDescent(a.Descent)},
			{Name: "A_F^orient is the Higgs-oriented stabilizer, not full A_F", Passed: a.Stabilizer.PreservesHPlusHMinus && a.Stabilizer.PreservesProjectors && !a.Stabilizer.IsFullH && !a.Stabilizer.IsFullAF && !a.Stabilizer.NativeDescentCertified && containsAll(a.Stabilizer.Supports, []string{SupportAFOrientIsStabilizer, SupportFullToOrientRestriction}) && containsAll(a.Stabilizer.Failures, []string{FailureStabilizerNotFullNativeAF, FailureAFOrientNotFullAF}), Detail: FormatStabilizer(a.Stabilizer)},
			{Name: "Higgs-orientation source candidates audited but no native source certified", Passed: len(a.OrientationSources.Candidates) >= 7 && !a.OrientationSources.AnyNativeSourceCertified && a.OrientationSources.RequiresOrientationSeal && containsAll(a.OrientationSources.Supports, []string{SupportOrientationSourceCandidatesAudited, SupportR3CandidateRequiresOrientationSeal}) && containsAll(a.OrientationSources.Failures, []string{FailureNoNativeHiggsOrientationSource, FailureNoNativeDescentFullToOrient}), Detail: FormatOrientationSources(a.OrientationSources)},
			{Name: "official diagnostics remain frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4 physical generation flavor and official-update firewalls preserved", Passed: firewallsOK(a.Firewalls) && !hasPhysicalLeak(a) && containsAll(a.FirewallsList(), []string{FailureNotNativeR3, FailureNoNativeHiggsOrientationSource, FailureNoNativeR3SectorLedger}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatProjectors(a.Projectors), FormatStabilizer(a.Stabilizer), FormatOrientationSources(a.OrientationSources), FormatDescent(a.Descent), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureNoNativeIncidenceFunctor, FailureAFOrientNotFullAF, FailureFullHMixesWeakSockets, FailureNoNativeDescentFullToOrient, FailureNoNativeHiggsOrientationSource, FailurePostOrientationNotFullAF, FailureNoNativeFiniteSectorProjectorTheorem, FailureNoNativeR3SectorLedger, FailureNoPhysicalParticleAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func hasPhysicalLeak(a Audit) bool {
	for _, p := range a.Projectors {
		if p.PhysicalSector || p.GenerationResolved || p.FlavorResolved || p.IndividualYukawaValue {
			return true
		}
	}
	return false
}
