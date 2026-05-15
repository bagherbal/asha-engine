package generation2topologicaldeformation

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate483TopologicalNoGo(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate482SourceAbsent || a.Inheritance.IKVac != 0.5 {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.TopologicalAudit.SectorSeparatorsFound || !a.TopologicalAudit.QuarkLeptonOnly || a.TopologicalAudit.NativeFullSourceFound {
		t.Fatalf("bad topological audit: %+v", a.TopologicalAudit)
	}
	if !a.GenerationAwareness.ColorDistinguishesQuarkLepton || a.GenerationAwareness.ColorDistinguishesGenerations || a.GenerationAwareness.CandidatesPassingGenerationAwareness != 0 {
		t.Fatalf("bad generation awareness: %+v", a.GenerationAwareness)
	}
	if !a.DeformationMap.TopologicalStressNative || a.DeformationMap.DeltaAlphaMapNative || a.DeformationMap.DeltaPhiMapNative || a.DeformationMap.NumericCoordinateMapNative {
		t.Fatalf("bad deformation map: %+v", a.DeformationMap)
	}
	if !a.BridgeSlot.AllowsTopologicalLabels || !a.BridgeSlot.RequiresAirlock || a.BridgeSlot.CanComputePhysicalResidual {
		t.Fatalf("bad bridge slot: %+v", a.BridgeSlot)
	}
	if a.Firewall.TopologicalNativeSourceFound || a.Firewall.SectorPerturbationsNative || a.Firewall.PhysicalDUDComputed || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}

func TestRenderAuditGate483(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusFailedNativeTopologicalSourceAbsent, "SU(3) color representation", "generation-awareness", "topological-sector-perturbation-ledger", "physical d_ud = undefined"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
