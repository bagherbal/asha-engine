package generation2scalaroneeighthproxyloopmatchingaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate622Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.Verdict != StatusGate621Inherited {
		t.Fatalf("bad inherited verdict %s", a.Inherited.Verdict)
	}
	if math.Abs(a.OneEighth.ProxyMinusOneEighth+0.00009689763984987998) > 1e-15 {
		t.Fatalf("bad proxy-one-eighth %.18g", a.OneEighth.ProxyMinusOneEighth)
	}
	if math.Abs(a.Inherited.RelativeToProxy-0.03802517792256989) > 1e-15 {
		t.Fatalf("bad relative gap %.18g", a.Inherited.RelativeToProxy)
	}
	if math.Abs(a.Inherited.RelativeToProxy-1/(8*math.Pi))/(1/(8*math.Pi)) > 0.05 {
		t.Fatalf("relative gap not close to loop unit")
	}
	if math.Abs(a.Inherited.DeltaLambdaMatch-1/(64*math.Pi))/(1/(64*math.Pi)) > 0.05 {
		t.Fatalf("absolute gap not close to 1/(64pi)")
	}
	if a.RelativeLoops.ClosestName != "1/(8*pi)" {
		t.Fatalf("unexpected closest relative %q", a.RelativeLoops.ClosestName)
	}
	if a.AbsoluteLoops.ClosestName != "lambda_proxy/(8*pi)" {
		t.Fatalf("unexpected closest absolute %q", a.AbsoluteLoops.ClosestName)
	}
	if !(a.HiggsDiagnostic.LambdaAnsatz > a.HiggsDiagnostic.LambdaProxy) {
		t.Fatalf("ansatz should increase proxy")
	}
	if a.HiggsDiagnostic.ClaimsHiggsPrediction {
		t.Fatalf("must not claim Higgs prediction")
	}
	if !a.Sign.PositiveCorrection {
		t.Fatalf("positive correction required")
	}
	if a.NativeStatus.NativeOneOver8PiScalarMatching {
		t.Fatalf("no native loop theorem")
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2ScalarOneEighthProxyLoopMatchingAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusOneEighthProxyAudited, StatusLoopSized, StatusOneOver8PiClose, StatusNoLoopMatchingTheorem, StatusGate622Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
