package generation2gravitationalindexetaairlock

import (
	"strings"
	"testing"
)

func TestGate517Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate516Inherited || !a.Inheritance.Gate516EulerSocket || !a.Inheritance.Gate516PontryaginSocket || !a.Inheritance.Gate516CharacteristicScaleFree || !a.Inheritance.Gate516ChiralIndexSocket || !a.Inheritance.Gate516MixedGaugeGravityTraceZero || !a.Inheritance.Gate516GlobalIntegersBlocked || !a.Inheritance.Gate516EtaBlocked || a.Inheritance.Gate516ObservedTopologyImported {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Index.LocalIndexDensitySocketPresent || !a.Index.ClosedManifoldSocketConsistent || a.Index.GlobalIndexIntegerDerived || a.Index.BoundaryEtaDerived || a.Index.BoundaryKernelDimensionDerived || a.Index.BoundarySpectrumSelected || a.Index.ClosedManifoldSelected {
		t.Fatalf("bad index ledger: %+v", a.Index)
	}
	if !a.Eta.BoundaryOperatorRequired || !a.Eta.BoundarySpectrumRequired || !a.Eta.EtaInvariantRequired || !a.Eta.KernelCorrectionRequired || !a.Eta.BoundaryConditionRequired || !a.Eta.GlobalTopologyRequired || a.Eta.BoundaryDataImported || a.Eta.BoundaryEtaNativeDerived || a.Eta.BoundaryEtaNativeWrite || !a.Eta.ClosedManifoldIsAllowedBridge || a.Eta.ClosedManifoldIsNativeSelected {
		t.Fatalf("bad eta airlock: %+v", a.Eta)
	}
	if !a.Inflow.PontryaginDescentSocketPresent || !a.Inflow.ChernSimonsBoundarySocketPresent || !a.Inflow.ChiralIndexAnomalySocketPresent || !a.Inflow.MixedGaugeGravityTraceZero || !a.Inflow.BoundaryEtaPairsWithInflow || a.Inflow.PhysicalThetaCoefficientDerived || a.Inflow.BoundaryTheorySelected {
		t.Fatalf("bad inflow audit: %+v", a.Inflow)
	}
	if a.Firewall.UsesLambdaCutoff || a.Firewall.UsesF2Moment || a.Firewall.UsesF4Moment || a.Firewall.UsesNewtonConstant || a.Firewall.UsesCosmologicalConstant || a.Firewall.UsesHiggsOrElectroweakScale || a.Firewall.UsesFlavorYukawaData || a.Firewall.ObservedTopologyImported || a.Firewall.ObservedBoundarySpectrumImported || a.Firewall.GlobalIndexIntegerNativeWrite || a.Firewall.BoundaryEtaNativeWrite || a.Firewall.PhysicalGravitationalThetaWritten || a.Firewall.NativeGravityNormalizationWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 518 {
		t.Fatalf("unexpected next gate: %+v", a.Next)
	}
}

func TestGate517Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 517 Registry Audit", StatusAPSIndexLedgerDefined, StatusFailedBoundaryEtaNotDerived, StatusFirewallIndexEtaNativeWriteBlocked, "ind_APS", "Gate 518"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate517Theorem(t *testing.T) {
	result := Generation2GravitationalIndexBoundaryEtaAirlockTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
