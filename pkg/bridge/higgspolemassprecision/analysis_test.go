package higgspolemassprecision

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || a.Inputs.AddsFit {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestTreeProxy(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !nearlyEqual(a.Tree.LambdaH, 1197.0/(4624.0*2.0), 1e-15) {
		t.Fatalf("bad lambda: %s", FormatTree(a.Tree))
	}
	if a.Tree.MassGeV < 125.0 || a.Tree.MassGeV > 125.6 {
		t.Fatalf("bad mass proxy: %s", FormatTree(a.Tree))
	}
}

func TestPrecisionGap(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Capacity.GapIsSubGeV || a.Comparison.RelativeMassErrorPct <= 0 || a.Comparison.RelativeMassErrorPct >= 0.2 {
		t.Fatalf("bad gap: %s %s", FormatComparison(a.Comparison), FormatCapacity(a.Capacity))
	}
}

func TestPoleLedgerFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Pole.Executed || !a.Pole.RequiresTopSelfEnergy || !a.Pole.RequiresWeakBosonSelfEnergy || !a.Pole.RequiresRenormalizationScheme {
		t.Fatalf("bad pole ledger: %s", FormatPole(a.Pole))
	}
	if !a.Audit.NoExactColliderClaim || !a.Audit.NoSelfEnergiesComputed || !a.Audit.NoTwoLoopClaim {
		t.Fatalf("bad audit: %s", FormatAudit(a.Audit))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	statuses := Statuses(a)
	required := []string{StatusTreeProxyRecomputed, StatusPoleConversionLedgerFormalized, StatusFailedSelfEnergiesNotComputed, StatusFailedExactColliderMassNotClaimed}
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
	res := HiggsPoleMassConversionPrecisionGapLedgerAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
