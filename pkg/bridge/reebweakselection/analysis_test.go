package reebweakselection

import "testing"

func TestGate241ReebWeakSelectionAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Previous.Summary.U1RejectsTemporalPlanes || a.Previous.Summary.PureSpatialPlanesRemain != 3 {
		t.Fatalf("Gate 240 inheritance should leave exactly three pure-spatial planes: %+v", a.Previous.Summary)
	}
	if !a.Contact.ContactProjectorExists || a.Contact.ContactDimension != 7 || !a.Contact.FiniteContactActionOnly {
		t.Fatalf("contact K should be available as finite projector geometry: %+v", a.Contact)
	}
	if a.Contact.EtaOneFormDerived || a.Contact.DEtaTwoFormDerived || a.Contact.ReebVectorDerived {
		t.Fatalf("eta/deta/Reeb must not be claimed as derived: %+v", a.Contact)
	}
	if a.Reeb.CandidateAvailable || a.Reeb.MappedToSpatialFockAxes || a.Reeb.ManualAxisChoice {
		t.Fatalf("Reeb vector/axis must not be forced: %+v", a.Reeb)
	}
	if a.Projection.KToWProjectionDerived || a.Projection.UniqueSpatialAxisTagged || a.Projection.S3PermutationBroken {
		t.Fatalf("no K-to-Fock projection or S3 breaking should be derived: %+v", a.Projection)
	}
	if len(a.Planes) != 3 || a.Sieve.CandidatePlaneCount != 3 || len(a.Sieve.SelectedPlanes) != 0 || a.Sieve.UniqueWeakPlaneSelected {
		t.Fatalf("three planes should remain unselected: sieve=%+v planes=%+v", a.Sieve, a.Planes)
	}
	for _, p := range a.Planes {
		if !p.SurvivesU1Twist || !p.RequiresTaggedAxis || p.SelectedByReeb {
			t.Fatalf("pure-spatial plane audit inconsistent: %+v", p)
		}
	}
	if a.Summary.UniqueWeakPlaneDerived || a.Summary.GlobalHDerived || a.Summary.PhysicalChiralityDerived {
		t.Fatalf("summary overclaims weak plane/H/chirality: %+v", a.Summary)
	}
}
