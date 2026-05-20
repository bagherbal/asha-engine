package generation2flavorbranchcompatibilityselectoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2FlavorBranchCompatibilitySelectorAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 flavor branch-compatibility selector audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate601 flavor branch-compatibility selector audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate600 branch/chamber monodromy boundary", Passed: a.Inherited.BranchSealDefined && a.Inherited.EpsilonBranchAlgebraic && !a.Inherited.NativeBranchTheorem && !a.Inherited.NativeFourthRoot && !a.Inherited.NativeChamberSelector && !a.Inherited.BFlavNative, Detail: FormatInherited(a.Inherited)},
			{Name: "define branch balance function", Passed: a.Definition.Environmental && !a.Definition.Native && a.Definition.Formula != "", Detail: FormatDefinition(a.Definition)},
			{Name: "enumerate charged-lepton branch permutations", Passed: len(a.LeptonBranches) == 6 && allLeptonBranchesPositive(a.LeptonBranches), Detail: FormatLeptonBranches(a.LeptonBranches)},
			{Name: "enumerate PMNS projector overlaps", Passed: len(a.PMNSOverlaps) == 3 && a.PMNSOverlaps[2].Index == 3 && a.PMNSOverlaps[2].Li > 0, Detail: FormatPMNSTable(a.PMNSOverlaps)},
			{Name: "enumerate CKM orientation signs", Passed: len(a.CKMSigns) == 2 && a.CKMSigns[0].Sign == +1 && a.CKMSigns[1].Sign == -1, Detail: FormatCKMSigns(a.CKMSigns)},
			{Name: "compute full branch balance table", Passed: len(a.BalanceTable) == 36 && a.BalanceTable[0].NeutrinoI == 3 && a.BalanceTable[0].CKMSign == +1, Detail: FormatTopBranchBalances(a.BalanceTable, 8)},
			{Name: "rank observed branch", Passed: a.ObservedRank.Rank == 1 && a.ObservedRank.MinimalClassSize == 6 && !a.ObservedRank.Unique, Detail: FormatObservedRank(a.ObservedRank)},
			{Name: "audit gap to next distinct branch", Passed: a.Gap.GapLarge && a.Gap.GapToNextDistinct > 1e-5, Detail: FormatGap(a.Gap)},
			{Name: "compile selector verdict", Passed: a.SelectorVerdict.ObservedInMinimalClass && a.SelectorVerdict.SelectsNeutrinoThirdProjector && a.SelectorVerdict.SelectsPositiveCKMSign && !a.SelectorVerdict.SelectsChargedLeptonOrdering && !a.SelectorVerdict.UniqueBranchSelector && !a.SelectorVerdict.NativeBranchSelector, Detail: FormatSelectorVerdict(a.SelectorVerdict)},
			{Name: "preserve flavor branch firewalls", Passed: !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesChargedLeptonMasses && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesCKM && !a.Firewalls.DerivesNeutrinoData && !a.Firewalls.DerivesBFlavZeroNative && !a.Firewalls.PromotesObservedData && !a.Firewalls.AddsCarrier && !a.Firewalls.AddsSelector && a.Firewalls.PreservesGate352 && a.Firewalls.PreservesGate596 && a.Firewalls.PreservesGate600, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth, a.SelectorVerdict.Decision)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func allLeptonBranchesPositive(rows []ChargedLeptonBranch) bool {
	for _, r := range rows {
		if !r.PositiveChamber {
			return false
		}
	}
	return true
}
