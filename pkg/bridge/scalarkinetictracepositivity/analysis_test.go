package scalarkinetictracepositivity

import "testing"

func TestGate301Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Input.ZHDefined || !a.Input.ScalarKineticSelectorDefined || !a.Input.RescalingDefined {
		t.Fatalf("Gate 300 inheritance incomplete: %s", FormatInput(a.Input))
	}
	if a.Input.PositiveZHNumericallyProved || a.Input.NumericalDynamicsDerived {
		t.Fatalf("Gate 301 must inherit Gate 300 firewalls: %s", FormatInput(a.Input))
	}
}

func TestKineticTraceFunctionalEdgeLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Trace.Formalized || !a.Trace.UsesHilbertSchmidtNorm || len(a.Trace.EdgeTerms) != 4 {
		t.Fatalf("trace functional incomplete: %s", FormatTrace(a.Trace))
	}
	seen := map[string]bool{}
	for _, edge := range a.Trace.EdgeTerms {
		if !edge.AppearsWithAdjoint || !edge.PreservesHermiticity || edge.Multiplicity <= 0 || !edge.SealedAmplitude || edge.NumericallyProvided {
			t.Fatalf("bad scalar edge: %s", FormatScalarEdge(edge))
		}
		seen[edge.AmplitudeSymbol] = true
	}
	for _, symbol := range []string{"Y_u", "Y_d", "Y_e", "Y_ν"} {
		if !seen[symbol] {
			t.Fatalf("missing amplitude symbol %s in %s", symbol, FormatTrace(a.Trace))
		}
	}
}

func TestDoubledSpaceEvaluation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Doubled.DoubleCountingHandled || !a.Doubled.PositivePairingPreserved || a.Doubled.QuarkEdgesMapped != 2 || a.Doubled.LeptonEdgesMapped != 2 || a.Doubled.TotalEdgesMapped != 4 {
		t.Fatalf("bad doubled carrier evaluation: %s", FormatDoubled(a.Doubled))
	}
}

func TestPositivitySieve(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Positivity.PositiveSemidefinite || a.Positivity.NegativeTermsPermitted || a.Positivity.ImaginaryKineticPermitted || !a.Positivity.GhostRiskEliminatedStructurally {
		t.Fatalf("positivity sieve failed: %s", FormatPositivity(a.Positivity))
	}
	if !a.Positivity.StrictPositiveConditional || a.Positivity.StrictPositiveProved {
		t.Fatalf("strict positivity must be conditional, not numeric/native: %s", FormatPositivity(a.Positivity))
	}
}

func TestAmplitudeSealLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Seals.LedgerBuilt || !a.Seals.AtLeastOneNonzeroNeeded || !a.Seals.AllNumericalValuesSealed || !a.Seals.NoEmpiricalValuesInserted || a.Seals.ReducibleToNumericZH || len(a.Seals.Seals) != 4 {
		t.Fatalf("bad amplitude seal ledger: %s", FormatSeals(a.Seals))
	}
	for _, seal := range a.Seals.Seals {
		if seal.NativeValueDerived || !seal.RequiredForNumericZH || !seal.RequiredForStrictZH {
			t.Fatalf("bad seal: %s", FormatAmplitudeSeal(seal))
		}
	}
}

func TestZHCarrierMapAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.ZH.EvaluableAfterAmplitudeSeal || !a.ZH.RequiresPositiveF0 || !a.ZH.RequiresPositiveTraceNorm || !a.ZH.RequiresEuclideanSignLedger || a.ZH.NumericalZHComputed {
		t.Fatalf("bad ZH map: %s", FormatZH(a.ZH))
	}
	if !a.Firewalls.NoYukawaNumbersInserted || !a.Firewalls.NoObservedMassesInserted || !a.Firewalls.NoCutoffMomentInserted || !a.Firewalls.NoSubtractionSchemeInvented || !a.Firewalls.NoBGapInstantonClaimed || !a.Firewalls.NoMassQuarticClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	res := ScalarKineticTraceFunctionalPositiveZHEvaluableCarrierAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
