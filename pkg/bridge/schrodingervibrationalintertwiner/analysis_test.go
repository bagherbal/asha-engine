package schrodingervibrationalintertwiner

import "testing"

func TestGate370Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Executed || !a.Inheritance.AllCurrentNativeMapsI3 || !a.Inheritance.TauEtaManualMapCircular || !a.Inheritance.NoEmpiricalFlavorData {
		t.Fatalf("bad inheritance:\n%s", FormatInheritance(a.Inheritance))
	}
}
func TestFockFormalizationIsHypothesis(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Formalization.Executed || len(a.Formalization.FockBasis) != 3 || len(a.Formalization.NumberOperator) != 3 || a.Formalization.BasisSelectedByASHA {
		t.Fatalf("bad formalization:\n%s", FormatFormalization(a.Formalization))
	}
}
func TestNumberOperatorBreaksU3ButIsNotTau(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	n := laneByID(a.Lanes, "B")
	if !n.NativeToChosenFock || n.NativeToCurrentASHA || n.Central || !n.NonCentral || !n.BreaksFlavorOrbit || n.Decomposition.TargetReached || n.PromotableHamiltonian {
		t.Fatalf("number lane should be capacity witness only:\n%s", FormatLane(n))
	}
}
func TestSupportDefectNumberCouplingIsNewStructure(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := laneByID(a.Lanes, "C")
	if !c.RequiresNewCoupling || !c.TopologicalPullback || c.NativeToCurrentASHA || c.Decomposition.TargetReached || c.PromotableHamiltonian {
		t.Fatalf("support-defect times N should be new and unpromoted:\n%s", FormatLane(c))
	}
}
func TestEntropyOperatorDependsOnChosenNumberHamiltonian(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	e := laneByID(a.Lanes, "G")
	if !e.NonCentral || !e.KMS.Faithful || !e.KMS.NontrivialFrequencies || e.Decomposition.TargetReached || e.NativeToCurrentASHA {
		t.Fatalf("entropy lane should be noncentral but not derived:\n%s", FormatLane(e))
	}
}
func TestTauPolynomialWitnessIsCircular(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	h := laneByID(a.Lanes, "I")
	if !h.Circular || !h.NonCentral || !h.Decomposition.TargetReached || !h.BreaksFlavorOrbit || h.PromotableHamiltonian {
		t.Fatalf("tau polynomial witness should be exact but circular:\n%s", FormatLane(h))
	}
}
func TestPullbackActivationAndCensus(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Pullback.AnyNoncentralFockWitness || a.Pullback.NativeASHAFockMapDerived || a.Pullback.AnyPromotableHamiltonian || !a.Pullback.PolynomialCircular {
		t.Fatalf("bad pullback:\n%s", FormatPullback(a.Pullback))
	}
	if !a.Activation.NoncentralCapacityWitnessed || a.Activation.InternalThermalTimeActivated || a.Activation.VacuumCoordinatesReduced {
		t.Fatalf("bad activation:\n%s", FormatActivation(a.Activation))
	}
	if a.Census.Reduction != 0 || a.Census.RemainingInputs != 15 {
		t.Fatalf("bad census:\n%s", FormatCensus(a.Census))
	}
}
func TestLandscapeAndKinetic(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Landscape.WeakMixingPreserved || !a.Landscape.QuarticRatioPreserved || !a.Landscape.AlphaGUTPreserved || !a.Landscape.MoritaSplitPreserved || a.Landscape.FiniteCorePolluted {
		t.Fatalf("bad landscape:\n%s", FormatLandscape(a.Landscape))
	}
	if !a.Kinetic.AllOperatorsSelf || !a.Kinetic.FaithfulKMSStates || !a.Kinetic.NoNonunitaryPush {
		t.Fatalf("bad kinetic:\n%s", FormatKinetic(a.Kinetic))
	}
}
func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusFockSpaceFormalized, StatusNumberOperatorSieveExecuted, StatusVibrationalCapacityWitnessed, StatusTensionFockBasisNotSelected, StatusTensionPolynomialCircular, StatusFailedVibrationalIntertwinerNotDerived, StatusFailedThermalTimeNotActivated, StatusFailedCensusNotReduced}
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
	res := SchrodingerVibrationalModesQuantumInformationIntertwinerAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit did not pass:\n%s", res.Details())
	}
}
