package generation2firsttracefunctionalselectionandspectralorderaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate689Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ResponseOperatorSpectrumInherited || a.Inherited.Operator != "R_split = S_split P_K7" || a.Inherited.PriorSpectrumTraceSelectsK7 || !a.Inherited.PriorSupportSelectsK7 {
		t.Fatalf("bad Gate688 inheritance: %+v", a.Inherited)
	}
	if a.Discipline.ClaimsNativeFirstTraceTheorem || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.ClaimsK7IdentityFromTrace {
		t.Fatalf("discipline firewall violated: %+v", a.Discipline)
	}
}

func TestSpectralFunctionalCandidates(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Functionals.Candidates) != 6 || a.Functionals.CandidateCount != 6 {
		t.Fatalf("expected six candidate functionals: %+v", a.Functionals)
	}
	if math.Abs(a.Functionals.FirstOrdinaryTrace-(7.0/72.0)*auditedSSplit) > residualTolerance {
		t.Fatalf("bad first trace: %+v", a.Functionals)
	}
	if math.Abs(a.Functionals.QuadraticTrace-(7.0/72.0)*auditedSSplit*auditedSSplit) > residualTolerance || math.Abs(a.Functionals.FrobeniusNorm-a.Functionals.QuadraticTrace) > residualTolerance {
		t.Fatalf("bad quadratic/Frobenius response: %+v", a.Functionals)
	}
	if math.Abs(a.Functionals.HodgeSignedTrace-(1.0/72.0)*auditedSSplit) > residualTolerance || math.Abs(a.Functionals.FullIdentityTrace-auditedSSplit) > residualTolerance {
		t.Fatalf("bad signed/full response: %+v", a.Functionals)
	}
}

func TestResidualComparisonSelectsFirstTrace(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Residuals.BestCandidate != "F_1 = Tr(R_split)/72" || math.Abs(a.Residuals.BestResidual-math.Abs(firstTraceResidual)) > residualTolerance {
		t.Fatalf("first trace should be best residual candidate: %+v", a.Residuals)
	}
	if !a.Residuals.QuadraticTooSmall || !a.Residuals.FrobeniusTooSmall || !a.Residuals.HodgeSignedTooSmall || !a.Residuals.FullIdentityTooLarge {
		t.Fatalf("expected inactive candidate scale diagnostics: %+v", a.Residuals)
	}
	if !strings.Contains(a.Residuals.Verdict, StatusQuadraticTraceOrFrobeniusNotActive) || !strings.Contains(a.Residuals.Verdict, StatusHodgeSignedTraceNotActive) || !strings.Contains(a.Residuals.Verdict, StatusFullIdentityTraceNotActive) {
		t.Fatalf("missing inactive-route verdicts: %+v", a.Residuals)
	}
}

func TestLinearOrderAudit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Order.DBaseLinearInWallCoords || !a.Order.SSplitLinearInWallCoords || a.Order.RequiredFunctionalOrder != 1 {
		t.Fatalf("bad linear order audit: %+v", a.Order)
	}
	if a.Order.FirstTraceOrder != 1 || a.Order.QuadraticTraceOrder != 2 || a.Order.CubicTraceOrder != 3 || !a.Order.HigherPowersAreInactive {
		t.Fatalf("bad spectral order classification: %+v", a.Order)
	}
}

func TestTraceTypeAudit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.TraceType.K7PlusDimension != 4 || a.TraceType.K7MinusDimension != 3 || a.TraceType.OrdinaryMultiplicity != 7 || a.TraceType.HodgeSignedMultiplicity != 1 {
		t.Fatalf("bad trace type dimensions: %+v", a.TraceType)
	}
	if !a.TraceType.ActiveUsesTotalSupport || a.TraceType.ActiveUsesSignedPolarity {
		t.Fatalf("active bridge should use ordinary total support trace, not signed polarity: %+v", a.TraceType)
	}
}

func TestComparativeSelectionAndMissingTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Selection.SelectionIsComparative || !strings.Contains(a.Selection.SelectedFunctional, "Tr(R_split)/72") {
		t.Fatalf("bad functional selection: %+v", a.Selection)
	}
	if a.Selection.NativeFirstTraceProved || a.Selection.NativeSevenOver72Proved {
		t.Fatalf("Gate689 must not prove native first-trace or 7/72 theorem: %+v", a.Selection)
	}
	if !strings.Contains(a.Missing.PreciseGap, "HistoryResponseFirstTraceTheorem") {
		t.Fatalf("missing theorem target not sharpened: %+v", a.Missing)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2FirstTraceFunctionalSelectionAndSpectralOrderAuditTheorem().Verify()
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
