package generation2masterenvironmentalhistorysealvectoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2MasterEnvironmentalHistorySealVectorAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 master environmental history seal vector audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate605 master environmental history seal vector audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate604 minimal flavor seal", Passed: a.Inherited.MinimalFlavorSealDefined && a.Inherited.SigmaGaugeForBFlav && a.Inherited.OptionalFullOrderSeal, Detail: FormatInherited(a.Inherited)},
			{Name: "construct master seal table", Passed: len(a.MasterSealTable) >= 25 && containsRow(a.MasterSealTable, "Cℓ(1,7)") && containsRow(a.MasterSealTable, "B_flav≈0") && containsRow(a.MasterSealTable, "RG thresholds and matching"), Detail: FormatMasterSealTable(a.MasterSealTable)},
			{Name: "clarify native/environmental boundary", Passed: a.Summary.BoundaryClear && a.Summary.NativeCount >= 10 && a.Summary.EnvironmentalSealCount >= 7 && a.Summary.ObservedLedgerCount >= 5, Detail: FormatSummary(a.Summary)},
			{Name: "write master history transport formula", Passed: a.Formula.Formula != "" && containsText(a.Formula.NativeLawInputs, "Cℓ(1,7)") && containsText(a.Formula.HistorySeals, "MinimalFlavorHistoryBranchSeal") && containsText(a.Formula.ObservedEndpointLedgers, "CKM"), Detail: FormatFormula(a.Formula)},
			{Name: "build solved/unsolved ledger", Passed: len(a.SolvedUnsolved) >= 12 && hasSolvedStatus(a.SolvedUnsolved, "B-L 4=1+3", "solved/native") && hasSolvedStatus(a.SolvedUnsolved, "H_e^(1/4)", "unsolved/native gap"), Detail: FormatSolvedUnsolved(a.SolvedUnsolved)},
			{Name: "rank next targets", Passed: len(a.Ranking) >= 5 && a.Ranking[0].Path == "RG / threshold transport" && a.Ranking[0].Verdict == StatusRGThresholdNextActionable, Detail: FormatRanking(a.Ranking)},
			{Name: "preserve master seal firewalls", Passed: !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesFlavor && !a.Firewalls.DerivesEWMasses && !a.Firewalls.DerivesCosmology && !a.Firewalls.DerivesObservedEndpoint && !a.Firewalls.SearchesNewConstants && a.Firewalls.PreservesGate352 && a.Firewalls.PreservesGate596 && a.Firewalls.PreservesGate604, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth, a.Formula.Formula)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func hasSolvedStatus(rows []SolvedUnsolvedRow, item string, status string) bool {
	for _, r := range rows {
		if r.Item == item && r.Status == status {
			return true
		}
	}
	return false
}
