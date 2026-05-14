package branchselector

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactModeBranchSelectorConstructionAttemptTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-MODE-BRANCH-SELECTOR-CONSTRUCTION-ATTEMPT"
	const name = "contact-mode branch selector / finite constraint complex or local bundle construction attempt"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact-mode branch selector attempt", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 112 beta-permission firewall inherited", Passed: a.BetaPermission.BetaPermissionFirewallConstructed && a.BetaPermission.ContactRows == 7 && a.BetaPermission.UnresolvedContactRows == 7 && a.BetaPermission.BetaCorrectionRowsAllowed == 0, Detail: fmt.Sprintf("contact=%d; unresolved=%d; beta rows=%d", a.BetaPermission.ContactRows, a.BetaPermission.UnresolvedContactRows, a.BetaPermission.BetaCorrectionRowsAllowed)},
			{Name: "seven positive contact modes remain branch-open", Passed: a.ContactRows == 7 && a.PositiveFiniteContactRows == 7 && a.OpenContactRows == 7 && a.ResolvedContactRows == 0 && a.UnresolvedContactRows == 7, Detail: FormatRows(a.Rows, 12)},
			{Name: "local-bundle branch attempted but incomplete", Passed: a.LocalBundleAttemptConstructed && a.LocalBundleBranchCompleteRows == 0 && !a.BaseSpaceDerived && !a.FiberRepresentationDerived && !a.TransitionFunctionsDerived && !a.SectionMapDerived && !a.LorentzKineticForContactDerived && !a.PoleResidueForContactDerived, Detail: FormatAttempts(a.ConstructionAttempts)},
			{Name: "constraint-complex branch attempted but incomplete", Passed: a.ConstraintComplexAttemptConstructed && a.ConstraintComplexCompleteRows == 0 && !a.ChainGroupsDerived && !a.DifferentialDerived && !a.NilpotentDifferentialDerived && !a.GhostGradingDerived && !a.PairingDerived && !a.ExactnessOrCohomologyDerived && !a.CancellationLedgerDerived, Detail: FormatAttempts(a.ConstructionAttempts)},
			{Name: "branch selector not derived", Passed: a.BranchSelectorAttempted && !a.BranchSelectorDerived && a.PhysicalSelectedRows == 0 && a.ConstraintSelectedRows == 0 && !a.RepresentationOrConstraintDichotomyDerived, Detail: FormatCriteria(a.SelectorCriteria)},
			{Name: "ambiguous completions blocked", Passed: len(a.AmbiguityWitnesses) >= 5 && a.BetaCorrectionRowsAllowed == 0 && a.ZeroContributionRowsProved == 0, Detail: FormatWitnesses(a.AmbiguityWitnesses)},
			{Name: "threshold beta tensor remains sealed", Passed: a.BetaCorrectionRowsAllowed == 0 && a.ZeroContributionRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived, Detail: fmt.Sprintf("beta rows allowed=%d; zero rows proved=%d", a.BetaCorrectionRowsAllowed, a.ZeroContributionRowsProved)},
			{Name: "residual physical-flow nullity unchanged", Passed: a.ResidualNullityAfter == a.ResidualNullityBefore && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken, Detail: fmt.Sprintf("nullity before=%d; after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)},
			{Name: "physical predictions remain sealed", Passed: !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, M*, g_*, or observed thresholds are used"},
		}, Notes: []string{a.TruthStatement, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
