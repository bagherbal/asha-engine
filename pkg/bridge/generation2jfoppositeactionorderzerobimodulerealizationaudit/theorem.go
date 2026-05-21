package generation2jfoppositeactionorderzerobimodulerealizationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_JF_OPPOSITE_ACTION_ORDER_ZERO_BIMODULE_REALIZATION_AUDIT"
	theoremName = "Gate 858 — J_F-Opposite Action and Order-Zero Bimodule Realization Audit"
)

func Generation2JFOppositeActionOrderZeroBimoduleRealizationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit A_F^orient stabilizer layer and keep full-H firewall", Passed: a.Algebra.OrientedAlgebra == "A_F^orient=C_R plus C_H plus M_3(C)" && !a.Algebra.ContainsFullH && a.Algebra.ContainsCH && a.Algebra.ContainsRightC && a.Algebra.ContainsM3C && a.Algebra.PostOrientationLayer && containsAll(a.Algebra.Failures, []string{FailureAForientNotFullAF, FailureFullHSocketFirewall}), Detail: FormatAlgebra(a.Algebra)},
			{Name: "verify oriented left action preserves support blocks without becoming operator theorem", Passed: a.Left.PreservesHPlusHMinus && a.Left.PreservesEPlusEMinus && a.Left.PreservesP1P3 && a.Left.PreservesHRMin && a.Left.PreservesHF && !a.Left.LeftActionOperatorCertified && containsAll(a.Left.Supports, []string{StatusAForientLeftAction, SupportAForientSupportBimodule}), Detail: FormatLeftAction(a.Left)},
			{Name: "define rho_F^op as formal J-opposite support seal", Passed: a.Opposite.Expression == "rho_F^op(b)=J_F rho_F(b) J_F^{-1}" && a.Opposite.FormalJExchangeDefined && a.Opposite.OppositeSupportDefined && !a.Opposite.OppositeOperatorCertified && a.Opposite.OrderZeroTargetTyped && a.Opposite.MinimalCarrierClosedUnderJSeal && !a.Opposite.AmbientCellReintroduced && containsAll(a.Opposite.Failures, []string{FailureJOppositeSealOnly, FailureNoOperatorJOppositeProof}), Detail: FormatOpposite(a.Opposite)},
			{Name: "audit order-zero target before first-order retry", Passed: a.OrderZero.Expression == "[rho_F(a),rho_F^op(b)]=0" && a.OrderZero.SupportAuditable && a.OrderZero.BlockSupportCompatible && !a.OrderZero.OperatorTheoremCertified && a.OrderZero.RequiresOperatorJOpposite && containsAll(a.OrderZero.Supports, []string{SupportOrderZeroBlockSupport, SupportAForientSupportBimodule}) && containsAll(a.OrderZero.Failures, []string{FailureOrderZeroSupportOnly, FailureNoOrderZeroOperatorTheorem}), Detail: FormatOrderZero(a.OrderZero)},
			{Name: "audit active edge bimodule support without Yukawa promotion", Passed: len(a.Edges) == 3 && allEdgesSupportOnly(a.Edges), Detail: FormatEdges(a.Edges)},
			{Name: "preserve minimal 15/30 carrier under formal J copy without restoring ambient puncture", Passed: a.Carrier.HLRank == HLRank && a.Carrier.HRMinRank == HRMinRank && a.Carrier.HPartMinRank == HPartMinRank && a.Carrier.HFMinRank == HFMinRank && a.Carrier.AmbientPartRank == AmbientPartRank && a.Carrier.AmbientFRank == AmbientFRank && a.Carrier.RightPunctureOutsideMinimal && a.Carrier.LeftKernelPresent && !a.Carrier.JCopyRestoresAmbientPuncture && containsAll(a.Carrier.Failures, []string{FailureMinimalCarrierJSealOnly, FailureJCopyDoesNotRestoreAmbientCell}), Detail: FormatCarrier(a.Carrier)},
			{Name: "prepare next first-order target but do not prove it", Passed: a.FirstOrderNext.OrderZeroPrerequisiteAudited && a.FirstOrderNext.ReadyForOperatorFirstOrderAttempt && !a.FirstOrderNext.FirstOrderOperatorCertified && containsAll(a.FirstOrderNext.Failures, []string{FailureNoFirstOrderOperatorTheorem, FailureNoBimoduleCommutantProof}), Detail: FormatFirstOrderBoundary(a.FirstOrderNext)},
			{Name: "freeze ledgers and block R3/R4 promotion", Passed: a.Ledger.OfficialFrozen && !a.Ledger.AlphaNative && !a.Ledger.R3 && !a.Ledger.R4 && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 858 support-bimodule firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.AForientNotFullAF && a.Firewalls.FullHSocketFirewall && a.Firewalls.JOppositeSealOnly && a.Firewalls.NoOperatorJOpposite && a.Firewalls.NoOrderZeroOperator && a.Firewalls.OrderZeroSupportOnly && a.Firewalls.NoFirstOrderOperator && a.Firewalls.NoBimoduleProof && a.Firewalls.NoNativeTriple && a.Firewalls.MinimalJClosureSealOnly && a.Firewalls.JCopyNoAmbientRestore && a.Firewalls.EdgeSupportOnly && a.Firewalls.EdgeNoMagnitude && a.Firewalls.YSymbolicOnly && a.Firewalls.NoAlphaSource && a.Firewalls.NoTraceReadout && a.Firewalls.NoOfficialNEffUpdate && a.Firewalls.NoCYukawaCHiggsUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoParticleAssignment && a.Firewalls.NoNeutrinoTheorem && a.Firewalls.NoThreeGenerationTheorem && a.Firewalls.Verdict == StatusFirewallVerdict, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatAlgebra(a.Algebra), FormatLeftAction(a.Left), FormatOpposite(a.Opposite), FormatOrderZero(a.OrderZero), FormatEdges(a.Edges), FormatCarrier(a.Carrier), FormatFirstOrderBoundary(a.FirstOrderNext), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func allEdgesSupportOnly(edges []EdgeBimodule) bool {
	for _, e := range edges {
		if !e.LeftSupportCompatible || !e.RightSupportCompatible || e.OperatorIntertwinerCertified || e.YukawaMagnitude {
			return false
		}
	}
	return true
}
