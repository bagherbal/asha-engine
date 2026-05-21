package generation2finitebodyaggregatetracecompressionshadowmapaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_FINITE_BODY_AGGREGATE_TRACE_COMPRESSION_SHADOW_MAP_AUDIT"
	theoremName = "Gate 845 — Finite-Body Aggregate Trace-Compression Shadow Map Audit"
)

func Generation2FiniteBodyAggregateTraceCompressionShadowMapAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		checks := []theorem.Check{}
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks = append(checks,
			theorem.Check{Name: "inherit Gate 844 finite edge-domain and split into top/rest", Passed: a.Domain.InheritedGate844 && a.Domain.Orthogonal && a.Domain.CompleteOnHRMin && a.Domain.HRMinRank == 7 && a.Domain.Top.Rank == 3 && a.Domain.Rest.Rank == 4 && containsAll(a.Domain.Supports, []string{SupportHRMinAsTopPlusRest, SupportPiTopRankThree, SupportPiRestRankFour}), Detail: FormatDomain(a.Domain)},
			theorem.Check{Name: "exclude puncture from aggregate support", Passed: a.Domain.PunctureExcluded && !a.Domain.Puncture.Included && a.Domain.Puncture.Rank == 1 && a.Domain.Puncture.BMinusLTrace == -1 && containsAll(a.Domain.Supports, []string{SupportPunctureExcluded}), Detail: FormatDomain(a.Domain)},
			theorem.Check{Name: "place aggregate operator on finite-body support at seal level", Passed: a.Operator.FiniteBodyLocationAtSealLevel && !a.Operator.NativeCompressionTheorem && !a.Operator.NativeTraceCompressionFunctional && containsAll(a.Operator.Supports, []string{SupportFiniteBodyLocation, SupportRestBLTransferOnW}) && containsAll(a.Operator.Failures, []string{FailureCompressionSealNotNative, FailureNoNativeCompressionMap, FailureNoNativeShadowFunctional}), Detail: FormatOperator(a.Operator)},
			theorem.Check{Name: "reconstruct Gate 829 trace and square trace", Passed: nearly(a.Operator.TotalTrace, 3+3*AlphaB) && nearly(a.Operator.TotalSquareTrace, 3+3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) && containsAll(a.Operator.Supports, []string{SupportTraceReconstructionMatchesGate829}), Detail: FormatOperator(a.Operator)},
			theorem.Check{Name: "reproduce diagnostic operator N_eff without official update", Passed: nearly(a.Operator.OperatorNEff, a.Ledger.OperatorNEff) && a.Operator.OperatorNEff != a.Operator.OfficialNEff && a.Ledger.OfficialFrozen && containsAll(a.Operator.Supports, []string{SupportOperatorNEffDiagnostic}), Detail: FormatOperator(a.Operator) + " | " + FormatLedger(a.Ledger)},
			theorem.Check{Name: "verify compatibility with Gate 844 symbolic edge support", Passed: a.Edges.CompatibleWithGate844 && a.Edges.SupportOnly && !a.Edges.ExplicitDFMatrix && !a.Edges.FirstOrderCertified && !a.Edges.BimoduleCommutantCertified && !a.Edges.Magnitudes && containsAll(a.Edges.Supports, []string{SupportEdgeCompatibleDomain}) && containsAll(a.Edges.Failures, []string{FailureDFSupportNotMatrix, FailureNoExplicitDFMatrix, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof, FailureDFEdgeSupportNotYukawa}), Detail: FormatEdges(a.Edges)},
			theorem.Check{Name: "keep alpha and magnitude firewalls closed", Passed: !a.Operator.AlphaDerived && !a.Operator.TraceMagnitudeReadout && containsAll(a.Operator.Failures, []string{FailureAlphaStillSealed, FailureCompressionDoesNotDeriveAlpha, FailureNoTraceMagnitudeReadout}), Detail: FormatOperator(a.Operator)},
			theorem.Check{Name: "preserve official ledger freeze and no promotions", Passed: a.Ledger.OfficialFrozen && a.Ledger.R2PlusPlusPlus && !a.Ledger.R3 && !a.Ledger.R4 && !a.Ledger.AlphaNative && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			theorem.Check{Name: "preserve Gate 845 firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.CompressionSealNotNative && a.Firewalls.NoNativeCompressionMap && a.Firewalls.NoNativeShadowFunctional && a.Firewalls.AlphaStillSealed && a.Firewalls.CompressionDoesNotDeriveAlpha && a.Firewalls.DFSupportNotMatrix && a.Firewalls.NoExplicitDFMatrix && a.Firewalls.NoFirstOrderProof && a.Firewalls.NoBimoduleCommutantProof && a.Firewalls.DFEdgeSupportNotYukawa && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.NoNumericalYukawaValues && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoPhysicalParticleAssignment && a.Firewalls.NoRightNeutrinoTheorem && a.Firewalls.NoThreeGenerationTheorem && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.Verdict == StatusFirewallGate845, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatDomain(a.Domain), FormatOperator(a.Operator), FormatEdges(a.Edges), FormatImpact(a.Impact), a.Final}
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
