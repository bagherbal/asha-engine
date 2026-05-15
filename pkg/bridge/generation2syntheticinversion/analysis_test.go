package generation2syntheticinversion

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate468(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Harness.AcceptedSyntheticCases != 1 || !a.Harness.ValidSyntheticDUDComputed || !a.Harness.UncertaintyPropagationExecuted {
		t.Fatalf("expected one accepted synthetic d_ud dry run with uncertainty propagation: %+v", a.Harness)
	}
	c := a.Harness.Cases[0]
	if !c.Accepted || c.Distance.DUD <= 0 || c.Distance.DUDMin <= 0 || c.Distance.DUDMax <= c.Distance.DUDMin {
		t.Fatalf("invalid accepted synthetic distance: %+v", c)
	}
	if c.Distance.CabibboCompared || c.Distance.CKMMatrixConstructed || c.Distance.CKMEntryComputed || c.Distance.NativePrediction {
		t.Fatalf("synthetic d_ud must not become CKM/native: %+v", c.Distance)
	}
	if a.Firewall.CKMNativePrediction || a.Firewall.ObservedMassesImported || a.Firewall.CabibboUsedAsRayInput || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestEvaluateCaseRejectsUnsafeRoutes(t *testing.T) {
	bad := canonicalU()
	bad.ObservedData = true
	res := EvaluateCase("bad", bad, canonicalD())
	if res.Accepted || !contains(res.Failures, StatusFailedObservedDataRejected) {
		t.Fatalf("observed data should fail: %+v", res)
	}
	bad = canonicalU()
	bad.IK = 1
	res = EvaluateCase("bad", bad, canonicalD())
	if res.Accepted || !contains(res.Failures, StatusFailedProjectiveDomainRejected) {
		t.Fatalf("projective boundary should fail: %+v", res)
	}
	bad = canonicalU()
	bad.CabibboAsRayInput = true
	res = EvaluateCase("bad", bad, canonicalD())
	if res.Accepted || !contains(res.Failures, StatusFailedCabibboAsRayRejected) {
		t.Fatalf("Cabibbo-as-ray should fail: %+v", res)
	}
}

func TestRenderAuditContainsGate468Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 468 Registry Audit", StatusSyntheticInversionValidated, "d_ud", "alpha", StatusFailedObservedDataRejected, "not `V_us`"} {
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
