package hilbertspacetracecapacity

import "github.com/bagherbal/asha-engine/pkg/theorem"

func HilbertSpaceDimensionTraceCapacityLedgerAuditTheorem() theorem.Theorem {
	const id = "AUDIT-HILBERT-SPACE-DIMENSION-TRACE-CAPACITY-LEDGER"
	const name = "Hilbert Space Dimension / Trace Capacity Ledger Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 317 trace-capacity ledger", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "physical one-generation H_F state ledger is counted as 16 including nu_R", Passed: a.Physical.IncludesRightNeutrino && a.Physical.LeptonSlotsPerGen == 4 && a.Physical.QuarkSlotsPerGen == 12 && a.Physical.SlotsPerGeneration == 16 && a.Physical.ThreeGenerationSlots == 48, Detail: FormatPhysicalLedger(a.Physical)},
			{Name: "Gate-293 doubled space gives 32 per generation and 96 for three generations", Passed: a.Doubled.DoubledSpaceMandated && a.Doubled.DoubledSlotsPerGen == 32 && a.Doubled.DoubledSlotsThreeGen == 96 && !a.Doubled.Equals25, Detail: FormatDoubled(a.Doubled)},
			{Name: "target C_trace=25 is audited against canonical Hilbert counts", Passed: a.Target.TargetCapacity == 25 && len(a.Target.CanonicalValues) >= 4 && !a.Target.HasCanonical25 && !a.Target.NativeTraceCapacityDerived, Detail: FormatTarget(a.Target)},
			{Name: "25-shaped candidates exist only as noncanonical category mixtures or empirical target echoes", Passed: a.Target.HasAny25Coincidence && !a.Target.HasCanonical25, Detail: FormatTarget(a.Target)},
			{Name: "firewalls reject forced alpha_GUT derivation and mixed-category promotion", Passed: a.Firewalls.NoAlphaGUTDerivationClaimed && a.Firewalls.NoForced25Selection && a.Firewalls.NoMixedCategoryPromotion && a.Firewalls.NoContinuumPrefactorInvented && a.Firewalls.NoHiggsProxyUpgradeClaimed && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "summary localizes the next theorem as weighted trace capacity, not raw dimension count", Passed: a.Summary.PhysicalLedgerBuilt && a.Summary.DoubledSpaceAudited && a.Summary.CanonicalCountsAudited && a.Summary.Target25Audited && !a.Summary.NativeTrace25Derived && !a.Summary.AlphaGUTDerived && a.Summary.FirewallsPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 317 answers the Gate-316 question negatively for raw Hilbert dimensions: 25 is not produced by the canonical completed finite carrier count.", "The absolute unified coupling remains sealed until a weighted trace-capacity or heat-kernel normalization theorem derives the missing capacity."}}
	}}
}
