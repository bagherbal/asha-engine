package generation2puncturedsocketresponsefunctionalaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_PUNCTURED_SOCKET_RESPONSE_FUNCTIONAL_AUDIT"
	theoremName = "Gate 846 — Punctured Socket Response Functional Audit"
)

func Generation2PuncturedSocketResponseFunctionalAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		checks := []theorem.Check{}
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks = append(checks,
			theorem.Check{Name: "inherit Gate 845 finite-body support and reconstruct response table", Passed: a.Table.ReconstructsGate845 && a.Table.ActiveRank == 7 && a.Table.PunctureRank == 1 && a.Table.RightFullRank == 8 && containsAll(a.Table.Supports, []string{SupportPuncturedTableShape, SupportTopColorIdentityCell, SupportRestBLTransferCell, SupportPunctureAbsentCell}), Detail: FormatTable(a.Table)},
			theorem.Check{Name: "verify punctured socket cell eigenvalues", Passed: a.Table.TopColor.Included && a.Table.TopColor.Rank == 3 && nearly(a.Table.TopColor.Eigenvalue, 1) && a.Table.Puncture.Rank == 1 && !a.Table.Puncture.Included && a.Table.RestColor.Rank == 3 && nearly(a.Table.RestColor.Eigenvalue, AlphaB*(1-AlphaB)) && a.Table.RestLepton.Rank == 1 && nearly(a.Table.RestLepton.Eigenvalue, 3*AlphaB*AlphaB), Detail: FormatTable(a.Table)},
			theorem.Check{Name: "audit B-L transfer identity on rest socket", Passed: a.Functional.BLTransferTraceZero && a.Functional.UsesRestBLTransfer && containsAll(a.Functional.Supports, []string{SupportBLIdentity, SupportRestBLTransferCell}), Detail: FormatFunctional(a.Functional)},
			theorem.Check{Name: "formal compression functional reconstructs but remains seal", Passed: a.Functional.ReconstructsTable && !a.Functional.NativeFunctional && !a.Functional.AlphaDerived && !a.Functional.PunctureDerived && containsAll(a.Functional.Supports, []string{SupportFunctionalProjectorExpression}) && containsAll(a.Functional.Failures, []string{FailureFunctionalSealNotNative, FailureNoNativeCompressionFunctional, FailureFunctionalDoesNotDeriveAlpha, FailurePunctureStillSealed}), Detail: FormatFunctional(a.Functional)},
			theorem.Check{Name: "reproduce trace, square trace, and diagnostic operator N_eff", Passed: nearly(a.Table.Trace, 3+3*AlphaB) && nearly(a.Table.SquareTrace, 3+3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) && nearly(a.Table.OperatorNEff, a.Ledger.OperatorNEff) && a.Ledger.OperatorNEff != a.Ledger.OfficialNEff, Detail: FormatTable(a.Table) + " | " + FormatLedger(a.Ledger)},
			theorem.Check{Name: "active response cells are compatible with Gate 844 support-only edge graph", Passed: a.Edges.CompatibleWithGate844 && a.Edges.SupportOnly && a.Edges.ActiveCells == 3 && a.Edges.ActiveCellsHaveTargets && !a.Edges.PunctureHasTarget && !a.Edges.ExplicitDFMatrix && !a.Edges.FirstOrderCertified && !a.Edges.BimoduleCommutantCertified && !a.Edges.Magnitudes && containsAll(a.Edges.Supports, []string{SupportActiveCellsEdgeCompatible}) && containsAll(a.Edges.Failures, []string{FailureDFSupportNotMatrix, FailureNoExplicitDFMatrix, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof, FailureDFEdgeSupportNotYukawa}), Detail: FormatEdges(a.Edges)},
			theorem.Check{Name: "preserve alpha, puncture, and magnitude firewalls", Passed: !a.Ledger.AlphaNative && a.Impact.AlphaStillSealed && a.Impact.PunctureStillSealed && a.Impact.MagnitudesStillMissing && containsAll(a.Table.Failures, []string{FailureNoSourceForTable, FailureAlphaStillSealed, FailurePunctureStillSealed}) && containsAll(a.Functional.Failures, []string{FailureFunctionalDoesNotDeriveAlpha}), Detail: FormatImpact(a.Impact)},
			theorem.Check{Name: "preserve official ledger freeze and no promotions", Passed: a.Ledger.OfficialFrozen && a.Ledger.R2PlusPlusPlusPlus && !a.Ledger.R3 && !a.Ledger.R4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			theorem.Check{Name: "preserve Gate 846 firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.FunctionalSealNotNative && a.Firewalls.NoNativeCompressionFunctional && a.Firewalls.NoSourceForTable && a.Firewalls.AlphaStillSealed && a.Firewalls.FunctionalDoesNotDeriveAlpha && a.Firewalls.PunctureStillSealed && a.Firewalls.DFSupportNotMatrix && a.Firewalls.NoExplicitDFMatrix && a.Firewalls.NoFirstOrderProof && a.Firewalls.NoBimoduleCommutantProof && a.Firewalls.DFEdgeSupportNotYukawa && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.NoNumericalYukawaValues && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoPhysicalParticleAssignment && a.Firewalls.NoRightNeutrinoTheorem && a.Firewalls.NoThreeGenerationTheorem && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.Verdict == StatusFirewallGate846, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatTable(a.Table), FormatFunctional(a.Functional), FormatEdges(a.Edges), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func nearly(a, b float64) bool {
	if a > b {
		return a-b < 1e-12
	}
	return b-a < 1e-12
}
