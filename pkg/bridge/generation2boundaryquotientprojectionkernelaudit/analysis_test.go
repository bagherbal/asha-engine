package generation2boundaryquotientprojectionkernelaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate679Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.AugmentedDiagramInherited || !a.Inherited.WeakerDiagramLawful || a.Inherited.StrictExactnessCertified || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.Projection.DomainDimension != 72 || a.Projection.CodomainDimension != 1 || !a.Projection.Surjective {
		t.Fatalf("bad projection: %+v", a.Projection)
	}
	if a.Kernel.TotalKernelDimension != 71 || a.Kernel.K7IsFullKernel || !a.Kernel.K7InsideKernel {
		t.Fatalf("bad kernel: %+v", a.Kernel)
	}
	if a.Defect.Rank != 7 || a.Defect.FullKernel || math.Abs(a.Defect.RelativeToAmbient-7.0/72.0) > 1e-15 || math.Abs(a.Defect.RelativeToKernel-7.0/71.0) > 1e-15 {
		t.Fatalf("bad defect: %+v", a.Defect)
	}
	if math.Abs(a.Trace.Residual-8.52583439801e-10) > 1e-14 || !a.Trace.UsesGlobalH72 {
		t.Fatalf("bad trace: %+v", a.Trace)
	}
	if a.Discipline.ClaimsK7KernelOfPiSplit || a.Discipline.ClaimsLiteralExactSequence || a.Discipline.ClaimsNativeGlobalTraceTheorem || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.Verdict != StatusGate679Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestK7IsNotProjectionKernel(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Kernel.TotalKernelDimension == a.Kernel.K7Rank {
		t.Fatal("K7 must not be the full kernel")
	}
	if a.Kernel.TotalKernelDimension != 70+1 {
		t.Fatalf("expected 71-dim kernel, got %d", a.Kernel.TotalKernelDimension)
	}
	if !strings.Contains(a.Kernel.Verdict, StatusLiteralExactSequenceWithK7KernelBlocked) {
		t.Fatalf("missing blocked exact-sequence verdict: %s", a.Kernel.Verdict)
	}
}

func TestGlobalTraceBeatsAlternatives(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	global := a.Alternatives[0]
	if global.Name != "tau_global" {
		t.Fatalf("unexpected first alternative: %+v", global)
	}
	for _, alt := range a.Alternatives[1:] {
		if !(global.AbsResidual < alt.AbsResidual) {
			t.Fatalf("global should beat %s: global=%g alt=%g", alt.Name, global.AbsResidual, alt.AbsResidual)
		}
	}
	if !strings.Contains(a.Missing.Verdict, StatusNoNativeReasonForGlobalH72Trace) {
		t.Fatal("missing global H72 trace theorem firewall")
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2BoundaryQuotientProjectionKernelAndRelativeTraceResponseAuditTheorem().Verify()
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
