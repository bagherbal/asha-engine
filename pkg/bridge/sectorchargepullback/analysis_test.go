package sectorchargepullback

import "testing"

func TestWeakIsospinPullbackNotDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Pullback.Formalized || a.Pullback.NativeGeneratorSwap {
		t.Fatalf("bad pullback: %s", FormatPullback(a.Pullback))
	}
}

func TestCandidateCKMOverlapsComputedButNotDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CKM.Executed || len(a.CKM.Candidates) < 3 || !a.CKM.AnyCKMCapacity {
		t.Fatalf("bad ckm sieve: %s", FormatCKM(a.CKM))
	}
	if a.CKM.NativeAssignment || a.CKM.NativeCKMDerived {
		t.Fatalf("ckm should remain unpromoted: %s", FormatCKM(a.CKM))
	}
}

func TestColorTraceNormNotDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Color.Audited || a.Color.PullsGlobalAmplifier {
		t.Fatalf("bad color trace norm: %s", FormatColor(a.Color))
	}
}

func TestParameterCensusNoReduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Census.TotalReduction != 0 || a.Census.RemainingInputs != 15 || a.Census.SevenSealReached {
		t.Fatalf("bad census: %s", FormatCensus(a.Census))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusWeakIsospinPullbackFormalized, StatusCKMOverlapExtracted, StatusColorTraceNormAudited, StatusTensionT3DoesNotSelectGenerator, StatusTensionColorDoesNotImplyTraceNorm, StatusFailedSectorChargePullbackDerived, StatusFailedCKMTextureDerived, StatusFailedVacuumReduced}
	for _, req := range required {
		found := false
		for _, got := range statuses {
			if got == req {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing status %s in %v", req, statuses)
		}
	}
}

func TestTheoremPasses(t *testing.T) {
	res := SectorChargePullbackCKMMoritaMisalignmentSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
