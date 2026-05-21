package generation2stabilizerbranchfirstordermatrixedgeintertwineraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_STABILIZER_BRANCH_FIRST_ORDER_MATRIX_EDGE_INTERTWINER_AUDIT"
	theoremName = "Gate 857 — Stabilizer-Branch First-Order Matrix and Edge-Intertwiner Audit"
)

func Generation2StabilizerBranchFirstOrderMatrixEdgeIntertwinerAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit A_F^orient stabilizer layer from Gate 856", Passed: a.Impact.Gate856Inherited && a.Algebra.OrientedAlgebra == "A_F^orient=C_R plus C_H plus M_3(C)" && !a.Algebra.ContainsFullH && a.Algebra.ContainsCH && a.Algebra.ContainsRightC && a.Algebra.ContainsM3C && a.Algebra.PostOrientationLayer && containsAll(a.Algebra.Failures, []string{FailureAForientNotFullAF, FailureFullHSocketFirewall}), Detail: FormatAlgebra(a.Algebra)},
			{Name: "verify A_F^orient preserves active blocks, puncture exclusion, and left kernel candidate at support level", Passed: a.Preserve.PreservesHPlusHMinus && a.Preserve.PreservesEPlusEMinus && a.Preserve.PreservesP1P3 && a.Preserve.PreservesHRMin && a.Preserve.PreservesHF && a.Preserve.RightPunctureOutside && a.Preserve.LeftKernelCandidate && containsAll(a.Preserve.Supports, []string{StatusSupportPreservation, SupportPunctureKernelStable}), Detail: FormatPreservation(a.Preserve)},
			{Name: "audit active edge intertwiners blockwise in the stabilizer branch", Passed: len(a.Edges) == 3 && a.Edges[0].Name == "Y_+3" && a.Edges[1].Name == "Y_-3" && a.Edges[2].Name == "Y_-1" && allEdgesBlockwiseOnly(a.Edges), Detail: FormatEdges(a.Edges)},
			{Name: "separate allowed nonzero D-rho commutator from first-order obstruction", Passed: a.FirstOrder.DCommutatorExpectedNonzero && a.FirstOrder.NonzeroCommutatorAllowed && a.FirstOrder.SupportFirstOrderCompatible && !a.FirstOrder.OperatorFirstOrderCertified && containsAll(a.FirstOrder.Supports, []string{SupportOneFormCommutatorAllowed, SupportStabilizerFirstOrderSupportCompatibility}), Detail: FormatFirstOrder(a.FirstOrder)},
			{Name: "keep J-opposite and bimodule proof at support firewall level", Passed: a.FirstOrder.OppositeSupportAuditable && !a.FirstOrder.OppositeOperatorCertified && !a.FirstOrder.BimoduleCertified && containsAll(a.FirstOrder.Failures, []string{FailureNoCompleteJOppositeProof, FailureNoBimoduleProof, FailureNoFullOperatorFirstOrderTheorem}), Detail: FormatFirstOrder(a.FirstOrder)},
			{Name: "preserve chiral neutral puncture/kernel pair without physical promotion", Passed: a.NeutralPair.RightPuncture == "e_+ tensor P_1" && a.NeutralPair.LeftKernel == "h_+ tensor P_1" && a.NeutralPair.RightPuncturePreserved && a.NeutralPair.LeftKernelPreserved && !a.NeutralPair.LeftKernelOperatorStable && !a.NeutralPair.PhysicalParticleTheorem && containsAll(a.NeutralPair.Failures, []string{FailureNoParticleAssignment, FailureNoNeutrinoTheorem}), Detail: FormatNeutralPair(a.NeutralPair)},
			{Name: "preserve carrier ranks and symbolic D_F kernel ledger", Passed: a.Carrier.HLRank == HLRank && a.Carrier.HRMinRank == HRMinRank && a.Carrier.HPartMinRank == HPartMinRank && a.Carrier.HFMinRank == HFMinRank && a.Carrier.AmbientPartRank == AmbientPartRank && a.Carrier.AmbientFRank == AmbientFRank && a.Carrier.DSymRank == DSymRank && a.Carrier.KernelRank == DSymKernelRank, Detail: FormatCarrier(a.Carrier)},
			{Name: "freeze official ledgers and block R3/R4 promotion", Passed: a.Ledger.OfficialFrozen && !a.Ledger.AlphaNative && !a.Ledger.R3 && !a.Ledger.R4 && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 857 stabilizer-support first-order firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.FullAFNotTarget && a.Firewalls.AForientNotFullAF && a.Firewalls.FullHSocketFirewall && a.Firewalls.SupportOnly && a.Firewalls.NoOperatorFirstOrder && a.Firewalls.NoJOppositeProof && a.Firewalls.NoBimoduleProof && a.Firewalls.NoNativeTriple && a.Firewalls.CharacterMatchSupportOnly && a.Firewalls.IntertwinerNoValue && a.Firewalls.YSymbolicOnly && a.Firewalls.NoAlphaSource && a.Firewalls.NoTraceReadout && a.Firewalls.NoOfficialNEffUpdate && a.Firewalls.NoCYukawaCHiggsUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoParticleAssignment && a.Firewalls.NoNeutrinoTheorem && a.Firewalls.NoThreeGenerationTheorem && a.Firewalls.Verdict == StatusFirewallVerdict, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatAlgebra(a.Algebra), FormatPreservation(a.Preserve), FormatEdges(a.Edges), FormatFirstOrder(a.FirstOrder), FormatNeutralPair(a.NeutralPair), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func allEdgesBlockwiseOnly(edges []EdgeIntertwiner) bool {
	for _, e := range edges {
		if !e.BlockwiseCompatible || e.OperatorIntertwiner || e.CharacterMatchCertified {
			return false
		}
	}
	return true
}
