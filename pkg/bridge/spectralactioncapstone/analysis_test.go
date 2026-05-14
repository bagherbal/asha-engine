package spectralactioncapstone

import "testing"

func TestGate282SpectralActionCapstone(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	if !a.Scaffold.MoritaMultiplicityRecorded || !a.Scaffold.ScalarMoritaShapeConstraintRecorded || !a.Scaffold.TwoBranchRConstraintRecorded {
		t.Fatalf("expected scaffold to record Morita and scalar-Morita achievements: %+v", a.Scaffold)
	}
	if len(a.Obstructions.Obstructions) != 6 || !a.Obstructions.AllUnsatisfied || !a.Obstructions.HiggsPredictionBlocked {
		t.Fatalf("expected six unsatisfied obstructions blocking Higgs prediction: %+v", a.Obstructions)
	}
	if !a.Seal.Active || a.Seal.CanClaimFiniteDerivedHiggsRatio {
		t.Fatalf("expected active Higgs firewall with no finite Higgs claim: %+v", a.Seal)
	}
	if a.Summary.HiggsRatioDerived || !a.Summary.PathBClosed {
		t.Fatalf("expected Path B closure without Higgs derivation: %+v", a.Summary)
	}
	if a.Firewall.FiniteCorePolluted || !a.Firewall.NoObservedMassesUsed || !a.Firewall.NoEmpiricalYukawaInserted {
		t.Fatalf("firewall failure: %+v", a.Firewall)
	}
}

func TestGate282TheoremPassesChecks(t *testing.T) {
	res := SpectralActionEpistemologicalCapstoneHiggsPredictionFirewallAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem checks failed:\n%s", res.Details())
	}
	if res.Status == "EXACT_FINITE" {
		t.Fatalf("capstone must not be exact finite: %s", res.Status)
	}
}
