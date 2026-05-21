package generation2minimalrightneutralabsencesealedgeskeletonaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_MINIMAL_RIGHT_NEUTRAL_ABSENCE_SEAL_EDGE_SKELETON_AUDIT"
	theoremName = "Gate 843 — Minimal RightNeutral Absence Seal and Edge-Skeleton Audit"
)

func Generation2MinimalRightNeutralAbsenceSealAndEdgeSkeletonAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		checks := []theorem.Check{}
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks = append(checks,
			theorem.Check{Name: "inherit Gate 842 four-cell right rectangle", Passed: a.Rectangle.InheritedFromGate842 && a.Rectangle.FullRank == 8 && a.Rectangle.RankPattern == "8=3+1+3+1" && a.Rectangle.Complete && a.Rectangle.Orthogonal && containsAll(a.Rectangle.Supports, []string{SupportGate842Inherited}), Detail: FormatRectangle(a.Rectangle)},
			theorem.Check{Name: "certify active seven as eight minus neutral singleton", Passed: a.Rectangle.ActiveRank == 7 && a.Rectangle.PunctureRank == 1 && a.Rectangle.ActiveIsFullMinusPuncture && a.Rectangle.MinimalPattern == "7=8-1=3+3+1" && containsAll(a.Rectangle.Supports, []string{SupportHRMinRankSeven, SupportPunctureIsCompensatingSingleton}), Detail: FormatRectangle(a.Rectangle)},
			theorem.Check{Name: "preserve B-L compensation", Passed: a.Rectangle.BMinusLActive == 1 && a.Rectangle.BMinusLPuncture == -1 && a.Rectangle.BMinusLFull == 0, Detail: FormatRectangle(a.Rectangle)},
			theorem.Check{Name: "compare minimal absent-cell and extended neutral-inclusive branches", Passed: a.Branches.MinimalBranchAdmittedAsSeal && !a.Branches.MinimalBranchNative && a.Branches.MinimalRank == 7 && a.Branches.ExtendedBranchAvailable && a.Branches.ExtendedRank == 8 && !a.Branches.ExtendedBranchMatchesR2Support && a.Branches.ExtendedBranchNeedsExtraProjectionOrExclusion && a.Branches.R2PlusPlusPrefersMinimalBranch && containsAll(a.Branches.Supports, []string{SupportMinimalRightNeutralAbsenceSeal, SupportExtendedNeutralInclusiveRankEight, SupportR2PrefersMinimalBranch}) && containsAll(a.Branches.Failures, []string{FailureAbsenceSealNotNative, FailureExtendedBranchNeedsProjectionLaw, FailureNoNativeMinimalAbsenceTheorem}), Detail: FormatBranches(a.Branches)},
			theorem.Check{Name: "register dominant/rest finite-body location at seal level", Passed: a.Orientation.DominantLocationSealed && a.Orientation.RestLocationSealed && a.Orientation.DominantRank == 3 && a.Orientation.RestRank == 4 && a.Orientation.TotalRank == 7 && !a.Orientation.NativeOrientationTheorem && !a.Orientation.PhysicalParticleAssignment && containsAll(a.Orientation.Supports, []string{SupportDominantTripletAtSealLevel, SupportRestQuartetAtSealLevel}), Detail: FormatOrientation(a.Orientation)},
			theorem.Check{Name: "audit D_F edge skeleton without null-edge over-certification", Passed: !a.Edge.DFEdgeGraphAvailable && !a.Edge.ExplicitDFMatrixAvailable && !a.Edge.NullEdgeCertified && !a.Edge.MinimalAbsenceEdgeCertified && a.Edge.AbsentNullEdgeCandidateOnly && !a.Edge.PhysicalRightNeutrinoTheorem && containsAll(a.Edge.Failures, []string{FailureNoDFEdgeGraph, FailureNoNullEdgeTheorem, FailureNoNativeMinimalAbsenceTheorem, FailureNoExplicitDFMatrix, FailureNoPhysicalParticleAssignment, FailureNoRightNeutrinoTheorem}), Detail: FormatEdge(a.Edge)},
			theorem.Check{Name: "classify aggregate as sealed finite-body shadow, not native compression", Passed: a.Placement.FiniteBodyLocationAtSealLevel && a.Placement.TraceCompressionShadowAtSealLevel && !a.Placement.NativeCompressionTheorem && !a.Placement.AlphaDerived && !a.Placement.TraceMagnitudeReadout && !a.Placement.R3 && !a.Placement.R4 && containsAll(a.Placement.Failures, []string{FailureCompressionSealNotNativeMap, FailureNoAggregateCompressionTheorem, FailureNoAlphaDerivation, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureNotR3, FailureNotR4}), Detail: FormatPlacement(a.Placement)},
			theorem.Check{Name: "preserve official ledger freeze", Passed: a.Ledger.OfficialFrozen && a.Ledger.R2PlusPlus && !a.Ledger.R3 && !a.Ledger.R4 && !a.Ledger.AlphaNative && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			theorem.Check{Name: "preserve Gate 843 firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.AbsenceSealNotNative && a.Firewalls.NoDFEdgeGraph && a.Firewalls.NoNullEdgeTheorem && a.Firewalls.NoNativeMinimalAbsenceTheorem && a.Firewalls.NoPhysicalParticleAssignment && a.Firewalls.NoRightNeutrinoTheorem && a.Firewalls.ExtendedBranchNeedsProjectionLaw && a.Firewalls.NoFullRhoFActionLedger && a.Firewalls.NoExplicitDFMatrix && a.Firewalls.CompressionSealNotNativeMap && a.Firewalls.NoAggregateCompressionTheorem && a.Firewalls.NoAlphaDerivation && a.Firewalls.AlphaStillSealed && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoObservedYukawaFit && a.Firewalls.NoThreeGenerationTheorem && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.Verdict == StatusFirewallGate843, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatRectangle(a.Rectangle), FormatBranches(a.Branches), FormatOrientation(a.Orientation), FormatEdge(a.Edge), FormatPlacement(a.Placement), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
