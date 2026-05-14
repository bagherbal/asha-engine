package su2l

import "testing"

func TestDoubletAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ChargeLevelSU2LDoubletsDerived {
		t.Fatalf("expected charge-level SU(2)_L doublets to be derived")
	}
	if !a.Standard.MatchesStandardOrientation {
		t.Fatalf("expected standard orientation to match Q_L/L_L charges: %+v", a.Standard)
	}
	if !a.Conjugate.MatchesConjugateOrientation {
		t.Fatalf("expected conjugate orientation to match mirror Q_L/L_L charges: %+v", a.Conjugate)
	}
	if a.Standard.QuarkDoubletDim != 6 || a.Standard.LeptonDoubletDim != 2 {
		t.Fatalf("bad standard doublet dimensions: Q=%d L=%d", a.Standard.QuarkDoubletDim, a.Standard.LeptonDoubletDim)
	}
	if !a.NeutralSeedAmbiguity {
		t.Fatalf("expected neutral seed ambiguity to remain exposed")
	}
}
