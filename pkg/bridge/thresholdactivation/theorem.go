package thresholdactivation

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ThresholdActivationDecouplingAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-THRESHOLD-ACTIVATION-DECOUPLING-AUDIT"
	const name = "finite threshold activation and decoupling audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct threshold activation audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "threshold candidate inventory", Passed: a.CandidateCount > 0, Detail: fmt.Sprintf("candidates=%d", a.CandidateCount)},
			{Name: "scalar/contact continuum-field candidate", Passed: a.ScalarSectorRemainsContinuumCandidate, Detail: "active scalar/contact sector has a sector-level doublet representation, but no heavy threshold mass or decoupling rule"},
			{Name: "leakage classified as vacuum-only", Passed: a.LeakageClassifiedAsVacuumOnly, Detail: "bare contact leakage remains a vacuum-frustration invariant, not a threshold"},
			{Name: "B-sector activation rule", Passed: a.BGapActivationDerived, Detail: "not derived; B-sector gap cannot be integrated into RG thresholds yet"},
			{Name: "contact partial-overlap activation rule", Passed: a.ContactOverlapActivationDerived, Detail: "not derived; seven partial-overlap modes remain activation-open"},
			{Name: "physical threshold mass unit", Passed: a.PhysicalMassUnitDerived, Detail: "not derived; every candidate mass remains a scale family M_i(μ)"},
			{Name: "decoupling/matching rule", Passed: a.DecouplingRuleDerived, Detail: "not derived; no continuum-active heavy mode has a matching prescription"},
			{Name: "threshold-corrected beta coefficients", Passed: a.ThresholdCorrectedBetaDerived, Detail: fmt.Sprintf("not derived; beta-eligible threshold modes=%d", a.BetaCorrectionAllowedCount)},
			{Name: "hidden scale insertion", Passed: !a.HiddenScaleInserted, Detail: "no observed mass, GUT scale, or threshold scale was inserted"},
			{Name: "activation summary", Passed: true, Detail: fmt.Sprintf("continuum-field candidates=%d, derived active=%d, integrated-out=%d, heavy thresholds=%d, vacuum-only=%d, unclassified/open=%d; %s", a.ContinuumFieldCandidateCount, a.ContinuumActiveDerivedCount, a.IntegratedOutDerivedCount, a.HeavyThresholdDerivedCount, a.VacuumOnlyCount, a.UnclassifiedCount, FormatDecisions(a.Decisions, 8))},
		}, Notes: []string{a.TruthStatement, fmt.Sprintf("minimum missing data: %v", a.MinimumMissingData)}}
	}}
}
