package generation2observedpreflight

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate469(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Preflight.AcceptedSchemaCases != 1 || a.Preflight.ReadyNumericCases != 0 || a.Preflight.DUDComputed {
		t.Fatalf("unexpected preflight counts: %+v", a.Preflight)
	}
	if !a.Preflight.MissingNumericRejected || !a.Preflight.CabibboRayRejected || !a.Preflight.NativePromotionRejected || !a.Preflight.CKMNativePredictionRejected {
		t.Fatalf("expected fail-closed routes: %+v", a.Preflight)
	}
	if a.Firewall.CKMNativePrediction || a.Firewall.CKMMatrixConstructed || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestEvaluateCasePreflightRejectsMissingIK(t *testing.T) {
	u := completeRedacted("u")
	u.HasIK = false
	c := EvaluateCase("bad", u, completeRedacted("d"))
	if c.Accepted || !contains(c.Failures, StatusFailedMissingIK) {
		t.Fatalf("missing IK should fail: %+v", c)
	}
}

func TestRenderAuditContainsGate469Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 469 Registry Audit", StatusPreflightValidated, "I_spec", "I_K", StatusFailedMissingNumeric, "Cabibbo remains a residual target"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func contains(xs []string, want string) bool {
	for _, x := range xs {
		if x == want {
			return true
		}
	}
	return false
}
