package yukawaamplitudesource

import "testing"

func TestBuildDefaultYukawaAmplitudeSourceObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Generation.GenerationBlind || !a.Generation.ProjectsToIdentity || !a.Generation.PermutationInvariant {
		t.Fatalf("generation functional should be blind: %s", FormatGeneration(a.Generation))
	}
	if a.Triality.ExactTrialitySelectsTexture || a.Triality.CouplingsDerived || a.Triality.CKMDerived || a.Triality.PMNSDerived {
		t.Fatalf("triality unexpectedly selected texture: %s", FormatTriality(a.Triality))
	}
	if a.Curvature.NonCommutingTexturePairInduced || a.Curvature.T1T2FlavorOffDiagonal {
		t.Fatalf("curvature should not induce flavor texture: %s", FormatCurvature(a.Curvature))
	}
	if !a.SourceSearch.NoCanonicalAmplitudeSource || a.SourceSearch.SelectedAmplitudeSources != 0 {
		t.Fatalf("amplitude source unexpectedly found: %s", FormatSourceSearch(a.SourceSearch))
	}
	if a.Firewall.YukawaAmplitudesDerived || a.Firewall.FermionMassesDerived || a.Firewall.CKMMatrixDerived || a.Firewall.PMNSMatrixDerived || a.Firewall.ObservedMassRatiosImported || a.Firewall.CabibboAngleImported {
		t.Fatalf("firewall leaked: %s", FormatFirewall(a.Firewall))
	}
}

func TestGenerationTraceIdentity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	m := a.Generation.GenerationTraceMatrix
	if m[0][0] != 16 || m[1][1] != 16 || m[2][2] != 16 {
		t.Fatalf("expected native one-generation support 16 on all diagonal entries: %s", FormatGeneration(a.Generation))
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i != j && m[i][j] != 0 {
				t.Fatalf("off-diagonal generation entry selected: %s", FormatGeneration(a.Generation))
			}
		}
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := FiniteYukawaTextureOperatorAmplitudeSourceObstructionTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
