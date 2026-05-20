package generation2augmentedchamberdefecttraceresponseaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate674Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.InheritedLinePullback || !a.Inherited.FullK7BoundaryMapFailed || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.Chamber.Lambda4Dimension != 70 || a.Chamber.BoundaryDimension != 2 || a.Chamber.TotalDimension != 72 || math.Abs(a.Chamber.TraceWeight-sevenOver72) > 1e-15 {
		t.Fatalf("bad chamber: %+v", a.Chamber)
	}
	if a.RankSeven.DimK7 != 7 || a.RankSeven.DimKernelA != 7 || a.RankSeven.DimCokernelA != 7 || a.RankSeven.FanoHitchinCarrierDimension != 7 || len(a.RankSeven.CandidateSources) != 4 {
		t.Fatalf("bad rank-seven audit: %+v", a.RankSeven)
	}
	if math.Abs(a.Trace.QTrace-sevenOver72) > 1e-15 || math.Abs(a.Trace.TraceResidual-8.52583439801e-10) > 1e-14 || a.Trace.RequiresVectorMap || !a.Trace.RequiresScalarTraceMap {
		t.Fatalf("bad trace response: %+v", a.Trace)
	}
	if a.Alternatives.BestName != "7/72" || len(a.Alternatives.Alternatives) != 4 {
		t.Fatalf("bad alternatives: %+v", a.Alternatives)
	}
	if len(a.Missing.NativeTheoremTargets) != 2 || len(a.Missing.MissingTheorems) != 4 || len(a.Missing.AllowedSupport) != 4 {
		t.Fatalf("bad missing theorem audit: %+v", a.Missing)
	}
	if a.Discipline.ClaimsNativeTraceResponse || a.Discipline.ClaimsNativeStressSplitPullback || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.ClaimsFullK7BoundaryMap || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsHiggsMassPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.Verdict != StatusGate674Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestDenominatorAlternativeOrdering(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	best := a.Alternatives.Alternatives[0]
	for _, c := range a.Alternatives.Alternatives[1:] {
		if c.AbsResidual < best.AbsResidual {
			best = c
		}
	}
	if best.Name != "7/72" {
		t.Fatalf("expected 7/72 best, got %+v", best)
	}
}

func TestTraceResponseIdentity(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	got := a.Trace.DBase - (float64(a.RankSeven.NumeratorCandidate)/float64(a.Chamber.TotalDimension))*a.Trace.SSplit
	if math.Abs(got-a.Trace.TraceResidual) > 1e-15 {
		t.Fatalf("identity mismatch: got %.17g residual %.17g", got, a.Trace.TraceResidual)
	}
	if math.Abs(got-a.Inherited.SevenOver72Residual) > 1e-15 {
		t.Fatalf("inheritance mismatch: got %.17g inherited %.17g", got, a.Inherited.SevenOver72Residual)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2AugmentedChamberDefectTraceResponseCoefficientAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
