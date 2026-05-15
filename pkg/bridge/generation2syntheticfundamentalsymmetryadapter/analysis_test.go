package generation2syntheticfundamentalsymmetryadapter

import (
	"strings"
	"testing"
)

func TestGate532Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate531AirlockDefined || !a.Inheritance.Gate531SchemaRowsEnumerated || !a.Inheritance.Gate531RequiresKreinMetric || !a.Inheritance.Gate531RequiresTheta || !a.Inheritance.Gate531RequiresProjectorCompat || !a.Inheritance.Gate531RejectsNativePromotion || !a.Inheritance.Gate531ComparatorBlocked || !a.Inheritance.Gate531HilbertWickOSBlocked || !a.Inheritance.Gate532SyntheticRedirect {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Import.Loaded || a.Import.AcceptedRows != 1 || a.Import.RejectedRows != 0 || !a.Import.MetadataComplete || !a.Import.AllRowsBridgeOnly || !a.Import.AllRowsComparatorOnly || !a.Import.AllRowsMatrixPositivityOnly || !a.Import.AllRowsNoTheoremInput || !a.Import.AllRowsSynthetic || a.Import.AnyObservedClaim {
		t.Fatalf("bad import: %+v", a.Import)
	}
	if !a.Output.FiniteMatrixPlumbingVerified || !a.Output.PositiveHilbertMatrixVerified || !a.Output.GThetaPositiveDefinite || a.Output.GThetaPositiveEigenvalues != 8 || a.Output.GThetaNegativeEigenvalues != 0 || a.Output.GThetaZeroEigenvalues != 0 || a.Output.PhysicalHilbertSpaceGranted || a.Output.WickRotationGranted || a.Output.ReflectionPositivityGranted || a.Output.PositiveEnergyGranted || a.Output.UnitaryRealTimeGranted || a.Output.GlobalHyperbolicityGranted || a.Output.ArrowOfTimeSelected {
		t.Fatalf("bad output: %+v", a.Output)
	}
	if a.Firewall.ObservedHilbertDataImported || a.Firewall.ObservedWickDataImported || a.Firewall.ObservedBoundaryDataImported || !a.Firewall.SyntheticFixtureOnly || a.Firewall.NativeFundamentalSymmetryWrite || a.Firewall.NativeHilbertProductWrite || a.Firewall.NativePhysicalStateSpaceWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeReflectionWrite || a.Firewall.NativePositiveEnergyWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 533 {
		t.Fatalf("unexpected next gate: %+v", a.Next)
	}
}

func TestGate532Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 532 Registry Audit", StatusGThetaPositiveDefinite, StatusFailedPositiveMatrixNotOS, StatusFirewallNativeWriteBlocked, "Θ", "Gate 533"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate532Theorem(t *testing.T) {
	result := Generation2SyntheticFundamentalSymmetryLedgerAdapterPositivityDryRunTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
