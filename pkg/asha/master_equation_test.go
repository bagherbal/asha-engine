package asha

import "testing"

func TestMasterEquationLedgerPreservesNativeEnvironmentalBoundary(t *testing.T) {
	ledger := BuildMasterEquationLedger()
	if got := ValidateMasterEquationLedger(ledger); len(got) != 0 {
		t.Fatalf("master equation ledger validation failed: %v", got)
	}
	if ledger.FormulaGitHubLaTeX != MasterEquationGitHubLaTeX {
		t.Fatalf("unexpected master equation: %s", ledger.FormulaGitHubLaTeX)
	}
	for _, term := range append(append([]MasterEquationTerm{}, ledger.EnvironmentalTerms...), ledger.BridgeTerms...) {
		if term.NativeWriteAllowed {
			t.Fatalf("%s must not allow native writes", term.Symbol)
		}
	}
}
