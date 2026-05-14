package modeclass

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteMassActivationClassClassifierTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-MASS-ACTIVATION-CLASS-CLASSIFIER"
	const name = "finite mass/activation class classifier for B-sector and contact-overlap modes"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build finite mode-class classifier", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 108 beta-matching obstruction inherited", Passed: a.BetaMatching.ResidualNullityAfter == 3 && !a.BetaMatching.ThresholdCorrectedBetaDerived && a.BetaMatching.IncompleteOpenRows == 8, Detail: fmt.Sprintf("nullity=%d; incomplete open rows=%d", a.BetaMatching.ResidualNullityAfter, a.BetaMatching.IncompleteOpenRows)},
			{Name: "mode-class ledger constructed", Passed: a.TotalRows >= 15 && a.BaselineSectorRows == 1 && a.UnorientedScalarRows == 4 && a.VacuumFrustrationRows == 1, Detail: fmt.Sprintf("rows=%d; baseline=%d; scalar constituents=%d; vacuum=%d; %s", a.TotalRows, a.BaselineSectorRows, a.UnorientedScalarRows, a.VacuumFrustrationRows, FormatRows(a.Rows, 9))},
			{Name: "B-sector gap classified as constrained finite vacuum mode", Passed: a.BGapClassifiedAsConstrainedFinite && a.BGapExcludedFromThresholdBeta && a.BGapRows == 1, Detail: fmt.Sprintf("B-gap rows=%d; constrained=%d; excluded=%t", a.BGapRows, a.ConstrainedFiniteRows, a.BGapExcludedFromThresholdBeta)},
			{Name: "contact partial-overlap modes remain class-open", Passed: a.ContactOpenRows == 7 && !a.ContactOverlapClassDerived && a.ContactOverlapAmbiguityWitnessed, Detail: fmt.Sprintf("open contact modes=%d; witnesses=%s", a.ContactOpenRows, FormatWitnesses(a.AmbiguityWitnesses))},
			{Name: "no physical heavy-threshold class selected", Passed: a.PhysicalHeavyThresholdRows == 0 && a.RegulatorRowsDerived == 0 && a.BetaCorrectionRowsAllowed == 0, Detail: fmt.Sprintf("physical-heavy=%d; regulators=%d; beta rows allowed=%d", a.PhysicalHeavyThresholdRows, a.RegulatorRowsDerived, a.BetaCorrectionRowsAllowed)},
			{Name: "classifier attempts audited", Passed: len(a.ClassifierAttempts) >= 6, Detail: FormatAttempts(a.ClassifierAttempts)},
			{Name: "activation, decoupling, mass spectrum still absent", Passed: !a.ActivationPredicateDerived && !a.DecouplingClassDerived && !a.PhysicalMassUnitDerived && !a.ThresholdMassSpectrumDerived, Detail: "no finite predicate supplies mass unit, active shell, scalar/Weyl/vector class, or decoupling rule"},
			{Name: "threshold beta tensor remains sealed", Passed: !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived, Detail: "B-gap is excluded; contact modes are still class-open; no heavy Δb_i rows are allowed"},
			{Name: "residual nullity unchanged", Passed: a.ResidualNullityAfter == a.ResidualNullityBefore && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken, Detail: fmt.Sprintf("nullity before=%d; after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)},
			{Name: "physical predictions remain sealed", Passed: !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, observed scales, or fitted thresholds are used"},
		}, Notes: []string{a.TruthStatement, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
