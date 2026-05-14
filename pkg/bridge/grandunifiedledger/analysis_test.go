package grandunifiedledger

import "testing"

func TestSpan(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Span.HighestGateInherited != highestInheritedGate || a.Span.AddsNewPhysicsFit || a.Span.RewritesHistory {
		t.Fatalf("bad span: %s", FormatSpan(a.Span))
	}
}

func TestAbsoluteTriumphs(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Triumphs.Cataloged || len(a.Triumphs.Items) < 8 || a.Triumphs.NativeCount < 8 {
		t.Fatalf("triumph ledger too small: %s", FormatTriumphs(a.Triumphs))
	}
	if !a.Triumphs.ContainsWeakMixing || !a.Triumphs.ContainsMoritaColorSplit || !a.Triumphs.ContainsGenerationTriality || !a.Triumphs.ContainsTrueBimodule || !a.Triumphs.ContainsTopologicalResonance || !a.Triumphs.ContainsTraceEquivalence || !a.Triumphs.ContainsThresholdJump {
		t.Fatalf("missing named triumph: %s", FormatTriumphs(a.Triumphs))
	}
}

func TestProxyAlignments(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Proxies.Cataloged || len(a.Proxies.Alignments) < 4 || !a.Proxies.ContainsTreeLevel125Proxy || !a.Proxies.ContainsThresholdTransport125 || !a.Proxies.Contains331Diagnostic || !a.Proxies.Contains157ContinuousFloor || !a.Proxies.EmpiricalInputsQuarantined || a.Proxies.FinalMassClaimed {
		t.Fatalf("bad proxy ledger: %s", FormatProxies(a.Proxies))
	}
}

func TestFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.Cataloged || len(a.Firewalls.Items) < 6 || !a.Firewalls.ContainsAlphaGUTOrigin || !a.Firewalls.ContainsWeightedTrace25 || !a.Firewalls.ContainsFlavorVacuumSelection || !a.Firewalls.ContainsProjectionMetricSelection || !a.Firewalls.ContainsTwoLoopPolePrecision || !a.Firewalls.ContainsExactColliderMass || a.Firewalls.AnyClosed {
		t.Fatalf("bad firewall ledger: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestPhaseIIIAndIntegrity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.PhaseIII.Formalized || len(a.PhaseIII.Targets) < 5 || !a.PhaseIII.IncludesWeightedTrace || !a.PhaseIII.IncludesFlavorVacuum || !a.PhaseIII.IncludesProjectionMetric || !a.PhaseIII.IncludesPrecisionTransport || !a.PhaseIII.IncludesFullSigmaPotential || !a.PhaseIII.RequiresNoEmpiricalTuning {
		t.Fatalf("bad Phase III ledger: %s", FormatPhaseIII(a.PhaseIII))
	}
	if !a.Audit.NoAlphaGUTFitPromoted || !a.Audit.NoCKMTextureInvented || !a.Audit.NoFlavorMetricForced || !a.Audit.NoObservedHiggsFitInserted || !a.Audit.NoObservedTopFitInserted || !a.Audit.NoTwoLoopClaimed || !a.Audit.NoPoleMassClaimed || !a.Audit.NoFinalTOEClaimed || !a.Audit.NoExactColliderMassClaimed || a.Audit.FiniteCorePolluted {
		t.Fatalf("bad audit: %s", FormatAudit(a.Audit))
	}
}

func TestSummaryAndStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Summary.LedgerCompiled || !a.Summary.ProjectCapstone || !a.Summary.TriumphsReady || !a.Summary.ProxiesReady || !a.Summary.FirewallsReady || !a.Summary.PhaseIIIReady || !a.Summary.FirewallsPreserved || a.Summary.FinalTOEClaimed || a.Summary.ExactColliderClaimed {
		t.Fatalf("bad summary: %s", FormatSummary(a.Summary))
	}
	statuses := Statuses(a)
	required := []string{StatusGrandUnifiedLedgerCompiled, StatusProjectCapstoneAchieved, StatusFailedAlphaGUTNotDerived, StatusFailedFlavorVacuumNotSelected, StatusFailedColliderMassNotClaimed}
	for _, req := range required {
		found := false
		for _, s := range statuses {
			if s == req {
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
	res := GrandUnifiedLedgerProjectCapstoneAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
