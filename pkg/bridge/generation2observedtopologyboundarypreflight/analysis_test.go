package generation2observedtopologyboundarypreflight

import (
	"strings"
	"testing"
)

func TestGate519Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate518Inherited || !a.Inheritance.Gate518SyntheticAPSDryRun || !a.Inheritance.Gate518BridgeOnly || !a.Inheritance.Gate518GlobalTopologyBlocked || !a.Inheritance.Gate518BoundaryEtaBlocked || !a.Inheritance.Gate518NativeWriteBlocked || a.Inheritance.ObservedDataImported {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if a.Topology.RequiredRows != 7 || !a.Topology.RequiresEulerCharacteristic || !a.Topology.RequiresPontryaginClasses || !a.Topology.RequiresSignature || !a.Topology.RequiresGlobalAPSIndex || !a.Topology.RequiresManifoldDimension || !a.Topology.RequiresOrientationAndClosedness || !a.Topology.RequiresModelID || !a.Topology.RejectsNativePromotion || !a.Topology.RedactedSchemaAccepted || a.Topology.ObservedNumbersImported || !a.Topology.ComparatorTargetOnly {
		t.Fatalf("bad topology schema: %+v", a.Topology)
	}
	if a.Boundary.RequiredRows != 7 || !a.Boundary.RequiresBoundaryConditionType || !a.Boundary.RequiresEtaInvariantValue || !a.Boundary.RequiresKernelDimensionH || !a.Boundary.RequiresBoundarySpectrumMetadata || !a.Boundary.RequiresBoundaryOrientation || !a.Boundary.RequiresBoundaryComponentCount || !a.Boundary.RequiresModelID || !a.Boundary.RejectsNativePromotion || !a.Boundary.RedactedSchemaAccepted || a.Boundary.ObservedNumbersImported || !a.Boundary.ComparatorTargetOnly {
		t.Fatalf("bad boundary schema: %+v", a.Boundary)
	}
	if !a.Policy.RequiresSource || !a.Policy.RequiresSourceVersion || !a.Policy.RequiresUncertainty || !a.Policy.RequiresScheme || !a.Policy.RequiresScaleOrTopologyContext || !a.Policy.RequiresBridgeOnlyTrue || !a.Policy.RequiresNativePromotionFalse || !a.Policy.RequiresComparatorOnlyPurpose || !a.Policy.RequiresNoTheoremInputFlag || !a.Policy.RejectsMissingSource || !a.Policy.RejectsMissingUncertainty || !a.Policy.RejectsBridgeOnlyFalse || !a.Policy.RejectsNativePromotionTrue || a.Policy.AcceptedRedactedSchemaCases != 1 || a.Policy.RejectedFailClosedCases < 10 {
		t.Fatalf("bad policy: %+v", a.Policy)
	}
	if !a.Rejection.TopologyNativePredictionBlocked || !a.Rejection.BoundaryEtaNativePredictionBlock || !a.Rejection.GlobalAPSIndexNativeWriteBlocked || !a.Rejection.EulerCharacteristicNativeBlocked || !a.Rejection.PontryaginNumberNativeBlocked || !a.Rejection.SignatureNativeBlocked || !a.Rejection.BoundarySpectrumNativeBlocked || !a.Rejection.ClosedManifoldConditionBlocked || !a.Rejection.ComparatorExecutionBlockedNow || !a.Rejection.ResidualComputationBlockedNow {
		t.Fatalf("bad rejection: %+v", a.Rejection)
	}
	if a.Firewall.ObservedTopologyImported || a.Firewall.ObservedBoundaryDataImported || a.Firewall.ObservedBoundarySpectrumImported || a.Firewall.UsesNewtonConstant || a.Firewall.UsesPlanckScale || a.Firewall.UsesLambdaCutoff || a.Firewall.UsesCosmologicalConstant || a.Firewall.UsesElectroweakScale || a.Firewall.UsesFlavorYukawaData || a.Firewall.NativeTopologyWrite || a.Firewall.NativeBoundaryWrite || a.Firewall.NativeGlobalIndexWrite || a.Firewall.ComparatorExecuted {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 520 {
		t.Fatalf("unexpected next gate: %+v", a.Next)
	}
}

func TestGate519Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 519 Registry Audit", StatusTopologyAirlockDefined, StatusBoundaryAirlockDefined, StatusFirewallNativeWriteBlocked, "euler_characteristic", "eta_invariant_value", "bridge_only=true", "Gate 520"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate519Theorem(t *testing.T) {
	result := Generation2ObservedTopologyBoundaryComparatorPreflightTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
