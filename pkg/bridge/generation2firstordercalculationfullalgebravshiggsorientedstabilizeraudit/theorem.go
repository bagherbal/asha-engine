package generation2firstordercalculationfullalgebravshiggsorientedstabilizeraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_FIRST_ORDER_CALCULATION_FULL_ALGEBRA_VS_HIGGS_ORIENTED_STABILIZER_AUDIT"
	theoremName = "Gate 855 — First-Order Calculation: Full Algebra vs Higgs-Oriented Stabilizer Audit"
)

func Generation2FirstOrderCalculationFullAlgebraVsHiggsOrientedStabilizerAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 854 matrix seal and type first-order target", Passed: a.Impact.Gate854Inherited && a.Target.WellTyped && a.Target.SupportExecutable && !a.Target.Certified && containsAll(a.Target.Supports, []string{SupportFirstOrderTargetExecutable}), Detail: FormatTarget(a.Target)},
			{Name: "audit minimal carrier preservation and ambient/active separation", Passed: a.Carrier.HLRank == HLRank && a.Carrier.HRMinRank == HRMinRank && a.Carrier.HPartMinRank == HPartMinRank && a.Carrier.HFMinRank == HFMinRank && a.Carrier.AmbientPartRank == AmbientPartRank && a.Carrier.AmbientFRank == AmbientFRank && a.Carrier.RightPunctureOutside && a.Carrier.LeftKernelInsideHL && a.Carrier.MinimalPreservedInStabilizer && !a.Carrier.MinimalPreservedUnderFullAF && !a.Carrier.AbsentCellForcedBack, Detail: FormatCarrier(a.Carrier)},
			{Name: "block full A_F first-order branch by generic H weak-socket mixing", Passed: a.Full.GenericHMixesHPlusHMinus && a.Full.DRequiresOrientedWeakSockets && a.Full.CarrierPreservedByFullAFAtCoarseLevel && !a.Full.OrientedSupportPreservedByFullAF && !a.Full.FirstOrderSupportZero && !a.Full.FirstOrderCertified && len(a.Full.ObstructionTerms) >= 2 && containsAll(a.Full.Failures, []string{FailureFullAFNoFirstOrderTheorem, FailureFullHActionMixesWeakSockets, FailureFullAFBlockedByHiggsOrientation, FailureWeakFrameNotNativeHInvariant}), Detail: FormatFullBranch(a.Full)},
			{Name: "classify oriented stabilizer branch as support-compatible but not full theorem", Passed: a.Stabilizer.PreservesHPlusHMinus && a.Stabilizer.PreservesLeptoColor && a.Stabilizer.PreservesMinimalCarrier && a.Stabilizer.PunctureRemainsOutside && a.Stabilizer.KernelRemainsStableCandidate && a.Stabilizer.FirstOrderSupportCompatible && !a.Stabilizer.FirstOrderOperatorTheorem && !a.Stabilizer.FullUnbrokenAFTheorem && containsAll(a.Stabilizer.Failures, []string{FailureStabilizerNotFullUnbrokenAFTheorem, FailureStabilizerOnlySupportLevel}), Detail: FormatStabilizer(a.Stabilizer)},
			{Name: "audit puncture/kernel stability distinction", Passed: a.Kernel.RightPuncture == "e_+ tensor P_1" && a.Kernel.LeftKernel == "h_+ tensor P_1" && a.Kernel.RightPunctureRank == 1 && a.Kernel.LeftKernelRank == 1 && !a.Kernel.LeftKernelStableFullAF && a.Kernel.LeftKernelStableStabilizer && !a.Kernel.PhysicalNeutrinoTheorem && !a.Kernel.MasslessnessTheorem, Detail: FormatKernel(a.Kernel)},
			{Name: "classify D_F^sym as post-orientation support object", Passed: a.Impact.PostOrientationSupportObject && a.Impact.StabilizerSupportPass && !a.Impact.FullAFPass && !a.Impact.NativeFiniteTripleProof && !a.Impact.FirstOrderCertified && containsAll(a.Stabilizer.Supports, []string{SupportDFPostOrientation, SupportStabilizerCompatibilityAtSeal}), Detail: FormatImpact(a.Impact)},
			{Name: "preserve ledgers and R3/R4 firewalls", Passed: a.Ledger.OfficialFrozen && !a.Ledger.AlphaNative && !a.Ledger.R3 && !a.Ledger.R4 && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 855 first-order compatibility firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.FullAFNoFirstOrder && a.Firewalls.FullHMixesWeakSockets && a.Firewalls.FullAFBlockedByOrientation && a.Firewalls.WeakFrameNotNativeH && a.Firewalls.StabilizerNotFullAF && a.Firewalls.StabilizerSupportOnly && a.Firewalls.NoJOppositeProof && a.Firewalls.NoBimoduleProof && a.Firewalls.NoFullFirstOrder && a.Firewalls.NoKOProof && a.Firewalls.NoNativeFiniteTriple && a.Firewalls.DSymbolicOnly && a.Firewalls.YSymbolicOnly && a.Firewalls.NoAlphaSource && a.Firewalls.NoTraceReadout && a.Firewalls.NoOfficialNEffUpdate && a.Firewalls.NoCYukawaCHiggsUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoParticleAssignment && a.Firewalls.NoNeutrinoTheorem && a.Firewalls.NoThreeGenerationTheorem && a.Firewalls.Verdict == StatusFirewallVerdict, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatCarrier(a.Carrier), FormatTarget(a.Target), FormatFullBranch(a.Full), FormatStabilizer(a.Stabilizer), FormatKernel(a.Kernel), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
