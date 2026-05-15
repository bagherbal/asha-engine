package generation2a4curvaturesquaredledger

import (
	"strings"
	"testing"
)

func TestGate511Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate510A2AuditInherited || !a.Inheritance.Gate510NewtonNormalizationBlocked || !a.Inheritance.ProductA4ChannelsDeclared {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Basis.TopologicalCounterterm || !a.Basis.DynamicalCurvatureSocket || a.Basis.UniqueMetricDynamics || a.Basis.BasisRank != 3 {
		t.Fatalf("bad basis: %+v", a.Basis)
	}
	if !a.A4.DimensionlessChannel || a.A4.UsesF2LambdaSquared || a.A4.UsesF4LambdaFourth || a.A4.NewtonConstantDerived || a.A4.PhysicalGravityCouplingDerived {
		t.Fatalf("bad a4 channel: %+v", a.A4)
	}
	if !nearly(a.A4.F0Moment, 7, 1e-12) || !nearly(a.A4.FiniteTraceDimension, 96, 1e-12) {
		t.Fatalf("bad finite data: %+v", a.A4)
	}
	if !a.Topological.IntegralTopologicalInFourD || !a.Topological.TopologicalSocketNative || a.Topological.EulerCharacteristicNumeric || a.Topological.TopologicalCoefficientPhysical {
		t.Fatalf("bad topological counterterm: %+v", a.Topological)
	}
	if !a.Dynamical.WeylSquaredSocketPresent || a.Dynamical.RenormalizationSchemeSelected || a.Dynamical.BoundaryConditionsSelected || a.Dynamical.LowEnergyEinsteinLimitDerived || a.Dynamical.PhysicalA4DynamicsClosed {
		t.Fatalf("bad dynamics firewall: %+v", a.Dynamical)
	}
	if a.Firewall.NewtonConstantImported || a.Firewall.NewtonConstantDerived || a.Firewall.CosmologicalConstantDerived || a.Firewall.PhysicalA4DynamicsWritten || a.Firewall.NativeGravityNormalizationWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 512 {
		t.Fatalf("unexpected next gate: %+v", a.Next)
	}
}

func TestGate511Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{
		"# Gate 511 Registry Audit",
		StatusA4CurvatureSquaredSocketDefined,
		StatusGaussBonnetTopologicalCountertermFound,
		StatusWeylSquaredDynamicalSocketFound,
		StatusFailedCosmologicalF4StillUnsolved,
		StatusFirewallA4NativeDynamicsBlocked,
		"E₄ = Riem² - 4 Ric² + R²",
		"Gate 512",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate511Theorem(t *testing.T) {
	result := Generation2A4CurvatureSquaredTopologicalCountertermAuditTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
