package generation2scalarproxyruntimematchinggapaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate621Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.Verdict != StatusGate620Inherited {
		t.Fatalf("bad inherited verdict %s", a.Inherited.Verdict)
	}
	if math.Abs(a.MatchingGap.DeltaLambdaMatch-0.0047494626903258) > 1e-15 {
		t.Fatalf("bad delta lambda %.18g", a.MatchingGap.DeltaLambdaMatch)
	}
	if !(a.MatchingGap.RelativeToProxy > 0.037 && a.MatchingGap.RelativeToProxy < 0.039) {
		t.Fatalf("unexpected rel proxy %.18g", a.MatchingGap.RelativeToProxy)
	}
	if math.Abs(a.EffectiveCLambda.CNeededMZ-0.389259441720964) > 1e-15 {
		t.Fatalf("bad c needed %.18g", a.EffectiveCLambda.CNeededMZ)
	}
	if !(a.HiggsProxyGap.MassRuntimeGeV > a.HiggsProxyGap.MassProxyGeV) {
		t.Fatalf("expected runtime mass diagnostic above proxy")
	}
	if a.HiggsProxyGap.ClaimsHiggsDerivation || a.HiggsProxyGap.ClaimsPoleMassTheorem {
		t.Fatalf("must not claim Higgs theorem")
	}
	if !a.Sign.PositiveDeltaLambda {
		t.Fatalf("positive matching correction expected")
	}
	if a.NeutrinoTrace.ValuesInserted {
		t.Fatalf("must not insert neutrino values")
	}
	if !a.StressImpact.StressStillUsesLambdaRuntimeL12 || a.StressImpact.CanReplaceStressLambdaWithProxy {
		t.Fatalf("bad stress impact %+v", a.StressImpact)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2ScalarTreeProxyToRuntimeMatchingGapAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusMatchingGapComputed, StatusPositiveMatchingRequired, StatusProxyRuntimeChainDefined, StatusNoProxyRuntimeTheorem, StatusGate621Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
