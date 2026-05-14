package finiteanchordm

import "testing"

func TestBuildDefaultFiniteAnchorDMObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.Gate224.HeavySectorDMAbsent {
		t.Fatalf("expected Gate 224 heavy-sector DM absence inheritance")
	}
	if a.Summary.FiniteAnchorDMViable {
		t.Fatalf("finite anchor DM should not be marked viable without shift/stability/scale")
	}
	if a.ALP.GenericALPStructureSupported || a.ALP.QCDAxionStructureSupported {
		t.Fatalf("ALP support should be obstructed")
	}
	if a.Contact.StrictDarkSectorSupported {
		t.Fatalf("contact modes should not be promoted to strict dark sector")
	}
	if a.Misalign.OmegaComputed {
		t.Fatalf("misalignment relic density should not be computed")
	}
}

func TestLoopScaledBGapNearMissRemainsUnpromoted(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if LoopScaledBGap() <= 0 {
		t.Fatalf("expected positive loop-scaled B-gap diagnostic")
	}
	if a.Firewall.BGapUsedAsAxionScale || a.Firewall.AxionDecayConstantInvented {
		t.Fatalf("B-gap must not be used as an axion scale or f_a")
	}
}
