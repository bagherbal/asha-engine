package contactincidence

import "testing"

func TestContactIncidenceFiberFunctorSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.ContactRows != 7 || a.PositiveFiniteContactRows != 7 || a.SurvivingCohomologyRows != 7 {
		t.Fatalf("expected seven positive surviving contact rows, got contact=%d positive=%d surviving=%d", a.ContactRows, a.PositiveFiniteContactRows, a.SurvivingCohomologyRows)
	}
	if !a.FanoIncidenceAvailable || a.FanoPointCount != 7 || a.FanoLineCount != 7 || !a.EveryFanoPointDegreeThree || !a.EveryFanoLineSizeThree {
		t.Fatalf("expected exact 7-point/7-line Fano incidence with degree three")
	}
	if !a.FanoContactCardinalityMatch || !a.FanoIncidenceResonance {
		t.Fatalf("expected Fano/contact cardinality resonance")
	}
	if a.CanonicalContactToFanoMap || a.FiberFunctorDerived || a.ChartAtlasDerived || a.TransitionCocycleDerived || a.SectionMapDerived {
		t.Fatalf("unexpected canonical contact-Fano map/fiber functor/chart atlas derived")
	}
	if a.GaugeRepresentationDerived || a.LorentzKineticDerived || a.MassActivationDerived || a.DecouplingRuleDerived {
		t.Fatalf("unexpected representation/kinetic/mass/decoupling data derived")
	}
	if a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived || a.HiddenObservedInputUsed {
		t.Fatalf("physical bridge leaked")
	}
}
