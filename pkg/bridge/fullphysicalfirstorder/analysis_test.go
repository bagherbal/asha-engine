package fullphysicalfirstorder

import "testing"

func TestGate297ZeroOrderVerified(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ZeroOrder.ZeroOrderVerified || a.ZeroOrder.WeakColorCommutatorNorm > 1e-12 {
		t.Fatalf("zero-order failed: %+v", a.ZeroOrder)
	}
}

func TestGate297FullStructuralFirstOrderVerified(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FirstOrder.FullSweepVerified || a.FirstOrder.MaxLegalResidual > 1e-12 {
		t.Fatalf("first-order legal sweep failed: %+v", a.FirstOrder)
	}
	if a.FirstOrder.MinRejectedResidual <= 0.1 {
		t.Fatalf("rejected edges should have nonzero obstruction: %+v", a.FirstOrder)
	}
	if len(a.FirstOrder.LegalEdges) != 4 || len(a.FirstOrder.RejectedEdges) < 4 {
		t.Fatalf("unexpected edge ledger: %+v", a.FirstOrder)
	}
}

func TestGate297SkeletonNotDynamics(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Triple.StructuralSkeletonComplete {
		t.Fatalf("expected structural skeleton complete: %+v", a.Triple)
	}
	if a.Triple.DynamicalTripleComplete || a.Triple.NumericalYukawas || a.Triple.BGapMajorana || a.Triple.HyperchargeAbsoluteNormalization {
		t.Fatalf("dynamical triple should remain firewalled: %+v", a.Triple)
	}
}

func TestGate297Firewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Firewalls.FiniteCorePolluted || !a.Firewalls.DoesNotInventHyperchargeNormalization || !a.Firewalls.DoesNotInventYukawaMatrices || !a.Firewalls.DoesNotActivateBGapMajorana || !a.Firewalls.DoesNotClaimDynamics || !a.Firewalls.DoesNotUnlockHiggs || !a.Firewalls.DoesNotUnlockBGap {
		t.Fatalf("firewall failure: %+v", a.Firewalls)
	}
}
