package generation2augmenteddefectexactsequenceaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate678Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.TraceOperatorInherited || !a.Inherited.OperatorSharperThanFit || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.Chamber.TotalDimension != 72 || a.Chamber.FiniteDimension != 70 || a.Chamber.BoundaryDimension != 2 {
		t.Fatalf("bad chamber: %+v", a.Chamber)
	}
	if a.Defect.Rank != 7 || a.Defect.VectorBoundaryMapCertified || math.Abs(a.Defect.TauDefect-7.0/72.0) > 1e-15 {
		t.Fatalf("bad defect: %+v", a.Defect)
	}
	if a.Boundary.Dimension != 1 || math.Abs(a.Boundary.SSplit-0.0012924448188163) > 1e-15 {
		t.Fatalf("bad boundary: %+v", a.Boundary)
	}
	if a.History.Dimension != 1 || math.Abs(a.History.DBase-0.0001256552099684) > 1e-15 {
		t.Fatalf("bad history: %+v", a.History)
	}
	if a.Sequence.StrictExactSequenceCertified || !a.Sequence.WeakerDiagramLawful || !a.Sequence.DiagramObjectsCompatible {
		t.Fatalf("bad sequence: %+v", a.Sequence)
	}
	if math.Abs(a.Trace.Residual-8.52583439801e-10) > 1e-14 {
		t.Fatalf("bad residual: %+v", a.Trace)
	}
	if a.Discipline.ClaimsNativeExactSequenceTheorem || a.Discipline.ClaimsNativeTraceToQuotientTheorem || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.Verdict != StatusGate678Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestDiagramCompatibilityIsWeakerThanStrictExactness(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Sequence.StrictExactSequenceCertified {
		t.Fatal("strict exact sequence must not be certified")
	}
	if !a.Sequence.WeakerDiagramLawful || !strings.Contains(a.Sequence.Verdict, StatusWeakerDiagramLawful) {
		t.Fatalf("expected weaker bridge diagram support: %+v", a.Sequence)
	}
	if a.Sequence.ProjectionH72ToQBoundaryTyped || a.Sequence.KernelCokernelExactnessCertified {
		t.Fatalf("should preserve exactness gaps: %+v", a.Sequence)
	}
}

func TestTraceCompatibility(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	pred := a.Trace.TauDefect * a.Trace.SSplit
	if math.Abs(pred-a.Trace.PredictedDBase) > 1e-15 {
		t.Fatal("prediction mismatch")
	}
	if math.Abs(a.Trace.DBase-pred-a.Trace.Residual) > 1e-15 {
		t.Fatal("residual mismatch")
	}
	if !a.Trace.QuotientNormalized {
		t.Fatal("expected quotient-normalized trace response")
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2AugmentedDefectExactSequenceCompatibilityAuditTheorem().Verify()
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
