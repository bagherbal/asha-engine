package generation2phasetransportdomaincodomainrepresentationactionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_PHASE_TRANSPORT_DOMAIN_CODOMAIN_REPRESENTATION_ACTION_AUDIT"
	theoremName = "Gate 904 — PhaseTransport Domain/Codomain and Representation Action Audit"
)

func Generation2PhaseTransportDomainCodomainRepresentationActionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 903 phase-shape support and transport-missing wound inherited", Passed: a.Inherited.ShapeSupported && a.Inherited.TransportMissing && a.Inherited.MissingObject == TransportMap && !a.Inherited.NativeTransport && containsAll(a.Inherited.Supports, []string{SupportGate903Inherited}) && containsAll(a.Inherited.Failures, []string{FailureShapeNotTransport, FailurePhaseAnchorSealed}), Detail: FormatInherited(a.Inherited)},
			{Name: "phase-transport domain typed as Hopf S1 plus Cl17 chirality but without C_R2 action", Passed: a.Domain.HopfS1Typed && a.Domain.CL17ChiralityTyped && !a.Domain.HopfActsOnCR2 && !a.Domain.GammaChiActsOnCR2 && !a.Domain.NativeDomainActionMap && containsAll(a.Domain.Supports, []string{SupportDomainTyped, SupportHopfDomainTyped, SupportCL17DomainTyped}) && containsAll(a.Domain.Failures, []string{FailureNoTypedHopfActionCR2, FailureNoTypedGammaChiActionCR2}), Detail: FormatDomain(a.Domain)},
			{Name: "codomain typed as right-character projector pair in End(C_R2)", Passed: a.Codomain.ProjectorPairTyped && a.Codomain.OutputsOrderedPair && !a.Codomain.NativeTransportToProjectors && a.Codomain.RightCharacterSplit == RightCharacterSplit && containsAll(a.Codomain.Supports, []string{SupportCodomainTyped, SupportRhoRTarget}) && containsAll(a.Codomain.Failures, []string{FailureNoTransportToCR2Projectors}), Detail: FormatCodomain(a.Codomain)},
			{Name: "required action is known but not action-compatible with rho_R yet", Passed: a.Action.PositiveToEPlus && a.Action.ConjugateToEMinus && !a.Action.ActionCompatibleWithRhoR && !a.Action.TypedActionOnRightPair && containsAll(a.Action.Supports, []string{SupportTransportSealSelectsEPlus}) && containsAll(a.Action.Failures, []string{FailureTransportNotActionCompatible, FailureNoTypedPhaseActionRightPair, FailureNoHopfChiralityRhoRAction}), Detail: FormatAction(a.Action)},
			{Name: "rho_R target labels cannot noncircularly define the transport", Passed: a.NonCircularity.RhoRLabelsSockets && !a.NonCircularity.RhoRExplainsOrdering && !a.NonCircularity.TransportDefinedByLabels && a.NonCircularity.TargetLabelRestatement && !a.NonCircularity.NonCircularSourceCertified && containsAll(a.NonCircularity.Failures, []string{FailureRhoRRestatesOrder, FailureTransportByTargetLabelsOnly}), Detail: FormatNonCircularity(a.NonCircularity)},
			{Name: "phase transport seal would order the airlock but does not promote to native R3", Passed: a.AirlockEffect.IfTransportSealed && a.AirlockEffect.SelectsEPlusAsLambda && a.AirlockEffect.OrdersNeutralPuncture && a.AirlockEffect.CollapsesLocalWounds && !a.AirlockEffect.NativeR3Promotion && containsAll(a.AirlockEffect.Supports, []string{SupportTransportOrdersAirlock, SupportTransportCollapsesWounds}) && containsAll(a.AirlockEffect.Failures, []string{FailureTransportSealNotNativeR3}), Detail: FormatAirlockEffect(a.AirlockEffect)},
			{Name: "missing object sharpened to typed phase action on C_R2", Passed: a.MissingObject.MissingObject == TransportMap && a.MissingObject.NowFullyTyped && !a.MissingObject.NativeMapCertified && containsAll(a.MissingObject.Supports, []string{SupportMasterWoundToActionCR2}) && containsAll(a.MissingObject.Failures, []string{FailureNoNativePhaseTransport, FailureNoHopfChiralityRhoRAction}), Detail: FormatMissingObject(a.MissingObject)},
			{Name: "operator diagnostics remain coherent and official ledgers frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.Alpha, AlphaB) && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4, action transport, alpha, Higgs orientation, physical-sector, and official-ledger firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNoNativePhaseTransport, FailureNoTypedHopfActionCR2, FailureNoTypedGammaChiActionCR2, FailureNoHopfChiralityRhoRAction, FailureRhoRRestatesOrder, FailureNotNativeR3}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatInherited(a.Inherited), FormatDomain(a.Domain), FormatCodomain(a.Codomain), FormatAction(a.Action), FormatNonCircularity(a.NonCircularity), FormatAirlockEffect(a.AirlockEffect), FormatMissingObject(a.MissingObject), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
