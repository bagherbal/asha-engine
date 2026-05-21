package generation2leftneutralkernelsingletonchiralpuncturepairaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_LEFT_NEUTRAL_KERNEL_SINGLETON_CHIRAL_PUNCTURE_PAIR_AUDIT"
	theoremName = "Gate 849 — LeftNeutral Kernel Singleton and Chiral Puncture Pair Audit"
)

func Generation2LeftNeutralKernelSingletonChiralPuncturePairAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		checks := []theorem.Check{}
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks = append(checks,
			theorem.Check{Name: "inherit Gate 848 symbolic D_F rank anatomy", Passed: a.Impact.Gate848Inherited && a.Image.RightDomainRank == HRMinRank && a.Image.LeftTargetRank == HLRank && a.Kernel.TotalRank == ChiralTotalDim && a.Ledger.R2PlusPlusPlusPlusPlusKernel, Detail: FormatImage(a.Image) + " | " + FormatKernel(a.Kernel)},
			theorem.Check{Name: "force left neutral complement from rank 7 image in rank 8 target", Passed: a.Image.FullSupportRank && a.Image.ImageRank == 7 && a.Image.ComplementRank == 1 && a.Image.LeftComplement.Name == "h_+ tensor P_1" && a.Image.LeftComplement.Rank == 1 && a.Image.LeftComplement.Kernel && !a.Image.LeftComplement.InYImage && containsAll(a.Image.Supports, []string{SupportLeftNeutralKernelSingleton, SupportOneForcedLeftNullMode}), Detail: FormatImage(a.Image)},
			theorem.Check{Name: "verify active image cells exclude h_+ tensor P_1", Passed: activeCellsRank(a.Image.ActiveTargetCells) == 7 && len(a.Image.ActiveTargetCells) == 3 && activeCellsExclude(a.Image.ActiveTargetCells, "h_+ tensor P_1") && activeCellsInclude(a.Image.ActiveTargetCells, []string{"h_+ tensor P_3", "h_- tensor P_3", "h_- tensor P_1"}), Detail: FormatImage(a.Image)},
			theorem.Check{Name: "audit symbolic D_F kernel singleton", Passed: a.Kernel.SupportOnly && !a.Kernel.NativeDFMatrix && !a.Kernel.NumericalDFMatrix && a.Kernel.YRank == 7 && a.Kernel.DFRank == 14 && a.Kernel.KernelDim == 1 && a.Kernel.RightKernelDim == 0 && a.Kernel.LeftKernelDim == 1 && a.Kernel.KernelSupport.Name == "h_+ tensor P_1" && containsAll(a.Kernel.Supports, []string{SupportDFRankFourteenKernelOne}), Detail: FormatKernel(a.Kernel)},
			theorem.Check{Name: "compare right puncture with left kernel singleton", Passed: a.Pair.PairCandidate && !a.Pair.PhysicalParticleTheorem && a.Pair.RightPuncture.Name == "e_+ tensor P_1" && a.Pair.LeftKernel.Name == "h_+ tensor P_1" && a.Pair.SameLeptonSupport && a.Pair.SamePlusSocket && a.Pair.DifferentChirality && containsAll(a.Pair.Supports, []string{SupportChiralNeutralPair, SupportRightPunctureLeftKernel}), Detail: FormatPair(a.Pair)},
			theorem.Check{Name: "preserve physical naming, masslessness, and magnitude firewalls", Passed: !a.Pair.PhysicalParticleTheorem && !a.Impact.PhysicalNeutrinoTheorem && !a.Impact.MasslessTheorem && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing && containsAll(a.Pair.Failures, []string{FailureNoPhysicalNeutrinoTheorem, FailureNoRightNeutrinoTheorem, FailureNoMasslessnessTheorem}) && containsAll(a.Kernel.Failures, []string{FailureKernelNotYukawaMagnitude, FailureNoNumericalYukawaValues, FailureNoFirstOrderProof}), Detail: FormatPair(a.Pair) + " | " + FormatImpact(a.Impact)},
			theorem.Check{Name: "preserve official ledgers and no R3/R4 promotion", Passed: a.Ledger.OfficialFrozen && !a.Ledger.R3 && !a.Ledger.R4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			theorem.Check{Name: "preserve Gate 849 firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.KernelSealNotNative && a.Firewalls.NoNativeNullEdge && a.Firewalls.NoPhysicalNeutrino && a.Firewalls.NoRightNeutrino && a.Firewalls.NoMasslessnessTheorem && a.Firewalls.KernelNotYukawaMagnitude && a.Firewalls.NoNumericalYukawaValues && a.Firewalls.NoExplicitDFMatrix && a.Firewalls.NoFirstOrderProof && a.Firewalls.NoBimoduleProof && a.Firewalls.NoJOppositeProof && a.Firewalls.AlphaStillSealed && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.NoOfficialNEffUpdate && a.Firewalls.NoCYukawaCHiggsUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.Verdict == StatusFirewallGate849, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatImage(a.Image), FormatKernel(a.Kernel), FormatPair(a.Pair), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func activeCellsRank(cells []Cell) int {
	total := 0
	for _, c := range cells {
		if c.InYImage {
			total += c.Rank
		}
	}
	return total
}

func activeCellsExclude(cells []Cell, name string) bool {
	for _, c := range cells {
		if c.Name == name {
			return false
		}
	}
	return true
}

func activeCellsInclude(cells []Cell, names []string) bool {
	seen := map[string]bool{}
	for _, c := range cells {
		seen[c.Name] = true
	}
	for _, name := range names {
		if !seen[name] {
			return false
		}
	}
	return true
}
