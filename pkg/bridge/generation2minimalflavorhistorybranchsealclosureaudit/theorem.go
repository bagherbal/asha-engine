package generation2minimalflavorhistorybranchsealclosureaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2MinimalFlavorHistoryBranchSealClosureAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 minimal flavor history branch seal closure audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate604 minimal flavor history branch seal audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate603 branch result", Passed: a.Inherited.SelectsElectronWall && a.Inherited.SelectsP3Nu && a.Inherited.SelectsPositiveJ && a.Inherited.SigmaGaugeForBFlav && a.Inherited.OptionalSignedDiscriminantSeal, Detail: FormatInherited(a.Inherited)},
			{Name: "construct flavor history branch stack", Passed: len(a.BranchStack) >= 8 && hasStackItem(a.BranchStack, "R_e=Q[Tr(H_e),Tr(H_e^2),Tr(H_e^3)]") && hasStackItem(a.BranchStack, "electron wall / epsilon_e"), Detail: FormatBranchStack(a.BranchStack)},
			{Name: "classify native, extension, seal, gauge and ledger layers", Passed: len(a.Classification) >= 8 && hasClassification(a.Classification, "sixfold sigma/cyclic Fourier presentation", false, true) && hasClassification(a.Classification, "charged-lepton trace ring R_e", true, false), Detail: FormatClassification(a.Classification)},
			{Name: "audit minimality for B_flav", Passed: requires(a.Minimality, "electron-wall coordinate epsilon_e", true) && requires(a.Minimality, "full charged-lepton sigma/cyclic order", false) && requiresFull(a.Minimality, "signed Vandermonde orientation", true), Detail: FormatMinimality(a.Minimality)},
			{Name: "define minimal flavor history branch seal", Passed: a.MinimalSeal.Name == "MinimalFlavorHistoryBranchSeal" && !a.MinimalSeal.IsNative && a.MinimalSeal.IsEnvironmental && contains(a.MinimalSeal.SelectedByBFlav, "electron wall / alpha=e") && contains(a.MinimalSeal.NotIncluded, "full sigma/cyclic charged-lepton Fourier presentation"), Detail: FormatMinimalSeal(a.MinimalSeal)},
			{Name: "define optional full-order discriminant seal", Passed: a.OptionalFullOrder.Name == "ChargedLeptonDiscriminantOrientationSeal" && !a.OptionalFullOrder.RequiredForBFlav && a.OptionalFullOrder.RequiredForFullOrder && !a.OptionalFullOrder.NativeTheoremPresent, Detail: FormatOptionalFullOrder(a.OptionalFullOrder)},
			{Name: "update history transport flavor formula", Passed: a.Formula.Formula != "" && contains(a.Formula.YCore, "epsilon_e") && contains(a.Formula.OmegaCore, "positive CKM orientation sign"), Detail: FormatFormula(a.Formula)},
			{Name: "preserve flavor branch firewalls", Passed: !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesChargedLeptonMasses && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesCKM && !a.Firewalls.DerivesBFlavZero && !a.Firewalls.AddsCarrier && !a.Firewalls.AddsSelector && a.Firewalls.PreservesGate352 && a.Firewalls.PreservesGate596 && a.Firewalls.PreservesGate599 && a.Firewalls.PreservesGate603, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth, a.MinimalSeal.Verdict, a.OptionalFullOrder.Statement)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func hasStackItem(rows []BranchStackRow, item string) bool {
	for _, r := range rows {
		if r.Item == item {
			return true
		}
	}
	return false
}
func hasClassification(rows []ClassificationRow, item string, native bool, gauge bool) bool {
	for _, r := range rows {
		if r.Item == item {
			return r.Native == native && r.GaugeConvention == gauge
		}
	}
	return false
}
func requires(rows []MinimalityRow, item string, want bool) bool {
	for _, r := range rows {
		if r.Item == item {
			return r.RequiredForBFlav == want
		}
	}
	return false
}
func requiresFull(rows []MinimalityRow, item string, want bool) bool {
	for _, r := range rows {
		if r.Item == item {
			return r.RequiredForFullOrderedHistory == want
		}
	}
	return false
}
func contains(xs []string, s string) bool {
	for _, x := range xs {
		if x == s {
			return true
		}
	}
	return false
}
