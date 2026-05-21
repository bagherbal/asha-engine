package generation2higgsorientedstabilizeralgebrapostorientationlayeraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_HIGGS_ORIENTED_STABILIZER_ALGEBRA_POST_ORIENTATION_LAYER_AUDIT"
	theoremName = "Gate 856 — Higgs-Oriented Stabilizer Algebra and Post-Orientation Layer Audit"
)

func Generation2HiggsOrientedStabilizerAlgebraPostOrientationLayerAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 855 post-orientation layer classification", Passed: a.Impact.Gate855Inherited && a.Impact.PostOrientationLayer && !a.Impact.FullAFPass && containsAll(a.DF.Supports, []string{SupportPostOrientationLayerClassification}), Detail: FormatImpact(a.Impact)},
			{Name: "audit Higgs-oriented stabilizer of weak socket frame", Passed: a.WeakFrame.FullHActsOnFullDoublet && !a.WeakFrame.FullHPreservesIndividualLines && a.WeakFrame.StabilizerPreservesHPlus && a.WeakFrame.StabilizerPreservesHMinus && a.WeakFrame.StabilizerIsComplexSubalgebra && !a.WeakFrame.StabilizerIsNativeFullH && containsAll(a.WeakFrame.Supports, []string{SupportStabilizerIsComplexSubalgebra}) && containsAll(a.WeakFrame.Failures, []string{FailureFullHPreservesSocketFrame, FailureFullHNativeSocketEigensplit, FailureStabilizerNotFullH}), Detail: FormatWeakFrame(a.WeakFrame)},
			{Name: "define A_F^orient as post-orientation stabilizer algebra", Passed: a.Algebra.OrientedAlgebra == "A_F^orient=C_R plus C_H plus M_3(C)" && !a.Algebra.ContainsFullH && a.Algebra.ContainsCH && a.Algebra.ContainsM3C && a.Algebra.ContainsRightC && a.Algebra.PostOrientationLayer && !a.Algebra.UnbrokenFullAFTheorem && !a.Algebra.PhysicalElectroweakTheorem && containsAll(a.Algebra.Supports, []string{SupportAForientDefinition, SupportPostOrientationLayerClassification}), Detail: FormatAlgebra(a.Algebra)},
			{Name: "verify oriented action preserves socket frame, lepto-color blocks, minimal carrier, puncture and kernel candidates", Passed: a.Action.PreservesHPlusHMinus && a.Action.PreservesP1P3 && a.Action.PreservesHRMin && a.Action.PreservesHFMin && a.Action.PunctureRemainsOutside && a.Action.LeftKernelStableCandidate && a.Action.MinimalCarrierClosureInAForient && !a.Action.MinimalCarrierClosureInFullAF && containsAll(a.Action.Supports, []string{SupportAForientPreservesSockets, SupportAForientPreservesPunctureKernel}), Detail: FormatAction(a.Action)},
			{Name: "classify D_F^sym support compatibility with A_F^orient only", Passed: a.DF.SupportCompatible && !a.DF.OperatorTheoremCompatible && !a.DF.FullAFCompatible && a.DF.PostOrientationObject && a.DF.FirstOrderReadyForGate857 && !a.DF.FirstOrderCalculatedThisGate && !a.DF.FirstOrderCertified && containsAll(a.DF.Supports, []string{SupportDFCompatibleWithAForient, SupportFirstOrderNextTarget}) && containsAll(a.DF.Failures, []string{FailureNoFirstOrderCalculationYet, FailureNoFullFirstOrderProof, FailureDStillSymbolic}), Detail: FormatDF(a.DF)},
			{Name: "preserve active/ambient carrier and kernel-puncture ledger", Passed: a.Carrier.HLRank == HLRank && a.Carrier.HRMinRank == HRMinRank && a.Carrier.HPartMinRank == HPartMinRank && a.Carrier.HFMinRank == HFMinRank && a.Carrier.AmbientPartRank == AmbientPartRank && a.Carrier.AmbientFRank == AmbientFRank && a.Carrier.DSymRank == DSymRank && a.Carrier.KernelRank == DSymKernelRank && a.Carrier.RightPuncture == "e_+ tensor P_1" && a.Carrier.LeftKernel == "h_+ tensor P_1", Detail: FormatCarrier(a.Carrier)},
			{Name: "freeze official ledgers and block R3/R4 promotion", Passed: a.Ledger.OfficialFrozen && !a.Ledger.AlphaNative && !a.Ledger.R3 && !a.Ledger.R4 && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 856 post-orientation stabilizer firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.FullHPreservesSocketFrame && a.Firewalls.NativeHEigensplit && a.Firewalls.FullAFCompatible && a.Firewalls.StabilizerNotFullH && a.Firewalls.AForientNotFullAF && a.Firewalls.PostOrientationNotEWB && a.Firewalls.NoHiggsVacuumTheorem && a.Firewalls.NoWeakMixingTheorem && a.Firewalls.NoFirstOrderThisGate && a.Firewalls.NoFullFirstOrderProof && a.Firewalls.NoJOppositeProof && a.Firewalls.NoBimoduleProof && a.Firewalls.SupportOnly && a.Firewalls.NoNativeTriple && a.Firewalls.DSymbolicOnly && a.Firewalls.YSymbolicOnly && a.Firewalls.NoAlphaSource && a.Firewalls.NoTraceReadout && a.Firewalls.NoOfficialNEffUpdate && a.Firewalls.NoCYukawaCHiggsUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoParticleAssignment && a.Firewalls.NoNeutrinoTheorem && a.Firewalls.NoThreeGenerationTheorem && a.Firewalls.Verdict == StatusFirewallVerdict, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatWeakFrame(a.WeakFrame), FormatAlgebra(a.Algebra), FormatAction(a.Action), FormatDF(a.DF), FormatCarrier(a.Carrier), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
