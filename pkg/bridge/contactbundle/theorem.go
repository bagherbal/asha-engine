package contactbundle

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactLocalBundleObstructionRepresentationRowConstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-LOCAL-BUNDLE-OBSTRUCTION-REPRESENTATION-ROW-CONSTRUCTION"
	const name = "contact local-bundle obstruction / representation-row construction attempt"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact local-bundle obstruction", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 114 constraint shortcut inherited as blocked", Passed: a.ConstraintShortcutBlocked && a.ContactCohomology.CohomologyObstructionDerived && a.ContactCohomology.ContactZeroRowsProved == 0 && a.ContactCohomology.ContactBetaRowsAllowed == 0, Detail: "zero differential leaves seven cohomology classes; no cancellation ledger exists"},
			{Name: "seven positive contact cohomology rows remain available", Passed: a.FiniteCarrierAvailable && a.ContactRows == 7 && a.PositiveFiniteContactRows == 7 && a.SurvivingCohomologyRows == 7, Detail: FormatRows(a.Rows, 10)},
			{Name: "local bundle construction attempted but missing base/fiber/sections", Passed: a.LocalBundleConstructionAttempted && !a.BaseSpaceMapDerived && !a.FiberDerived && !a.TransitionFunctionsDerived && !a.SectionMapDerived, Detail: FormatCriteria(a.Criteria)},
			{Name: "contact gauge representation row absent", Passed: !a.GaugeRepresentationForContactDerived && !a.HyperchargeForContactDerived && !a.ColorRepresentationForContactDerived && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7, Detail: fmt.Sprintf("complete=%d; open=%d", a.RepresentationCompleteRows, a.RepresentationOpenRows)},
			{Name: "Lorentz kinetic/mass/decoupling data absent", Passed: !a.LorentzKineticForContactDerived && !a.CanonicalNormalizationForContactDerived && !a.MassActivationForContactDerived && !a.DecouplingRuleForContactDerived, Detail: FormatCriteria(a.Criteria)},
			{Name: "representation attempts are audited but noncanonical", Passed: a.AnyRepresentationAttemptConstructed && !a.AnyRepresentationAttemptCanonical && a.LocalBundleObstructionDerived && a.RepresentationRowConstructionObstructed, Detail: FormatAttempts(a.BundleAttempts)},
			{Name: "contact beta firewall remains closed", Passed: a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived, Detail: "no representation-complete contact row and no proven zero-row cancellation"},
			{Name: "dichotomy still unresolved", Passed: !a.RepresentationOrConstraintDichotomyDerived && !a.BranchSelectorDerived && a.OpenContactRowsAfter == 7, Detail: fmt.Sprintf("open after=%d", a.OpenContactRowsAfter)},
			{Name: "residual physical-flow nullity unchanged", Passed: a.ResidualNullityAfter == a.ResidualNullityBefore && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken, Detail: fmt.Sprintf("nullity before=%d; after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)},
			{Name: "physical predictions remain sealed", Passed: !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, M*, g_*, or observed thresholds are used"},
		}, Notes: []string{a.TruthStatement, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
