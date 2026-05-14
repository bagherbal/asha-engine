package contactbundle

import "testing"

func TestContactLocalBundleObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.ContactRows != 7 || a.PositiveFiniteContactRows != 7 || a.SurvivingCohomologyRows != 7 {
		t.Fatalf("expected seven positive surviving contact rows, got contact=%d positive=%d surviving=%d", a.ContactRows, a.PositiveFiniteContactRows, a.SurvivingCohomologyRows)
	}
	if !a.LocalBundleConstructionAttempted || !a.FiniteCarrierAvailable {
		t.Fatalf("expected local bundle construction attempt over finite carrier")
	}
	if a.BaseSpaceMapDerived || a.FiberDerived || a.TransitionFunctionsDerived || a.SectionMapDerived {
		t.Fatalf("unexpected local bundle data derived")
	}
	if a.GaugeRepresentationForContactDerived || a.HyperchargeForContactDerived || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 {
		t.Fatalf("unexpected contact representation completion")
	}
	if a.LorentzKineticForContactDerived || a.MassActivationForContactDerived || a.DecouplingRuleForContactDerived {
		t.Fatalf("unexpected kinetic/mass/decoupling data derived")
	}
	if a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("physical bridge leaked")
	}
}
