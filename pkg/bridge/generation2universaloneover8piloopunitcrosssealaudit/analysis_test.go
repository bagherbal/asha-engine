package generation2universaloneover8piloopunitcrosssealaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate623Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	L := 1 / (8 * math.Pi)
	if a.ScalarInherited.Verdict != StatusGate622Inherited {
		t.Fatalf("bad scalar inheritance: %s", a.ScalarInherited.Verdict)
	}
	if math.Abs(a.ScalarInherited.RhoLambdaMatch-0.0380251779225699) > 1e-15 {
		t.Fatalf("bad rho %.18g", a.ScalarInherited.RhoLambdaMatch)
	}
	if math.Abs(a.NormalForm.ScalarKappaLambda-0.0443230430960771) > 1e-14 {
		t.Fatalf("bad kappa_lambda %.18g", a.NormalForm.ScalarKappaLambda)
	}
	if math.Abs(a.FlavorInherited.KappaE-0.00550355419157456) > 1e-16 {
		t.Fatalf("bad kappa_e %.18g", a.FlavorInherited.KappaE)
	}
	if math.Abs(a.ScalarQuality.LambdaAnsatz-0.129872838897183) > 1e-12 {
		t.Fatalf("bad scalar ansatz %.18g", a.ScalarQuality.LambdaAnsatz)
	}
	if math.Abs(a.ScalarQuality.AnsatzMinusRuntime) > 3e-4 {
		t.Fatalf("scalar ansatz not close enough %.18g", a.ScalarQuality.AnsatzMinusRuntime)
	}
	if math.Abs(a.FlavorQuality.EpsilonRawL-L) > 1e-18 {
		t.Fatalf("bad raw L %.18g", a.FlavorQuality.EpsilonRawL)
	}
	if !(a.FlavorQuality.ResidualImprovementFactor > 70) {
		t.Fatalf("expected large flavor improvement, got %.12g", a.FlavorQuality.ResidualImprovementFactor)
	}
	if !a.SignRole.OppositeSigns || a.SignRole.NativeTheoremClaimed {
		t.Fatalf("bad sign-role audit: %+v", a.SignRole)
	}
	if !a.CrossSealType.BridgeOnly {
		t.Fatalf("cross seal must remain bridge-only")
	}
	if a.NativeStatus.NativeCrossSealTheorem {
		t.Fatalf("must not certify native cross-seal theorem")
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2UniversalOneOver8PiLoopUnitCrossSealAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusNormalFormWritten, StatusAppearsInBoth, StatusScalarAnsatzClose, StatusFlavorBalanceClose, StatusNoCrossSealTheorem, StatusGate623Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
