package generation2lorentzianspinoradjointairlock

import (
	"strings"
	"testing"
)

func TestGate527Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate526SignatureInherited || !a.Inheritance.Gate526NullConeConfirmed || !a.Inheritance.Gate526EuclideanLedgerSeparated || !a.Inheritance.Gate526WickBlocked || !a.Inheritance.Gate526ReflectionPositivityOpen || !a.Inheritance.Gate526PositiveEnergyOpen || !a.Inheritance.Gate526UnitaryDynamicsOpen || !a.Inheritance.Gate5263Plus1Open || !a.Inheritance.Gate526NoObservedDataImported || !a.Inheritance.Gate526NativeWriteBlocked || a.Inheritance.Gate526ReopenedSealedFirewalls {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Adjoint.IndefiniteMetricSocket || !a.Adjoint.KreinAdjointDefined || !a.Adjoint.DiracAdjointSocketDefined || !a.Adjoint.CliffordCompatibility || !a.Adjoint.ChargeConjugationSocket || !a.Adjoint.GradingSocketPreserved || a.Adjoint.PositiveHilbertProductSelected || a.Adjoint.FundamentalSymmetrySelected || a.Adjoint.PhysicalStateSpaceSelected {
		t.Fatalf("bad adjoint audit: %+v", a.Adjoint)
	}
	if !a.Reflection.EuclideanLedgerAvailable || !a.Reflection.TimeReflectionRequired || a.Reflection.TimeReflectionSelected || a.Reflection.ReflectionPositivityProven || a.Reflection.OsterwalderSchraderAxiomsProven || a.Reflection.WickContinuationSelected || a.Reflection.PositiveEnergyHamiltonianDerived || a.Reflection.UnitaryRealTimeDynamicsDerived || a.Reflection.GlobalHyperbolicitySelected {
		t.Fatalf("bad reflection audit: %+v", a.Reflection)
	}
	if a.Projection.NativeDimension != 8 || a.Projection.CandidateExternalDimension != 4 || a.Projection.CandidateInternalComplement != 4 || !a.Projection.ProjectionRankArithmeticValid || a.Projection.ProjectionOperatorNativeSelected || a.Projection.SubalgebraEmbeddingNativeSelected || a.Projection.InternalComplementNativeSelected || a.Projection.Physical3Plus1Selected || a.Projection.TimeOrientationSelected {
		t.Fatalf("bad projection audit: %+v", a.Projection)
	}
	if a.Firewall.ObservedConstantsImported || a.Firewall.ObservedMassesImported || a.Firewall.ObservedTopologyImported || a.Firewall.ObservedBoundaryImported || a.Firewall.NativePositiveHilbertWrite || a.Firewall.NativeReflectionWrite || a.Firewall.NativeWickWrite || a.Firewall.NativePositiveEnergyWrite || a.Firewall.NativeUnitaryWrite || a.Firewall.Native3Plus1Write || a.Firewall.NativeInternal4Write || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
}

func TestGate527Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 527 Registry Audit", StatusKreinAdjointSocketDefined, StatusFailedNoReflectionPositivity, StatusFailedNoNative3Plus1Projection, StatusFirewallLorentzianDynamicsWriteBlocked, "Gate 528"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate527Theorem(t *testing.T) {
	result := Generation2LorentzianSpinorAdjointReflectionPositivityAirlockTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
