package yukawairfixedpoint

import (
	"math"
	"testing"
)

func TestSpiralQuasiFixedBasinButNoReduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.Spiral
	if !s.Audited || !s.QuasiFixedPoint || s.ContractionRatio >= 0.05 || s.ReductionProved || s.ParameterReduction != 0 {
		t.Fatalf("bad spiral audit: %s", FormatSpiral(s))
	}
	if math.Abs(s.Boundary.RPlusYtUV-0.906917857583) > 1e-6 {
		t.Fatalf("unexpected r+ top boundary: %s", FormatSpiral(s))
	}
}

func TestRPlusEndpointHighTopLane(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Spiral.RPlusEndpointYt < 0.95 || a.Spiral.RPlusEndpointMassGeV < 200 {
		t.Fatalf("r+ should flow to high-top/high-mass lane in this audit: %s", FormatSpiral(a.Spiral))
	}
}

func TestCriticalityNoZeroAtIntermediateScale(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.Criticality
	if !c.Formalized || c.PerturbativeSolution || c.MinLambdaAtTarget <= 0 || c.ReductionProved {
		t.Fatalf("bad criticality audit: %s", FormatCriticality(c))
	}
	if c.MinLambdaAtTarget > 0.13 {
		t.Fatalf("min lambda should remain positive but near native scale: %s", FormatCriticality(c))
	}
}

func TestBaryogenesisNeedsCPAsymmetryOperator(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	b := a.Baryogenesis
	if !b.Formalized || !b.StandardCKMInsufficient || !b.BGapLeptogenesisHasCapacity || b.CPAsymmetryOperatorDerived || b.ReductionProved {
		t.Fatalf("bad baryogenesis audit: %s", FormatBaryogenesis(b))
	}
}

func TestCensusStillFifteen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Census.StartingVacuumInputs != 15 || a.Census.TotalReduction != 0 || a.Census.RemainingInputs != 15 || a.Census.SevenSealReached {
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
		StatusRGAttractorFormalized,
		StatusQuasiFixedPointDetected,
		StatusCriticalityScanned,
		StatusBaryogenesisFormalized,
		StatusFailedDynamicalSelectionNotActive,
		StatusFailedTopYukawaNotReduced,
		StatusFailedCriticalityNoSolution,
		StatusFailedBaryogenesisPhaseNotDerived,
		StatusFailedNoParameterReduction,
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
	res := YukawaInfraredFixedPointBasinRGAttractorReductionAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
