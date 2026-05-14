package moritabimodulesearch

import "testing"

func TestBuildDefaultExtractsMoritaBimodule(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Summary.BimoduleExtracted || a.Bimodule.TotalComplexDimension != 16 || len(a.Bimodule.Summands) != 4 {
		t.Fatalf("expected four-summand finite bimodule: %s", FormatBimodule(a.Bimodule))
	}
	if !a.Bimodule.LeftActionFaithful || !a.Bimodule.RightOppositeActionFaithful || !a.Bimodule.LeftRightCommute {
		t.Fatalf("expected faithful commuting A-Aop representation: %s", FormatBimodule(a.Bimodule))
	}
	if a.Bimodule.FullFockCarrierUsed {
		t.Fatalf("Gate 272 must not reuse full second-quantized S_C as spectral triple carrier")
	}
}

func TestOrderOneEdgeSieveAllowsNonVacuousButDoesNotSelectDF(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.OrderOne.NonVacuousOneFormsAvailable || !a.OrderOne.OrderOneSatisfiedForAllowedEdges {
		t.Fatalf("expected non-vacuous order-one edges: %s", FormatOrderOne(a.OrderOne))
	}
	if a.OrderOne.NonVacuousAllowedEdges != 2 || a.OrderOne.RejectedEdges != 2 {
		t.Fatalf("unexpected edge classification: %s", FormatOrderOne(a.OrderOne))
	}
	if a.OrderOne.CanonicalDFSelected || a.OrderOne.XYRatioLocked {
		t.Fatalf("order-one must not lock D_F or x:y: %s", FormatOrderOne(a.OrderOne))
	}
}

func TestRatioStillBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Ratio.DependsOnXOverY || a.Ratio.A2A4Derived || a.Ratio.HiggsRatioDerived {
		t.Fatalf("expected a2/a4 to remain blocked: %s", FormatRatio(a.Ratio))
	}
	if a.Summary.HiggsRatioDerived || a.Summary.XYRatioLocked || a.Summary.CanonicalDFDerived {
		t.Fatalf("summary over-promoted result: %s", FormatSummary(a.Summary))
	}
}

func TestFirewallAndFutureCriteria(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewall.NoConnesSMAlgebraImported || !a.Firewall.BimoduleNotPromotedToSM || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
	missing := 0
	for _, c := range a.Future.Criteria {
		if c.Required && !c.Satisfied {
			missing++
		}
	}
	if missing < 6 || !a.Future.NeedAmplitudeSelector {
		t.Fatalf("expected all future criteria missing: %s", FormatFuture(a.Future))
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	res := FiniteAlgebraRepresentationObstructionClassificationMoritaBimoduleSearchAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
