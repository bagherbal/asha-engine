package vacuumcriticalityradiative

import (
	"math"
	"testing"
)

func TestCriticalityTargetComputed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.Criticality
	if !c.Formalized || c.CriticalYukawa <= 0 || c.CriticalYukawaSquared <= 0 || !c.RequiresSaturationAxiom || c.ReductionProved {
		t.Fatalf("bad criticality audit: %s", FormatCriticality(c))
	}
	if math.Abs(c.CriticalYukawa-0.36532653633481293) > 1e-12 {
		t.Fatalf("unexpected critical top Yukawa: %s", FormatCriticality(c))
	}
}

func TestNativeBoundaryNotTangency(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Criticality.NativeLambdaBoundary <= 0 || a.Criticality.NativeBoundaryHasBetaZero {
		t.Fatalf("native boundary should not provide beta-zero tangency: %s", FormatCriticality(a.Criticality))
	}
}

func TestRadiativeZeroesStayZero(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Radiative
	if !r.Formalized || !r.ZeroYukawaIsFixedPoint || r.GaugeLoopsGenerateYukawas || !r.RequiresFlavorBreakingOperator || r.LightMassesGenerated || r.ReductionProved {
		t.Fatalf("bad radiative audit: %s", FormatRadiative(r))
	}
}

func TestMatrixInvariantProgramOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Invariants.Identified || len(a.Invariants.CandidateInvariants) < 4 || a.Invariants.PromotedThisGate {
		t.Fatalf("bad invariant program: %s", FormatInvariants(a.Invariants))
	}
}

func TestCensusStillFifteen(t *testing.T) {
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
	required := []string{
		StatusVacuumCriticalitySieveExecuted,
		StatusCriticalityEquationFormalized,
		StatusCriticalTopYukawaComputed,
		StatusRadiativeHierarchyAnsatzFormalized,
		StatusSMYukawaZeroFixedPointAudited,
		StatusFailedCriticalityNotDerived,
		StatusFailedNativeLambdaNoBetaTangency,
		StatusFailedRadiativeMassesNotGenerated,
		StatusFailedSevenVacuumCoordinatesNotProved,
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
	res := VacuumCriticalityRadiativeHierarchySieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
