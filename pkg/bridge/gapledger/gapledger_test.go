package gapledger

import "testing"

func TestNJLGapKernelCriticalityLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error: %v", err)
	}
	if a.GenerationCount != 3 || a.KindCount != 4 {
		t.Fatalf("unexpected domain generations=%d kinds=%d", a.GenerationCount, a.KindCount)
	}
	if a.QuarkLeptonAmplification != 3 {
		t.Fatalf("expected quark/lepton amplification 3, got %v", a.QuarkLeptonAmplification)
	}
	if a.StrongestWeightedPressure <= 0 {
		t.Fatalf("expected positive strongest pressure skeleton")
	}
	if a.UpDownDegeneracyResidual != 0 {
		t.Fatalf("expected up/down tie to remain in this ledger")
	}
	if a.FourFermionKernelDerived || a.GapEquationSolved || a.CondensateScaleDerived {
		t.Fatalf("ledger must not claim kernel, gap solution, or condensate scale")
	}
	if a.HiddenObservedCouplingsUsed || a.HiddenObservedMassScalesUsed {
		t.Fatalf("must not use observed couplings or mass scales")
	}
}
