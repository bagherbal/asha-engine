package post444flavoratlasreconciliation

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Final.Reconciled || !a.Final.KGenPromoted || !a.Final.Gen2ZeroPromoted || !a.Final.XSupportPromoted {
		t.Fatalf("expected structural promotions: %s", FormatFinal(a.Final))
	}
	if !a.Final.YPhaseStillQuarantined || !a.Final.CoefficientsStillQuarantined || a.Final.NativeFlavorDim != NativeChargedFlavorDim || a.Final.KXYCoeffDim != KXYChargedCoeffDim {
		t.Fatalf("firewall not preserved: %s", FormatFinal(a.Final))
	}
}

func TestDeltaDoesNotPredictValues(t *testing.T) {
	d := buildDelta()
	if d.FlavorObservableValuesAdded != 0 || d.CoefficientSelectorsAdded != 0 || d.NativeDimAfter != NativeChargedFlavorDim || d.KXYCoeffDimAfter != KXYChargedCoeffDim {
		t.Fatalf("Gate 448 delta must be structural-only: %s", FormatDelta(d))
	}
	if d.PromotedObjects != 3 || d.QuarantinedObjects != 2 {
		t.Fatalf("unexpected reclassification counts: %s", FormatDelta(d))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := Post444FlavorFrontierAtlasReconciliationTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsKeyStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusPost444FlavorAtlasReconciled, StatusKGenPromotedGeometric, StatusXTriangleSupportPromoted, StatusNineCoefficientFirewallPreserved, "dim M_charged^native = 13"} {
		if !stringsContains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func stringsContains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
