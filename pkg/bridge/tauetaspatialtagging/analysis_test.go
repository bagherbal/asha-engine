package tauetaspatialtagging

import "testing"

func TestBuildDefaultTauEtaSpatialTagging(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.TauEta.StableNativeDegrees || !a.TauEta.TwoPlusOneMagnitudeSelector || !a.TauEta.OnePlusOnePlusOneSignedSpectrum {
		t.Fatalf("tau eta audit wrong: %s", FormatTauEta(a.TauEta))
	}
	if !a.Spatial.WeakPlaneConditionallySeen || a.Spatial.WeakPlaneDerived || a.Spatial.NativePullbackDerived {
		t.Fatalf("spatial audit should be conditional only: %s", FormatSpatial(a.Spatial))
	}
	if a.Spatial.UniqueAxisIfMapped != "a†_3" || a.Spatial.ComplementPlaneIfMapped != "U={a†_1,a†_2}" {
		t.Fatalf("unexpected conditional weak plane: %s", FormatSpatial(a.Spatial))
	}
	if !a.Generation.CapacitySupported || a.Generation.TextureDerived || a.Generation.TauEtaToGenerationPullback {
		t.Fatalf("generation audit should support capacity only: %s", FormatGeneration(a.Generation))
	}
	if a.Summary.WeakPlaneDerived || a.Summary.GlobalHDerived || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall/summary leaked: %s :: %s", FormatSummary(a.Summary), FormatFirewall(a.Firewall))
	}
}

func TestPlaneSelectionIsConditionalOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	selectedIfMapped := 0
	selectedNatively := 0
	for _, p := range a.Planes {
		if p.SelectedIfTauMapped {
			selectedIfMapped++
		}
		if p.SelectedNatively {
			selectedNatively++
		}
	}
	if selectedIfMapped != 1 || selectedNatively != 0 {
		t.Fatalf("expected one conditional plane and zero native planes, conditional=%d native=%d: %s", selectedIfMapped, selectedNatively, FormatPlanes(a.Planes))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := TauEtaSpatialTaggingGenerationBreakingAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
