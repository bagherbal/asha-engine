package contactquarticdichotomy

import "testing"

func TestGate157QuarticDichotomyKeepsFirewallClosed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.QuarticOrbitRows != 4 || a.QuarticBlockInvariants != 4 || !a.ExactRationalOverlapMatrix || !a.ExactCharacteristicCertified || !a.ExactRootIsolationCertified {
		t.Fatalf("Gate 157 must inherit exact quartic finite block: %+v", a.Summary)
	}
	if !a.Dichotomy.ExactFiniteSpectralBlock || !a.Dichotomy.PropagatorBranchAudited || !a.Dichotomy.ConstraintBRSTBranchAudited || a.Dichotomy.CompleteBranches != 0 || a.Dichotomy.DichotomyResolved || a.Dichotomy.BetaRowsPermitted != 0 || a.Dichotomy.ZeroRowsProved != 0 {
		t.Fatalf("Gate 157 must audit both branches and leave dichotomy unresolved: %+v", a.Dichotomy)
	}
	if a.PropagatorBranch.PropagatorBranchComplete || a.PropagatorBranch.LocalFieldMap || a.PropagatorBranch.KineticOperator || a.PropagatorBranch.PoleResidueTheorem || a.PropagatorBranch.GaugeRepresentation || a.PropagatorBranch.HyperchargeRow || a.PropagatorBranch.MassActivation || a.PropagatorBranch.DecouplingRule {
		t.Fatalf("Gate 157 must not complete propagator branch: %+v", a.PropagatorBranch)
	}
	if a.ConstraintBranch.ConstraintBranchComplete || a.ConstraintBranch.ConstraintEquations || a.ConstraintBranch.GhostGrading || a.ConstraintBranch.BRSTOperator || a.ConstraintBranch.NilpotentDifferential || a.ConstraintBranch.BRSTPairing || a.ConstraintBranch.SupertraceCancellation || a.ConstraintBranch.ZeroBetaLedger {
		t.Fatalf("Gate 157 must not complete constraint/BRST branch: %+v", a.ConstraintBranch)
	}
	if !a.Firewall.FirewallClosed || a.Firewall.ThresholdBetaRows != 0 || a.Firewall.ProvenZeroRows != 0 || a.QuarticBlockBetaRows != 0 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("Gate 157 must keep beta firewall closed: %+v", a.Firewall)
	}
	if a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("Gate 157 must not change physical-flow nullity or import observations: %+v", a.Summary)
	}
}
