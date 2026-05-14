package scalartrialitytexture

import "testing"

func TestGate246TauEtaHasFlavorCapacityButNoPullback(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error: %v", err)
	}
	if !a.ScalarFlavor.ScalarOriginKnown || !a.ScalarFlavor.NativeHiggsSectorObservable {
		t.Fatalf("expected scalar/Higgs-sector origin: %s", FormatScalarFlavor(a.ScalarFlavor))
	}
	if !a.GenerationTexture.BreaksS3Degeneracy || a.GenerationTexture.DistinctEigenvalues != 3 {
		t.Fatalf("expected 3 distinct tau_eta generation eigenvalues: %s", FormatGenerationTexture(a.GenerationTexture))
	}
	if !a.NonCommutingTexture.RawNonCommutingWithTriality || !a.NonCommutingTexture.PairWouldBeQualifiedIfPullbackHeld {
		t.Fatalf("expected conditional non-commuting texture capacity: %s", FormatNonCommuting(a.NonCommutingTexture))
	}
	if a.NonCommutingTexture.PairActuallyQualified || a.Summary.GenerationTextureDerived || a.PullbackObstruction.PullbackDerived {
		t.Fatalf("Gate 246 must not derive a qualified texture without scalar-to-triality pullback: %s / %s", FormatNonCommuting(a.NonCommutingTexture), FormatPullback(a.PullbackObstruction))
	}
	if a.Firewall.ForcedTauDiagonalTexture || a.Firewall.ImportedYukawaMasses || a.Firewall.ClaimedFermionMasses || a.Firewall.ClaimedFiniteFlavorTheorem {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}
