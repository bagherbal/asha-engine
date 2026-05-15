package generation2synthetic3plus1projectionadapter

import (
	"strings"
	"testing"
)

func TestGate530Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate529AirlockDefined || !a.Inheritance.Gate529ProjectorSchemaReady || !a.Inheritance.Gate529RequiresSourceConvention || !a.Inheritance.Gate529RequiresBridgeOnly || !a.Inheritance.Gate529RejectsNativePromotion || !a.Inheritance.Gate529ComparatorExecutionBlocked || !a.Inheritance.Gate529WickHilbertUnitaryBlocked || !a.Inheritance.Gate529InternalGaugeBlocked || !a.Inheritance.Gate529NoObservedDimensionData || !a.Inheritance.Gate529NativeRegistryBlocked || !a.Inheritance.Gate530FileAdapterRedirect {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Import.Loaded || a.Import.Rows != 1 || a.Import.AcceptedRows != 1 || a.Import.RejectedRows != 0 || !a.Import.BridgeOnlyLedger || !a.Import.SyntheticFixture || a.Import.ObservedDimensionLoaded || a.Import.NativeRegistryWriteRequested || !a.Import.MetadataComplete || !a.Import.AllRowsBridgeOnly || !a.Import.AllRowsComparatorOnly || !a.Import.AllRowsNoTheoremInput || !a.Import.AllRowsSynthetic || a.Import.AnyObservedClaim {
		t.Fatalf("bad import: %+v", a.Import)
	}
	if !a.Output.Ready || a.Output.ProjectorRank != 4 || a.Output.ComplementRank != 4 || !a.Output.ProjectorRankValid || !a.Output.ComplementRankValid || !nearly(a.Output.PIdempotencyResidual, 0, tolerance) || !nearly(a.Output.QIdempotencyResidual, 0, tolerance) || !nearly(a.Output.PQOrthogonalityNorm, 0, tolerance) || !nearly(a.Output.QPOrthogonalityNorm, 0, tolerance) || !nearly(a.Output.PPlusQIdentityNorm, 0, tolerance) || !nearly(a.Output.MetricCrossResidual, 0, tolerance) || a.Output.ExternalPositive != 1 || a.Output.ExternalNegative != 3 || !a.Output.ExternalSignatureOK || !a.Output.InternalRankOK || !a.Output.AllResidualsZero || !a.Output.CliffordCompatible || a.Output.NativePrediction {
		t.Fatalf("bad output: %+v", a.Output)
	}
	if a.Firewall.ObservedDimensionImported || !a.Firewall.SyntheticFixtureOnly || a.Firewall.FileRowsNative || a.Firewall.AdapterOutputsNative || a.Firewall.ProjectorNativePrediction || a.Firewall.External3Plus1NativePrediction || a.Firewall.InternalComplementNativePrediction || a.Firewall.WickRotationGranted || a.Firewall.PositiveHilbertGranted || a.Firewall.ReflectionPositivityGranted || a.Firewall.PositiveEnergyGranted || a.Firewall.UnitaryRealTimeGranted || a.Firewall.GlobalHyperbolicityGranted || a.Firewall.InternalGaugeNativeIdentification || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 531 {
		t.Fatalf("unexpected next gate: %+v", a.Next)
	}
}

func TestGate530Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 530 Registry Audit", StatusAdapterExecuted, StatusCl17ExternalSignatureConfirmed, StatusFailedZeroResidualsNotNative, StatusFirewallNativeWriteBlocked, "P^2-P", "Gate 531"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate530Theorem(t *testing.T) {
	result := Generation2Synthetic3Plus1ProjectionFileAdapterCliffordFirewallTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
