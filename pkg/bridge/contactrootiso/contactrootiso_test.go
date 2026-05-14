package contactrootiso

import "testing"

func TestGate150ExactRootIsolationButNoSemanticProjectors(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.ExactRootIsolationCertified || a.RootIsolationCertificates != 7 {
		t.Fatalf("expected seven exact root-isolation certificates: cert=%+v", a.Certificate)
	}
	if !a.Quartic.AllQuarticRootsCovered || a.Quartic.SignChanges != 4 || !a.Quartic.IntervalsDisjoint {
		t.Fatalf("quartic roots were not isolated: %+v", a.Quartic)
	}
	if a.Certificate.RationalRootsCertified != 3 || a.Certificate.QuarticRootsCertified != 4 || a.Certificate.UnitRootMultiplicity != 7 {
		t.Fatalf("unexpected root accounting: %+v", a.Certificate)
	}
	if a.ExactNumberFieldProjectors != 0 || a.RowwiseRootAssignmentProofs != 0 || a.ChargeSemanticRows != 0 || a.RepresentationCompleteRows != 0 || a.ContactBetaRowsAllowed != 0 {
		t.Fatalf("Gate 150 must not open projector semantics or beta rows: summary=%+v", a.Summary)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("hidden observed physics leaked into Gate 150")
	}
}
