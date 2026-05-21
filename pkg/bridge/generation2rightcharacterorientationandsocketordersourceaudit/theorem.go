package generation2rightcharacterorientationandsocketordersourceaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_RIGHT_CHARACTER_ORIENTATION_SOCKET_ORDER_SOURCE_AUDIT"
	theoremName = "Gate 898 — RightCharacter Orientation and SocketOrder Source Audit"
)

func Generation2RightCharacterOrientationAndSocketOrderSourceAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 897 Z2 airlock obstruction inherited", Passed: containsAll(Statuses(), []string{StatusGate897Inherited}) && !a.RightCharacter.SelectsPlusWithoutOrientation, Detail: FormatRightCharacter(a.RightCharacter)},
			{Name: "right character pair is typed but requires phase orientation", Passed: a.RightCharacter.PairTyped && a.RightCharacter.ComplexOrientationCandidate && !a.RightCharacter.NativePhaseOrientation && a.RightCharacter.RequiresPhaseOrientationSeal && containsAll(a.RightCharacter.Failures, []string{FailureNoNativeRightPhaseOrientation, FailureCharacterConjugationNeedsPhase}), Detail: FormatRightCharacter(a.RightCharacter)},
			{Name: "complex orientation can state socket order only if sealed", Passed: a.RightCharacter.EPlusAsLambda && a.RightCharacter.EMinusAsBarLambda && a.RightCharacter.SocketOrderStatedGivenOrientation && !a.RightCharacter.SelectsPlusWithoutOrientation && containsAll(a.RightCharacter.Supports, []string{SupportEPlusLambdaEMinusBar, SupportSocketOrderGivenComplexOrientation, SupportPhaseOrientationCanSelectPlus}), Detail: FormatRightCharacter(a.RightCharacter)},
			{Name: "finite one-form arrow pattern matches but restates socket order", Passed: a.OneFormArrow.MatchesEPlusPunctureOrder && !a.OneFormArrow.DerivesOrderIndependently && a.OneFormArrow.RestatesSocketOrder && !a.OneFormArrow.NativeArrowDirectionSelector && containsAll(a.OneFormArrow.Failures, []string{FailureOneFormArrowRestatesOrder, FailureEdgeDirectionNotNativeSelector}), Detail: FormatOneFormArrow(a.OneFormArrow)},
			{Name: "boundary degree order indexes flag levels but does not break Z2", Passed: a.BoundaryDegree.IndexesFlagLevels && !a.BoundaryDegree.SelectsSocketSign && !a.BoundaryDegree.BreaksZ2 && containsAll(a.BoundaryDegree.Failures, []string{FailureBoundaryExposureNoSocketSign, FailureBoundaryDegreeNoBreakZ2}), Detail: FormatBoundaryDegree(a.BoundaryDegree)},
			{Name: "J/chirality and B-L do not select plus", Passed: a.JChirality.CandidatesPossible && !a.JChirality.KOSignCertified && !a.JChirality.JMirrorBreaksZ2 && !a.JChirality.ChiralitySelectsPlus && a.BMinusL.CompensationWorks && !a.BMinusL.BreaksZ2 && containsAll(a.JChirality.Failures, []string{FailureNoKOSignSelectsPlus, FailureJMirrorNoBreakZ2, FailureChiralityNoPlus}) && containsAll(a.BMinusL.Failures, []string{FailureBMinusLNoBreakZ2}), Detail: FormatJChirality(a.JChirality) + " | " + FormatBMinusL(a.BMinusL)},
			{Name: "socket-order wound sharpens to right-character phase orientation source", Passed: a.Phase.Z2IsPhaseAmbiguity && a.Phase.RightPhaseSealCandidate && !a.Phase.NativeTheorem && a.Phase.SelectsPlusIfSealed && a.Phase.NextFrontier == NextFrontier && containsAll(a.Phase.Supports, []string{SupportPhaseOrientationSealRequired, SupportZ2IsPhaseAmbiguity, SupportWoundSharpensToPhaseSource}), Detail: FormatPhase(a.Phase)},
			{Name: "official diagnostics remain frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4, alpha, Higgs orientation, phase orientation, physical-sector, and official-ledger firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureHiggsOrientationStillSealed, FailureNoNativeSocketOrderSelector, FailureNoNativeRightPhaseOrientation}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatRightCharacter(a.RightCharacter), FormatOneFormArrow(a.OneFormArrow), FormatBoundaryDegree(a.BoundaryDegree), FormatJChirality(a.JChirality), FormatBMinusL(a.BMinusL), FormatPhase(a.Phase), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
