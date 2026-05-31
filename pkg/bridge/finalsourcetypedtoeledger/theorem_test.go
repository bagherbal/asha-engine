package finalsourcetypedtoeledger

import "testing"

func TestLedgerHasFiveRemainingWounds(t *testing.T) {
	if got := len(RemainingWounds()); got != 5 {
		t.Fatalf("expected five remaining wounds, got %d", got)
	}
}

func TestFormulaSanity(t *testing.T) {
	if ScaleBridgeVOverMP() <= 0 {
		t.Fatal("scale bridge must be positive")
	}
	if HiggsQuartic() <= 0 {
		t.Fatal("Higgs quartic must be positive")
	}
	if TopAction() >= BottomAction() {
		t.Fatal("top action should be smaller than bottom action in the scalar-aligned lane")
	}
	if NeutrinoM2OverM3() <= 0 || NeutrinoM2OverM3() >= 1 {
		t.Fatal("rank-2 normal neutrino ratio should be between 0 and 1")
	}
}
