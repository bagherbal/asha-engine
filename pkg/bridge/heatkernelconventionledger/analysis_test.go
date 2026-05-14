package heatkernelconventionledger

import "testing"

func TestGate302Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Input.PositiveKRawCarrierProved || !a.Input.HilbertSchmidtSumStructure || !a.Input.StrictPositiveNeedsNonzeroEdge {
		t.Fatalf("Gate 301 inheritance incomplete: %s", FormatInput(a.Input))
	}
	if a.Input.NumericalYukawasInserted || a.Input.NumericalZHComputed {
		t.Fatalf("Gate 302 must inherit numerical firewalls: %s", FormatInput(a.Input))
	}
}

func TestPrefactorLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Ledger.AllFactorsExplicit || !a.Ledger.AllEmpiricalInputsExcluded || !a.Ledger.CanChoosePositiveClass || a.Ledger.AbsoluteN4Derived || len(a.Ledger.Factors) != 6 {
		t.Fatalf("bad prefactor ledger: %s", FormatLedger(a.Ledger))
	}
	if !allFactorsSignAudited(a.Ledger) {
		t.Fatalf("not all factors are sign audited: %s", FormatLedger(a.Ledger))
	}
	seen := map[string]bool{}
	for _, f := range a.Ledger.Factors {
		seen[f.Symbol] = true
		if f.NumericallyFixed || f.EmpiricalInput || f.SignCondition == "" {
			t.Fatalf("bad factor: %s", FormatPrefactorFactor(f))
		}
	}
	for _, s := range []string{"s_SD", "s_Tr", "m_J", "c_H", "σ_W", "f_0"} {
		if !seen[s] {
			t.Fatalf("missing factor %s in %s", s, FormatLedger(a.Ledger))
		}
	}
}

func TestWickConventionAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Wick.PositiveEnergyMapped || a.Wick.SignAmbiguityHidden || a.Wick.ConventionNativeToFinite || len(a.Wick.SignLedger) < 4 {
		t.Fatalf("bad Wick audit: %s", FormatWick(a.Wick))
	}
}

func TestF0Requirement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.F0.ConditionallyPositive || !a.F0.CanBePositiveWithoutEmpirics || a.F0.NumericalValueDerived || a.F0.ContactSpectralGate288Used {
		t.Fatalf("bad f0 requirement: %s", FormatF0(a.F0))
	}
}

func TestCanonicalMatchingRule(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Matching.RuleFormalized || a.Matching.PhysicalZHComputed || len(a.Matching.AbsorbedFactors) != len(a.Ledger.Factors) {
		t.Fatalf("bad matching rule: %s", FormatMatching(a.Matching))
	}
}

func TestPrefactorPositivitySieve(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Positivity.KRawPositiveSemidefinite || !a.Positivity.PositivePrefactorAvailable || !a.Positivity.StrictZHGuaranteedConditionally || a.Positivity.NumericalStrictZHProved {
		t.Fatalf("bad positivity sieve: %s", FormatPositivity(a.Positivity))
	}
}

func TestFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoF0NumberInserted || !a.Firewalls.NoCutoffGateActivated || !a.Firewalls.NoYukawaNumbersInserted || !a.Firewalls.NoObservedMassesInserted || !a.Firewalls.NoSubtractionSchemeInvented || !a.Firewalls.NoBGapInstantonClaimed || !a.Firewalls.NoHiggsPredictionClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	res := HeatKernelConventionLedgerPositivePrefactorNormalizationAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
