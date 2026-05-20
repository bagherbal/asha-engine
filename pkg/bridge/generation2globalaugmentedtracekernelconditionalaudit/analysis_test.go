package generation2globalaugmentedtracekernelconditionalaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate680Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.RelativeTraceInherited || a.Inherited.H72Dimension != 72 || a.Inherited.KernelDimension != 71 || a.Inherited.K7Rank != 7 || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Sequence.ExactByDimension || a.Sequence.KernelDimension+a.Sequence.QuotientDimension != a.Sequence.AmbientDimension {
		t.Fatalf("bad sequence: %+v", a.Sequence)
	}
	if a.Defect.FullKernel || math.Abs(a.Defect.GlobalDensity-7.0/72.0) > 1e-15 || math.Abs(a.Defect.KernelConditionalDensity-7.0/71.0) > 1e-15 {
		t.Fatalf("bad defect: %+v", a.Defect)
	}
	if len(a.Normalizations) != 4 || a.Normalizations[0].Name != "tau_global" || a.Normalizations[0].AbsResidual > 1e-8 {
		t.Fatalf("bad normalizations: %+v", a.Normalizations)
	}
	if a.Discipline.ClaimsNativeGlobalPrinciple || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.Verdict != StatusGate680Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestGlobalTraceBeatsKernelAndFinite(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	global := a.Normalizations[0]
	kernel := a.Normalizations[1]
	finite := a.Normalizations[2]
	if !(global.AbsResidual < kernel.AbsResidual && global.AbsResidual < finite.AbsResidual) {
		t.Fatalf("global should beat kernel/finite: global=%g kernel=%g finite=%g", global.AbsResidual, kernel.AbsResidual, finite.AbsResidual)
	}
	if !kernel.KernelExcludesQuotientForTest() {
		// keep test intent readable without exporting implementation internals
		if kernel.IncludesQuotientLine {
			t.Fatal("kernel conditional normalization must exclude quotient input")
		}
	}
}

func (x TraceNormalization) KernelExcludesQuotientForTest() bool {
	return x.Name == "tau_kernel" && !x.IncludesQuotientLine
}

func TestMissingPrinciple(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Missing.Verdict, StatusNoNativeGlobalTraceResponsePrinciple) {
		t.Fatal("missing global trace principle firewall")
	}
	if !strings.Contains(a.Compatibility.Verdict, StatusSevenOver72FullExtensionDefectDensity) {
		t.Fatal("missing full extension support")
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2GlobalAugmentedTraceVersusKernelConditionalTraceAuditTheorem().Verify()
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
