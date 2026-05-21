package generation2rightcharacterrepresentationinductioncomplexphasemoduleaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_RIGHT_CHARACTER_REPRESENTATION_INDUCTION_COMPLEX_PHASE_MODULE_AUDIT"
	theoremName = "Gate 905 — RightCharacter Representation Induction from ComplexPhase Module Audit"
)

func Generation2RightCharacterRepresentationInductionComplexPhaseModuleAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 904 phase-action wound inherited", Passed: a.Inherited.NeedsPhaseActionCR2 && a.Inherited.DomainCodomainTyped && !a.Inherited.NativeActionCertified && containsAll(a.Inherited.Supports, []string{SupportGate904Inherited}), Detail: FormatInherited(a.Inherited)},
			{Name: "C_R2 has minimal conjugation-closed lambda/barlambda module shape but not native identification", Passed: a.MinimalModule.TwoCharacterMinimal && a.MinimalModule.MatchesRhoRShape && !a.MinimalModule.NativeASHAIdentification && !a.MinimalModule.SelectsOrder && containsAll(a.MinimalModule.Supports, []string{SupportCR2MinimalModuleShape, SupportConjugationClosure, SupportRhoRMatchesTwoChar}) && containsAll(a.MinimalModule.Failures, []string{FailureMinimalModuleNotNative, FailureNoNativeCR2Identification, FailureModuleDoesNotOrder}), Detail: FormatMinimalModule(a.MinimalModule)},
			{Name: "phase-character supports reconstruct rho_R form but require orientation choice", Passed: a.ProjectorSupport.ProjectorsRealized && a.ProjectorSupport.RhoRFormReconstructed && !a.ProjectorSupport.IdentifyELambdaAsEPlus && a.ProjectorSupport.NeedsPhaseOrientationChoice && containsAll(a.ProjectorSupport.Supports, []string{SupportProjectorsAsCharacters, SupportPhaseActionReconstructs}) && containsAll(a.ProjectorSupport.Failures, []string{FailureELambdaEPlusNeedsChoice}), Detail: FormatProjectorSupport(a.ProjectorSupport)},
			{Name: "Hopf S1 acts on abstract two-character module but not natively on ASHA right sockets", Passed: a.HopfAction.S1ActsOnAbstractModule && a.HopfAction.ConjugateActionPresent && a.HopfAction.ReconstructsSplitIfIdentified && !a.HopfAction.NativeRightSocketAction && containsAll(a.HopfAction.Supports, []string{SupportHopfAbstractAction, SupportHopfReconstructsSplit}) && containsAll(a.HopfAction.Failures, []string{FailureHopfAbstractNotNative}), Detail: FormatHopfAction(a.HopfAction)},
			{Name: "Cl17 chirality supplies complex-structure candidate but no C_R2 eigensocket map", Passed: a.CL17Induction.SuppliesComplexStructure && a.CL17Induction.IMinusISplitMatchesPair && !a.CL17Induction.EigenSocketToCR2Map && !a.CL17Induction.InducesRhoRAction && containsAll(a.CL17Induction.Supports, []string{SupportCL17ComplexStructure, SupportIMinusIMatchesPair}) && containsAll(a.CL17Induction.Failures, []string{FailureNoGammaChiProjectorMap, FailureCL17NoRhoRAction}), Detail: FormatCL17Induction(a.CL17Induction)},
			{Name: "two-character module certifies pair, not lambda-over-barlambda order", Passed: a.Order.PairCertified && !a.Order.OrderCertified && !a.Order.PositivePhaseExposure && containsAll(a.Order.Supports, []string{SupportWoundReduced, SupportPairNoLongerArbitrary}) && containsAll(a.Order.Failures, []string{FailurePairNotOrder, FailureNoLambdaSelection, FailureNoPhaseOrder}), Detail: FormatOrder(a.Order)},
			{Name: "phase action wound reduced to identification and orientation", Passed: a.MissingObject.PairNoLongerArbitrary && !a.MissingObject.NativeSolved && containsAll(a.MissingObject.Supports, []string{SupportWoundReduced, SupportPairNoLongerArbitrary}) && containsAll(a.MissingObject.Failures, []string{FailureNoNativeCR2Identification, FailurePairNotOrder}), Detail: FormatMissingObject(a.MissingObject)},
			{Name: "operator diagnostics remain coherent and official ledgers frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.Alpha, AlphaB) && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4, phase order, alpha, Higgs orientation, physical-sector, generation/flavor, and official-ledger firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureMinimalModuleNotNative, FailureNoNativeCR2Identification, FailurePairNotOrder, FailureNoLambdaSelection, FailureNotNativeR3}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatInherited(a.Inherited), FormatMinimalModule(a.MinimalModule), FormatProjectorSupport(a.ProjectorSupport), FormatHopfAction(a.HopfAction), FormatCL17Induction(a.CL17Induction), FormatOrder(a.Order), FormatMissingObject(a.MissingObject), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
