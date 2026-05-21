package generation2neutralpunctureairlockunificationaudit

import (
	"strings"
	"testing"
)

func TestGate895PunctureIsUpstreamOfWeakFrame(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	p := a.PunctureIndependence
	if !p.RequiresRightCharacterSplit || !p.RequiresLeptoColorSplit || p.RequiresWeakSocketFrame || !p.DefinedBeforeWeakOrientation || !p.UpstreamOfAlphaFlagAndWeakKernel {
		t.Fatalf("bad puncture independence audit: %s", FormatPunctureIndependence(p))
	}
	if !containsAll(p.Supports, []string{SupportPunctureDefinedBeforeWeakOrientation, SupportCommonSourceOfAlphaAndOrientation}) {
		t.Fatalf("missing puncture support statuses: %s", FormatPunctureIndependence(p))
	}
}

func TestGate895AlphaFlagReconstructedFromPuncture(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.AlphaFlag
	if !f.F0SubsetF1SubsetF2 || f.RankF0 != 1 || f.RankF1 != 4 || f.RankF2 != 8 || f.RankQ1 != 3 || f.RankQ2 != 7 || !f.ReconstructsAlpha || f.NativeAlphaFunctor {
		t.Fatalf("bad alpha flag audit: %s", FormatAlphaFlag(f))
	}
	if !containsAll(f.Failures, []string{FailurePunctureFlagNotNativeAlphaFunctor, FailureNoNativeBoundaryIncidenceFunctor, FailureAlphaStillSealed}) {
		t.Fatalf("missing alpha flag failures: %s", FormatAlphaFlag(f))
	}
}

func TestGate895WeakKernelReconstructedButNotNative(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	w := a.WeakKernel
	if w.HLeftRank != 8 || w.ImageRank != 7 || w.KernelRank != 1 || w.Kernel != LeftKernel || !w.QuotientIsKernel || !w.CanReconstructWeakFrameCandidate || w.NativeMinimalImageRule || !w.DependsOnMinimalImageChoice {
		t.Fatalf("bad weak kernel audit: %s", FormatWeakKernel(w))
	}
	if !containsAll(w.Failures, []string{FailureNoNativeMinimalImageSelection, FailureWeakReconstructionDependsOnImage, FailureNoNativeWeakSocketSelector}) {
		t.Fatalf("missing weak kernel failures: %s", FormatWeakKernel(w))
	}
}

func TestGate895AirlockCollapsesTwoSealsButNotNatively(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	u := a.Airlock
	if u.CommonObject != Puncture || !u.ControlsAlphaFlag || !u.ControlsWeakKernel || !u.TwoSealProblemReducesToOne || u.NativeAirlockFunctor {
		t.Fatalf("bad airlock audit: %s", FormatAirlock(u))
	}
	if !strings.Contains(u.NextMissingObject, "NeutralPunctureAirlockFunctor") {
		t.Fatalf("wrong next missing object: %s", FormatAirlock(u))
	}
}

func TestGate895FreezeAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("freeze leak: %s", FormatFreeze(a.Freeze))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate895Theorem(t *testing.T) {
	res := Generation2NeutralPunctureAirlockUnificationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
