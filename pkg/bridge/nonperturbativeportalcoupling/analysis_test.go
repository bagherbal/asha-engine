package nonperturbativeportalcoupling

import "testing"

func TestInstantonActionFormalization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Instanton.Formalized || a.Instanton.Action < 12 || a.Instanton.DirectInstantonFactor > 1e-4 || a.Instanton.DirectExpCanHitTarget || a.Instanton.FunctionalDeterminantDerived {
		t.Fatalf("bad instanton ledger: %s", FormatInstanton(a.Instanton))
	}
}

func TestPortalTargetImportedFromGate314(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Target.RequiredDeltaLambda > -0.09 || a.Target.RequiredDeltaLambda < -0.11 || a.Target.RequiredRatio < 0.38 || a.Target.RequiredRatio > 0.40 || !a.Target.ModerateMagnitude {
		t.Fatalf("bad target: %s", FormatTarget(a.Target))
	}
}

func TestMoritaOverlapWitnessNearTargetButNotDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Portal.HasMagnitudeWitness || a.Portal.BestCandidate.Name != "Morita quark-color overlap witness" || !a.Portal.BestCandidate.WithinOnePercent || a.Portal.MagnitudeWitnessDerivedAsPortal {
		t.Fatalf("bad portal witness: %s", FormatPortal(a.Portal))
	}
	if !a.Sieve.TheoreticalCapacity || a.Sieve.NativePortalMapped || a.Sieve.KappaQFourOverPiBGapWitness < 0.39 || a.Sieve.KappaQFourOverPiBGapWitness > 0.392 {
		t.Fatalf("bad sieve: %s", FormatSieve(a.Sieve))
	}
}

func TestFirewallsPreserved(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoPortalCouplingClaimed || !a.Firewalls.NoThresholdJumpClaimed || !a.Firewalls.NoFunctionalDeterminantClaimed || !a.Firewalls.NoHeavyQuarticClaimed || !a.Firewalls.NoFinalMassClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := NonPerturbativeInstantonMappingHeavyPortalCouplingSieveAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
