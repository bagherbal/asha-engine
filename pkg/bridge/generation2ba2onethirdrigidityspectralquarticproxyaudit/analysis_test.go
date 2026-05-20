package generation2ba2onethirdrigidityspectralquarticproxyaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate620Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.Verdict != StatusGate619Inherited {
		t.Fatalf("bad inherited verdict %s", a.Inherited.Verdict)
	}
	if len(a.RigidityRows) != 2 {
		t.Fatalf("expected two rigidity rows")
	}
	if math.Abs(a.RigidityRows[0].BOverA2-0.33307493962706697) > 1e-15 {
		t.Fatalf("bad MZ b/a2 %.18g", a.RigidityRows[0].BOverA2)
	}
	if math.Abs(a.RigidityRows[1].BOverA2-0.3330764110541872) > 1e-15 {
		t.Fatalf("bad L12 b/a2 %.18g", a.RigidityRows[1].BOverA2)
	}
	if !a.RigiditySummary.NearlyInvariant {
		t.Fatalf("expected near invariant b/a2")
	}
	if len(a.ProxyRows) != 2 || a.ProxyRows[0].CLambdaCandidate != sin2Theta {
		t.Fatalf("bad proxy rows")
	}
	if !(a.ProxyRows[0].AbsResidual < 0.006) {
		t.Fatalf("expected low-scale proxy to be close, got %+v", a.ProxyRows[0])
	}
	if a.ProxyRows[1].SignCompatible {
		t.Fatalf("expected high-scale sign failure")
	}
	if a.HiggsProxy.ClaimsHiggsDerivation {
		t.Fatalf("must not claim Higgs derivation")
	}
	if !a.StressImpact.StressUsesRuntimeLambda || a.StressImpact.CanReplaceRuntimeWithProxy {
		t.Fatalf("bad stress impact %+v", a.StressImpact)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2BA2OneThirdRigiditySpectralQuarticProxyAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusBA2NearOneThird, StatusProxyCloseAtMZ, StatusProxyFailsAtL12, StatusSeparateScalarLanes, StatusGate620Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
