package generation2cosmologicalf4vacuumairlock

import (
	"math"
	"strings"
	"testing"
)

func TestGate512Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate511Inherited || !a.Inheritance.ProductA0ChannelDeclared || !a.Inheritance.ProductA0Computed || a.Inheritance.ProductA0PhysicalPrediction {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.A0.NativeDimensionlessTraceWeight || !a.A0.UsesF4LambdaFourth || a.A0.UsesF2LambdaSquared || a.A0.UsesF0Moment || a.A0.PhysicalCosmologicalConstant {
		t.Fatalf("bad a0 audit: %+v", a.A0)
	}
	if math.Abs(a.A0.PrefactorPerF4Lambda4-6.0/(math.Pi*math.Pi)) > 1e-12 {
		t.Fatalf("bad a0 prefactor: %.18g", a.A0.PrefactorPerF4Lambda4)
	}
	if !a.Cancellation.RawTracePositive || a.Cancellation.NativeZeroCancellationFound || a.Cancellation.VacuumEnergyCancelled || a.Cancellation.SupersymmetricPairingPresent {
		t.Fatalf("bad cancellation audit: %+v", a.Cancellation)
	}
	if a.Airlock.F4MomentSelected || a.Airlock.CutoffLambdaSelected || a.Airlock.VacuumSubtractionSelected || a.Airlock.PhysicalLambdaCosmoDerived || a.Airlock.NativeCosmologicalWriteAllowed {
		t.Fatalf("bad airlock: %+v", a.Airlock)
	}
	if a.Firewall.CosmologicalConstantImported || a.Firewall.DarkEnergyDensityImported || a.Firewall.NativeCosmologicalConstantWritten || a.Firewall.VacuumSubtractionWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 513 {
		t.Fatalf("unexpected next gate: %+v", a.Next)
	}
}

func TestGate512Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{
		"# Gate 512 Registry Audit",
		StatusA0VolumePrefactorComputed,
		StatusFailedCosmologicalConstantNotDerived,
		StatusFailedFiniteTraceDoesNotCancelVolume,
		StatusFirewallCosmologicalNativeWriteBlocked,
		"6/π²",
		"Gate 513",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate512Theorem(t *testing.T) {
	result := Generation2CosmologicalF4VacuumEnergySubtractionAirlockAuditTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
