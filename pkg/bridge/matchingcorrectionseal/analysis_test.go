package matchingcorrectionseal

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if a.Summary.Status != StatusConditionalPhenomenology {
		t.Fatalf("status=%s", a.Summary.Status)
	}
	if !a.Seal.Active || a.Seal.ResidualPromotedAsDerived {
		t.Fatalf("bad matching seal: %+v", a.Seal)
	}
	if !a.Fit.MatchingPlausible || a.Fit.ResidualOverEpsilon >= 1 {
		t.Fatalf("fit not plausible: %s", FormatFit(a.Fit))
	}
	if a.Firewall.MatchingCorrectionsDerived || a.Firewall.PhysicalPredictionClaimed {
		t.Fatalf("firewall broken: %s", FormatFirewall(a.Firewall))
	}
}

func TestTheorem(t *testing.T) {
	r := MatchingCorrectionSealFullSMYukawaTwoLoopIntegrationAuditTheorem().Run()
	if !r.Passed() {
		t.Fatalf("theorem failed:\n%s", r.Details())
	}
}
