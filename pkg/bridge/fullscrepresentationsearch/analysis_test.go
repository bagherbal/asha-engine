package fullscrepresentationsearch

import "testing"

func TestBuildDefaultAuditsFullCarrierButDoesNotPromoteRepresentation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Carrier.CARPassed || a.Carrier.BaseComplexDimension != 16 || a.Carrier.DoubledComplexDimension != 32 {
		t.Fatalf("expected full Fock carrier and CAR preflight: %+v", a.Carrier)
	}
	if a.Representation.ValidFullAssociativeRepFound {
		t.Fatalf("Gate 271 must not derive a full associative S_C representation: %+v", a.Representation)
	}
	if !a.Representation.FullSCPromotionBlocked {
		t.Fatalf("expected full-S_C promotion to be blocked: %+v", a.Representation)
	}
}

func TestLiftFailuresAreSeparated(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !liftFailureProfileOK(a.Representation.Candidates) {
		t.Fatalf("expected Γ, dΓ, and one-particle failures to be distinct: %s", FormatRepresentation(a.Representation))
	}
	for _, c := range a.Representation.Candidates {
		if c.Name == "Γ exterior functor lift" && c.DiagnosticDefect <= 0 {
			t.Fatalf("Γ lift should expose a positive additivity defect: %+v", c)
		}
		if c.Name == "dΓ creation-annihilation bilinear lift" && c.DiagnosticDefect <= 0 {
			t.Fatalf("dΓ lift should expose positive unital/multiplicative defects: %+v", c)
		}
	}
}

func TestOppositeOrderOneAndHiggsRemainBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Opposite.OppositeActionConstructed || a.Opposite.CandidateJPhysicalSemantics {
		t.Fatalf("physical J/opposite action must remain blocked: %+v", a.Opposite)
	}
	if a.OrderOne.ReevaluatedAsSpectralTriple || a.OrderOne.OrderOneSatisfied || a.OrderOne.NonVacuousOneFormsDerived {
		t.Fatalf("full order-one theorem must not be promoted: %+v", a.OrderOne)
	}
	if a.Ratio.HiggsRatioDerived || a.Ratio.XYRatioSelected {
		t.Fatalf("Higgs ratio and x:y selector must remain blocked: %+v", a.Ratio)
	}
}

func TestFirewallsAndFutureMap(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewall.EmpiricalYukawaSealPreserved || !a.Firewall.NoCandidatePromoted || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}
	missing := 0
	for _, o := range a.Future.Obligations {
		if o.Required && !o.Satisfied {
			missing++
		}
	}
	if missing < 6 {
		t.Fatalf("expected all representation obligations to remain unsatisfied, got %d: %+v", missing, a.Future)
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	th := FullSCFiniteAlgebraRepresentationSearchOppositeActionConstructionAuditTheorem()
	res := th.Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
