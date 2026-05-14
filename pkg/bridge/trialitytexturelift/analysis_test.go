package trialitytexturelift

import "testing"

func TestTrialityLiftDimensions(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.LiftAudit.GenerationCount != 3 || a.LiftAudit.FermionKindBlocks != 4 || a.LiftAudit.YukawaTextureMatrices != 4 {
		t.Fatalf("bad lift dimensions: %+v", a.LiftAudit)
	}
	if a.LiftAudit.GeneralEntriesPerMatrix != 9 || a.LiftAudit.TotalGeneralRealEntries != 36 {
		t.Fatalf("bad general texture dimensions: %+v", a.LiftAudit)
	}
	if a.LiftAudit.FullMixingMaps != 72 || a.LiftAudit.DiagonalTrialityChannels != 24 {
		t.Fatalf("bad triality channel counts: %+v", a.LiftAudit)
	}
}

func TestCandidateInventorySeparatesSymmetrySpurionAndAnsatz(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Candidates) != 6 {
		t.Fatalf("expected six candidates, got %d: %s", len(a.Candidates), FormatCandidates(a.Candidates))
	}
	exact := a.Candidates[0]
	if !exact.Canonical || exact.BreaksAllThreeGenerations || exact.ProducesMixing || exact.MatchesScalarShapeCondition {
		t.Fatalf("exact triality candidate should be canonical but too symmetric: %+v", exact)
	}
	spurion := a.Candidates[1]
	if !spurion.BreaksAllThreeGenerations || spurion.ProducesMixing || spurion.Canonical {
		t.Fatalf("generation spurion should split but not be canonical/mixing: %+v", spurion)
	}
	separable := a.Candidates[3]
	if !separable.MatchesScalarShapeCondition || !separable.RequiresBranchChoice || !separable.BreaksAllThreeGenerations || separable.ProducesMixing || separable.Canonical {
		t.Fatalf("separable ansatz audit wrong: %+v", separable)
	}
}

func TestNoCanonicalNoncommutingTextureSelected(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.OperatorAudit.UniqueTextureSelected || a.OperatorAudit.CanonicalBreakingOperatorsFound != 0 || a.OperatorAudit.NonCommutingPairsFound != 0 {
		t.Fatalf("no canonical non-commuting texture should be selected: %+v", a.OperatorAudit)
	}
	if a.OperatorAudit.ScalarShapeConditionalCandidates != 2 || a.OperatorAudit.BranchChoiceCandidates < 4 {
		t.Fatalf("expected conditional scalar ansaetze with branch choices: %+v", a.OperatorAudit)
	}
}

func TestMassFirewallRemainsClosed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.MassAudit.FourYukawaMatricesRecognized || !a.MassAudit.MassesAreSingularValues || !a.MassAudit.MixingNeedsRelativeLeftEigenbasis {
		t.Fatalf("mass texture recognition missing: %+v", a.MassAudit)
	}
	if a.MassAudit.YukawaMatricesDerived || a.MassAudit.FermionMassesDerived || a.MassAudit.CKMPMNSDerived {
		t.Fatalf("mass and mixing should remain open: %+v", a.MassAudit)
	}
	if !a.Firewall.GaugeRatioClosed || !a.Firewall.ScalarShapeTargetAvailable || !a.Firewall.TrialityLiftPerformed || a.Firewall.CanonicalTextureOperatorSelected || a.Firewall.ResidualNullityBefore != 3 || a.Firewall.ResidualNullityAfter != 3 {
		t.Fatalf("bad firewall: %+v", a.Firewall)
	}
}
