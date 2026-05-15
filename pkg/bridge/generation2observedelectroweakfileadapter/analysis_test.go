package generation2observedelectroweakfileadapter

import (
	"strings"
	"testing"
)

func TestGate507ObservedElectroweakFileAdapter(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Inheritance.Gate506PreflightValidated || a.Inheritance.Gate506NumericalAdapterExecuted || a.Inheritance.Gate506ObservedNumbersImported || !a.Inheritance.Gate506NativeRegistryWriteBlocked || !a.Inheritance.Gate507RedirectDefined {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Import.Loaded || a.Import.Rows != 6 || a.Import.AcceptedRows != 6 || a.Import.RejectedRows != 0 || a.Import.InputRows != 3 || a.Import.ComparatorRows != 3 || !a.Import.EmpiricalImport || !a.Import.BridgeOnlyLedger || !a.Import.SyntheticFixture || a.Import.ObservedValuesLoaded || a.Import.NativeRegistryWriteRequested || !a.Import.MetadataComplete || len(a.Import.Failures) != 0 {
		t.Fatalf("bad import: %+v", a.Import)
	}
	if !a.Input.HasV || !a.Input.HasG2 || !a.Input.HasGY || !nearly(a.Input.V, 2, 1e-12) || !nearly(a.Input.G2, 3, 1e-12) || !nearly(a.Input.GY, 4, 1e-12) || !a.Input.SyntheticFixture || a.Input.ObservedValuesLoaded || !a.Input.BridgeOnly || !a.Input.MetadataComplete || a.Input.NativePromotion {
		t.Fatalf("bad input: %+v", a.Input)
	}
	if !a.Output.Ready || !nearly(a.Output.Sin2ThetaW, 16.0/25.0, 1e-12) || !nearly(a.Output.Cos2ThetaW, 9.0/25.0, 1e-12) || !nearly(a.Output.MW, 3, 1e-12) || !nearly(a.Output.MZ, 5, 1e-12) || !nearly(a.Output.MGamma, 0, 1e-12) || !nearly(a.Output.RhoTree, 1, 1e-12) || !a.Output.PhotonZeroPreserved || !a.Output.RhoIdentityConfirmed {
		t.Fatalf("bad output: %+v", a.Output)
	}
	if !a.Residuals.ComparatorRowsAvailable || !a.Residuals.WeakAngleResidualComputed || !a.Residuals.MWResidualComputed || !a.Residuals.MZResidualComputed || !a.Residuals.AllResidualsZero || !a.Residuals.BridgeOnly || a.Residuals.NativePrediction {
		t.Fatalf("bad residuals: %+v", a.Residuals)
	}
	if a.Firewall.ObservedValuesImported || !a.Firewall.SyntheticFixtureOnly || a.Firewall.FileRowsNative || a.Firewall.AdapterOutputsNative || a.Firewall.WeakAngleNativePrediction || a.Firewall.WZMassNativePrediction || a.Firewall.GaugeCouplingsNativePrediction || a.Firewall.VEVNativePrediction || a.Firewall.KappaNativePromotion || a.Firewall.NativeRegistryWritten || a.Firewall.PhysicalElectroweakPredictionMade {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 508 {
		t.Fatalf("expected Gate508 redirect: %+v", a.Next)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 507 Registry Audit", StatusAdapterExecutedBridgeOnly, StatusComparatorResidualsComputed, StatusFirewallNativeWriteBlocked, "Gate 508"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate507RejectsMissingFile(t *testing.T) {
	_, err := BuildFromFile("data/does_not_exist_gate507.json")
	if err == nil {
		t.Fatal("expected missing file validation error")
	}
}

func TestGate507TheoremPasses(t *testing.T) {
	res := Generation2ObservedElectroweakComparatorFileAdapterFirewallTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
