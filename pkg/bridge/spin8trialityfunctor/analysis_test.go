package spin8trialityfunctor

import "testing"

func TestGate247Spin8TrialityDoesNotBypassScalarTraceDomain(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error: %v", err)
	}
	if !a.Spin8Triality.AbstractSpin8TrialityAvailable || !a.Spin8Triality.VectorToSpinorFunctorKnown {
		t.Fatalf("expected abstract Spin(8) triality preflight: %s", FormatSpin8(a.Spin8Triality))
	}
	if a.Spin8Triality.ScalarTraceTripleIsVector || a.ScalarSpinor.ExteriorOrVectorRepresentativeKnown || a.ScalarSpinor.PullbackFunctorDerived {
		t.Fatalf("tau_eta must remain outside triality action domain: %s / %s", FormatSpin8(a.Spin8Triality), FormatScalarSpinor(a.ScalarSpinor))
	}
	if !a.Texture.RawNonCommutingCapacity || a.Texture.YukawaTextureDerived || a.Texture.DiagonalOperatorConstructed {
		t.Fatalf("expected raw texture capacity only: %s", FormatTexture(a.Texture))
	}
	if a.Firewall.ForcedScalarToSpinorMap || a.Firewall.InsertedDTauAsTexture || a.Firewall.ImportedYukawaMasses || a.Firewall.ClaimedFiniteFlavorTheorem {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}
