package unificationtraceledger

import "testing"

func TestGate307Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Input.TraceEquivalenceProved || !a.Input.ProjectedScalarCarrierPromoted || a.Input.ShapeNumerator != rawTraceRatioNumerator || a.Input.ShapeDenominator != rawTraceRatioDenominator {
		t.Fatalf("bad Gate 307 inheritance: %s", FormatGate307Inheritance(a.Input))
	}
	if !a.Input.RequiresTraceIndex || !a.Input.RequiresQuarticSign || a.Input.AbsoluteLambdaHDerived || a.Input.AbsoluteGaugeCouplingDerived || a.Input.LowEnergyMassClaimed {
		t.Fatalf("Gate 308 inherited polluted prediction state: %s", FormatGate307Inheritance(a.Input))
	}
}

func TestUnificationTraceIndex(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Trace.UniversalIndexFormalized || !a.Trace.AssumesGaugeUnification || a.Trace.CanonicalTraceIndexValue != "1" || a.Trace.ComputesAbsoluteCoupling || a.Trace.UsesObservedCouplings {
		t.Fatalf("trace ledger failed: %s", FormatTraceIndex(a.Trace))
	}
	if len(a.Trace.GaugeFactors) != 3 {
		t.Fatalf("expected three gauge factor rows, got %d: %s", len(a.Trace.GaugeFactors), FormatTraceIndex(a.Trace))
	}
	for _, g := range a.Trace.GaugeFactors {
		if !g.IncludedInGUTLedger || !g.NormalizedToUniversal || g.CanonicalTraceIndex != "1" {
			t.Fatalf("gauge factor not normalized to universal ledger: %s", FormatGaugeFactor(g))
		}
	}
}

func TestSignConventionLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Sign.LedgerFormalized || !a.Sign.WickConventionDeclared || !a.Sign.PositivePotentialConvention || a.Sign.SignValue != sign4Canonical || a.Sign.DerivedFromFiniteCore || !a.Sign.BlocksIfNegative {
		t.Fatalf("bad sign ledger: %s", FormatSign(a.Sign))
	}
}

func TestQuarticBoundaryEquation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Boundary.AnalyticBoundaryDerived || a.Boundary.UnifiedBoundaryEquation != "λ_H(Λ_GUT) = (1197/4624) · g_*^2" || a.Boundary.ExactCoefficient != "1197/4624" {
		t.Fatalf("boundary equation not derived: %s", FormatBoundary(a.Boundary))
	}
	if !a.Boundary.DependsOnUnifiedGaugeValue || a.Boundary.DependsOnF2Moment || a.Boundary.DependsOnN4F0Ledger || a.Boundary.DependsOnCutoffProfileShape || a.Boundary.LowEnergyPredictionMade {
		t.Fatalf("boundary equation leaked forbidden dependencies/prediction: %s", FormatBoundary(a.Boundary))
	}
}

func TestFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoAbsoluteGaugeValueInserted || !a.Firewalls.NoBoundaryScaleInserted || !a.Firewalls.NoObservedCouplingsInserted || !a.Firewalls.NoRGERunningExecuted || !a.Firewalls.NoThresholdMatchingInserted || !a.Firewalls.NoLowEnergyHiggsMassClaimed || !a.Firewalls.NoYukawaNumbersInserted || !a.Firewalls.F2MassFirewallPreserved || !a.Firewalls.BGapInstantonFirewallPreserved || !a.Firewalls.AnalyticBoundaryOnly || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPassesWithAnalyticBoundaryOnly(t *testing.T) {
	res := UnificationTraceLedgerHiggsQuarticUnificationBoundaryAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
