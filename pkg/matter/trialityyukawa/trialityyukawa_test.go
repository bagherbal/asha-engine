package trialityyukawa

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.GenerationCount != 3 {
		t.Fatalf("generation count=%d", a.GenerationCount)
	}
	if a.OneGenerationChannels != 8 {
		t.Fatalf("one-generation channels=%d", a.OneGenerationChannels)
	}
	if a.DiagonalChannelCount != 24 {
		t.Fatalf("diagonal channels=%d", a.DiagonalChannelCount)
	}
	if a.FullMixingMapCount != 72 {
		t.Fatalf("full mixing maps=%d", a.FullMixingMapCount)
	}
	if a.DiagonalFiberEntries != 48 || a.FullMixingFiberEntries != 144 {
		t.Fatalf("fiber entries diagonal=%d full=%d", a.DiagonalFiberEntries, a.FullMixingFiberEntries)
	}
	if a.TextureSelectedByFiniteData || a.CouplingsDerived || a.CKMDerived || a.PMNSDerived {
		t.Fatal("Gate 26 must not claim texture/coupling/mixing derivation")
	}
}
