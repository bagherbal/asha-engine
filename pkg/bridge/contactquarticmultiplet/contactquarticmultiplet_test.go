package contactquarticmultiplet

import "testing"

func TestGate155QuarticMultipletAuditKeepsBetaFirewallClosed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.QuarticOrbitRows != 4 || a.QuarticBlockInvariants != 4 || a.QuarticCompressedBlocks != 1 {
		t.Fatalf("Gate 155 must inherit exact quartic block compression: %+v", a.Summary)
	}
	if a.MultipletAudit.DimensionMatches != 4 || a.MultipletAudit.RepresentationComplete != 0 || a.MultipletAudit.BetaPermitted != 0 {
		t.Fatalf("Gate 155 must audit dimension matches without completing representation: %+v", a.MultipletAudit)
	}
	if a.DynkinAudit.DynkinIndexRows != 0 || a.DynkinAudit.BetaIndexRows != 0 || a.DynkinAudit.AllRequirementsSatisfied {
		t.Fatalf("Gate 155 must not derive Dynkin/beta rows: %+v", a.DynkinAudit)
	}
	if a.InvariantUse.RepresentationSemantics || a.InvariantUse.DynkinIndexSemantics || a.InvariantUse.ChargeSemantics {
		t.Fatalf("Gate 155 must keep quartic invariants spectral-only: %+v", a.InvariantUse)
	}
	if !a.BetaPermissionFirewallClosed || a.QuarticBlockBetaRows != 0 || a.ContactBetaRowsAllowed != 0 || a.RepresentationCompleteRows != 0 {
		t.Fatalf("Gate 155 must keep beta firewall closed: %+v", a.Summary)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("Gate 155 must not import observed physics: %+v", a.Summary)
	}
}
