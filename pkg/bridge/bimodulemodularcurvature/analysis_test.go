package bimodulemodularcurvature

import "testing"

func TestFrameworkFormalized(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Formalization.Executed || !a.Formalization.NeedsEtaTrace || !a.Formalization.ForbidsManualTauPick {
		t.Fatalf("bad formalization:\n%s", FormatFormalization(a.Formalization))
	}
}

func TestScalarAndSupportLanesRemainCentral(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"A", "B", "C"} {
		lane := laneByID(a.Lanes, id)
		if !lane.Central || lane.BreaksFlavorOrbit || !lane.NativeSource {
			t.Fatalf("lane %s should be native but central:\n%s", id, FormatLane(lane))
		}
	}
}

func TestEtaLaneIsNoncentralButCircular(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	lane := laneByID(a.Lanes, "D")
	if !lane.NonCentral || !lane.BreaksFlavorOrbit || !lane.TauEtaInserted || lane.TauEtaDerived || lane.NativeSource {
		t.Fatalf("eta lane should witness capacity but remain circular:\n%s", FormatLane(lane))
	}
}

func TestKMSCapacityNotPromoted(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.KMS.Executed || !a.KMS.NontrivialFrequencies || a.KMS.PromotedNative || a.KMS.EnergyConstraint {
		t.Fatalf("bad KMS audit:\n%s", FormatKMS(a.KMS))
	}
}

func TestLandscapeAndKineticSafety(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Landscape.WeakMixingPreserved || !a.Landscape.QuarticRatioPreserved || !a.Landscape.AlphaGUTPreserved || !a.Landscape.MoritaSplitPreserved || a.Landscape.FiniteCorePolluted {
		t.Fatalf("bad landscape audit:\n%s", FormatLandscape(a.Landscape))
	}
	if !a.Kinetic.AllCandidatesSelf || !a.Kinetic.FaithfulStateSafe || !a.Kinetic.NoGhostMetric {
		t.Fatalf("bad kinetic audit:\n%s", FormatKinetic(a.Kinetic))
	}
}

func TestCensusUnchanged(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Census.StartingInputs != 15 || a.Census.Reduction != 0 || a.Census.RemainingInputs != 15 || a.Flow.SelectsVacuum || a.Flow.PromotedNative {
		t.Fatalf("bad census/flow:\n%s\n%s", FormatFlow(a.Flow), FormatCensus(a.Census))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusBimoduleCurvatureFormalized, StatusLRCommutantFrameworkAudited, StatusTensionTauInsertionCircular, StatusFailedOriginNotDerived, StatusFailedTauStillNotSelected, StatusFailedCensusNotReduced}
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
	res := BimoduleModularCurvatureInternalThermalTimeOriginSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
