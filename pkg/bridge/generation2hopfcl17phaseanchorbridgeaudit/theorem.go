package generation2hopfcl17phaseanchorbridgeaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_HOPF_CL17_PHASE_ANCHOR_BRIDGE_AUDIT"
	theoremName = "Gate 902 — Hopf–Cl(1,7) PhaseAnchor Bridge Audit"
)

func Generation2HopfCL17PhaseAnchorBridgeAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 901 phase-anchored airlock seal inherited without native anchor", Passed: a.Inherited.PhaseAnchorOrganizesChain && !a.Inherited.NativeAnchor && containsAll(a.Inherited.Supports, []string{SupportPhaseAnchoredAirlockInherited, SupportRightCharacterLambdaBarShape}) && containsAll(a.Inherited.Failures, []string{FailureNoNativeRightPhaseOrientation, FailureNoNativeSelectionLambdaBar}), Detail: FormatInherited(a.Inherited)},
			{Name: "Hopf/S1 phase orientation is strongest phase-anchor candidate but lacks typed socket-order map", Passed: a.HopfS1.HasOrientedPhaseCircle && a.HopfS1.MatchesLambdaBarShape && a.HopfS1.CanReadAirlockIfSealed && !a.HopfS1.NativeSocketOrderMapCertified && containsAll(a.HopfS1.Supports, []string{SupportHopfStrongestPhaseAnchor, SupportRightPairMatchesHopfConjugation, SupportAirlockReadableAsHopfIfSealed}) && containsAll(a.HopfS1.Failures, []string{FailureNoHopfToSocketOrderMap, FailureHopfNotNativeSelector}), Detail: FormatHopfS1(a.HopfS1)},
			{Name: "Cl(1,7) complex chirality airlock is strong candidate but lacks typed map to right-character phase", Passed: a.CL17Chirality.OmegaSquaredMinusOne && a.CL17Chirality.RequiresComplexChirality && a.CL17Chirality.CanOrientIOverMinusI && !a.CL17Chirality.TypedToRightCharacterPhase && !a.CL17Chirality.NativeSocketOrderMapCertified && containsAll(a.CL17Chirality.Supports, []string{SupportCL17StrongCandidate, SupportIVsMinusI, SupportCL17RealFormFirewall}) && containsAll(a.CL17Chirality.Failures, []string{FailureNoCL17ToRightPhaseMap, FailureChiralityNotSocketOrderTheorem}), Detail: FormatCL17Chirality(a.CL17Chirality)},
			{Name: "Hopf–Cl(1,7) bridge is the sharpened obstruction, not a certified theorem", Passed: a.HopfChiralityBridge.HopfAndChiralityCompatible && a.HopfChiralityBridge.PointsToSamePhaseWound && !a.HopfChiralityBridge.NativeBridgeCertified && !a.HopfChiralityBridge.CanAnchorRightCharacters && containsAll(a.HopfChiralityBridge.Supports, []string{SupportHopfCL17CompatibleSources, SupportPhaseAnchorMayRequireBridge, SupportHopfAndCL17SameWound}) && containsAll(a.HopfChiralityBridge.Failures, []string{FailureNoNativeHopfChiralityBridge}), Detail: FormatHopfChiralityBridge(a.HopfChiralityBridge)},
			{Name: "J/KO and boundary-pair orientation remain relevant but do not select lambda over bar(lambda)", Passed: a.JKO.RelevantToConjugacy && !a.JKO.SelectsLambda && !a.JKO.BreaksZ2 && !a.JKO.NativeTheorem && a.BoundaryPair.HasExteriorOrientation && a.BoundaryPair.SelectsDegreeOrder && !a.BoundaryPair.SelectsRightCharacterPhase && containsAll(a.JKO.Failures, []string{FailureJKONoLambdaSelection, FailureJMirrorNoBreak}) && containsAll(a.BoundaryPair.Failures, []string{FailureBoundaryNoLambdaSelection, FailureBoundaryDegreeNotPhaseOrder}), Detail: FormatJKO(a.JKO) + " | " + FormatBoundaryPair(a.BoundaryPair)},
			{Name: "finite spectral orientation cycle remains only a deep candidate", Passed: a.SpectralOrientation.DeepCandidate && !a.SpectralOrientation.CycleCertified && !a.SpectralOrientation.MapsToSocketOrder && !a.SpectralOrientation.NativeTheorem && containsAll(a.SpectralOrientation.Supports, []string{SupportSpectralDeepCandidate}) && containsAll(a.SpectralOrientation.Failures, []string{FailureNoSpectralCycleToSocketOrder}), Detail: FormatSpectralOrientation(a.SpectralOrientation)},
			{Name: "phase-anchor source reduced to Hopf–Cl17 bridge obstruction", Passed: containsAll(a.Ranking.StrongestCandidates, []string{HopfS1PhaseOrientation, CL17ComplexChirality}) && a.Ranking.BridgeCandidate == HopfChiralityBridge && !a.Ranking.NativeSourceFound && containsAll(a.Ranking.Supports, []string{SupportHopfCL17CompatibleSources, SupportHopfAndCL17SameWound, SupportPhaseAnchoredAirlockCoherent}) && containsAll(a.Ranking.Failures, []string{FailureNoNativeRightPhaseOrientation, FailureNoNativeHopfChiralityBridge, FailureNoNativeSelectionLambdaBar}), Detail: FormatRanking(a.Ranking)},
			{Name: "operator diagnostics remain coherent and official ledgers frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.Alpha, AlphaB) && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4, phase-anchor, alpha, Higgs orientation, physical-sector, and official-ledger firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNoNativeRightPhaseOrientation, FailureNoHopfToSocketOrderMap, FailureNoCL17ToRightPhaseMap, FailureNoNativeHopfChiralityBridge, FailureNotNativeR3, FailureNoR4NativeYukawaTheorem}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatInherited(a.Inherited), FormatHopfS1(a.HopfS1), FormatCL17Chirality(a.CL17Chirality), FormatHopfChiralityBridge(a.HopfChiralityBridge), FormatJKO(a.JKO), FormatBoundaryPair(a.BoundaryPair), FormatSpectralOrientation(a.SpectralOrientation), FormatRanking(a.Ranking), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
