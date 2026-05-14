package contactmatrixcert

import "testing"

func TestGate149ExactMatrixAndCharpolyCertificate(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.ExactRationalOverlapMatrix || !a.ExactDeterminantComputed || !a.ExactCharacteristicCertified || !a.ExactAnnihilationCertified {
		t.Fatalf("expected exact rational matrix/determinant/charpoly/annihilation certificates: %+v", a.Certificate)
	}
	if a.MatrixLift.Trace != "163/15" || a.MatrixLift.Determinant != "271/29160" {
		t.Fatalf("unexpected exact trace/determinant: trace=%s det=%s", a.MatrixLift.Trace, a.MatrixLift.Determinant)
	}
	if a.MatrixLift.UnitEigenspaceDimension != 7 || a.Certificate.PartialDegree != 7 || a.Certificate.Degree != 14 {
		t.Fatalf("unexpected spectral dimensions: lift=%+v cert=%+v", a.MatrixLift, a.Certificate)
	}
	if a.RootIsolationCertificates != 0 || a.RowwiseRootAssignmentProofs != 0 || a.ChargeSemanticRows != 0 || a.RepresentationCompleteRows != 0 || a.ContactBetaRowsAllowed != 0 {
		t.Fatalf("Gate 149 must not open semantics or beta rows: summary=%+v", a.Summary)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("hidden observed physics leaked into Gate 149")
	}
}
