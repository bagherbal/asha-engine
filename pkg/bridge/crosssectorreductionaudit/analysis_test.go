package crosssectorreductionaudit

import "testing"

func TestSeesawDependency(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Seesaw.Formalized || !a.Seesaw.CommonMajoranaScaleCancelsRatios || !a.Seesaw.RequiresDiracSingularValues || a.Seesaw.ReductionProved || a.Seesaw.RatioPredicted {
		t.Fatalf("bad seesaw audit: %s", FormatSeesaw(a.Seesaw))
	}
	if a.Seesaw.ObservedDeltaRatio < 30 || a.Seesaw.ObservedDeltaRatio > 40 {
		t.Fatalf("unexpected neutrino delta ratio: %s", FormatSeesaw(a.Seesaw))
	}
}

func TestStabilityInequality(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Stability.Formalized || !a.Stability.DependsOnTopYukawa || !a.Stability.BoundIsInequality || !a.Stability.RequiresSaturationAxiom || a.Stability.PredictsTopMass || a.Stability.ReductionProved {
		t.Fatalf("bad stability audit: %s", FormatStability(a.Stability))
	}
}

func TestBGapPowerLawRejected(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.PowerLaw.Formalized || len(a.PowerLaw.Data) < 4 || a.PowerLaw.UniversalSimplePowerLawFound || a.PowerLaw.ReductionProved {
		t.Fatalf("bad power law audit: %s", FormatPowerLaw(a.PowerLaw))
	}
}

func TestCensusUnchanged(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Census.StartingVacuumInputs != 15 || a.Census.TotalAdditionalReduction != 0 || a.Census.RemainingVacuumInputs != 15 || a.Census.SevenSealTargetReached {
		t.Fatalf("bad census: %s", FormatCensus(a.Census))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusCrossSectorReductionAuditExecuted, StatusSeesawDependencyFormalized, StatusStabilityBoundFormalized, StatusBGapPowerLawTested, StatusParameterCensusUpdated, StatusFailedNoParameterReductionProved, StatusFailedNeutrinoMassesNotDerived, StatusFailedTopYukawaNotPredicted, StatusFailedMassPowerLawNotDerived, StatusFailedSevenVacuumCoordinatesNotProved}
	for _, req := range required {
		found := false
		for _, s := range statuses {
			if s == req {
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
	res := CrossSectorReductionAuditVacuumParameterCompressionSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
