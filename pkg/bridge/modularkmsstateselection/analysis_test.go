package modularkmsstateselection

import (
	"math"
	"testing"
)

func TestEntropyPrinciple(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Principle.Formalized || a.Principle.EulerLagrange == "" {
		t.Fatalf("bad principle: %s", FormatPrinciple(a.Principle))
	}
}

func TestMaxEntropyIsTracial(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.MaxEntropy.State
	if !s.Faithful || !s.Tracial || a.MaxEntropy.SelectsNontracial {
		t.Fatalf("expected tracial max entropy: %s", FormatEntropyLane(a.MaxEntropy))
	}
	if math.Abs(s.Entropy-math.Log(3)) > 1e-12 {
		t.Fatalf("unexpected entropy: %.15f", s.Entropy)
	}
}

func TestKMSStateIsNontracialButNotPromoted(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.KMS.NonTrivialFlow || !allPairFrequenciesNonZero(a.KMS.State) || a.KMS.PromotedNative || a.KMS.EnergyConstraintDerived {
		t.Fatalf("bad KMS lane: %s", FormatKMS(a.KMS))
	}
}

func TestFlowCapacityNoVacuumSelection(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Flow.KMSFlowNontrivial || !a.Flow.PreservesLandscape || !a.Flow.KineticSafe || a.Flow.SelectsUniqueVacuum {
		t.Fatalf("bad flow: %s", FormatFlow(a.Flow))
	}
}

func TestCensusUnchanged(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Census.RemainingInputs != 15 || a.Census.Reduction != 0 {
		t.Fatalf("bad census: %s", FormatCensus(a.Census))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusEntropyPrincipleFormalized, StatusKMSStateSolved, StatusTauHamiltonianCapacityAudited, StatusFailedKMSSelectionNotNative, StatusFailedEnergyConstraintNotDerived, StatusFailedVacuumPointNotSelected}
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
	res := ModularKMSStateSelectionEntropyVariationalPrincipleAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
