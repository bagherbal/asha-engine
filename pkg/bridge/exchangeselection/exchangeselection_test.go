package exchangeselection

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.CandidateRuleCount != 4 {
		t.Fatalf("candidate rule count = %d, want 4", a.CandidateRuleCount)
	}
	if !a.AllRulesPositive {
		t.Fatalf("candidate rules should be positive diagnostics")
	}
	if !a.DirectInverseDisagree {
		t.Fatalf("expected direct/inverse diagnostics to disagree and force action-selection firewall")
	}
	if a.AnyRuleSelectedByAction || a.PropagatorRuleDerived || a.ExchangeKernelUpdated {
		t.Fatalf("Gate 68 must not select a propagator rule without an action Hessian")
	}
	if a.CondensationClaimAllowed {
		t.Fatalf("Gate 68 must not allow condensation claims")
	}
}
