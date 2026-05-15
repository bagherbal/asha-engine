package generation2leptonpmnsnullresidual

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate476ComputesSyntheticPMNSNullResidual(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate475LeptonPreflightDefined || !a.Map.StructurallyIdenticalToQuarkMap {
		t.Fatalf("bad inheritance/map: %+v %+v", a.Inheritance, a.Map)
	}
	if a.Sieve.AcceptedCaseCount != 1 || !a.Sieve.ValidSyntheticResidualAccepted || !a.Sieve.AllAcceptedBridgeOnlySynthetic {
		t.Fatalf("bad sieve: %+v", a.Sieve)
	}
	if a.Output.SyntheticDENu <= 0 || a.Output.SyntheticPMNSTarget <= 0 || a.Output.SyntheticResidual < 0 {
		t.Fatalf("bad output: %+v", a.Output)
	}
	if !a.Sieve.ObservedPMNSRejected || !a.Sieve.PMNSAsRayInputRejected || !a.Sieve.PMNSNativePredictionRejected || !a.Sieve.PMNSMatrixExportRejected || !a.Sieve.NativeResidualPromotionRejected {
		t.Fatalf("unsafe routes not rejected: %+v", a.Sieve)
	}
	if a.Firewall.PMNSMatrixConstructed || a.Firewall.PMNSEntryComputed || a.Firewall.PMNSNativePrediction || a.Firewall.DENuNativePrediction || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestEvaluateCaseFormula(t *testing.T) {
	c := EvaluateCase("formula", syntheticComparator("e", 0.85, 0.30), syntheticComparator("nu", 1.05, 0.90), 0.620)
	if !c.Accepted {
		t.Fatalf("case rejected: %+v", c)
	}
	if c.Residual.DENu < 0.623 || c.Residual.DENu > 0.625 {
		t.Fatalf("unexpected d_e_nu %.15f", c.Residual.DENu)
	}
	if c.Residual.PMNSMatrixConstructed || c.Residual.PMNSEntryComputed || c.Residual.ExportsNativeObservable {
		t.Fatalf("residual escaped bridge mode: %+v", c.Residual)
	}
}

func TestRenderAuditContainsGate476Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 476 Registry Audit", StatusPMNSNullResidualFirewallValidated, "d_e_nu", "structurally identical", "PMNS may be a residual target"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
