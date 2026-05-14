package actionscale

import "testing"

func TestActionNormalizationAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.UnitIndexAvailable {
		t.Fatalf("expected unit contact index")
	}
	if !a.DimensionlessActionDerived || a.TopologicalActionSeal <= 0 {
		t.Fatalf("expected positive dimensionless action seal")
	}
	if a.DimensionfulUnitDerived || a.GravityScaleDerived || a.ScalarScaleFixed {
		t.Fatalf("action normalization must not pretend to derive a physical mass scale")
	}
	if a.HiddenObservedScaleInserted {
		t.Fatalf("hidden observed scale insertion detected")
	}
}
