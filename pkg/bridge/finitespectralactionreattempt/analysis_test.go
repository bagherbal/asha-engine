package finitespectralactionreattempt

import "testing"

func TestBuildDefaultComputesRawMomentsButNotHiggsRatio(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Summary.RawMomentsEvaluated || !a.Trace.MomentsComputed {
		t.Fatalf("expected raw spectral moments to be evaluated: %+v", a.Trace)
	}
	if a.Summary.SeeleyDeWittDerived || a.Trace.SeeleyDeWittMapDerived {
		t.Fatalf("raw moments must not be promoted to Seeley-de Witt coefficients: %+v", a.Trace)
	}
	if a.Summary.HiggsRatioDerived || a.Higgs.HiggsRatioDerived || a.Higgs.HiggsMassPredicted {
		t.Fatalf("Gate 268 must not derive Higgs mass ratio: %+v", a.Higgs)
	}
}

func TestMomentRatioDependsOnUnselectedDF(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Trace.Rows) < 2 {
		t.Fatalf("expected unit and deformed diagnostic rows")
	}
	unit := a.Trace.Rows[0]
	deformed := a.Trace.Rows[1]
	if unit.TraceD2 != 16 || unit.TraceD4 != 16 || unit.RawA2OverA4 != 1 {
		t.Fatalf("unexpected unit diagnostic moments: %+v", unit)
	}
	if deformed.RawA2OverA4 == unit.RawA2OverA4 {
		t.Fatalf("deformed diagnostic should expose moment-ratio dependence: unit=%+v deformed=%+v", unit, deformed)
	}
	if !a.Trace.DependsOnDFSingularValues || a.Trace.RawMomentRatioInvariant {
		t.Fatalf("expected dependence on unselected D_F singular values: %+v", a.Trace)
	}
}

func TestFirewallsAndFutureObligations(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewall.EmpiricalYukawaSealPreserved || !a.Firewall.RawMomentsNotPromoted || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}
	if a.Future.CanDeriveHiggsRatioNow {
		t.Fatalf("future map must not allow Higgs ratio now: %+v", a.Future)
	}
	unsatisfied := 0
	for _, o := range a.Future.Obligations {
		if o.Required && !o.Satisfied {
			unsatisfied++
		}
	}
	if unsatisfied < 8 {
		t.Fatalf("expected all spectral obligations unsatisfied, got %d", unsatisfied)
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	th := FiniteSpectralActionReAttemptSeeleyDeWittCoefficientAuditTheorem()
	res := th.Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
