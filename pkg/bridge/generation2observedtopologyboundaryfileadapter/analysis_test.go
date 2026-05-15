package generation2observedtopologyboundaryfileadapter

import (
	"strings"
	"testing"
)

func TestGate520Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate519PreflightDefined || a.Inheritance.Gate519TopologyRows != 7 || a.Inheritance.Gate519BoundaryRows != 7 || !a.Inheritance.Gate519RequiresBridgeOnly || !a.Inheritance.Gate519RejectsNativePromotion || a.Inheritance.Gate519ComparatorExecuted || a.Inheritance.Gate519ObservedDataImported || !a.Inheritance.Gate519NativeRegistryBlocked || !a.Inheritance.Gate520FileAdapterRedirect {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Import.Loaded || a.Import.Rows != 15 || a.Import.AcceptedRows != 15 || a.Import.RejectedRows != 0 || a.Import.TopologyRows != 7 || a.Import.BoundaryRows != 7 || a.Import.AdapterRows != 1 || !a.Import.EmpiricalImport || !a.Import.BridgeOnlyLedger || !a.Import.SyntheticFixture || a.Import.ObservedValuesLoaded || a.Import.NativeRegistryWriteRequested || !a.Import.MetadataComplete || !a.Import.AllRowsBridgeOnly || !a.Import.AllRowsComparatorOnly || !a.Import.AllRowsNoTheoremInput {
		t.Fatalf("bad import: %+v", a.Import)
	}
	if !a.Output.Ready || !a.Output.UsesAPSBoundaryCorrection || !nearly(a.Output.BoundaryCorrection, 2, 1e-12) || !nearly(a.Output.ComputedAPSIndex, 9, 1e-12) || !nearly(a.Output.APSResidual, 0, 1e-12) || !nearly(a.Output.ComputedSignatureFromP1, 1, 1e-12) || !nearly(a.Output.SignatureResidual, 0, 1e-12) || !a.Output.BoundaryMode || !a.Output.AllResidualsZero || !a.Output.BridgeOnly || a.Output.NativePrediction {
		t.Fatalf("bad output: %+v", a.Output)
	}
	if a.Firewall.ObservedTopologyImported || a.Firewall.ObservedBoundaryDataImported || a.Firewall.ObservedBoundarySpectrumImported || !a.Firewall.SyntheticFixtureOnly || a.Firewall.FileRowsNative || a.Firewall.AdapterOutputsNative || a.Firewall.EulerNativePrediction || a.Firewall.PontryaginNativePrediction || a.Firewall.SignatureNativePrediction || a.Firewall.GlobalAPSIndexNativePrediction || a.Firewall.EtaNativePrediction || a.Firewall.BoundarySpectrumNativePrediction || a.Firewall.BoundaryConditionNativeSelected || a.Firewall.NativeRegistryWritten || a.Firewall.NewtonPlanckCosmologyImported {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 521 {
		t.Fatalf("unexpected next gate: %+v", a.Next)
	}
}

func TestGate520Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 520 Registry Audit", StatusExplicitTopologyFileLoaded, StatusAPSFormulaComputedFromFile, StatusFirewallNativeWriteBlocked, "computed_APS", "signature_residual", "Gate 521"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate520Theorem(t *testing.T) {
	result := Generation2ObservedTopologyBoundaryFileAdapterFirewallTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
