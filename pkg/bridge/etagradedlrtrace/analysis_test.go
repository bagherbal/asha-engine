package etagradedlrtrace

import "testing"

func TestGate368TargetInherited(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Executed || !a.Inheritance.CircularityFirewall || a.Inheritance.TargetEquation == "" {
		t.Fatalf("bad inheritance:\n%s", FormatInheritance(a.Inheritance))
	}
}

func TestEtaFormalizationSeparatesSupportFromGeneration(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Formalization.Executed || a.Formalization.NativeEtaActsOnGeneration || len(a.Formalization.NativeSupportEta) != 2 || len(a.Formalization.GenerationEtaCandidate) != 3 {
		t.Fatalf("bad formalization:\n%s", FormatFormalization(a.Formalization))
	}
}

func TestNativeEtaTraceLanesRemainCentral(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"A", "B", "C"} {
		lane := laneByID(a.Lanes, id)
		if !lane.Native || lane.Circular || !lane.Central || lane.Decomposition.HasNonzeroB || lane.BreaksFlavorOrbit || lane.PromotedHamiltonian {
			t.Fatalf("native lane %s should remain central and unpromoted:\n%s", id, FormatLane(lane))
		}
	}
}

func TestGenerationEtaWitnessIsCircular(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	lane := laneByID(a.Lanes, "D")
	if lane.Native || !lane.Circular || !lane.NonCentral || !lane.Decomposition.TargetReached || !lane.BreaksFlavorOrbit || lane.PromotedHamiltonian {
		t.Fatalf("generation eta witness should be noncentral but circular:\n%s", FormatLane(lane))
	}
}

func TestActivationRefused(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Activation.NativeTargetReached || !a.Activation.CircularCapacityWitnessed || a.Activation.PromotedNative || a.Activation.InternalTimeActivated || a.Activation.EnergyConstraintDerived {
		t.Fatalf("bad activation audit:\n%s", FormatActivation(a.Activation))
	}
}

func TestLandscapeKineticAndCensus(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Landscape.WeakMixingPreserved || !a.Landscape.QuarticRatioPreserved || !a.Landscape.AlphaGUTPreserved || !a.Landscape.MoritaSplitPreserved || a.Landscape.FiniteCorePolluted {
		t.Fatalf("bad landscape:\n%s", FormatLandscape(a.Landscape))
	}
	if !a.Kinetic.AllCandidatesSelf || !a.Kinetic.FaithfulStates || !a.Kinetic.NoGhostMetric {
		t.Fatalf("bad kinetic:\n%s", FormatKinetic(a.Kinetic))
	}
	if a.Census.Reduction != 0 || a.Census.RemainingInputs != 15 {
		t.Fatalf("bad census:\n%s", FormatCensus(a.Census))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusEtaGradedTraceExecuted, StatusTensionNativeSupportEtaCentral, StatusTensionGenerationEtaInsertionCircular, StatusFailedOriginNotDerived, StatusFailedTauStillNotSelected, StatusFailedCensusNotReduced}
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

func TestTheoremPassesAsFailedRouteAudit(t *testing.T) {
	res := EtaGradedLeftRightTraceNoncentralHamiltonianExtractionSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit did not pass:\n%s", res.Details())
	}
}
