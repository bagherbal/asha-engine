package generation2osreflectionpositivityairlock

import (
	"strings"
	"testing"
)

func TestGate533Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate532AdapterExecuted || !a.Inheritance.Gate532ThetaResidualsZero || !a.Inheritance.Gate532KreinAdjointResidualZero || !a.Inheritance.Gate532GThetaPositiveDefinite || !a.Inheritance.Gate532ProjectorCompatible || !a.Inheritance.Gate532TimeReflectionInvolution || !a.Inheritance.Gate532FinitePlumbingVerified || !a.Inheritance.Gate532PhysicalHilbertBlocked || !a.Inheritance.Gate532WickBlocked || !a.Inheritance.Gate532OSBlocked || !a.Inheritance.Gate532PositiveEnergyBlocked || !a.Inheritance.Gate532UnitaryBlocked || !a.Inheritance.Gate532GlobalCausalBlocked || !a.Inheritance.Gate532ArrowBlocked || !a.Inheritance.Gate532NativeWriteBlocked || !a.Inheritance.Gate532NoObservedDataImported || !a.Inheritance.Gate533OSAirlockRedirect {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if a.Schema.RequiredRowCount < 19 || !a.Schema.EuclideanReflectionOperatorRequired || !a.Schema.TestFunctionDomainRequired || !a.Schema.ReflectionActionRequired || !a.Schema.CorrelationKernelRequired || !a.Schema.KernelHermiticityCheckRequired || !a.Schema.ReflectionPositiveConeRequired || !a.Schema.OSQuadraticFormCheckRequired || !a.Schema.NullSpaceQuotientRequired || !a.Schema.ReconstructionMapRequired || !a.Schema.CompatibilityWithThetaRequired || !a.Schema.WickMapReferenceRequired || !a.Schema.IepsilonConventionRequired || !a.Schema.SourceRequired || !a.Schema.ConventionRequired || !a.Schema.BridgeOnlyRequired || !a.Schema.ComparatorOnlyRequired || !a.Schema.NoTheoremInputRequired || !a.Schema.NativePromotionRejected || !a.Schema.RedactedSchemaAccepted {
		t.Fatalf("bad schema: %+v", a.Schema)
	}
	if a.Guard.ComparatorExecutionPerformed || a.Guard.ReflectionOperatorEvaluated || a.Guard.TestFunctionDomainEvaluated || a.Guard.ReflectionActionEvaluated || a.Guard.KernelHermiticityEvaluated || a.Guard.OSQuadraticFormEvaluated || a.Guard.PositiveConeEvaluated || a.Guard.NullSpaceQuotientEvaluated || a.Guard.ReconstructionPerformed || a.Guard.CompatibilityWithThetaEvaluated || a.Guard.WickContinuationEvaluated || a.Guard.PositiveEnergyEvaluated || a.Guard.UnitaryDynamicsEvaluated || a.Guard.GlobalHyperbolicityEvaluated || a.Guard.ReflectionPositivityProven || a.Guard.WickRotationSelected || a.Guard.PhysicalHilbertSpaceSelected || a.Guard.PositiveEnergyHamiltonianDerived || a.Guard.UnitaryRealTimeDynamicsDerived || a.Guard.GlobalHyperbolicitySelected {
		t.Fatalf("guard should block execution/promotion: %+v", a.Guard)
	}
	if a.Firewall.ObservedOSDataImported || a.Firewall.ObservedWickDataImported || a.Firewall.ObservedCorrelationDataImported || a.Firewall.ObservedHamiltonianDataImported || a.Firewall.NativeOSKernelWrite || a.Firewall.NativeReflectionWrite || a.Firewall.NativeCorrelationWrite || a.Firewall.NativeHilbertProductWrite || a.Firewall.NativePhysicalStateSpaceWrite || a.Firewall.NativeWickWrite || a.Firewall.NativePositiveEnergyWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 534 {
		t.Fatalf("unexpected next gate: %+v", a.Next)
	}
}

func TestGate533Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 533 Registry Audit", StatusOSKernelAirlockDefined, StatusFailedPositiveMatrixNotOS, StatusFailedComparatorNotPerformed, StatusFirewallNativeWriteBlocked, "Osterwalder", "Gate 534"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate533Theorem(t *testing.T) {
	result := Generation2OSReflectionPositivityKernelAirlockPreflightTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
