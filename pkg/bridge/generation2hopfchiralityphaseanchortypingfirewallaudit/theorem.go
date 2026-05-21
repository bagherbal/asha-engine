package generation2hopfchiralityphaseanchortypingfirewallaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_HOPF_CHIRALITY_PHASE_ANCHOR_TYPING_FIREWALL_AUDIT"
	theoremName = "Gate 903 — HopfChirality PhaseAnchor Typing and Firewall Audit"
)

func Generation2HopfChiralityPhaseAnchorTypingFirewallAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 902 Hopf–Cl17 phase-anchor bridge obstruction inherited", Passed: a.Inherited.PhaseAnchorMissing && !a.Inherited.NativeBridgeCertified && containsAll(a.Inherited.StrongSources, []string{HopfPhaseCircle, CL17Chirality}) && containsAll(a.Inherited.Supports, []string{SupportGate902Inherited}) && containsAll(a.Inherited.Failures, []string{FailureNoNativeRightPhaseOrientation, FailurePhaseAnchorStillSealed}), Detail: FormatInherited(a.Inherited)},
			{Name: "Hopf phase has lambda/bar(lambda) shape but no typed action on C_R2", Passed: a.HopfTyping.HasS1Phase && a.HopfTyping.HasPositiveConjugatePair && a.HopfTyping.LabelsEPlusIfSealed && a.HopfTyping.LabelsEMinusIfSealed && !a.HopfTyping.TypedActionOnCR2Certified && !a.HopfTyping.NativeRepresentationMap && containsAll(a.HopfTyping.Supports, []string{SupportHopfHasRightShape, SupportHopfPositiveLabelsEPlus, SupportHopfConjugateLabelsEMinus}) && containsAll(a.HopfTyping.Failures, []string{FailureNoHopfToRightCharacterMap, FailureHopfLabelingStillSeal}), Detail: FormatHopfTyping(a.HopfTyping)},
			{Name: "Cl(1,7) complex chirality has i/-i shape but no right-character action map", Passed: a.CL17Typing.OmegaSquaredMinusOne && a.CL17Typing.SuppliesIOverMinusI && a.CL17Typing.CorrectConjugationShape && !a.CL17Typing.TypedToRightCharacter && !a.CL17Typing.SelectsSocketOrder && containsAll(a.CL17Typing.Supports, []string{SupportCL17HasConjugationShape, SupportCL17IOrientation, SupportIVsMinusIShape}) && containsAll(a.CL17Typing.Failures, []string{FailureNoGammaChiToRightAction, FailureChiralityDoesNotSelectSocket}), Detail: FormatCL17Typing(a.CL17Typing)},
			{Name: "Hopf–chirality alignment is strongest candidate but lacks native alignment/transport", Passed: a.Alignment.CompatiblePhaseTypes && a.Alignment.StrongestBridgeCandidate && !a.Alignment.NativeAlignmentMap && !a.Alignment.TransportToRhoR && !a.Alignment.CanSourceRightCharacterAnchor && containsAll(a.Alignment.Supports, []string{SupportHopfCL17Align, SupportCompatibleOrientationTypes, SupportAlignmentStrongestCandidate, SupportRightAnchorIfTransportTyped}) && containsAll(a.Alignment.Failures, []string{FailureNoHopfChiralityAlignment, FailureNoGammaTransportToRhoR, FailurePhaseAnchorStillSealed}), Detail: FormatAlignment(a.Alignment)},
			{Name: "shape match is not typed phase transport", Passed: a.ShapeFirewall.SamePhaseShape && !a.ShapeFirewall.TypedTransportCertified && a.ShapeFirewall.ConjugationResonanceOnly && !a.ShapeFirewall.NativeLambdaSelection && containsAll(a.ShapeFirewall.Supports, []string{SupportSourceSharpenedToTransport, SupportR3SealReducesToTransport}) && containsAll(a.ShapeFirewall.Failures, []string{FailureShapeMatchNotTheorem, FailureConjugationResonanceNotMap, FailureNoNativeLambdaSelection}), Detail: FormatShapeFirewall(a.ShapeFirewall)},
			{Name: "new missing object is HopfChiralityRightCharacterTransportMap", Passed: a.TransportTarget.MissingObject == PhaseTransportMap && a.TransportTarget.Sharpened && !a.TransportTarget.NativeMapCertified && containsAll(a.TransportTarget.Supports, []string{SupportSourceSharpenedToTransport, SupportR3SealReducesToTransport}) && containsAll(a.TransportTarget.Failures, []string{FailureNoTypedPhaseTransport, FailureNoNativeRightPhaseOrientation}), Detail: FormatTransportTarget(a.TransportTarget)},
			{Name: "operator diagnostics remain coherent and official ledgers frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.Alpha, AlphaB) && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4, phase-transport, alpha, Higgs orientation, physical-sector, and official-ledger firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNoHopfToRightCharacterMap, FailureNoGammaChiToRightAction, FailureNoHopfChiralityAlignment, FailureNoTypedPhaseTransport, FailureNotNativeR3, FailureNoR4NativeYukawaTheorem}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatInherited(a.Inherited), FormatHopfTyping(a.HopfTyping), FormatCL17Typing(a.CL17Typing), FormatAlignment(a.Alignment), FormatShapeFirewall(a.ShapeFirewall), FormatTransportTarget(a.TransportTarget), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
