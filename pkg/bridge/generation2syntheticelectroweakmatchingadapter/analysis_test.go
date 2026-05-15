package generation2syntheticelectroweakmatchingadapter

import (
	"strings"
	"testing"
)

func TestGate505SyntheticElectroweakMatchingAdapter(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Inheritance.PermissionLedgerAccepted || !a.Inheritance.BridgeInputSchemaDefined || a.Inheritance.NativeRows != 0 || a.Inheritance.BridgeRows != 6 || !a.Inheritance.FormulaBridgeOnly || !a.Inheritance.PermissionAllowsExplicitAdapter || a.Inheritance.Gate504NumericalAdapterExecuted || a.Inheritance.Gate504ObservedEWDataImported || !a.Inheritance.Gate504NativeWriteBlocked {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Input.Synthetic || a.Input.Observed || a.Input.Native || a.Input.V != 2 || a.Input.G2 != 3 || a.Input.GY != 4 || a.Input.RenormalizationScaleTag == "" || a.Input.Scheme == "" {
		t.Fatalf("bad synthetic input: %+v", a.Input)
	}
	if !a.Output.Executed || !a.Output.UsedTreeLevelFormula || !nearly(a.Output.MW, 3, 1e-12) || !nearly(a.Output.MZ, 5, 1e-12) || !nearly(a.Output.Sin2ThetaW, 16.0/25.0, 1e-12) || !nearly(a.Output.Cos2ThetaW, 9.0/25.0, 1e-12) || a.Output.MGamma != 0 || !nearly(a.Output.RhoTree, 1, 1e-12) {
		t.Fatalf("bad adapter output: %+v", a.Output)
	}
	if !a.Adapter.SyntheticOnly || a.Adapter.ObservedDataImported || a.Adapter.NativeDataImported || !a.Adapter.InputsFinite || !a.Adapter.ScaleSchemeMetadataPresent || !a.Adapter.ComputedWithExplicitInputs || !a.Adapter.WeakAngleBridgeOutputOnly || !a.Adapter.WZBridgeOutputOnly || !a.Adapter.PhotonZeroPreserved || !a.Adapter.RhoTreeIdentityConfirmed || a.Adapter.ObservedMassesClaimed || a.Adapter.NativeWeakAngleDerived || a.Adapter.NativeWZMassesDerived || a.Adapter.NativeGaugeCouplingsDerived || a.Adapter.NativeVEVDerived || a.Adapter.NativeKappaPromoted || a.Adapter.NativeYukawaTraceDerived {
		t.Fatalf("adapter over-promoted: %+v", a.Adapter)
	}
	if a.Firewall.ObservedVEVImported || a.Firewall.ObservedGaugeCouplingsImported || a.Firewall.ObservedWeakAngleImported || a.Firewall.ObservedWMassImported || a.Firewall.ObservedZMassImported || a.Firewall.ObservedYukawaImported || a.Firewall.NativeVEVWritten || a.Firewall.NativeGaugeCouplingWritten || a.Firewall.NativeWeakAngleWritten || a.Firewall.NativeWZMassWritten || a.Firewall.NativeKappaWritten || a.Firewall.SyntheticOutputWrittenNative {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 506 {
		t.Fatalf("expected Gate506 redirect, got %+v", a.Next)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 505 Registry Audit", "m_W = g2 v / 2 = 3", "sin^2(theta_W)", StatusSyntheticAdapterExecuted, StatusFailedSyntheticNotNativePrediction, "Gate 506"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate505TheoremPasses(t *testing.T) {
	res := Generation2SyntheticElectroweakMatchingAdapterDryRunTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
