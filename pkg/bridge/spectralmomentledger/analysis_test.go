package spectralmomentledger

import (
	"math"
	"testing"
)

func close(a, b, tol float64) bool {
	if a > b {
		return a-b <= tol
	}
	return b-a <= tol
}

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || a.Inputs.NGen != 3 {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestMomentLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Ledger.Gauge.Value != 7 || !a.Ledger.Gravity.Derived || a.Ledger.Cosmological.Derived {
		t.Fatalf("bad ledger: %s", FormatLedger(a.Ledger))
	}
	wantGravity := math.Pi / 64 * unreducedPlanckGeV * unreducedPlanckGeV
	if !close(a.Ledger.Gravity.Value, wantGravity, wantGravity*1e-15) {
		t.Fatalf("bad gravity product: %s", FormatLedger(a.Ledger))
	}
}

func TestGaugeGravityRatio(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	want := math.Pi / (64 * f0Contact) / (a.Inputs.HierarchyRho * a.Inputs.HierarchyRho)
	if !close(a.GaugeGravityRatio.Value, want, want*1e-15) {
		t.Fatalf("bad ratio: %s", FormatRatio(a.GaugeGravityRatio))
	}
	if a.GaugeGravityRatio.Value < 1e31 || a.GaugeGravityRatio.Value > 1e32 {
		t.Fatalf("ratio outside hierarchy range: %s", FormatRatio(a.GaugeGravityRatio))
	}
}

func TestCosmologicalTarget(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Target.TargetRatioToMP4 != observedCosmologicalRatio {
		t.Fatalf("bad target: %s", FormatTarget(a.Target))
	}
	if a.Target.RhoFourth < 1e-68 || a.Target.RhoFourth > 1e-66 {
		t.Fatalf("rho^4 should be around 1e-67: %s", FormatTarget(a.Target))
	}
	if a.Target.RequiredHalfActionCount < 7 || a.Target.RequiredHalfActionCount > 8 {
		t.Fatalf("unexpected half-action count: %s", FormatTarget(a.Target))
	}
}

func TestCosmologicalCandidatesPreserveFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Cosmological.Derived {
		t.Fatalf("unexpected cosmological derivation: %s", FormatCosmologicalAudit(a.Cosmological))
	}
	for _, c := range a.Cosmological.Candidates {
		if c.Name != "required target" && c.Promoted {
			t.Fatalf("unexpected promoted candidate: %s", FormatCandidate(c))
		}
	}
}

func TestStatusesAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	statuses := Statuses(a)
	required := []string{StatusMomentLedgerFormalized, StatusDarkEnergyTargetExtracted, StatusFailedCosmologicalConstantNotDerived, StatusFailedF4Lambda4NotLocked, StatusFailedArbitraryExponentExtensionRejected}
	for _, req := range required {
		found := false
		for _, got := range statuses {
			if got == req {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing status %s in %v", req, statuses)
		}
	}
	res := CompleteSpectralMomentLedgerCosmologicalConstantTripleHierarchyAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
