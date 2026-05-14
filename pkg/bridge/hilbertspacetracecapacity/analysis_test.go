package hilbertspacetracecapacity

import "testing"

func TestPhysicalStateLedgerCounts(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Physical.IncludesRightNeutrino || a.Physical.LeptonSlotsPerGen != 4 || a.Physical.QuarkSlotsPerGen != 12 || a.Physical.SlotsPerGeneration != 16 || a.Physical.ThreeGenerationSlots != 48 {
		t.Fatalf("bad physical ledger: %s", FormatPhysicalLedger(a.Physical))
	}
}

func TestDoubledSpaceCapacity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Doubled.DoubledSpaceMandated || a.Doubled.DoubledSlotsPerGen != 32 || a.Doubled.DoubledSlotsThreeGen != 96 || a.Doubled.Equals25 {
		t.Fatalf("bad doubled ledger: %s", FormatDoubled(a.Doubled))
	}
}

func TestTarget25NotCanonicallyDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Target.TargetCapacity != 25 || a.Target.HasCanonical25 || a.Target.NativeTraceCapacityDerived || !a.Target.HasAny25Coincidence {
		t.Fatalf("bad target audit: %s", FormatTarget(a.Target))
	}
}

func TestFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoAlphaGUTDerivationClaimed || !a.Firewalls.NoForced25Selection || !a.Firewalls.NoMixedCategoryPromotion || !a.Firewalls.NoContinuumPrefactorInvented || !a.Firewalls.NoHiggsProxyUpgradeClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := HilbertSpaceDimensionTraceCapacityLedgerAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
