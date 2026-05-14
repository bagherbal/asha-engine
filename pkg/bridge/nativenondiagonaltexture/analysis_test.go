package nativenondiagonaltexture

import "testing"

func TestRotationSearchCandidates(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Search.Formalized || len(a.Search.Candidates) != 3 {
		t.Fatalf("bad search: %s", FormatSearch(a.Search))
	}
	dft := a.Search.Candidates[0]
	if dft.Name != "normalized DFT3 flavor rotation" || !dft.Unitary || !dft.OffDiagonal || !dft.VisibleSignInterference {
		t.Fatalf("bad DFT candidate: %s", FormatCandidate(dft))
	}
	cyc := a.Search.Candidates[1]
	if cyc.Name != "Z3 cyclic permutation" || !cyc.Unitary || cyc.OffDiagonal || cyc.HierarchyBroken {
		t.Fatalf("bad cyclic candidate: %s", FormatCandidate(cyc))
	}
}

func TestUnitaryRotationsPreserveSpectrum(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Invariance.Proved || a.Invariance.SeedSpectrum != [3]float64{2, 2, 1} {
		t.Fatalf("bad invariance proof: %s", FormatInvariance(a.Invariance))
	}
	for _, c := range a.Search.Candidates {
		if c.Unitary && c.SingularValues != [3]float64{2, 2, 1} {
			t.Fatalf("unitary candidate changed singular values: %s", FormatCandidate(c))
		}
		if c.FirstSecondSplit != 0 {
			t.Fatalf("unexpected first/second split: %s", FormatCandidate(c))
		}
	}
}

func TestNoHierarchyReductionProved(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Search.AnyHierarchyBroken || a.Search.BestHighLowRatio != 2 {
		t.Fatalf("hierarchy should not be broken: %s", FormatSearch(a.Search))
	}
	if !a.Requirement.NeedsNonUnitaryProjector || !a.Requirement.NeedsAdditionalOperator {
		t.Fatalf("missing texture requirement: %s", FormatRequirement(a.Requirement))
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
	required := []string{
		StatusRotationSearchFormalized,
		StatusDFTAudited,
		StatusInterferenceComputed,
		StatusSingularInvarianceProved,
		StatusTensionUnitaryRotationsInsufficient,
		StatusFailedNativeTextureNotDerived,
		StatusFailedHierarchyNotBroken,
		StatusFailedFirstSecondNotSplit,
		StatusFailedNoReduction,
	}
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
	res := NativeNonDiagonalTextureFlavorOrientationSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
