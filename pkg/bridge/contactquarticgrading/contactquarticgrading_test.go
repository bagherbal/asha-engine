package contactquarticgrading

import "testing"

func TestGate159QuarticGhostGradingGaloisObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ExactRationalOverlapMatrix || !a.ExactCharacteristicCertified || !a.ExactRootIsolationCertified || a.QuarticOrbitRows != 4 {
		t.Fatalf("Gate 159 must inherit exact quartic block: %+v", a.Summary)
	}
	if a.GaloisAction.GaloisOrderCandidate != 24 || !a.GaloisAction.TransitiveOrbit || a.GaloisAction.InvariantParityFunctions != 2 || a.GaloisAction.NontrivialInvariantParity != 0 {
		t.Fatalf("transitive quartic orbit should have only constant invariant parities: %+v", a.GaloisAction)
	}
	if a.GaloisAction.AllParityAssignments != 16 || a.GaloisAction.NontrivialParityAssignments != 14 || a.GaloisAction.ZeroSignedCountAssignments != 6 || a.GaloisAction.InvariantZeroSignedCount != 0 {
		t.Fatalf("parity-count audit mismatch: %+v", a.GaloisAction)
	}
	if len(a.ParityClasses) != 4 {
		t.Fatalf("expected four parity classes, got %d", len(a.ParityClasses))
	}
	uniformEven := a.ParityClasses[0]
	uniformOdd := a.ParityClasses[1]
	if !uniformEven.GaloisInvariant || uniformEven.Nontrivial || uniformEven.ZeroSignedCount || uniformEven.SupertraceLedger || !uniformOdd.GaloisInvariant || uniformOdd.Nontrivial || uniformOdd.ZeroSignedCount || uniformOdd.SupertraceLedger {
		t.Fatalf("uniform gradings must be invariant but non-cancelling: %+v %+v", uniformEven, uniformOdd)
	}
	twoTwo := a.ParityClasses[3]
	if !twoTwo.ZeroSignedCount || twoTwo.GaloisInvariant || !twoTwo.RequiresBranchChoice || twoTwo.Assignments != 6 || twoTwo.StabilizerOrder != 4 || twoTwo.OrbitSize != 6 || twoTwo.SupertraceLedger || twoTwo.ZeroBetaLedger || twoTwo.CancellationComplete {
		t.Fatalf("two/two split must be branch-dependent and non-permissive: %+v", twoTwo)
	}
	if a.Obstruction.CanonicalZeroSupertrace || !a.Obstruction.BranchChoiceRequired || a.Obstruction.CompleteGhostGrading || a.Obstruction.ZeroBetaLedger {
		t.Fatalf("Gate 159 must block canonical zero-supertrace grading: %+v", a.Obstruction)
	}
	if !a.Firewall.FirewallClosed || a.Firewall.GhostGrading || a.Firewall.NontrivialParity || a.Firewall.SupertraceCancellation || a.Firewall.ZeroBetaLedger || a.Firewall.ThresholdBetaRows != 0 || a.Firewall.ProvenZeroRows != 0 {
		t.Fatalf("Gate 159 must keep beta firewall closed: %+v", a.Firewall)
	}
	if a.QuarticZeroBetaRows != 0 || a.QuarticBlockBetaRows != 0 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("Gate 159 must not derive physical rows/constants: %+v", a.Summary)
	}
}
