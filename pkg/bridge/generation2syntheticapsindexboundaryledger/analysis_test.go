package generation2syntheticapsindexboundaryledger

import (
	"strings"
	"testing"
)

func TestGate518Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate517Inherited || !a.Inheritance.Gate517LocalIndexDensitySocket || !a.Inheritance.Gate517APSSocket || !a.Inheritance.Gate517BoundaryEtaAirlock || !a.Inheritance.Gate517AnomalyInflowSocket || !a.Inheritance.Gate517GlobalIndexBlocked || !a.Inheritance.Gate517EtaBlocked || !a.Inheritance.Gate517BoundarySpectrumBlocked || a.Inheritance.Gate517ObservedBoundaryImported {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Ledger.BridgeOnly || !a.Ledger.SyntheticOnly || a.Ledger.UsesObservedTopology || a.Ledger.UsesBoundarySpectrum || !nearly(a.Ledger.BoundaryCorrection, 2, eps) || !nearly(a.Ledger.APSIndex, 9, eps) || !nearly(a.Ledger.ClosedManifoldIndex, 11, eps) || !nearly(a.Ledger.APSResidual, 0, eps) || !nearly(a.Ledger.ClosedResidual, 0, eps) || !a.Ledger.APSIndexIntegerLike || !a.Ledger.ClosedIndexIntegerLike {
		t.Fatalf("bad synthetic ledger: %+v", a.Ledger)
	}
	if !a.Policy.RequiresBridgeOnlyTag || !a.Policy.RequiresSyntheticOrExternalTag || !a.Policy.RequiresSourceMetadata || !a.Policy.RequiresTopologyMetadata || !a.Policy.RequiresBoundaryMetadata || !a.Policy.RejectsNativePromotion || !a.Policy.RejectsObservedByDefault || !a.Policy.RejectsMissingEtaKernelRows || !a.Policy.RejectsMissingBoundaryCondition || !a.Policy.RejectsMissingUncertaintyMetadata || a.Policy.NativeIndexPredictionMade || a.Policy.NativeEtaPredictionMade || a.Policy.BoundaryConditionSelected || a.Policy.BoundarySpectrumDerived || a.Policy.ClosedManifoldNativelySelected {
		t.Fatalf("bad policy: %+v", a.Policy)
	}
	if a.Firewall.UsesLambdaCutoff || a.Firewall.UsesF2Moment || a.Firewall.UsesF4Moment || a.Firewall.UsesNewtonConstant || a.Firewall.UsesCosmologicalConstant || a.Firewall.UsesPlanckScale || a.Firewall.UsesHiggsOrElectroweakScale || a.Firewall.UsesFlavorYukawaData || a.Firewall.ObservedTopologyImported || a.Firewall.ObservedBoundarySpectrumImported || a.Firewall.SyntheticOutputNativeWrite || a.Firewall.GlobalIndexNativePrediction || a.Firewall.BoundaryEtaNativePrediction || a.Firewall.PhysicalGravitationalThetaWritten || a.Firewall.GravityCosmologyNormalizationWrite {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 519 {
		t.Fatalf("unexpected next gate: %+v", a.Next)
	}
}

func TestGate518Markdown(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 518 Registry Audit", StatusSyntheticAPSLedgerExecuted, StatusFirewallSyntheticAPSWritesBlocked, "ind_APS", "11 - (3+1)/2 = 9", "Gate 519"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate518Theorem(t *testing.T) {
	result := Generation2SyntheticAPSIndexBoundaryLedgerDryRunTheorem().Verify()
	for _, c := range result.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
}
