package bminuslweakselector

import "testing"

func TestGate258BMinusLWeakPlaneSelectorAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Summary.Gate257Inherited || !a.BMinusL.DerivedFiniteFockLedger || a.BMinusL.UsesObservedInput {
		t.Fatalf("expected inherited Gate 257 and native B-L ledger: %s / %s", FormatSummary(a.Summary), FormatBMinusL(a.BMinusL))
	}
	if a.ScalarSieve.InputCount != 8 || a.ScalarSieve.SurvivorCount != 2 || !a.ScalarSieve.Reduced || a.ScalarSieve.UniqueSelected {
		t.Fatalf("unexpected scalar B-L sieve: %s", FormatScalarSieve(a.ScalarSieve))
	}
	if a.WeakSieve.InputCount != 12 || a.WeakSieve.SurvivorCount != 6 || !a.WeakSieve.Reduced || a.WeakSieve.UniqueSelected {
		t.Fatalf("unexpected weak B-L sieve: %s", FormatWeakSieve(a.WeakSieve))
	}
	if a.CombinedSieve.InputWitnessCount != 96 || a.CombinedSieve.SurvivingWitnessCount != 12 || !a.CombinedSieve.Reduced || a.CombinedSieve.UniqueOrientation {
		t.Fatalf("unexpected combined witness sieve: %s", FormatCombined(a.CombinedSieve))
	}
	if !a.RestrictedScan.AllSurvivorsScanned || a.RestrictedScan.ResultCount != 36 || a.RestrictedScan.BranchCount != 3 {
		t.Fatalf("unexpected restricted scan inventory: %s", FormatRestrictedScan(a.RestrictedScan))
	}
	if a.RestrictedScan.ExactPolarized3PlaneResults != 0 || a.RestrictedScan.ExactFull3KernelResults != 0 || a.Summary.Neutral3PlaneDerived {
		t.Fatalf("B-L sieve must not force a three-plane: %s", FormatRestrictedScan(a.RestrictedScan))
	}
	if a.RestrictedScan.MaxPolarizedKernelComplexDim != 1 || a.RestrictedScan.MaxFullQ8vCKernelComplexDim != 2 {
		t.Fatalf("unexpected restricted kernel maxima: %s", FormatRestrictedScan(a.RestrictedScan))
	}
	if a.Firewall.ForcedWeakPlane || a.Firewall.ForcedScalarOrientation || a.Firewall.SelectedTrialityByHand || a.Firewall.ForcedKernelDim3 || a.Firewall.PollutedFiniteCore {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}
