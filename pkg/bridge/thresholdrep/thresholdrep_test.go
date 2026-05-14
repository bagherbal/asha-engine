package thresholdrep

import "testing"

func TestThresholdRepresentationAssignmentAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.CandidateCount == 0 {
		t.Fatal("expected threshold candidates")
	}
	if !a.ScalarDoubletSectorDerived {
		t.Fatal("expected scalar doublet sector-level assignment")
	}
	if a.ThresholdCorrectedBetaDerived {
		t.Fatal("threshold-corrected beta coefficients must not be derived yet")
	}
	if a.ContactOverlapRepresentationsDerived || a.BGapRepresentationDerived {
		t.Fatal("B/contact partial threshold representations must remain open")
	}
}
