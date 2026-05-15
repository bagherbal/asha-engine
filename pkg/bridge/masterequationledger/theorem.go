package masterequationledger

import "github.com/bagherbal/asha-engine/pkg/theorem"

func AshaMasterEquationGeometryHistoryBoundaryLedgerTheorem() theorem.Theorem {
	const id = AuditID
	const name = "ASHA master equation geometry/history boundary ledger"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a := BuildDefault()
		checks := []theorem.Check{
			{Name: "compile GitHub-safe master equation ledger", Passed: len(a.Problems) == 0 && a.Ledger.FormulaGitHubLaTeX != "" && a.Ledger.FormulaPlain != "", Detail: FormatLedger(a.Ledger)},
			{Name: "classify native geometric law terms", Passed: len(a.Ledger.NativeTerms) >= 4, Detail: FormatTermClasses(a.Ledger)},
			{Name: "quarantine environmental and bridge moduli", Passed: len(a.Ledger.EnvironmentalTerms) >= 4 && len(a.Ledger.BridgeTerms) >= 2, Detail: "environmental and bridge terms require airlocks and forbid native writes"},
			{Name: "preserve OS/Wick/Hilbert/Hamiltonian/firewall boundary", Passed: len(a.Ledger.Firewalls) >= 7 && a.Ledger.NativeDeltaZero, Detail: "native_delta_zero=true; OS/Wick/Hilbert/Hamiltonian and Schwinger functions remain bridge/environmental"},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(Statuses(), a.Truth)}
	}}
}
