package lorentziantimepullback

import "testing"

func TestLorentzianTimeFormalized(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Time.NativeClifford || !a.Time.Lorentzian || !a.Time.ActsOnSpinor {
		t.Fatalf("bad time generator: %s", FormatTime(a.Time))
	}
}

func TestFlavorCentrality(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Time.FlavorCentral || a.Time.BreaksFlavorOrbit || a.Time.ActsOnFlavor {
		t.Fatalf("e0 should be flavor-central: %s", FormatTime(a.Time))
	}
}

func TestCommutatorVanishes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Commutator.CommutesFlavor || a.Commutator.CommutatorNorm != 0 {
		t.Fatalf("expected zero flavor commutator: %s", FormatCommutator(a.Commutator))
	}
}

func TestFlowPreservesButDoesNotSelect(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Flow.NontrivialPhysicalTime || a.Flow.NontrivialFlavorTime || !a.Flow.PreservesLandscape || !a.Flow.KineticSafe || a.Flow.SelectsVacuum {
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
	required := []string{StatusLorentzianTimeFormalized, StatusCliffordTimePullbackAudited, StatusFlavorCommutatorComputed, StatusFailedTimeKernelNotFlavorBreaking, StatusFailedVacuumNotSelected, StatusFailedCensusNotReduced}
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
	res := LorentzianTimePullbackE0ModularKernelSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
