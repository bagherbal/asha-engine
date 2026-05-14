package eftdecayportal

import "testing"

func TestBuildDefaultGate222(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Gate221.FatalCosmologicalPathology {
		t.Fatalf("expected Gate221 pathology inheritance")
	}
	if !a.Operators.TripletPortalFound {
		t.Fatalf("expected sealed triplet portal support")
	}
	if !a.Operators.OctetMassMixingRejected {
		t.Fatalf("expected false octet-Q mass mixing to be rejected")
	}
	if a.Operators.OctetPureSMPortalFound {
		t.Fatalf("did not expect pure-SM octet portal through audited basis")
	}
	if a.RelicSeal.SealGranted {
		t.Fatalf("full RelicDecaySeal should not be granted")
	}
	if !a.RelicSeal.PartialTripletSubseal {
		t.Fatalf("expected partial triplet sub-seal support")
	}
	if a.Kinematics.TripletYukawaMin <= 0 || a.Kinematics.TripletYukawaMin >= 1e-12 {
		t.Fatalf("unexpected triplet Yukawa bound: %.12g", a.Kinematics.TripletYukawaMin)
	}
	if a.Summary.FatalPathologyCleared {
		t.Fatalf("full cosmological pathology should not be cleared")
	}
}

func TestTheoremGate222(t *testing.T) {
	r := EFTDecayPortalRelicDecaySealActivationAuditTheorem().Run()
	if !r.Passed() {
		t.Fatalf("theorem did not pass checks:\n%s", r.Details())
	}
}
