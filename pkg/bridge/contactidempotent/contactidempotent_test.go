package contactidempotent

import "testing"

func TestGate151BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Decomposition.RationalBlockProjectors != 5 || a.Decomposition.TotalSpectralDimension != 14 || a.Decomposition.IndividualQuarticProjectors != 0 {
		t.Fatalf("unexpected decomposition: %+v", a.Decomposition)
	}
	if a.NumberField.RootChosenCanonically || a.NumberField.ExactEigenprojectorFormula || a.NumberField.IndividualQuarticProjectors != 0 {
		t.Fatalf("Gate 151 must not pretend to split quartic branches: %+v", a.NumberField)
	}
	if a.ContactBetaRowsAllowed != 0 || a.ChargeSemanticRows != 0 || a.HyperchargeRowsDerived != 0 || a.RepresentationCompleteRows != 0 {
		t.Fatalf("Gate 151 must keep contact physics firewall closed: summary=%+v", a.Summary)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("hidden observed physics leaked into Gate 151")
	}
}
