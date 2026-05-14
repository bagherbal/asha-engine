package supportgenerationintertwiner

import "testing"

func TestGate369ObstructionInherited(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Executed || !a.Inheritance.NativeEtaTraceCentral || !a.Inheritance.TauEtaInsertionCircular || a.Inheritance.RequiredNewObject == "" {
		t.Fatalf("bad inheritance:\n%s", FormatInheritance(a.Inheritance))
	}
}

func TestIntertwinerFormalization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Formalization.Executed || a.Formalization.TargetFormula == "" || len(a.Formalization.NativeAdmissibility) < 3 || len(a.Formalization.ForbiddenMoves) < 3 {
		t.Fatalf("bad formalization:\n%s", FormatFormalization(a.Formalization))
	}
}

func TestNativeCandidatesFactorThroughIdentity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"A", "B", "C", "D", "E"} {
		c := candidateByID(a.Candidates, id)
		if !c.Native || c.Circular || !c.U3Equivariant || c.GenerationAddressed || !c.Central || c.NonCentral || c.Decomposition.HasNonzeroB || c.Promotable {
			t.Fatalf("native candidate %s should factor through I3:\n%s", id, FormatCandidate(c))
		}
	}
}

func TestTauEtaWitnessIsCircular(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := candidateByID(a.Candidates, "F")
	if c.Native || !c.Circular || !c.GenerationAddressed || c.U3Equivariant || !c.NonCentral || !c.Decomposition.TargetReached || !c.BreaksFlavorOrbit || c.Promotable {
		t.Fatalf("tau eta witness should be noncentral but circular:\n%s", FormatCandidate(c))
	}
}

func TestNoGoAndActivation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.NoGo.AllNativeMapsFactorThroughI3 || a.NoGo.NativeNoncentralCount != 0 || a.NoGo.NativeGenerationAddressCount != 0 || a.NoGo.CircularNoncentralWitnessCount != 1 {
		t.Fatalf("bad no-go audit:\n%s", FormatNoGo(a.NoGo))
	}
	if a.Activation.NativeIntertwinerDerived || a.Activation.InternalThermalTimeActivated || a.Activation.TauEtaHamiltonianSelected || !a.Activation.CircularCapacityWitnessed {
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
	if !a.Kinetic.AllCandidatesSelf || !a.Kinetic.NoNonunitaryPush || !a.Kinetic.FaithfulCarrier {
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
	required := []string{StatusIntertwinerSieveFormalized, StatusEquivarianceNoGoAudited, StatusTensionNoNativeGenerationAddress, StatusTensionTauIntertwinerCircular, StatusFailedIntertwinerNotDerived, StatusFailedThermalTimeNotActivated, StatusFailedCensusNotReduced}
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
	res := SupportToGenerationIntertwinerTopologicalIndexMapSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit did not pass:\n%s", res.Details())
	}
}
