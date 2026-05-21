package generation2firstorderjoppositecompatibilitycalculationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_FIRST_ORDER_J_OPPOSITE_COMPATIBILITY_CALCULATION_AUDIT"
	theoremName = "Gate 852 — First-Order / J-Opposite Compatibility Calculation Audit"
)

func Generation2FirstOrderJOppositeCompatibilityCalculationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 851 data seal and type the first-order target", Passed: a.Impact.DataSealInherited && a.FirstOrder.TypedAfterDataSeal && a.FirstOrder.HasRhoSeal && a.FirstOrder.HasJSeal && a.FirstOrder.HasGammaSeal && a.FirstOrder.HasDFSupport && containsAll(a.FirstOrder.Supports, []string{SupportFirstOrderTargetWellTyped}), Detail: FormatFirstOrder(a.FirstOrder)},
			{Name: "audit minimal carrier closure without native proof", Passed: a.CarrierClosure.PreservedByBlockActionSeal && !a.CarrierClosure.ClosedNatively && !a.CarrierClosure.AbsentCellForcedBackBySchematicRhoF && a.CarrierClosure.HPartMinRank == HPartMinRank && a.CarrierClosure.HFMinRank == HFMinRank && containsAll(a.CarrierClosure.Failures, []string{FailureMinimalCarrierClosureNotNative}), Detail: FormatCarrierClosure(a.CarrierClosure)},
			{Name: "isolate weak h_+/h_- orientation as fragile blocker", Passed: a.WeakOrientation.SplitDefinedAtSealLevel && a.WeakOrientation.RequiresHiggsOrientation && a.WeakOrientation.PrimaryFragilePoint && !a.WeakOrientation.StableUnderFullHAction && !a.WeakOrientation.NativeHEigensplit && containsAll(a.WeakOrientation.Failures, []string{FailureWeakSocketSplitNotNative, FailureWeakOrientationNeedsHiggsSeal}), Detail: FormatWeakOrientation(a.WeakOrientation)},
			{Name: "block J-opposite compatibility without operator-level J", Passed: a.JOpposite.JSealAvailable && !a.JOpposite.OperatorLevelJ && !a.JOpposite.OppositeCertified && !a.JOpposite.CanBuildOppositeCommutator && containsAll(a.JOpposite.Failures, []string{FailureNoOperatorLevelJOppositeAction, FailureNoJOppositeCompatibilityProof}), Detail: FormatJOpposite(a.JOpposite)},
			{Name: "audit first-order expression as non-executable with seal-level data", Passed: a.FirstOrder.TypedAfterDataSeal && !a.FirstOrder.ExecutableNow && !a.FirstOrder.Certified && !a.FirstOrder.HasOperatorRho && !a.FirstOrder.HasOperatorJ && !a.FirstOrder.HasOperatorGamma && !a.FirstOrder.HasOperatorDF && len(a.FirstOrder.ObstructionSources) >= 5 && containsAll(a.FirstOrder.Failures, []string{FailureFirstOrderNotExecutableWithSealData, FailureNoFullFirstOrderConditionProof, FailureNoBimoduleCommutantProof, FailureNoOperatorValuedDFMatrix}), Detail: FormatFirstOrder(a.FirstOrder)},
			{Name: "keep chirality oddness support-level only", Passed: a.Chirality.OddnessSupportLevel && !a.Chirality.OperatorGamma && !a.Chirality.KOExtensionSet && a.Chirality.LeftSign == 1 && a.Chirality.RightSign == -1 && containsAll(a.Chirality.Supports, []string{SupportChiralityOddnessSupportLevel}), Detail: FormatChirality(a.Chirality)},
			{Name: "audit kernel stability without physical promotion", Passed: a.Kernel.Kernel == "h_+ tensor P_1" && a.Kernel.RightPuncture == "e_+ tensor P_1" && a.Kernel.KernelRank == 1 && a.Kernel.DFSymKernelDim == 1 && a.Kernel.KernelInsideMinimalCarrier && a.Kernel.RightPunctureOutside && a.Kernel.StableUnderSchematicBlocks && !a.Kernel.StableUnderFullRhoJ && !a.Kernel.PhysicalNeutrinoTheorem && !a.Kernel.MasslessnessTheorem && containsAll(a.Kernel.Failures, []string{FailureKernelStabilityNotCertified, FailureNoPhysicalNeutrinoTheorem, FailureNoMasslessnessTheorem}), Detail: FormatKernel(a.Kernel)},
			{Name: "preserve ledgers and R3/R4 firewalls", Passed: a.Ledger.OfficialFrozen && !a.Ledger.AlphaNative && !a.Ledger.R3 && !a.Ledger.R4 && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && !a.Impact.FirstOrderCertified && !a.Impact.NativeFiniteTriple, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 852 compatibility firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.DataSealNotNative && a.Firewalls.NoCompleteRhoF && a.Firewalls.NoCompletePackage && a.Firewalls.WeakOrientationSealOnly && a.Firewalls.NoOperatorJ && a.Firewalls.NoJOppositeProof && a.Firewalls.FirstOrderNotExecutable && a.Firewalls.NoFirstOrderProof && a.Firewalls.NoBimoduleProof && a.Firewalls.NoOperatorDF && a.Firewalls.SymbolicDFNotNative && a.Firewalls.KernelStabilityNotCertified && a.Firewalls.DFSymNotMagnitudeSource && a.Firewalls.NoNumericalYukawas && a.Firewalls.AlphaStillSealed && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.NoOfficialNEffUpdate && a.Firewalls.NoCYukawaCHiggsUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoPhysicalNeutrino && a.Firewalls.NoRightNeutrino && a.Firewalls.NoMasslessness && a.Firewalls.Verdict == StatusCompatibilityFirewallVerdict, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatCarrierClosure(a.CarrierClosure), FormatWeakOrientation(a.WeakOrientation), FormatJOpposite(a.JOpposite), FormatFirstOrder(a.FirstOrder), FormatChirality(a.Chirality), FormatKernel(a.Kernel), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
