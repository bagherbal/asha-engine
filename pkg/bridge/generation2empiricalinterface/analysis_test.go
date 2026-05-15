package generation2empiricalinterface

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Native.NativeOnlyPredictsGST {
		t.Fatalf("native ledger must not restore GST: %s", FormatNativeLedger(a.Native))
	}
	if a.Contract.AllowsNativeClaim || !a.Contract.RequiresExplicitLabel || !a.Contract.RequiresSectorTag || !a.Contract.RequiresRenormalizationTag {
		t.Fatalf("import contract is not strict enough: %s", FormatImportContract(a.Contract))
	}
	if a.Sieve.AnyForbiddenAccepted || !a.Sieve.PromotionRejected {
		t.Fatalf("interface sieve accepted a forbidden promotion: %s", FormatInterfaceSieve(a.Sieve))
	}
	if !a.Firewall.GSTFritzschRelationsQuarantined || !a.Firewall.NoGSTPromotion {
		t.Fatalf("GST branch must remain quarantined: %s", FormatFirewall(a.Firewall))
	}
}

func TestForbiddenInputsRejected(t *testing.T) {
	c := buildImportContract()
	for _, input := range c.Inputs {
		if input.NativePromotion && input.Allowed {
			t.Fatalf("native promotion input was allowed: %s", FormatEmpiricalInput(input))
		}
	}
	if c.RejectedPromotionCount < 2 {
		t.Fatalf("expected at least two rejected native promotions, got %d", c.RejectedPromotionCount)
	}
}

func TestResidualLedgerQuarantinesNativeGST(t *testing.T) {
	r := buildResidualLedger()
	if !r.AllowsGSTResidual {
		t.Fatal("labelled GST residual should be allowed as empirical comparator")
	}
	if r.AllowsNativeGSTRatioClaim || r.AllowsCoefficientFittingAsNative {
		t.Fatalf("residual ledger leaked into native claims: %s", FormatResidualLedger(r))
	}
}

func TestInterfaceRequests(t *testing.T) {
	s := buildInterfaceSieve()
	if !s.NativeOnlyAllowed || !s.EmpiricalFitAllowed {
		t.Fatalf("legal interface requests not accepted: %s", FormatInterfaceSieve(s))
	}
	for _, req := range s.Requests {
		if req.AttemptsPromotion && req.Allowed {
			t.Fatalf("promotion request was accepted: %s", FormatInterfaceRequest(req))
		}
	}
}

func TestTheoremPasses(t *testing.T) {
	res := Generation2TextureZeroInvariantLedgerAllowedEmpiricalInterfaceTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit checks did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsKeyStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusEmpiricalInterfaceDefined, StatusFailedGSTRequiresEmpiricalBranchInput, "M_22=0 spectral sum rule", "renormalization scale/scheme"} {
		if !stringsContains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func TestContainsHelper(t *testing.T) {
	if !contains([]string{"a", "b"}, "b") {
		t.Fatal("contains helper failed")
	}
	if contains([]string{"a", "b"}, "c") {
		t.Fatal("contains helper false positive")
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
