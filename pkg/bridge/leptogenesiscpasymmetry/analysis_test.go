package leptogenesiscpasymmetry

import (
	"math"
	"testing"
)

func TestGate354LeptogenesisCPAsymmetryAudit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	if a.Span.InheritedGate != 353 || a.Span.AddsFit {
		t.Fatalf("bad span: %+v", a.Span)
	}
	if !a.Decay.Formalized || a.Decay.BGap <= 0 {
		t.Fatalf("decay not formalized: %+v", a.Decay)
	}
	if !a.Sakharov.Formalized || a.Sakharov.ASHAHasCPPhaseOperator {
		t.Fatalf("bad sakharov ledger: %+v", a.Sakharov)
	}
	if math.Abs(a.Target.RequiredEpsKappa-2.618277698863636e-8) > 1e-18 {
		t.Fatalf("unexpected epsilon*kappa target: %.18e", a.Target.RequiredEpsKappa)
	}
	if math.Abs(a.Capacity.PortalOverlap-0.39138716882561015) > 1e-12 {
		t.Fatalf("unexpected portal overlap: %.15f", a.Capacity.PortalOverlap)
	}
	if math.Abs(a.Capacity.InstantonOverlapEps-1.5704319682290527e-6) > 1e-16 {
		t.Fatalf("unexpected epsilon witness: %.18e", a.Capacity.InstantonOverlapEps)
	}
	if !a.Capacity.ViableEfficiencyWindow {
		t.Fatalf("expected viable efficiency window: %+v", a.Capacity)
	}
	if a.Leptogenesis.DerivedCPInvariant || a.Leptogenesis.DerivedEfficiency {
		t.Fatalf("unexpected derived leptogenesis inputs: %+v", a.Leptogenesis)
	}
	if a.CKMShadow.ParameterReduction != 0 || a.Census.RemainingInputs != 15 {
		t.Fatalf("parameter count changed unexpectedly: %+v %+v", a.CKMShadow, a.Census)
	}
	if a.Summary.BaryonAsymmetryPredicted || a.Summary.AnyReductionProved {
		t.Fatalf("summary overclaimed: %+v", a.Summary)
	}
}

func TestStatusLedgerIncludesFirewalls(t *testing.T) {
	statuses := StatusLedger()
	want := []string{StatusTensionCapacityNearTarget, StatusFailedCPAsymmetryOperatorNotDerived, StatusFailedCKMPhaseNotDerived, StatusFailedNoParameterReduction}
	for _, w := range want {
		found := false
		for _, s := range statuses {
			if s == w {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing status %s", w)
		}
	}
}
