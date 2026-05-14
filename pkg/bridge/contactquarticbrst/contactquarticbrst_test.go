package contactquarticbrst

import "testing"

func TestGate158QuarticBRSTAttemptKeepsFirewallClosed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ExactRationalOverlapMatrix || !a.ExactCharacteristicCertified || !a.ExactRootIsolationCertified || a.QuarticOrbitRows != 4 {
		t.Fatalf("Gate 158 must inherit exact quartic block: %+v", a.Summary)
	}
	if a.Construction.DifferentialsAudited != 4 || a.Construction.NilpotentCandidates != 3 || a.Construction.CanonicalNilpotent != 1 || a.Construction.NonzeroCanonicalBRST != 0 || a.Construction.CompleteBRSTCancellations != 0 || a.Construction.ZeroSupertraceLedgers != 0 || a.Construction.ZeroBetaLedgers != 0 || a.Construction.ConstructionComplete {
		t.Fatalf("Gate 158 must audit BRST candidates without completing cancellation: %+v", a.Construction)
	}
	zero := a.DifferentialCandidates[0]
	if !zero.Nilpotent || !zero.SquareZero || !zero.Canonical || !zero.GaloisInvariant || zero.CohomologyDimension != 4 || zero.ZeroSupertraceLedger || zero.ZeroBetaLedger || zero.BRSTCancellationComplete {
		t.Fatalf("zero differential must be canonical but inert: %+v", zero)
	}
	if a.Supertrace.GaloisInvariantGradings != 2 || a.Supertrace.NontrivialInvariantGradings != 0 || a.Supertrace.ZeroSupertraceGradings != 0 || a.Supertrace.CanonicalZeroSupertrace || a.Supertrace.ZeroBetaLedger {
		t.Fatalf("Gate 158 must block Galois-invariant zero-supertrace: %+v", a.Supertrace)
	}
	branchSplit := a.GhostGradings[2]
	if !branchSplit.TraceZeroPossible || branchSplit.GaloisInvariant || !branchSplit.RequiresBranchChoices || branchSplit.ZeroSupertraceLedger || branchSplit.BRSTCancellationComplete {
		t.Fatalf("two/two split must remain branch-dependent and non-permissive: %+v", branchSplit)
	}
	if !a.Firewall.FirewallClosed || a.Firewall.BRSTOperator || a.Firewall.GhostGrading || a.Firewall.NilpotentDifferential || a.Firewall.SupertraceCancellation || a.Firewall.ZeroBetaLedger || a.Firewall.ThresholdBetaRows != 0 || a.Firewall.ProvenZeroRows != 0 {
		t.Fatalf("Gate 158 must keep BRST beta firewall closed: %+v", a.Firewall)
	}
	if a.QuarticZeroBetaRows != 0 || a.QuarticBlockBetaRows != 0 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("Gate 158 must not derive physical rows/constants: %+v", a.Summary)
	}
}
