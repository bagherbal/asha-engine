package generation2coefficientledger

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Closure.AmplitudeFirewallClosed || a.Closure.CoefficientReductionBelow9 || a.Closure.AnyCoefficientAxiomPromoted {
		t.Fatalf("Gate 447 must close, not reduce, the coefficient firewall: %s", FormatClosure(a.Closure))
	}
	if a.Ledger.NativeCoefficientValues != 0 || a.Ledger.QuarantinedCoefficientDim != KXYCoeffDim {
		t.Fatalf("unexpected coefficient ledger: %s", FormatCoefficientLedger(a.Ledger))
	}
}

func TestCounterLedgersProveNonUniqueness(t *testing.T) {
	s := buildCounterLedgerSieve()
	if s.UniqueCoefficientLedger || s.SurvivingLedgers < 3 || s.DistinctSurvivors < 3 {
		t.Fatalf("counter-ledger sieve should prove underdetermination: %s", FormatSieve(s))
	}
	for _, l := range s.Ledgers {
		if !l.Hermitian || !l.TraceNeutral || !l.GaugeCompatible || !l.KMSCompatible || !l.MassLiftCompatible || l.ImportsEmpiricalData {
			t.Fatalf("bad counter-ledger witness: %s", FormatCounterLedger(l))
		}
	}
}

func TestBoundariesDoNotSelectCoefficientValues(t *testing.T) {
	bs := buildBoundaries()
	if !boundariesPassWithoutSelection(bs) {
		t.Fatalf("boundary stack unexpectedly selected values")
	}
}

func TestTheoremPassesAsFailedRouteAudit(t *testing.T) {
	res := Generation2SectorCoefficientSourceLedgerAmplitudeFirewallClosureTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit checks did not pass:\n%s", res.Details())
	}
	if string(res.Status) == "" {
		t.Fatalf("empty theorem status")
	}
}

func TestRenderAuditContainsKeyStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusCoefficientFirewallClosed, StatusFailedMultipleLedgersSurvive, StatusNineCoefficientsRemainQuarantined, "dim C_KXY^charged = 3 sectors × 3 coefficients = 9"} {
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
