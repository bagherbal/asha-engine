package generation2airlockadmissiblesupportlatticesourceaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE931-AIRLOCK-ADMISSIBLESUPPORT-LATTICE-SOURCE-AUDIT"
	theoremName = "Gate 931: Airlock AdmissibleSupport Lattice Source Audit"
)

func Generation2AirlockAdmissibleSupportLatticeSourceAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 931 audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}
		checks := []theorem.Check{
			{Name: "admissible support chain is puncture-rooted", Passed: a.PunctureRoot.RootedAtPuncture && a.PunctureRoot.F0Rank == RankF0 && a.PunctureRoot.RelativeTargets && !a.PunctureRoot.NativeTheorem && containsAll(a.PunctureRoot.Supports, []string{SupportChainRootedAtPuncture, SupportF0EqualsPuncture, SupportActiveTargetsRelativeToF0}) && containsAll(a.PunctureRoot.Failures, []string{FailurePunctureRootednessNotNative}), Detail: FormatPunctureRoot(a.PunctureRoot)},
			{Name: "same-socket completion forces F1 and rank-three exposed quotient", Passed: a.SameSocket.CompletesSameSocket && a.SameSocket.F1EqualsPhaseW && a.SameSocket.F1Rank == RankF1 && a.SameSocket.F1OverF0Rank == RankF1OverF0 && !a.SameSocket.NativeTheorem && containsAll(a.SameSocket.Supports, []string{SupportSameSocketForcesF1, SupportF1OverF0EqualsP3, SupportRankThreeFromSameSocket}) && containsAll(a.SameSocket.Failures, []string{FailureSameSocketCompletionNotNative}), Detail: FormatSameSocket(a.SameSocket)},
			{Name: "admissible supports are tensor-structured completions", Passed: a.TensorIntegrity.StructuredCompletions && !a.TensorIntegrity.ArbitrarySubspaces && a.TensorIntegrity.PreservesSocketW && !a.TensorIntegrity.NativeTheorem && containsAll(a.TensorIntegrity.Supports, []string{SupportTensorStructuredNotArbitrary, SupportPartialOppositeFragmentsExcluded, SupportSocketWIntegrity}) && containsAll(a.TensorIntegrity.Failures, []string{FailureTensorFactorIntegrityNotNative}), Detail: FormatTensorIntegrity(a.TensorIntegrity)},
			{Name: "orphan opposite-socket fragments are excluded", Passed: a.NoOrphan.ExcludesOppositeLeptonSingleton && a.NoOrphan.ExcludesOppositeColorFragment && a.NoOrphan.OppositeOnlyAtFullSaturation && !a.NoOrphan.NativeTheorem && containsAll(a.NoOrphan.Supports, []string{SupportOrphanOppositeExcluded, SupportNoPartialOppositeSocketLevel, SupportOppositeOnlyAtFullSaturation}) && containsAll(a.NoOrphan.Failures, []string{FailureNoOrphanRuleNotNative}), Detail: FormatNoOrphan(a.NoOrphan)},
			{Name: "full boundary-pair activation forces F2 and rank-seven quotient", Passed: a.Saturation.FullPairForcesF2 && a.Saturation.FullRectangle && a.Saturation.RejectsStopAtF1 && a.Saturation.F2Rank == RankF2 && a.Saturation.F2OverF0Rank == RankF2OverF0 && !a.Saturation.NativeTheorem && containsAll(a.Saturation.Supports, []string{SupportFullPairForcesF2, SupportF2FullRightRectangle, SupportF2OverF0RankSeven, SupportRankSevenFromSaturation}) && containsAll(a.Saturation.Failures, []string{FailureFullRectangleSaturationNotNative}), Detail: FormatSaturation(a.Saturation)},
			{Name: "F0 F1 F2 form a minimal sufficient airlock support chain", Passed: a.MinimalChain.MinimalSufficient && a.MinimalChain.NoExtraIntermediate && a.MinimalChain.ThreeLevelCollapse && !a.MinimalChain.NativeUniquenessProof && containsAll(a.MinimalChain.Supports, []string{SupportMinimalSufficientChain, SupportNoExtraIntermediateLevel, SupportAdmissibleLatticeThreeLevel}) && containsAll(a.MinimalChain.Failures, []string{FailureMinimalSufficientChainNotNative}), Detail: FormatMinimalChain(a.MinimalChain)},
			{Name: "admissible support lattice descends to Z2 class", Passed: a.Z2Lattice.DescendsToZ2Class && a.Z2Lattice.PhaseFlipExchangesF0F1 && a.Z2Lattice.PhaseFlipFixesF2 && a.Z2Lattice.RanksRepresentativeFree && !a.Z2Lattice.NativePhaseTheorem && containsAll(a.Z2Lattice.Supports, []string{SupportAdmissibleLatticeDescendsToZ2, SupportPhaseFlipExchangesRepresentatives, SupportRanksZ2Independent}) && containsAll(a.Z2Lattice.Failures, []string{FailureZ2SupportLatticeNotNative}), Detail: FormatZ2Lattice(a.Z2Lattice)},
			{Name: "Theta_B^Z2 and mu_B are reconstructed from sourced support lattice while native alpha remains blocked", Passed: a.Measure.ThetaRecovered && a.Measure.ClosureDomainSourced && a.Measure.MuBReconstructed && !a.Measure.NativeAlpha && a.Measure.ThetaOneRank == RankF1OverF0 && a.Measure.ThetaTwoRank == RankF2OverF0 && a.Measure.H10Rank == RankH10 && a.Measure.H72Rank == RankH72 && containsAll(a.Measure.Supports, []string{SupportThetaMuBReconstructed, SupportClosureDomainFromLattice}) && containsAll(a.Measure.Failures, []string{FailureNoNativeAdmissibleSupportLattice, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}) && firewallsOK(a.Firewalls), Detail: FormatMeasure(a.Measure) + " | " + FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, a.Inherited, AmbientRightSupportRectangle, CR2Decomposition, WDecomposition, AtomicCells, Z2PunctureClass, RepresentativePuncture, AdmissibleSupportChain, Z2AdmissibleSupportLattice, ClosureOperatorConsequence, ThetaFromSupportLattice, MuBFromSupportLattice, NextGate, FormatPunctureRoot(a.PunctureRoot), FormatSameSocket(a.SameSocket), FormatTensorIntegrity(a.TensorIntegrity), FormatNoOrphan(a.NoOrphan), FormatSaturation(a.Saturation), FormatMinimalChain(a.MinimalChain), FormatZ2Lattice(a.Z2Lattice), FormatMeasure(a.Measure), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
