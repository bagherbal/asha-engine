package hierarchyrankpromotion

import "testing"

func close(a, b, tol float64) bool {
	if a > b {
		return a-b <= tol
	}
	return b-a <= tol
}

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || a.Inputs.BooleanProjectorRank != 56 || a.Inputs.STop <= 78 {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestEffectiveExponents(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !close(a.Targets.Log2Unreduced, 55.46076288096928, 1e-12) {
		t.Fatalf("unexpected unreduced exponent: %s", FormatTargets(a.Targets))
	}
	if !close(a.Targets.Log2Reduced, 53.13501481623311, 1e-12) {
		t.Fatalf("unexpected reduced exponent: %s", FormatTargets(a.Targets))
	}
}

func TestRank56AndHalfAction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Ledger.Rank56.Promotable || !close(a.Ledger.Rank56.RatioToUnreduced, 0.6881346907920799, 1e-15) {
		t.Fatalf("bad rank56 candidate: %s", FormatCandidate(a.Ledger.Rank56))
	}
	if a.Ledger.HalfTopological.Promotable || !close(a.Ledger.HalfTopological.RatioToUnreduced, 0.35489043118025776, 1e-15) {
		t.Fatalf("bad half action candidate: %s", FormatCandidate(a.Ledger.HalfTopological))
	}
}

func TestPrefactorSieve(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Prefactors.NativePrefactorDerived {
		t.Fatalf("prefactor should not be derived: %s", FormatPrefactors(a.Prefactors))
	}
	if !close(a.Prefactors.RequiredForRank56, 1.4532038761902069, 1e-12) {
		t.Fatalf("unexpected required rank prefactor: %s", FormatPrefactors(a.Prefactors))
	}
	if a.Prefactors.BestAccidentalExpression != "sqrt(2)·2^-56" {
		t.Fatalf("unexpected best accidental expression: %s", FormatPrefactors(a.Prefactors))
	}
}

func TestFirewallsAndStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Firewalls.RankExponentControlsMassScale || a.Firewalls.HalfTopologicalActionControlsVEV || !a.Firewalls.ArbitraryExponentFittingRejected {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
	statuses := Statuses(a)
	required := []string{StatusFailedRank56ScaleLawNotDerived, StatusFailedHalfInstantonRuleNotDerived, StatusFailedHierarchyStillNotDerived}
	for _, req := range required {
		found := false
		for _, got := range statuses {
			if got == req {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing status %s in %v", req, statuses)
		}
	}
}

func TestTheoremPasses(t *testing.T) {
	res := Rank56HalfInstantonHierarchyPromotionSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
