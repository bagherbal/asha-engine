package generation2minimalrightmodulefinitediracedgeskeletonaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_MINIMAL_RIGHT_MODULE_FINITE_DIRAC_EDGE_SKELETON_AUDIT"
	theoremName = "Gate 844 — Minimal RightModule Finite-Dirac Edge-Skeleton Audit"
)

func Generation2MinimalRightModuleFiniteDiracEdgeSkeletonAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		checks := []theorem.Check{}
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks = append(checks,
			theorem.Check{Name: "inherit Gate 843 minimal active right domain", Passed: a.Domain.InheritedFromGate843 && a.Domain.MinimalAbsenceSeal && a.Domain.ActiveRank == 7 && a.Domain.PunctureRank == 1 && a.Domain.ActiveIsFullMinusPuncture && a.Domain.ActivePattern == "7=3+3+1=8-1" && containsAll(a.Domain.Supports, []string{SupportHRMinAsRightEdgeDomain, SupportActiveSevenAsDomain}), Detail: FormatDomain(a.Domain)},
			theorem.Check{Name: "preserve B-L compensation on right domain", Passed: a.Domain.BMinusLActive == 1 && a.Domain.BMinusLPuncture == -1 && a.Domain.BMinusLFull == 0, Detail: FormatDomain(a.Domain)},
			theorem.Check{Name: "define left lepto-color target", Passed: a.Target.Complete && a.Target.LeptoColorPreserved && a.Target.Rank == 8 && a.Target.ColorRank == 6 && a.Target.LeptonRank == 2 && containsAll(a.Target.Supports, []string{SupportLeftTargetHL}), Detail: FormatTarget(a.Target)},
			theorem.Check{Name: "construct symbolic D_F edge support at seal level", Passed: a.Edges.ConstructedAtSealLevel && a.Edges.EdgeSupportOnly && a.Edges.DomainRank == 7 && a.Edges.TargetRank == 8 && len(a.Edges.Edges) == 3 && containsAll(a.Edges.Supports, []string{SupportSymbolicDFEdgeSkeleton, SupportDFSuppCouplingGraphOnly}), Detail: FormatEdges(a.Edges)},
			theorem.Check{Name: "preserve lepto-color support in symbolic edges", Passed: a.Edges.Edges[0].ColorLeptonPreserving && a.Edges.Edges[1].ColorLeptonPreserving && a.Edges.Edges[2].ColorLeptonPreserving && a.Edges.Edges[0].TargetExpression == "C_L^2 tensor P_3" && a.Edges.Edges[1].TargetExpression == "C_L^2 tensor P_3" && a.Edges.Edges[2].TargetExpression == "C_L^2 tensor P_1" && containsAll(a.Edges.Supports, []string{SupportColorSupportPreserved, SupportLeptonSupportPreserved}), Detail: FormatEdges(a.Edges)},
			theorem.Check{Name: "keep puncture absent but not null-edge certified", Passed: !a.Edges.PunctureInDomain && a.Edges.PunctureAbsenceCompatible && !a.Edges.PunctureNullEdgeCertified && !a.Edges.PunctureAbsenceDerivedFromDFNullEdge && containsAll(a.Edges.Supports, []string{SupportPunctureAbsenceCompatible, SupportPunctureAsAbsentOnly}) && containsAll(a.Edges.Failures, []string{FailurePunctureAbsenceNotFromNullEdge, FailureNoNativeNullEdgeTheorem}), Detail: FormatEdges(a.Edges)},
			theorem.Check{Name: "preserve D_F and first-order firewalls", Passed: !a.Edges.NativeDFMatrixCertified && !a.Edges.ExplicitDFMatrixCertified && !a.Edges.FirstOrderConditionCertified && !a.Edges.BimoduleCommutantCertified && containsAll(a.Edges.Failures, []string{FailureDFSupportSealNotNative, FailureNoExplicitDFMatrix, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof}), Detail: FormatEdges(a.Edges)},
			theorem.Check{Name: "classify aggregate as sealed edge-support shadow only", Passed: a.Shadow.FiniteBodyLocationAtSealLevel && a.Shadow.EdgeSupportSealLevel && !a.Shadow.NativeCompressionTheorem && !a.Shadow.AlphaDerived && !a.Shadow.TraceMagnitudeReadout && !a.Shadow.R3 && !a.Shadow.R4 && containsAll(a.Shadow.Supports, []string{SupportFiniteBodyShadowStrengthened}) && containsAll(a.Shadow.Failures, []string{FailureNoAggregateCompressionNative, FailureNoAlphaDerivation, FailureNoTraceMagnitudeReadout, FailureNotR3, FailureNotR4}), Detail: FormatShadow(a.Shadow)},
			theorem.Check{Name: "preserve official ledger freeze", Passed: a.Ledger.OfficialFrozen && a.Ledger.R2PlusPlus && !a.Ledger.R3 && !a.Ledger.R4 && !a.Ledger.AlphaNative && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			theorem.Check{Name: "preserve Gate 844 firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.DFSupportSealNotNative && a.Firewalls.NoExplicitDFMatrix && a.Firewalls.NoFirstOrderProof && a.Firewalls.NoBimoduleCommutantProof && a.Firewalls.PunctureAbsenceNotFromNullEdge && a.Firewalls.NoNativeNullEdgeTheorem && a.Firewalls.NoNativeMinimalAbsenceTheorem && a.Firewalls.NoPhysicalParticleAssignment && a.Firewalls.NoRightNeutrinoTheorem && a.Firewalls.NoFullRhoFActionLedger && a.Firewalls.NoGammaFJFPackage && a.Firewalls.DFEdgeSupportNotYukawa && a.Firewalls.NoNumericalYukawaValues && a.Firewalls.NoAlphaDerivation && a.Firewalls.AlphaStillSealed && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.NoAggregateCompressionNative && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoObservedYukawaFit && a.Firewalls.NoThreeGenerationTheorem && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.Verdict == StatusFirewallGate844, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatDomain(a.Domain), FormatTarget(a.Target), FormatEdges(a.Edges), FormatShadow(a.Shadow), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
