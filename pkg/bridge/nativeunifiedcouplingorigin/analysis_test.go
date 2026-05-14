package nativeunifiedcouplingorigin

import "testing"

func TestGaugeKineticAbsoluteMap(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Map.ContactMomentPromoted || !a.Map.TauLedgerApplied || a.Map.F0 != 7 || a.Map.TauGUT != 1 {
		t.Fatalf("bad gauge map: %s", FormatGaugeKineticMap(a.Map))
	}
}

func TestAlphaTargetReconstructedNotDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Target.ReconstructedFromGate315Input || a.Target.TargetAlphaInverse != 25 || a.Target.DerivedFromFiniteCore {
		t.Fatalf("target derivation firewall failure: %s", FormatTarget(a.Target))
	}
	if a.Target.TargetGStarSquared <= 0.50 || a.Target.TargetGStarSquared >= 0.51 {
		t.Fatalf("bad gstar target: %s", FormatTarget(a.Target))
	}
}

func TestRequiredPrefactor(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Requirement.RequiredN4 <= 0.28 || a.Requirement.RequiredN4 >= 0.29 || !a.Requirement.MissingPrefactorMatchesTarget {
		t.Fatalf("bad required prefactor: %s", FormatRequirement(a.Requirement))
	}
	if a.Requirement.ContactF0AloneMatchesTarget || a.Requirement.AlphaInverseIfN4EqualsOne <= 80 {
		t.Fatalf("f0-alone insufficiency not detected: %s", FormatRequirement(a.Requirement))
	}
}

func TestTraceCapacityAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Capacity.HasInteger25Candidate || a.Capacity.HasCanonicalNativeDerivation || len(a.Capacity.Candidates) < 3 {
		t.Fatalf("bad trace capacity audit: %s", FormatCapacity(a.Capacity))
	}
}

func TestHiggsProxyRemainsConditional(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.HiggsProxy.SameAsGate315Proxy || a.HiggsProxy.UpgradedToNativeDerivation {
		t.Fatalf("higgs proxy overclaimed: %s", FormatHiggsProxy(a.HiggsProxy))
	}
	if a.HiggsProxy.LambdaFromTargetAlpha <= 0.129 || a.HiggsProxy.LambdaFromTargetAlpha >= 0.132 {
		t.Fatalf("bad lambda proxy: %s", FormatHiggsProxy(a.HiggsProxy))
	}
}

func TestFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoAlphaGUTDerivationClaimed || !a.Firewalls.NoForcedCapacitySelection || !a.Firewalls.NoContinuumPrefactorInvented || !a.Firewalls.NoHiggsMassDerivationClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := NativeUnifiedCouplingOriginAbsoluteGaugeCouplingTraceCapacityAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
