package generation2syntheticgravitycosmologyadapter

import (
	"math"
	"strings"
	"testing"
)

func TestGate515Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate514Inherited || !a.Inheritance.RedactedSchemaAccepted || a.Inheritance.RequiredRows != 10 || a.Inheritance.AcceptedCases != 1 || a.Inheritance.RejectedCases != 8 || !a.Inheritance.Gate514NoAdapterExecuted || !a.Inheritance.Gate514NativeWriteBlocked {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Inputs.AllInputsSynthetic || !a.Inputs.AllRowsBridgeOnly || !a.Inputs.AllNativePromotionBlocked || a.Inputs.ObservedDataImported || a.Inputs.LambdaCutoff != 2 || a.Inputs.F2Moment != 3 || a.Inputs.F4Moment != 5 || a.Inputs.F0Moment != 7 || a.Inputs.VacuumSubtraction != 11 {
		t.Fatalf("bad inputs: %+v", a.Inputs)
	}
	if !nearly(a.Output.F2LambdaSquared, 12, 1e-12) || !nearly(a.Output.F4LambdaFourth, 80, 1e-12) || !nearly(a.Output.EinsteinHilbertCoefficient, 6/(math.Pi*math.Pi), 1e-12) || !nearly(a.Output.CosmologicalVolumeRaw, 480/(math.Pi*math.Pi), 1e-12) || !nearly(a.Output.CosmologicalAfterSubtraction, 480/(math.Pi*math.Pi)-11, 1e-12) || a.Output.NativeGravityPrediction || a.Output.NativeCosmologyPrediction {
		t.Fatalf("bad output: %+v", a.Output)
	}
	if !a.Residuals.ResidualsZeroByConstruction || a.Residuals.ObservedComparatorUsed || !nearly(a.Residuals.EHComparatorResidual, 0, 1e-12) || !nearly(a.Residuals.CosmologicalComparatorResidual, 0, 1e-12) {
		t.Fatalf("bad residuals: %+v", a.Residuals)
	}
	if !a.Airlock.NumericalAdapterExecuted || !a.Airlock.SyntheticOnly || a.Airlock.ObservedComparatorImported || a.Airlock.LambdaNativeSelected || a.Airlock.F2NativeSelected || a.Airlock.F4NativeSelected || a.Airlock.NewtonConstantDerived || a.Airlock.CosmologicalConstantDerived || a.Airlock.NativeNormalizationWrite {
		t.Fatalf("bad airlock: %+v", a.Airlock)
	}
	if a.Firewall.NewtonConstantImported || a.Firewall.PlanckScaleImported || a.Firewall.CutoffLambdaImported || a.Firewall.F2MomentImported || a.Firewall.F4MomentImported || a.Firewall.CosmologicalConstantImported || a.Firewall.DarkEnergyImported || a.Firewall.SyntheticOutputNativeWrite {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 516 {
		t.Fatalf("unexpected next gate: %+v", a.Next)
	}
}

func TestGate515Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 515 Registry Audit", StatusSyntheticAdapterExecuted, StatusFailedSyntheticOutputsNotNative, StatusFirewallSyntheticNativeWriteBlocked, "C_EH", "f2Λ²", "Gate 516"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate515Theorem(t *testing.T) {
	result := Generation2BridgeOnlyGravityCosmologyAdapterDryRunTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
