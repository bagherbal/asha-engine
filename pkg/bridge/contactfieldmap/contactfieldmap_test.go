package contactfieldmap

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.ContactRows != 7 || a.PositiveFiniteContactRows != 7 {
		t.Fatalf("contact rows=%d positive=%d, want 7/7", a.ContactRows, a.PositiveFiniteContactRows)
	}
	if !a.LocalFieldMapCandidateConstructed || a.LocalCoordinateDerived || a.SpacetimeSupportDerived || a.InvertibleFieldMapDerived {
		t.Fatalf("local-field map flags inconsistent")
	}
	if !a.ConstraintClassifierConstructed || a.ConstraintGeneratorDerived || a.GhostGradingDerived || a.NilpotentBRSTDerived || a.SupertraceCancellationDerived {
		t.Fatalf("BRST classifier flags inconsistent")
	}
	if a.PhysicalLocalContactFieldsDerived || a.ConstrainedContactClassDerived || a.RegulatorGhostContactClassDerived || a.VacuumFrustrationContactClassDerived || a.ContactFieldClassDerived {
		t.Fatalf("contact class must remain open")
	}
	if a.BetaCorrectionRowsAllowed != 0 || a.ThresholdCorrectedBetaDerived || a.FullFiniteBetaMatchingTensorDerived {
		t.Fatalf("threshold beta tensor must remain sealed")
	}
	if a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 || a.ResidualSymmetryBroken {
		t.Fatalf("residual nullity changed: before=%d after=%d broken=%t", a.ResidualNullityBefore, a.ResidualNullityAfter, a.ResidualSymmetryBroken)
	}
	if a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.HiddenObservedInputUsed {
		t.Fatalf("physical predictions/observed input must stay sealed")
	}
}

func TestTheoremPasses(t *testing.T) {
	r := ContactOverlapLocalFieldMapConstraintBRSTClassifierTheorem().Run()
	if string(r.Status) != "VARIATIONAL" {
		t.Fatalf("status=%s", r.Status)
	}
	for _, c := range r.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
