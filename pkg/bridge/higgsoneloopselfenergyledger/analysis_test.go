package higgsoneloopselfenergyledger

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || a.Inputs.RequiredRePiGeV2 <= 0 {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestComponentLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Ledger.Components) != 4 {
		t.Fatalf("bad component count: %s", FormatLedger(a.Ledger))
	}
	if a.Ledger.Components[0].ContributionGeV2 >= 0 {
		t.Fatalf("top loop should be negative: %s", FormatLedger(a.Ledger))
	}
	if a.Ledger.Components[1].ContributionGeV2 <= 0 || a.Ledger.Components[2].ContributionGeV2 <= 0 || a.Ledger.Components[3].ContributionGeV2 <= 0 {
		t.Fatalf("bosonic/scalar loops should be positive in this kernel: %s", FormatLedger(a.Ledger))
	}
}

func TestKernelAndCounterterm(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !nearlyEqual(a.Kernel.RawKernelGeV2, -991.5670298916102, 1e-6) {
		t.Fatalf("unexpected raw kernel: %s", FormatKernel(a.Kernel))
	}
	if a.Kernel.MatchesTarget {
		t.Fatalf("raw kernel must not match target: %s", FormatKernel(a.Kernel))
	}
	if !a.Counterterms.CountertermMandatory || a.Counterterms.RequiredFiniteCountertermGeV2 < 1000 {
		t.Fatalf("bad counterterm ledger: %s", FormatCounterterms(a.Counterterms))
	}
}

func TestSchemeAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Scheme.NeedsPVIntegrals || !a.Scheme.NeedsFiniteCounterterms || a.Scheme.ExactPoleMassComputed {
		t.Fatalf("bad scheme ledger: %s", FormatScheme(a.Scheme))
	}
	if !a.Audit.NoExactPVFunctions || !a.Audit.NoCountertermDerivation || !a.Audit.NoExactColliderClaim {
		t.Fatalf("bad firewall audit: %s", FormatAudit(a.Audit))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	statuses := Statuses(a)
	required := []string{StatusComponentLedgerFormalized, StatusVeltmanKernelAudited, StatusTensionRawKernelNotTarget, StatusFailedPassarinoVeltmanNotComputed, StatusFailedExactPoleMassNotClaimed}
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
	res := HiggsOneLoopSelfEnergyComponentLedgerRenormalizedPoleKernelAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
