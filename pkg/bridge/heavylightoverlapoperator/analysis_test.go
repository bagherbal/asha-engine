package heavylightoverlapoperator

import "testing"

func TestFunctionalDeterminantFormalized(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Determinant.Formalized || a.Determinant.PortalTermOrder != 2 || !a.Determinant.NeedsHeavyPropagator || !a.Determinant.NeedsOverlapInsertion {
		t.Fatalf("bad determinant ledger: %s", FormatDeterminant(a.Determinant))
	}
}

func TestDirectSumCrossTermsVanish(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	lane := a.Operator.DirectSumLane
	if !lane.CrossTermsVanish || lane.OverlapOperatorExists || lane.Coefficient != 0 || !lane.DerivedFromMatrices {
		t.Fatalf("direct-sum lane should vanish: %s", FormatLane(lane))
	}
}

func TestTrueBimoduleWitnessMatchesTargetButIsNotDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	lane := a.Operator.TrueBimoduleConditionalLane
	if !lane.OverlapOperatorExists || lane.CrossTermsVanish || !lane.WithinOnePercent || lane.Coefficient < 0.39 || lane.Coefficient > 0.392 || lane.DerivedFromMatrices {
		t.Fatalf("bad true-bimodule lane: %s", FormatLane(lane))
	}
	if a.Operator.ExplicitSigmaHMatrixDerived || a.Operator.OverlapIndexDerived || a.Operator.PortalPromoted {
		t.Fatalf("operator was over-promoted: %s", FormatOperator(a.Operator))
	}
}

func TestMultiplicativeSieveKeepsPromotionFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Sieve.WithinOnePercent || a.Sieve.FactorsForcedMultiplicative || a.Sieve.ConditionalDeltaLambda > -0.097 || a.Sieve.ConditionalDeltaLambda < -0.099 {
		t.Fatalf("bad multiplicative sieve: %s", FormatSieve(a.Sieve))
	}
	if a.Promotion.PromotionAuthorized || a.Promotion.ThresholdJumpDerived || a.Promotion.HeavySelfQuarticDerived {
		t.Fatalf("promotion firewall failed: %s", FormatPromotion(a.Promotion))
	}
}

func TestFirewallsPreserved(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoPortalCouplingClaimed || !a.Firewalls.NoThresholdJumpClaimed || !a.Firewalls.NoFinalMassClaimed || !a.Firewalls.NoExplicitMatrixClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewalls failed: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := FunctionalDeterminantSieveHeavyLightOverlapOperatorAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
