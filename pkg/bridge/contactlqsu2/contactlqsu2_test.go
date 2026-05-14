package contactlqsu2

import "testing"

func TestGate133SU2ActionSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.ContactRows != 7 || a.LeptoquarkRows != 6 || a.ColorPlanes != 3 || a.RealOrientationsPerColor != 2 || a.TotalCurrentLQSlots != 6 {
		t.Fatalf("unexpected dimensions: %+v", a.Summary)
	}
	if !a.OrientationSO2Available || !a.ColorWiseSO2Available || !a.DiagonalSO2Available || !a.OrientationActionAbelian {
		t.Fatalf("expected only abelian orientation SO(2) actions: %+v", a.Summary)
	}
	if a.NonAbelianSU2TripleDerived || a.SU2CommutationDerived || a.SU2WeakDoubletActionDerived || a.WeakDoubletSemanticsDerived {
		t.Fatalf("did not expect derived weak SU(2)L action: %+v", a.Summary)
	}
	if !a.BorrowedMatterActionRejected || !a.SemanticBridgeMissing || a.CurrentNaturalSU2Action {
		t.Fatalf("matter borrowing / semantic bridge audit wrong")
	}
	if a.RepresentationCompleteRows != 0 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("beta firewall should remain closed")
	}
	if a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("physical constants leaked or nullity changed")
	}
}

func TestOrientationPlanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if len(a.Planes) != 3 {
		t.Fatalf("expected three planes, got %d", len(a.Planes))
	}
	for _, p := range a.Planes {
		if p.RealDimension != 2 || !p.SO2GeneratorAvailable || !p.AbelianClosure {
			t.Fatalf("bad orientation plane: %+v", p)
		}
		if p.SU2TripleDerived || p.WeakDoubletDerived || p.RepresentationRows || p.BetaPermitted {
			t.Fatalf("plane should not be representation-complete: %+v", p)
		}
	}
}
