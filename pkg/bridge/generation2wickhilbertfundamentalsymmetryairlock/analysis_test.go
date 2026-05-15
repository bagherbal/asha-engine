package generation2wickhilbertfundamentalsymmetryairlock

import (
	"strings"
	"testing"
)

func TestGate531Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate530AdapterExecuted || !a.Inheritance.Gate530ProjectorResidualsZero || !a.Inheritance.Gate530Rank44Confirmed || !a.Inheritance.Gate530ExternalSignature13 || !a.Inheritance.Gate530WickBlocked || !a.Inheritance.Gate530HilbertBlocked || !a.Inheritance.Gate530UnitaryBlocked || !a.Inheritance.Gate530InternalGaugeBlocked || !a.Inheritance.Gate530NoObservedDimensionData || !a.Inheritance.Gate530NativeWriteBlocked || a.Inheritance.Gate530ReopenedSealedFirewalls || !a.Inheritance.Gate531FundamentalAirlockRedirect {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if a.Schema.RequiredRowCount < 15 || !a.Schema.KreinMetricMatrixRequired || !a.Schema.FundamentalSymmetryMatrixRequired || !a.Schema.ThetaInvolutionCheckRequired || !a.Schema.ThetaKreinSelfAdjointCheckRequired || !a.Schema.PositiveHilbertFormCheckRequired || !a.Schema.ProjectorCompatibilityCheckRequired || !a.Schema.TimeReflectionOperatorRequired || !a.Schema.WickMapRequired || !a.Schema.IepsilonPrescriptionRequired || !a.Schema.ReflectionPositivityProofRequired || !a.Schema.PositiveEnergySpectrumRequired || !a.Schema.GlobalHyperbolicityDataRequired || !a.Schema.SourceRequired || !a.Schema.ConventionRequired || !a.Schema.BridgeOnlyRequired || !a.Schema.NoTheoremInputRequired || !a.Schema.NativePromotionRejected || !a.Schema.RedactedSchemaAccepted {
		t.Fatalf("bad schema: %+v", a.Schema)
	}
	if a.Guard.ComparatorExecutionPerformed || a.Guard.ThetaSquaredIdentityEvaluated || a.Guard.ThetaKreinSelfAdjointEvaluated || a.Guard.HilbertFormPositiveEvaluated || a.Guard.ProjectorCommutationEvaluated || a.Guard.TimeReflectionEvaluated || a.Guard.WickContinuationEvaluated || a.Guard.ReflectionPositivityEvaluated || a.Guard.PositiveEnergyEvaluated || a.Guard.UnitaryDynamicsEvaluated || a.Guard.GlobalHyperbolicityEvaluated || a.Guard.PositiveHilbertProductGranted || a.Guard.PhysicalStateSpaceSelected || a.Guard.WickRotationSelected || a.Guard.ReflectionPositivityProven || a.Guard.PositiveEnergyHamiltonianDerived || a.Guard.UnitaryRealTimeDynamicsDerived || a.Guard.GlobalHyperbolicitySelected {
		t.Fatalf("guard should block execution/promotion: %+v", a.Guard)
	}
	if a.Firewall.ObservedHilbertDataImported || a.Firewall.ObservedWickDataImported || a.Firewall.ObservedBoundaryDataImported || a.Firewall.ObservedHamiltonianDataImported || a.Firewall.NativeFundamentalSymmetryWrite || a.Firewall.NativeHilbertProductWrite || a.Firewall.NativePhysicalStateSpaceWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeReflectionWrite || a.Firewall.NativePositiveEnergyWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.Native3Plus1UpgradeWrite || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 532 {
		t.Fatalf("unexpected next gate: %+v", a.Next)
	}
}

func TestGate531Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 531 Registry Audit", StatusFundamentalSymmetryAirlockDefined, StatusFailedThetaDoesNotGrantHilbert, StatusFailedComparatorNotPerformed, StatusFirewallNativeWriteBlocked, "Θ", "Gate 532"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate531Theorem(t *testing.T) {
	result := Generation2WickHilbertFundamentalSymmetryAirlockPreflightTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
