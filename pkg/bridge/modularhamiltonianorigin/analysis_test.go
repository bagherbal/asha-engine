package modularhamiltonianorigin

import "testing"

func TestCriteriaFormalized(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Criteria.Formalized || !a.Criteria.MustBeNative || !a.Criteria.MustBeNonCircular {
		t.Fatalf("bad criteria: %s", FormatCriteria(a.Criteria))
	}
}

func TestCandidateLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	identity := candidateByName(a.Candidates, "identity")
	mag := candidateByName(a.Candidates, "tau magnitude")
	tau := candidateByName(a.Candidates, "triality signed tau_eta")
	if !identity.Central || identity.BreaksAllDegeneracy {
		t.Fatalf("identity should freeze flow: %s", FormatCandidate(identity))
	}
	if mag.Central || mag.BreaksAllDegeneracy {
		t.Fatalf("magnitude lane should retain degeneracy: %s", FormatCandidate(mag))
	}
	if tau.Central || !tau.BreaksAllDegeneracy || tau.PromotedNative {
		t.Fatalf("signed tau should have capacity but not promotion: %s", FormatCandidate(tau))
	}
}

func TestEnergyConstraintCircular(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Energy.Formalized || !a.Energy.Circular || a.Energy.ConstraintNative || a.Energy.PromotesKMSNative {
		t.Fatalf("energy constraint should remain circular: %s", FormatEnergy(a.Energy))
	}
}

func TestFlowCapacityNoSelection(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Flow.NontrivialCapacity || !a.Flow.PreservesLandscape || !a.Flow.KineticSafe || a.Flow.SelectsVacuum {
		t.Fatalf("bad flow audit: %s", FormatFlow(a.Flow))
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
	required := []string{StatusHamiltonianOriginFormalized, StatusTrialityHamiltonianAudited, StatusEnergyConstraintAudited, StatusFailedHamiltonianNotDerived, StatusFailedEnergyConstraintNotNative, StatusFailedVacuumNotSelected}
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
	res := ModularHamiltonianOriginTrialityEnergyConstraintDerivationAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
