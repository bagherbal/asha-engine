package generation2symbolicdffirstorderjoppositecompatibilityaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_SYMBOLIC_D_F_FIRST_ORDER_J_OPPOSITE_COMPATIBILITY_AUDIT"
	theoremName = "Gate 850 — Symbolic D_F First-Order and J-Opposite Compatibility Audit"
)

func Generation2SymbolicDFFirstOrderJOppositeCompatibilityAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		checks := []theorem.Check{}
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks = append(checks,
			theorem.Check{Name: "inherit Gate 849 symbolic D_F and puncture/kernel anatomy", Passed: a.Impact.Gate849Inherited && a.Impact.SymbolicDFInherited && a.SymbolicDF.LeftRank == HLRank && a.SymbolicDF.RightRank == HRMinRank && a.SymbolicDF.TotalRank == ChiralTotalDim && a.SymbolicDF.YRank == 7 && a.SymbolicDF.DFRank == 14 && a.SymbolicDF.KernelDim == 1 && containsAll(a.SymbolicDF.Supports, []string{SupportCorrectChiralSupportForm, SupportPunctureKernelPairInherited}), Detail: FormatSymbolicDF(a.SymbolicDF)},
			theorem.Check{Name: "audit representation action data availability", Passed: !a.Representation.RhoFAvailable && !a.Representation.CompleteActionLedger && !a.Representation.CompletePackage && !a.Representation.CentralSupportLedger && containsAll(a.Representation.Failures, []string{FailureNoCompleteRhoFActionLedger, FailureNoCompleteFiniteTriplePackage}), Detail: FormatRepresentation(a.Representation)},
			theorem.Check{Name: "audit J-opposite and first-order obstruction", Passed: !a.FirstOrder.CanFormCommutator && !a.FirstOrder.CanFormOppositeAction && !a.FirstOrder.FirstOrderProven && !a.FirstOrder.BimoduleStable && !a.FirstOrder.JOppositeCompatible && len(a.FirstOrder.MissingData) >= 5 && containsAll(a.FirstOrder.Failures, []string{FailureNoJOppositeCompatibilityProof, FailureNoFullFirstOrderConditionProof, FailureNoBimoduleCommutantProof}), Detail: FormatFirstOrder(a.FirstOrder)},
			theorem.Check{Name: "preserve chirality oddness only at support level", Passed: a.Chirality.SupportLevelOddness && !a.Chirality.NativeGammaCertified && a.Chirality.KernelChiralityConsistent && a.Chirality.PunctureChiralitySet && containsAll(a.Chirality.Failures, []string{FailureChiralityOnlySupportLevel, FailureNoCompleteFiniteTriplePackage}), Detail: FormatChirality(a.Chirality)},
			theorem.Check{Name: "audit left neutral kernel stability candidate without certifying stability", Passed: a.Representation.Kernel.Name == "h_+ tensor P_1" && a.Representation.Kernel.Kernel && a.Representation.Kernel.StableCandidate && !a.Representation.Kernel.PhysicalAssignment && !a.Representation.KernelStabilityCertified && !a.FirstOrder.KernelStable && containsAll(a.Representation.Supports, []string{SupportKernelStableCandidate}) && containsAll(a.FirstOrder.Failures, []string{FailureKernelStabilityNotCertified}), Detail: FormatRepresentation(a.Representation) + " | " + FormatFirstOrder(a.FirstOrder)},
			theorem.Check{Name: "preserve symbolic D_F and Yukawa magnitude firewalls", Passed: a.SymbolicDF.SupportOnly && !a.SymbolicDF.NativeFiniteTriple && !a.SymbolicDF.ExplicitDFOperator && !a.SymbolicDF.NumericalDFMatrix && !a.Impact.NativeFiniteTriple && !a.Impact.FirstOrderCertified && !a.Impact.JOppositeCertified && a.Impact.MagnitudesStillMissing && containsAll(a.SymbolicDF.Failures, []string{FailureSymbolicDFNotNativeFiniteTriple, FailureNoExplicitDFOperator, FailureSymbolicYNotYukawaMagnitude}), Detail: FormatSymbolicDF(a.SymbolicDF) + " | " + FormatImpact(a.Impact)},
			theorem.Check{Name: "preserve physical naming and official ledger firewalls", Passed: !a.Impact.PhysicalNeutrinoTheorem && !a.Impact.MasslessTheorem && a.Impact.AlphaStillSealed && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && a.Ledger.OfficialFrozen && !a.Ledger.R3 && !a.Ledger.R4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			theorem.Check{Name: "preserve Gate 850 compatibility firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoRhoFActionLedger && a.Firewalls.NoCompletePackage && a.Firewalls.NoJOppositeProof && a.Firewalls.NoFirstOrderProof && a.Firewalls.NoBimoduleProof && a.Firewalls.NoExplicitDFOperator && a.Firewalls.SymbolicDFNotNative && a.Firewalls.KernelStabilityNotCertified && a.Firewalls.ChiralitySupportOnly && a.Firewalls.NoPhysicalNeutrino && a.Firewalls.NoRightNeutrino && a.Firewalls.NoMasslessness && a.Firewalls.NoYukawaMagnitudes && a.Firewalls.AlphaStillSealed && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.NoOfficialNEffUpdate && a.Firewalls.NoCYukawaCHiggsUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.Verdict == StatusCompatibilityFirewall, Detail: FormatFirewalls(a.Firewalls)},
		)
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatSymbolicDF(a.SymbolicDF), FormatRepresentation(a.Representation), FormatFirstOrder(a.FirstOrder), FormatChirality(a.Chirality), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
