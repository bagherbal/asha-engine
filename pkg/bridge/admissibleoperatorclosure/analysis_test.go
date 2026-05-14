package admissibleoperatorclosure

import "testing"

func TestOperatorClassesEnumerated(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sieve.Executed || len(a.Sieve.Classes) < 7 || !a.Sieve.AllNativeAudited {
		t.Fatalf("bad closure sieve: %s", FormatSieve(a.Sieve))
	}
}

func TestNoNativeVacuumSelectorFound(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sieve.NoGoApplies || a.Sieve.AnyKineticSafeSelector || a.Sieve.AnyUniqueSelector {
		t.Fatalf("no-go should apply: %s", FormatSieve(a.Sieve))
	}
}

func TestNoGoTheoremFormalized(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.NoGo.Formalized || !a.NoGo.RequiresExtension || a.NoGo.VacuumInputsRemain != 15 {
		t.Fatalf("bad no-go theorem: %s", FormatNoGo(a.NoGo))
	}
}

func TestExtensionForkAndCensus(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Extension.Formalized || len(a.Extension.MinimalNewObjects) < 3 {
		t.Fatalf("bad extension fork: %s", FormatExtension(a.Extension))
	}
	if a.Census.ReductionFromClosure != 0 || a.Census.RemainingInputs != 15 || a.Census.SevenSealReached {
		t.Fatalf("bad census: %s", FormatCensus(a.Census))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusOperatorClassesEnumerated, StatusClosureSieveExecuted, StatusNoGoTheoremFormalized, StatusLandscapeTheoryComplete, StatusExtensionForkFormalized, StatusFailedVacuumSelectorInCore, StatusFailedPhaseIIICoordinates}
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
	res := AdmissibleOperatorClosureVacuumSelectionNoGoTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
