package contactquarticgalois

import "testing"

func TestGate152KeepsQuarticBranchesNonCanonical(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.QuarticField.GaloisOrderCandidate != 24 || !a.QuarticField.GaloisTransitive || a.CanonicalQuarticBranches != 0 {
		t.Fatalf("Gate 152 must expose Galois branch obstruction: %+v", a.QuarticField)
	}
	if a.IndividualQuarticProjectors != 0 || a.ProjectorAudit.RowwiseRootAssignments != 0 {
		t.Fatalf("Gate 152 must not construct individual quartic semantics: projector=%+v summary=%+v", a.ProjectorAudit, a.Summary)
	}
	if !a.BetaPermissionFirewallClosed || a.ContactBetaRowsAllowed != 0 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("Gate 152 must keep physical firewall closed: summary=%+v", a.Summary)
	}
}
