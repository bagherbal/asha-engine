package tracesynthesiscarrierequivalence

import "testing"

func TestGate306Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Input.QuarticChannelExtracted || !a.Input.LambdaOverGaugeRatioFormalized || !a.Input.RelativeRatioCancelsN4F0 || a.Input.Raw1197PromotedDirectly {
		t.Fatalf("bad Gate 306 inheritance: %s", FormatGate306Inheritance(a.Input))
	}
	if a.Input.RawTraceNumerator != rawTraceRatioNumerator || a.Input.RawTraceDenominator != rawTraceRatioDenominator || a.Input.NumericalLambdaHDerived || a.Input.HiggsMassPredictionClaimed {
		t.Fatalf("Gate 307 inherited polluted state: %s", FormatGate306Inheritance(a.Input))
	}
}

func TestPhysicalCarrierTraceParsing(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Carrier.TraceParsed || !a.Carrier.UsesDoubledSpace || !a.Carrier.UsesAllowedDiracEdges || !a.Carrier.UsesMoritaMultiplicities || !a.Carrier.RejectsVacuumTerms || !a.Carrier.RejectsGaugeCurvatureTerms || !a.Carrier.RejectsMixedDerivativeTerms {
		t.Fatalf("carrier parsing failed: %s", FormatCarrier(a.Carrier))
	}
	if len(a.Carrier.Edges) != 2 {
		t.Fatalf("expected two projected Morita edge carriers, got %d: %s", len(a.Carrier.Edges), FormatCarrier(a.Carrier))
	}
	if a.Carrier.Edges[0].MoritaMultiplicity != kappaC || a.Carrier.Edges[1].MoritaMultiplicity != kappaQ {
		t.Fatalf("bad Morita multiplicities: %s", FormatCarrier(a.Carrier))
	}
}

func TestQuarticKineticPolynomialConstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Polynomial.PolynomialConstructed || !a.Polynomial.ScaleCancels || !a.Polynomial.MoritaShapeUsed {
		t.Fatalf("polynomial was not constructed: %s", FormatPolynomial(a.Polynomial))
	}
	if a.Polynomial.PhysicalRatioPolynomial != "R_phys(r) := C4_raw/K_H_raw^2 = (1+3r^2)/(1+3r)^2" {
		t.Fatalf("unexpected ratio polynomial: %s", FormatPolynomial(a.Polynomial))
	}
}

func TestTraceEquivalenceSieve(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Equivalence.EquivalenceProved || !a.Equivalence.PromotesProjectedShapeBound || a.Equivalence.Numerator != rawTraceRatioNumerator || a.Equivalence.Denominator != rawTraceRatioDenominator {
		t.Fatalf("equivalence not proved: %s", FormatEquivalence(a.Equivalence))
	}
	if !a.Equivalence.ScalarProjectorRequired || a.Equivalence.UnprojectedGlobalTraceUsed || !a.Equivalence.VacuumTermsSeparated || !a.Equivalence.GaugeCrossTermsSeparated {
		t.Fatalf("equivalence violated projector/firewall obligations: %s", FormatEquivalence(a.Equivalence))
	}
}

func TestDimensionlessPhysicalRatioMap(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.RatioMap.MapFormalized || !a.RatioMap.UsesEquivalenceSeal || a.RatioMap.ProducesNumericalLambdaH || a.RatioMap.ProducesAbsoluteGaugeCoupling {
		t.Fatalf("bad ratio map: %s", FormatRatioMap(a.RatioMap))
	}
	if !a.RatioMap.RequiresTraceIndex || !a.RatioMap.RequiresQuarticSign || !a.RatioMap.RequiresYukawaOrigin {
		t.Fatalf("ratio map lost obligations: %s", FormatRatioMap(a.RatioMap))
	}
}

func TestFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoUnprojectedTracePromotion || !a.Firewalls.NoVacuumContamination || !a.Firewalls.NoGaugeCrossContamination || !a.Firewalls.NoYukawaNumbersInserted || !a.Firewalls.NoNumericalLambdaHComputed || !a.Firewalls.NoAbsoluteGaugeClaimed || !a.Firewalls.NoHiggsMassClaimed || !a.Firewalls.NoBGapInstantonClaimed || !a.Firewalls.ProjectedEquivalenceOnly || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	res := RawTraceSynthesisCarrierEquivalenceQuarticKineticRatioAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
