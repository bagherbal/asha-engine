package generation2spectralcutoffrenormalizationairlock

import (
	"strings"
	"testing"
)

func TestGate514Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate513Inherited || !a.Inheritance.StrippedHierarchyNative || !nearly(a.Inheritance.A2OverA0Ratio, 1.0/12.0, 1e-12) || !nearly(a.Inheritance.A4OverA0Ratio, 1.0/360.0, 1e-12) || !nearly(a.Inheritance.A4OverA2Ratio, 1.0/30.0, 1e-12) {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Schema.RowsBridgeOnly || !a.Schema.RowsRejectNativePromotion || !a.Schema.AllMetadataComplete || a.Schema.RequiredRowCount != 10 || a.Schema.NumericalRows != 0 || a.Schema.EmpiricalRows != 0 {
		t.Fatalf("bad schema: %+v", a.Schema)
	}
	if a.Preflight.AcceptedCases != 1 || a.Preflight.RejectedCases != 8 || a.Preflight.RejectedExecutionCases != 1 {
		t.Fatalf("bad preflight: %+v", a.Preflight)
	}
	if a.Airlock.LambdaCutoffSelected || a.Airlock.F2MomentSelected || a.Airlock.F4MomentSelected || a.Airlock.PlanckMatchingNative || a.Airlock.NewtonConstantDerived || a.Airlock.CosmologicalConstantDerived || a.Airlock.VacuumSubtractionSelectedNative || a.Airlock.NumericalAdapterExecuted || a.Airlock.NativeNormalizationWrite {
		t.Fatalf("bad airlock: %+v", a.Airlock)
	}
	if a.Firewall.NewtonConstantImported || a.Firewall.PlanckScaleImported || a.Firewall.CutoffLambdaImported || a.Firewall.F2MomentImported || a.Firewall.F4MomentImported || a.Firewall.CosmologicalConstantImported || a.Firewall.NativeCutoffRenormalizationWrite {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 515 {
		t.Fatalf("unexpected next gate: %+v", a.Next)
	}
}

func TestGate514Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 514 Registry Audit", StatusCutoffRenormalizationSchemaDefined, StatusFailedLambdaNativeSelectionRejected, StatusFirewallNativeWriteBlocked, "cutoff_lambda", "f2_lambda_squared", "Gate 515"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate514Theorem(t *testing.T) {
	result := Generation2SpectralCutoffRenormalizationAirlockComparatorTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
