package generation2unsealedleptonwallpmnsrowbranchselectoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2UnsealedLeptonWallPMNSRowBranchSelectorAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 unsealed lepton-wall / PMNS-row branch selector audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate602 unsealed lepton-wall PMNS-row branch selector audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate601 compatibility result", Passed: a.Inherited.ObservedMinimal && a.Inherited.SelectsP3 && a.Inherited.SelectsPositiveJ && !a.Inherited.SelectsLeptonOrdering && !a.Inherited.UniqueBranchSelector && !a.Inherited.NativeBranchSelector, Detail: FormatInherited(a.Inherited)},
			{Name: "define branch-row balance function", Passed: a.Definition.Environmental && !a.Definition.Native && a.Definition.Formula != "" && a.Definition.AlphaDomain != "", Detail: FormatDefinition(a.Definition)},
			{Name: "enumerate charged-lepton zero-wall candidates", Passed: len(a.WallCandidates) == 18 && allWallCandidatesPositive(a.WallCandidates), Detail: FormatWallCandidates(a.WallCandidates, 9)},
			{Name: "enumerate PMNS row/projector overlaps", Passed: len(a.PMNSOverlaps) == 9 && pmnsRowsNormalized(a.PMNSOverlaps), Detail: FormatPMNSTable(a.PMNSOverlaps)},
			{Name: "enumerate CKM signs", Passed: len(a.CKMSigns) == 2 && a.CKMSigns[0].Sign == +1 && a.CKMSigns[1].Sign == -1, Detail: FormatCKMSigns(a.CKMSigns)},
			{Name: "compute full branch-row balance table", Passed: len(a.BalanceTable) == 108 && a.BalanceTable[0].Alpha == "e" && a.BalanceTable[0].NeutrinoI == 3 && a.BalanceTable[0].CKMSign == +1, Detail: FormatTopBranchRows(a.BalanceTable, 10)},
			{Name: "rank observed tuple", Passed: a.ObservedRank.Rank == 1 && a.ObservedRank.MinimalClassSize == 6 && !a.ObservedRank.Unique && a.ObservedRank.ObservedAlpha == "e" && a.ObservedRank.ObservedNeutrinoI == 3 && a.ObservedRank.ObservedCKMSign == +1, Detail: FormatObservedRank(a.ObservedRank)},
			{Name: "audit gap to next distinct tuple", Passed: a.Gap.GapLarge && a.Gap.GapToNextDistinct > 1e-5, Detail: FormatGap(a.Gap)},
			{Name: "audit remaining degeneracy", Passed: a.Degeneracy.ElectronRowSelected && a.Degeneracy.P3Selected && a.Degeneracy.PositiveJSelected && a.Degeneracy.SigmaStillDegenerate && len(a.Degeneracy.DistinctSigmas) == 6, Detail: FormatDegeneracy(a.Degeneracy)},
			{Name: "compile selector verdict", Passed: a.SelectorVerdict.ObservedInMinimalClass && a.SelectorVerdict.SelectsElectronRow && a.SelectorVerdict.SelectsThirdNeutrinoProjector && a.SelectorVerdict.SelectsPositiveCKMSign && !a.SelectorVerdict.SelectsFullChargedLeptonSigma && !a.SelectorVerdict.UniqueSelector && !a.SelectorVerdict.NativeSelector, Detail: FormatSelectorVerdict(a.SelectorVerdict)},
			{Name: "preserve flavor branch-row firewalls", Passed: !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesChargedLeptonMasses && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesCKM && !a.Firewalls.DerivesNeutrinoData && !a.Firewalls.DerivesFlavor && !a.Firewalls.DerivesBFlavZeroNative && !a.Firewalls.PromotesObservedData && !a.Firewalls.AddsCarrier && !a.Firewalls.AddsSelector && a.Firewalls.PreservesGate352 && a.Firewalls.PreservesGate596 && a.Firewalls.PreservesGate600 && a.Firewalls.PreservesGate601, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth, a.SelectorVerdict.Decision)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func allWallCandidatesPositive(rows []ChargedLeptonWallCandidate) bool {
	for _, r := range rows {
		if !r.PositiveChamber {
			return false
		}
	}
	return true
}

func pmnsRowsNormalized(rows []PMNSRowProjectorOverlap) bool {
	sums := map[string]float64{}
	for _, r := range rows {
		sums[r.Alpha] += r.UAbs2
	}
	for _, s := range sums {
		if s < 0.999999999 || s > 1.000000001 {
			return false
		}
	}
	return len(sums) == 3
}
