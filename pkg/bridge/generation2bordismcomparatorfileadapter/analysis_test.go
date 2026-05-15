package generation2bordismcomparatorfileadapter

import (
	"strings"
	"testing"
)

func TestGate522Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate521ClassifierDefined || !a.Inheritance.Gate521OrientedSocket || !a.Inheritance.Gate521SpinSocket || !a.Inheritance.Gate521SpinCSocket || !a.Inheritance.Gate521BoundarySocket || !a.Inheritance.Gate521CharacteristicResidual || !a.Inheritance.Gate521ScaleFree || a.Inheritance.Gate521SpecificClassSelected || a.Inheritance.Gate521ManifoldSelected || a.Inheritance.Gate521ObservedDataImported || !a.Inheritance.Gate521NativeWriteBlocked || !a.Inheritance.Gate522FileAdapterRedirect {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Import.Loaded || a.Import.Rows != 12 || a.Import.AcceptedRows != 12 || a.Import.RejectedRows != 0 || a.Import.StiefelWhitneyRows != 4 || a.Import.CharacteristicRows != 4 || a.Import.BoundaryRows != 2 || a.Import.BordismRows != 1 || a.Import.AdapterRows != 1 || !a.Import.EmpiricalImport || !a.Import.BridgeOnlyLedger || !a.Import.SyntheticFixture || a.Import.ObservedValuesLoaded || a.Import.NativeRegistryWriteRequested || !a.Import.MetadataComplete || !a.Import.AllRowsBridgeOnly || !a.Import.AllRowsComparatorOnly || !a.Import.AllRowsNoTheoremInput {
		t.Fatalf("bad import: %+v", a.Import)
	}
	if !a.Output.Ready || !a.Output.OrientedAdmissible || !a.Output.SpinAdmissible || !a.Output.SpinCAdmissible || !a.Output.ClosedBoundary || !a.Output.CharacteristicAdmissible || !a.Output.OverallAdmissible || !nearly(a.Output.SignatureFromP1, -16, 1e-12) || !nearly(a.Output.SignatureP1Residual, 0, 1e-12) || !nearly(a.Output.AHatFromTau, 2, 1e-12) || !nearly(a.Output.AHatResidual, 0, 1e-12) || !a.Output.RokhlinDivisibilityPassed || !nearly(a.Output.C1Mod2W2Residual, 0, 1e-12) || !a.Output.AllResidualsZero || !a.Output.BridgeOnly || a.Output.NativePrediction {
		t.Fatalf("bad output: %+v", a.Output)
	}
	if a.Firewall.ObservedBordismImported || a.Firewall.ObservedTangentBundleImported || a.Firewall.ObservedBoundaryDataImported || !a.Firewall.SyntheticFixtureOnly || a.Firewall.FileRowsNative || a.Firewall.AdapterOutputsNative || a.Firewall.StiefelWhitneyNativePrediction || a.Firewall.SpinStructureNativePrediction || a.Firewall.SpinCStructureNativePrediction || a.Firewall.SpecificBordismClassNative || a.Firewall.ManifoldRepresentativeNative || a.Firewall.CharacteristicNumbersNative || a.Firewall.BoundaryConditionNativeSelected || a.Firewall.NativeRegistryWritten || a.Firewall.NewtonPlanckCosmologyImported {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 523 {
		t.Fatalf("bad next gate: %+v", a.Next)
	}
}

func TestGate522Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 522 Registry Audit", StatusExplicitBordismFileLoaded, StatusStiefelWhitneyMetadataAudited, StatusFirewallNativeWriteBlocked, "stiefel", "Ahat_residual", "Gate 523"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate522Theorem(t *testing.T) {
	result := Generation2BordismComparatorFileAdapterStiefelWhitneyFirewallTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
