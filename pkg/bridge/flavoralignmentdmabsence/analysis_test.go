package flavoralignmentdmabsence

import "testing"

func TestGate224GrantsFlavorSealAndRemovesHeavyDM(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Gate223.RelicDecaySealGranted {
		t.Fatalf("Gate223 relic seal should be inherited: %+v", a.Gate223)
	}
	if a.Flavor.GenericFlavorSafe || a.Flavor.ArbitraryFirstSecondAllowed {
		t.Fatalf("generic flavor should not be marked safe: %+v", a.Flavor)
	}
	if !a.Seal.SealGranted || !a.Seal.StillPhenomenological || a.Seal.NativeFlavorTheoremDerived {
		t.Fatalf("expected phenomenological FlavorAlignmentSeal: %+v", a.Seal)
	}
	if !a.DarkMatter.TripletDecaysBeforeBBN || !a.DarkMatter.OctetDecaysBeforeBBN {
		t.Fatalf("expected both carriers to decay before BBN under seals: %+v", a.DarkMatter)
	}
	if a.DarkMatter.OmegaHeavySectorH2 != 0 || a.DarkMatter.PresentDayStableFraction != 0 || a.DarkMatter.HeavySectorDMCandidate {
		t.Fatalf("heavy sector should not be dark matter: %+v", a.DarkMatter)
	}
	if a.Firewall.NativeFlavorClaimed || a.Firewall.ExactFCNCRatesClaimed || a.Firewall.WilsonCoefficientsDerived || a.Firewall.HeavySectorDMClaimed || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall polluted: %+v", a.Firewall)
	}
}
