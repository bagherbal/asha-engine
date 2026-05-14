package contactpropagator

import "testing"

func TestContactPropagatorClassifierKeepsContactModesOpen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.ContactRows != 7 || a.PositiveFiniteContactRows != 7 || !a.PositiveFiniteOverlapSpectrumDerived {
		t.Fatalf("expected seven positive finite contact overlap rows, got rows=%d positive=%d", a.ContactRows, a.PositiveFiniteContactRows)
	}
	if a.LorentzKineticSignDerived || a.LocalityDerived || a.PoleDenominatorDerived || a.ResidueSignDerived {
		t.Fatalf("propagator data leaked: kinetic=%t locality=%t pole=%t residue=%t", a.LorentzKineticSignDerived, a.LocalityDerived, a.PoleDenominatorDerived, a.ResidueSignDerived)
	}
	if a.ContactPropagatorClassDerived || a.PhysicalPositiveNormContactPropagatorDerived || a.RegulatorGhostContactClassDerived || a.ConstrainedContactClassDerived || a.VacuumFrustrationContactClassDerived {
		t.Fatalf("contact propagator class should remain open")
	}
	if a.BetaCorrectionRowsAllowed != 0 || a.ThresholdCorrectedBetaDerived || a.FullFiniteBetaMatchingTensorDerived {
		t.Fatalf("threshold beta correction leaked through")
	}
}

func TestDenominatorAmbiguityWitnessesBlockPhysicalPoleSelection(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.DenominatorAmbiguityWitnessed || len(a.DenominatorWitnesses) < 4 {
		t.Fatalf("expected denominator ambiguity witnesses")
	}
	for _, w := range a.DenominatorWitnesses {
		if w.Canonical || w.SelectsPole || w.SelectsPhysicalClass {
			t.Fatalf("denominator witness incorrectly selected physics: %+v", w)
		}
	}
}

func TestResidualNullityAndPhysicalPredictionsRemainSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.ResidualNullityAfter != 3 || a.ResidualNullityAfter != a.ResidualNullityBefore || a.ResidualSymmetryBroken {
		t.Fatalf("unexpected residual nullity: before=%d after=%d broken=%t", a.ResidualNullityBefore, a.ResidualNullityAfter, a.ResidualSymmetryBroken)
	}
	if a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.HiddenObservedInputUsed {
		t.Fatalf("physical prediction or hidden input leaked through")
	}
}
