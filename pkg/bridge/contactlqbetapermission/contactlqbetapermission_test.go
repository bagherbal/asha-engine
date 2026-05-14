package contactlqbetapermission

import "testing"

func TestGate134LeptoquarkBetaPermission(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Gate133SU2ObstructionInherited || a.LeptoquarkRows != 6 || a.CurrentLQSlots != 6 || a.ColorDirections != 3 || a.RealOrientations != 2 {
		t.Fatalf("Gate 133 inheritance failed: %+v", a.Summary)
	}
	if a.HyperchargeRowDerived || a.LocalFieldMapDerived || a.LorentzKineticRowDerived || a.PoleResidueTheoremDerived || a.MassActivationDerived || a.DecouplingRuleDerived {
		t.Fatalf("unexpected representation/propagation row derived: %+v", a.Summary)
	}
	if !a.BetaPermissionFirewallClosed || a.RepresentationCompleteRows != 0 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("beta firewall should remain closed: %+v", a.Summary)
	}
	if a.ResidualS6Choices != 720 || a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 {
		t.Fatalf("residual obstruction changed: %+v", a.Summary)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("physical input leaked")
	}
}

func TestCandidateRowsRemainUnpermitted(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if len(a.CandidateRows) != 6 {
		t.Fatalf("expected six leptoquark candidate rows, got %d", len(a.CandidateRows))
	}
	for _, r := range a.CandidateRows {
		if !r.ColorTriplet || !r.RealOrientationPair || !r.RequiresS6Choice {
			t.Fatalf("candidate row should retain only color/orientation diagnostics and S6 obstruction: %+v", r)
		}
		if r.WeakDoubletDerived || r.HyperchargeDerived || r.LocalFieldDerived || r.LorentzKineticDerived || r.PoleResidueDerived || r.MassActivationDerived || r.DecouplingDerived || r.RepresentationComplete || r.BetaPermitted || r.ZeroContributionProved {
			t.Fatalf("candidate row was over-promoted: %+v", r)
		}
	}
}
