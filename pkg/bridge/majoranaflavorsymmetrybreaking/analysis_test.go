package majoranaflavorsymmetrybreaking

import "testing"

func TestCrossTermLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.CrossTerms.Terms) < 5 || a.CrossTerms.NativeTerms < 3 || a.CrossTerms.UnitaryInvariantTerms < 3 || a.CrossTerms.BreakingTemplates < 2 || a.CrossTerms.NativeBreakingTerms != 0 {
		t.Fatalf("bad cross-term ledger: %s", FormatCrossTerms(a.CrossTerms))
	}
}

func TestUnitarySymmetryBreakingSieve(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Symmetry.StandardCrossTermsFlat || !nearlyZero(a.Symmetry.OmegaIndex-1) || a.Symmetry.OmegaAloneBreaksCKM || !a.Symmetry.MajoranaActsOnLeptonSlot || a.Symmetry.DirectQuarkCKMBridgeDerived {
		t.Fatalf("bad symmetry test: %s", FormatSymmetry(a.Symmetry))
	}
}

func TestDegeneracyLiftingStillUnpromoted(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Degeneracy.SignedNullity != signedNullity || !a.Degeneracy.TemplateCanLift || !a.Degeneracy.UniqueMinimumIfProjectorGiven || a.Degeneracy.NativeProjectorDerived || a.Degeneracy.UniqueVacuumDerived {
		t.Fatalf("bad degeneracy sieve: %s", FormatDegeneracy(a.Degeneracy))
	}
}

func TestVerdictAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict.NonUnitaryOperatorFoundNatively || a.Verdict.MajoranaBreaksFlavorNatively || a.Verdict.DegeneracyLifted || a.Verdict.CKMDerived {
		t.Fatalf("bad verdict: %s", FormatVerdict(a.Verdict))
	}
	if !a.Audit.NoCKMImported || !a.Audit.NoObservedYukawasImported || !a.Audit.NoTextureForced || !a.Audit.NoFinalVacuumClaim || !a.Audit.NoColliderMassClaim {
		t.Fatalf("bad audit: %s", FormatAudit(a.Audit))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusMajoranaDiracCrossTermsFormalized, StatusUnitarySymmetryBreakingTestExecuted, StatusDegeneracyLiftingSieveExecuted, StatusFailedMajoranaFlavorBreaking, StatusFailedUniqueVacuumDegeneracyLifted, StatusFailedNativeTextureOperator, StatusFailedCKMTexture}
	for _, req := range required {
		found := false
		for _, s := range statuses {
			if s == req {
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
	res := NonUnitaryInvariantTextureSieveMajoranaFlavorSymmetryBreakingAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
