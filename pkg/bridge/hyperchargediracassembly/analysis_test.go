package hyperchargediracassembly

import "testing"

func TestGate296HyperchargeRayRecoveredButNormalizationOpen(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Hypercharge.RayRecovered {
		t.Fatalf("expected hypercharge ray recovered: %+v", a.Hypercharge)
	}
	if a.Hypercharge.AbsoluteNormalizationFixed {
		t.Fatalf("absolute normalization should remain open: %+v", a.Hypercharge)
	}
	if a.Hypercharge.NormalizedWithQOneSixth.Q != 1.0/6.0 || a.Hypercharge.NormalizedWithQOneSixth.U != 2.0/3.0 || a.Hypercharge.NormalizedWithQOneSixth.E != -1.0 {
		t.Fatalf("expected SM normalized ray as conditional ledger: %+v", a.Hypercharge.NormalizedWithQOneSixth)
	}
}

func TestGate296DiracEdgeGraphAssembled(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Dirac.AllowedEdges) != 4 {
		t.Fatalf("expected four Dirac edge classes: %+v", a.Dirac.AllowedEdges)
	}
	for _, e := range a.Dirac.AllowedEdges {
		if !e.StructurallyLegal || !e.SharedRightModule {
			t.Fatalf("allowed edge should be legal/shared-right: %+v", e)
		}
	}
	if a.Dirac.NumericalYukawas || a.Dirac.BGapActivated || a.Dirac.IncludesMajorana {
		t.Fatalf("D_F amplitudes/B-gap should not be activated: %+v", a.Dirac)
	}
}

func TestGate296FirstOrderPreflightForbidsIllegalEdges(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FirstOrder.ColorIntertwinerVerified || a.FirstOrder.ColorIdentityResidual > 1e-12 || a.FirstOrder.ColorChangingResidual <= 0.1 {
		t.Fatalf("color intertwiner sieve failed: %+v", a.FirstOrder)
	}
	if !a.FirstOrder.LeptonQuarkForbiddenByModule || !a.FirstOrder.ChargeViolatingEdgesForbidden {
		t.Fatalf("illegal edge sieve failed: %+v", a.FirstOrder)
	}
	if a.FirstOrder.FullFirstOrderVerified {
		t.Fatalf("full first-order should remain unverified: %+v", a.FirstOrder)
	}
}

func TestGate296Firewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Firewalls.FiniteCorePolluted || !a.Firewalls.DoesNotInventHyperchargeNormalization || !a.Firewalls.DoesNotInventYukawaMatrices || !a.Firewalls.DoesNotActivateBGapMajorana || !a.Firewalls.DoesNotClaimFullSpectralTriple || !a.Firewalls.DoesNotUnlockHiggs || !a.Firewalls.DoesNotUnlockBGap {
		t.Fatalf("firewall failure: %+v", a.Firewalls)
	}
}
