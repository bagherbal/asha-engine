package generation2yukawatracescalarnormalization

import (
	"strings"
	"testing"
)

func TestGate501YukawaTraceScalarNormalizationAirlock(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Inheritance.SymbolicScalarKineticReadOff || !a.Inheritance.CoefficientDependsOnTraceA || a.Inheritance.Gate500TraceANativeNumeric || !a.Inheritance.Gate500WZMassBlocked {
		t.Fatalf("expected Gate500 symbolic kinetic channel with unresolved trace-a obstruction: %+v", a.Inheritance)
	}
	if !a.Trace.PositiveSemidefiniteNorm || !a.Trace.BasisInvariant || !a.Trace.RephasingInvariant || !a.Trace.CKMOrientationIndependent {
		t.Fatalf("expected trace a to be invariant symbolic norm: %+v", a.Trace)
	}
	if !a.Trace.DependsOnYukawaSingularVals || !a.Trace.DependsOnYukawaAmplitudes || a.Trace.DiscreteTopologicalCharge || a.Trace.NativeNumericValueDerived {
		t.Fatalf("trace a was over-promoted: %+v", a.Trace)
	}
	if !a.YukawaAirlock.NativeYukawaSelectorBranchClosed || !a.YukawaAirlock.TraceASealedByFirewall || a.YukawaAirlock.YukawaNativeSelectorsPassing != 0 || a.YukawaAirlock.RankThreeYukawaMatricesDerived {
		t.Fatalf("expected trace a to remain sealed by Yukawa airlock: %+v", a.YukawaAirlock)
	}
	if !a.Decision.TraceABridgeScalarNormAccepted || a.Decision.TraceANativeNumericAccepted || a.Decision.ScalarKineticCoefficientNative || a.Decision.CanonicalScalarMetricNative || a.Decision.KappaU1Native || a.Decision.WZMassMatrixNative {
		t.Fatalf("normalization decision over-promoted trace a: %+v", a.Decision)
	}
	if a.Firewall.ObservedYukawaImported || a.Firewall.ObservedFermionMassImported || a.Firewall.ObservedCKMPMNSImported || a.Firewall.ObservedWMassImported || a.Firewall.ObservedHiggsVEVImported || a.Firewall.NativeTraceAValueWritten || a.Firewall.NativeWZMassWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 502 {
		t.Fatalf("expected Gate502 redirect, got %+v", a.Next)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 501 Registry Audit", "a = Tr(Y†Y)", StatusFailedTraceValueNotNative, "Gate 502"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate501TheoremPasses(t *testing.T) {
	res := Generation2YukawaTraceScalarNormalizationAirlockAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
