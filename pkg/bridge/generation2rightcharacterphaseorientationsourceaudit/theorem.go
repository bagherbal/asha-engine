package generation2rightcharacterphaseorientationsourceaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_RIGHT_CHARACTER_PHASE_ORIENTATION_SOURCE_AUDIT"
	theoremName = "Gate 899 — RightCharacter PhaseOrientation Source Audit"
)

func Generation2RightCharacterPhaseOrientationSourceAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 898 phase-orientation obstruction inherited", Passed: a.Inherited.Z2Family && a.Inherited.RequiresPhaseOrientation && !a.Inherited.CanSelectSigmaPlusNatively && containsAll(a.Inherited.Failures, []string{FailureNoNativeSocketOrderSelector, FailureNoNativeRightPhaseOrientation}), Detail: FormatInherited(a.Inherited)},
			{Name: "Hopf/S1 phase orientation is strongest socket-order source candidate but not native", Passed: a.HopfS1.PhaseOrientation && a.HopfS1.CanOrientLambdaBar && a.HopfS1.StrongestCandidate && !a.HopfS1.NativeTheorem && a.HopfS1.SelectsPlusIfSealed && containsAll(a.HopfS1.Supports, []string{SupportHopfStrongestCandidate, SupportHopfCanSourceIfSealed}) && containsAll(a.HopfS1.Failures, []string{FailureNoHopfToSocketOrderTheorem, FailurePhaseOrientationSealNotNative}), Detail: FormatHopfS1(a.HopfS1)},
			{Name: "Cl(1,7) complex chirality airlock is a candidate but lacks typed map to right-character order", Passed: a.ComplexChirality.OmegaSquaredMinusOne && a.ComplexChirality.RequiresComplexAirlock && a.ComplexChirality.CanOrientIOverMinusI && !a.ComplexChirality.TypedToRightCharacters && !a.ComplexChirality.NativeSocketOrderTheorem && containsAll(a.ComplexChirality.Supports, []string{SupportChiralityCandidate, SupportCL17FirewallCandidate}) && containsAll(a.ComplexChirality.Failures, []string{FailureNoCL17ChiralityToRightOrderMap, FailureComplexChiralityNotSocketOrderTheorem}), Detail: FormatComplexChirality(a.ComplexChirality)},
			{Name: "J/KO and boundary-pair orientation do not select socket phase", Passed: a.JKO.RelevantToConjugation && !a.JKO.KOSignCertified && !a.JKO.SelectsPlus && !a.JKO.NativeTheorem && a.BoundaryPair.ExteriorOrientation && a.BoundaryPair.SelectsDegreeOrder && !a.BoundaryPair.SelectsSocketPhase && containsAll(a.JKO.Failures, []string{FailureJKODoesNotSelectPlus, FailureNoKOSignSocketOrderTheorem}) && containsAll(a.BoundaryPair.Failures, []string{FailureBoundaryPairNoLambdaBar, FailureBoundarySelectsDegreeNotSocketPhase}), Detail: FormatJKO(a.JKO) + " | " + FormatBoundaryPair(a.BoundaryPair)},
			{Name: "finite spectral orientation cycle is only a deep candidate", Passed: a.SpectralOrientation.DeepCandidate && !a.SpectralOrientation.CycleCertified && !a.SpectralOrientation.MapsToSocketOrder && !a.SpectralOrientation.NativeTheorem && containsAll(a.SpectralOrientation.Supports, []string{SupportSpectralOrientationCandidate}) && containsAll(a.SpectralOrientation.Failures, []string{FailureNoFiniteSpectralOrientationTheorem}), Detail: FormatSpectralOrientation(a.SpectralOrientation)},
			{Name: "socket-order wound is classified as phase-orientation seal", Passed: containsAll(a.Ranking.StrongestCandidates, []string{HopfS1PhaseOrientation, ComplexChiralityAirlock}) && a.Ranking.PhaseSealName == RightCharacterPhaseSeal && a.Ranking.SocketSelectorName == SocketOrderPhaseSelector && !a.Ranking.NativeSourceFound && a.Ranking.NextFrontier == NextFrontier && containsAll(a.Ranking.Supports, []string{SupportHopfAndChiralityStrongest, SupportSocketOrderReducedToPhaseSeal}) && containsAll(a.Ranking.Failures, []string{FailureNoNativeRightPhaseOrientation, FailureNoNativeSocketOrderSelector, FailureNoNativeSelectionSigmaPlus}), Detail: FormatRanking(a.Ranking)},
			{Name: "official diagnostics remain frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4, alpha, Higgs orientation, socket phase, physical-sector, and official-ledger firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureHiggsOrientationStillSealed, FailureNoNativeSocketOrderSelector, FailureNoNativeRightPhaseOrientation}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatInherited(a.Inherited), FormatHopfS1(a.HopfS1), FormatComplexChirality(a.ComplexChirality), FormatJKO(a.JKO), FormatBoundaryPair(a.BoundaryPair), FormatSpectralOrientation(a.SpectralOrientation), FormatRanking(a.Ranking), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
