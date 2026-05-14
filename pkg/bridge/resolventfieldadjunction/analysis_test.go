package resolventfieldadjunction

import "testing"

func TestGate280ResolventAdjunctionProjectors(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	if !a.Seal.Active || !a.Seal.GrantsConditionalProjectors || a.Seal.GrantsNativeBranchSelection {
		t.Fatalf("unexpected seal state: %+v", a.Seal)
	}
	if a.BranchSpace.BranchCount != 3 || !a.BranchSpace.AllBranchesProjectorsValid {
		t.Fatalf("expected 3 valid conditional branches: %+v", a.BranchSpace)
	}
	if !allProjectorResidualsOK(a.BranchSpace) {
		t.Fatalf("projector residuals too large: %s", FormatBranchSpace(a.BranchSpace))
	}
}

func TestGate280DoesNotOverselectSectorOrRBranch(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	if !a.SectorBijection.SectorPairingSelected || !a.SectorBijection.ConditionalProjectorsExist {
		t.Fatalf("expected inherited sector pairing and conditional projectors")
	}
	if a.SectorBijection.MappingDerived || a.RBranch.UniqueAmplitudeBranch || a.RBranch.ResolventToRMapDerived {
		t.Fatalf("overselected sector/r branch: sector=%+v r=%+v", a.SectorBijection, a.RBranch)
	}
	if err := AssertNoOverclaim(a); err != nil {
		t.Fatalf("overclaim: %v", err)
	}
}

func TestGate280TheoremChecksPass(t *testing.T) {
	th := ResolventFieldAdjunctionContactProjectorConstructionAuditTheorem()
	res := th.Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("theorem build failed: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
