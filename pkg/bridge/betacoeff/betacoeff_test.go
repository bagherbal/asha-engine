package betacoeff

import (
	"math"
	"testing"
)

func TestBuildDefaultBetaAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inventory.Generations != 3 {
		t.Fatalf("generations=%d", a.Inventory.Generations)
	}
	if a.Inventory.WeylStatesPerGeneration != 16 {
		t.Fatalf("Weyl states/gen=%d", a.Inventory.WeylStatesPerGeneration)
	}
	if math.Abs(a.B1GUTNormalized-4.1) > 1e-10 {
		t.Fatalf("b1=%.12g", a.B1GUTNormalized)
	}
	if math.Abs(a.B2-(-19.0/6.0)) > 1e-10 {
		t.Fatalf("b2=%.12g", a.B2)
	}
	if math.Abs(a.B3-(-7.0)) > 1e-10 {
		t.Fatalf("b3=%.12g", a.B3)
	}
	if a.ImportedSMBetaTable || a.HiddenObservedCouplingsUsed {
		t.Fatalf("hidden/imported input flags should be false")
	}
	if a.FiniteBetaTheoremDerived {
		t.Fatalf("finite beta theorem must remain open")
	}
}
