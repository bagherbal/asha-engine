package axionphenomenologyseal

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.Summary.SealGranted {
		t.Fatalf("expected AxionPhenomenologySeal granted")
	}
	if !a.Summary.MisalignmentComputed || a.Misalign.RequiredFAGeV <= 0 {
		t.Fatalf("expected positive sealed f_a calculation")
	}
	if a.Summary.ScaleResonanceFound || a.Resonance.ResonanceFound {
		t.Fatalf("unexpected scale resonance: %+v", a.Resonance)
	}
	if !a.DM.ALPAccountsForDMUnderSeal || a.DM.DarkMatterDerivedFromFinite {
		t.Fatalf("expected sealed-only dark matter parameterization, got %+v", a.DM)
	}
	if a.Firewall.FiniteCorePolluted || a.Firewall.BGapPromotedWithoutSeal || a.Firewall.ObservedDMUsedToRewriteCore {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
}

func TestMisalignmentScale(t *testing.T) {
	m := computeMisalignment(1)
	if m.RequiredFAGeV != referenceFA {
		t.Fatalf("theta=1 target should give reference f_a, got %.12g", m.RequiredFAGeV)
	}
	variant := auditBGapThetaVariant(bSectorFirstGap)
	if !variant.Evaluated || variant.Promoted || variant.RequiredFAGeV <= referenceFA {
		t.Fatalf("unexpected B-gap theta diagnostic: %+v", variant)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := AxionPhenomenologySealBGapMisalignmentScaleAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
