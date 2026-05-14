package nontracialmodularstate

import (
	"math"
	"testing"
)

func TestTopologicalSourcing(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Topological.Formalized || a.Topological.SignedTauDensityValid || a.Topological.NativeNonTracialFound {
		t.Fatalf("bad topological sourcing: %s", FormatTopological(a.Topological))
	}
}

func TestTauMagnitudeStateResidualDegeneracy(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.Topological.SquaredMagnitudeState
	if !s.Faithful || s.Tracial || !s.ResidualDeg12 {
		t.Fatalf("expected faithful nontracial residual 1-2 degeneracy: %s", FormatState(s))
	}
	if math.Abs(s.Rho[0]-4.0/9.0) > 1e-12 || math.Abs(s.Rho[2]-1.0/9.0) > 1e-12 {
		t.Fatalf("unexpected tau^2 rho: %v", s.Rho)
	}
}

func TestKMSStateActivatesAllPairFrequenciesButIsNotMandated(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.KMS.NonTrivial || a.KMS.Mandated || !allPairFrequenciesNonZero(a.KMS.State) {
		t.Fatalf("KMS state mismatch: %s", FormatKMS(a.KMS))
	}
}

func TestFlowCapacityNoSelection(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Flow.AnyNonTrivial || a.Flow.AnyMandatedNativeNontracial || a.Flow.SelectsUniqueVacuum || a.Flow.RemainingInputs != 15 {
		t.Fatalf("flow should have capacity but no native selection: %s", FormatFlow(a.Flow))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusTopologicalSourcingFormalized, StatusKMSStateFormalized, StatusFlowActivationSieveExecuted, StatusModularTimeCapacityIdentified, StatusFailedNonTracialStateNotDerived, StatusFailedModularTimeNotActivated, StatusFailedFlavorVacuumNotSelected}
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
	res := NontracialModularStateOriginVacuumDensityMatrixDerivationAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
