package higgspolemasseselfenergy

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || !a.Inputs.UsesObservedForTarget {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestPoleEquationTarget(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Equation.RequiredRePiGeV2 <= 0 || a.Equation.DeltaPoleMinusRunGeV2 >= 0 {
		t.Fatalf("bad pole equation: %s", FormatEquation(a.Equation))
	}
	if !nearlyEqual(a.Target.RequiredRePiGeV2, 43.604449567481424, 1e-9) {
		t.Fatalf("bad required RePi: %s", FormatTarget(a.Target))
	}
}

func TestLoopScaleCapacity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Capacity.RequiredIsOrderOneLoop || !a.Capacity.RequiredIsSmallFraction {
		t.Fatalf("bad loop capacity: %s", FormatCapacity(a.Capacity))
	}
	if a.Target.RequiredRePiOverLoopUnit < 0.8 || a.Target.RequiredRePiOverLoopUnit > 1.0 {
		t.Fatalf("unexpected loop units: %s", FormatTarget(a.Target))
	}
}

func TestPrecisionLedgerFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Ledger.FullCalculationExecuted || !a.Ledger.NeedsTopLoop || !a.Ledger.NeedsWZLoops || !a.Ledger.NeedsSchemeChoice {
		t.Fatalf("bad precision ledger: %s", FormatLedger(a.Ledger))
	}
	if !a.Audit.NoExactPoleMassClaim || !a.Audit.NoLoopIntegralsEvaluated || !a.Audit.NoFitParameterIntroduced {
		t.Fatalf("bad audit: %s", FormatAudit(a.Audit))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	statuses := Statuses(a)
	required := []string{StatusRequiredSelfEnergyComputed, StatusLoopScaleCapacityAudited, StatusFailedFullOneLoopNotComputed, StatusFailedExactPoleMassNotClaimed}
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
}

func TestTheoremPasses(t *testing.T) {
	res := HiggsPoleSelfEnergyTargetMinimalPrecisionCorrectionAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
