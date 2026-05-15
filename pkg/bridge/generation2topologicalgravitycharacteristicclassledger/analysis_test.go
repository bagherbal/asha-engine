package generation2topologicalgravitycharacteristicclassledger

import (
	"math"
	"strings"
	"testing"
)

func TestGate516Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate515Inherited || !a.Inheritance.Gate515SyntheticOnly || !a.Inheritance.Gate515NativeNormalizationBlocked || a.Inheritance.Gate515ObservedDataImported || !a.Inheritance.Gate511GaussBonnetSocket || !a.Inheritance.Gate511DimensionlessA4 || !a.Inheritance.Gate511A4DoesNotUseF2Lambda || !a.Inheritance.Gate490MixedGravityTraceCanceled {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Ledger.EulerSocketPresent || !a.Ledger.PontryaginSocketPresent || !a.Ledger.SignatureSocketPresent || !a.Ledger.A4CharacteristicSubspace || !nearly(a.Ledger.A4UnitPrefactor, 1/(60*math.Pi*math.Pi), 1e-12) || a.Ledger.PhysicalThetaAngleDerived {
		t.Fatalf("bad ledger: %+v", a.Ledger)
	}
	if a.Scale.UsesLambdaCutoff || a.Scale.UsesF2Moment || a.Scale.UsesF4Moment || a.Scale.UsesNewtonConstant || a.Scale.UsesCosmologicalConstant || a.Scale.UsesHiggsVEVOrEWScale || a.Scale.UsesFlavorYukawaData || a.Scale.UsesObservedManifoldData || !a.Scale.CharacteristicIntegralsScaleFree {
		t.Fatalf("bad scale audit: %+v", a.Scale)
	}
	if !a.Finite.ChiralIndexSocketPresent || !a.Finite.MixedGravitationalGaugeTraceZero || a.Finite.ContinuumEulerIntegerDerived || a.Finite.ContinuumSignatureIntegerDerived || a.Finite.ManifoldTopologySelected || a.Finite.BoundaryEtaInvariantClosed {
		t.Fatalf("bad finite signature audit: %+v", a.Finite)
	}
	if a.Firewall.NewtonConstantImported || a.Firewall.PlanckScaleImported || a.Firewall.CutoffLambdaImported || a.Firewall.CosmologicalConstantImported || a.Firewall.ManifoldEulerIntegerImported || a.Firewall.ManifoldSignatureImported || a.Firewall.ObservedTopologyImported || a.Firewall.ManifoldIntegerNativeWrite || a.Firewall.NativeGravityNormalizationWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 517 {
		t.Fatalf("unexpected next gate: %+v", a.Next)
	}
}

func TestGate516Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 516 Registry Audit", StatusNativeGravityTopologyConfirmed, StatusFailedManifoldEulerIntegerNotDerived, StatusFirewallTopologicalIntegerNativeWriteBlock, "Gauss-Bonnet", "Pontryagin", "Gate 517"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate516Theorem(t *testing.T) {
	result := Generation2TopologicalGravityCharacteristicClassLedgerTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
