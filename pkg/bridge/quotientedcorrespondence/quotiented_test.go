package quotientedcorrespondence

import "testing"

func TestGaugeQuotientedProtectedBrokenCorrespondence(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.CountCorrespondenceSurvivesQuotient {
		t.Fatalf("expected count/rank correspondence to survive O(3) quotient")
	}
	if !a.FrameComponentComparisonRejected {
		t.Fatalf("component-wise frame matching must be rejected after quotienting")
	}
	if a.QuotientSafeIntertwinerDerived {
		t.Fatalf("quotient-safe protected-to-broken intertwiner should remain open")
	}
	if a.BrokenMetricIsIsotropic {
		t.Fatalf("current broken-image metric should be anisotropic")
	}
}
