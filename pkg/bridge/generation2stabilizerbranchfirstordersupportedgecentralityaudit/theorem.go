package generation2stabilizerbranchfirstordersupportedgecentralityaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_STABILIZER_BRANCH_FIRST_ORDER_SUPPORT_EDGE_CENTRALITY_AUDIT"
	theoremName = "Gate 859 — Stabilizer-Branch First-Order Support Calculation and Edge-Centrality Audit"
)

func Generation2StabilizerBranchFirstOrderSupportEdgeCentralityAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit A_F^orient support layer and keep full-unbroken firewall", Passed: a.Algebra.Algebra == "A_F^orient=C_R plus C_H plus M_3(C)" && a.Algebra.PostOrientation && !a.Algebra.ContainsFullH && a.Algebra.SupportPreservesSockets && a.Algebra.SupportPreservesP && containsAll(a.Algebra.Failures, []string{FailureAForientNotFullAF}), Detail: FormatAlgebra(a.Algebra)},
			{Name: "audit first-order support target while separating one-form commutator", Passed: a.FirstOrder.Expression == "[[D_F^sym,rho_F(a)],rho_F^op(b)]=0" && a.FirstOrder.OrderZeroInherited && a.FirstOrder.SupportAuditable && a.FirstOrder.DRhoCommutatorAllowedOneFormSource && a.FirstOrder.FirstOrderSupportConditionAudited && !a.FirstOrder.OperatorTheoremCertified && containsAll(a.FirstOrder.Supports, []string{SupportFirstOrderIfCentral, SupportNonzeroDRhoOneForm}) && containsAll(a.FirstOrder.Failures, []string{FailureNoFullOperatorFirstOrder, FailureNoCompleteJOppositeProof}), Detail: FormatFirstOrder(a.FirstOrder)},
			{Name: "require color-edge centrality on the P3 factor", Passed: colorEdgesCentral(a.Edges), Detail: FormatEdges(a.Edges)},
			{Name: "audit lepton-edge triviality on P1 and preserve puncture zero edge", Passed: leptonAndPunctureOK(a.Edges), Detail: FormatEdges(a.Edges)},
			{Name: "preserve puncture and left kernel in oriented support branch", Passed: a.PunctureKernel.PunctureCoefficientZero && !a.PunctureKernel.PunctureReintroduced && a.PunctureKernel.RightPunctureOutsideMinimal && a.PunctureKernel.LeftKernelPresent && a.PunctureKernel.KernelRank == KernelRank, Detail: FormatPunctureKernel(a.PunctureKernel)},
			{Name: "freeze ledgers and block R3/R4 promotion", Passed: a.Ledger.OfficialFrozen && !a.Ledger.AlphaNative && !a.Ledger.R3 && !a.Ledger.R4 && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 859 edge-centrality firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.AForientNotFullAF && a.Firewalls.NoFullFirstOrder && a.Firewalls.NoCompleteJOpposite && a.Firewalls.NoBimoduleProof && a.Firewalls.EdgeCentralitySupportOnly && a.Firewalls.CharacterMatchSupportOnly && a.Firewalls.SupportIntertwinerNotOperator && a.Firewalls.YSymbolicOnly && a.Firewalls.NoAlphaSource && a.Firewalls.NoTraceReadout && a.Firewalls.NoOfficialNEffUpdate && a.Firewalls.NoCYukawaCHiggsUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoParticleAssignment && a.Firewalls.NoNeutrinoTheorem && a.Firewalls.NoThreeGenerationTheorem && a.Firewalls.Verdict == StatusFirewallVerdict, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatAlgebra(a.Algebra), FormatFirstOrder(a.FirstOrder), FormatEdges(a.Edges), FormatPunctureKernel(a.PunctureKernel), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func colorEdgesCentral(edges []EdgeCentrality) bool {
	seen := map[string]bool{}
	for _, e := range edges {
		if e.ColorEdge {
			seen[e.Name] = true
			if !e.Present || !e.CentralOnFactor || !e.IdentityOnFactor || e.OperatorIntertwinerCertified || e.YukawaMagnitude || !containsAll(e.Supports, []string{SupportColorEdgesCentral, SupportYPlusMinusIdentityOnColor}) {
				return false
			}
		}
	}
	return seen["Y_+3"] && seen["Y_-3"]
}

func leptonAndPunctureOK(edges []EdgeCentrality) bool {
	lepton, puncture := false, false
	for _, e := range edges {
		if e.LeptonEdge {
			lepton = e.Name == "Y_-1" && e.Present && e.CentralOnFactor && e.IdentityOnFactor && !e.OperatorIntertwinerCertified && !e.YukawaMagnitude && containsAll(e.Supports, []string{SupportLeptonEdgeP1Compatible})
		}
		if e.PunctureEdge {
			puncture = e.Name == "Y_+1" && !e.Present && e.RequiredForm == "0" && e.CentralOnFactor && !e.IdentityOnFactor
		}
	}
	return lepton && puncture
}
