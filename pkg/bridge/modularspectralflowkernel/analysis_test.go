package modularspectralflowkernel

import "testing"

func TestModularOperatorFormalized(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Operator.Formalized || !a.Operator.DeltaDefined || !a.Operator.State.Faithful {
		t.Fatalf("bad operator: %s", FormatOperator(a.Operator))
	}
}

func TestNativeTracialFlowIsTrivial(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Operator.State.Tracial || a.Operator.NonTrivial || a.Flavor.NativeBreaksFlavorOrbit || a.Flavor.DegeneracyBroken {
		t.Fatalf("native flow should be tracial/trivial: %s / %s", FormatOperator(a.Operator), FormatFlavor(a.Flavor))
	}
}

func TestNonTracialCandidateCapacityButNotSelection(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Flavor.CandidateBreaksFlavorOrbit || a.Flavor.CandidateSelectsUniquePoint || a.Flavor.RemainingInputs != 15 {
		t.Fatalf("candidate capacity/selection mismatch: %s", FormatFlavor(a.Flavor))
	}
}

func TestLandscapePreservation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Landscape.PreservesWeakAngle || !a.Landscape.PreservesQuarticRatio || !a.Landscape.PreservesAlphaGUT || !a.Landscape.PreservesMoritaSplit || !a.Landscape.PreservesKineticSafety {
		t.Fatalf("landscape not preserved: %s", FormatLandscape(a.Landscape))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusModularOperatorFormalized, StatusGradientFlowFormalized, StatusFlavorOrbitSieveExecuted, StatusLandscapePreservationAudited, StatusFailedKernelNotConstructed, StatusFailedDegeneracyNotBroken, StatusFailedVacuumNotSelected}
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
	res := ModularSpectralFlowKernelVacuumAddressOperatorConstructionAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
