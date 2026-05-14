package betapermission

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactOverlapRepresentationConstraintBetaPermissionFirewallTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-OVERLAP-REPRESENTATION-CONSTRAINT-BETA-PERMISSION-FIREWALL"
	const name = "contact-overlap representation-or-constraint dichotomy / beta-permission firewall"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact beta-permission firewall", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 111 unresolved contact-field state inherited", Passed: a.ContactFieldMap.ContactRows == 7 && a.ContactFieldMap.PositiveFiniteContactRows == 7 && !a.ContactFieldMap.ContactFieldClassDerived && a.ContactFieldMap.ResidualNullityAfter == 3, Detail: fmt.Sprintf("contact=%d; positive=%d; field-class-derived=%t; nullity=%d", a.ContactFieldMap.ContactRows, a.ContactFieldMap.PositiveFiniteContactRows, a.ContactFieldMap.ContactFieldClassDerived, a.ContactFieldMap.ResidualNullityAfter)},
			{Name: "beta-permission firewall constructed", Passed: a.BetaPermissionFirewallConstructed && a.PhysicalBranchRuleConstructed && a.ConstraintBranchRuleConstructed, Detail: FormatRules(a.BranchRules)},
			{Name: "seven contact modes are dichotomy-open", Passed: a.ContactRows == 7 && a.PositiveFiniteContactRows == 7 && a.DichotomyOpenRows == 7 && a.UnresolvedContactRows == 7 && a.ResolvedContactRows == 0, Detail: FormatRows(a.Rows, 12)},
			{Name: "physical representation branch incomplete", Passed: a.PhysicalBranchCompleteRows == 0 && !a.ActivationRuleDerived && !a.DecouplingMatchingRuleDerived && !a.PhysicalMassUnitDerived, Detail: "no local support, Lorentz kinetic action, gauge representation, pole/residue, mass unit, activation, or decoupling rule is complete for contact modes"},
			{Name: "constraint/BRST branch incomplete", Passed: a.ConstraintBranchCompleteRows == 0 && a.ZeroContributionRowsProved == 0, Detail: "no constraint generator, ghost grading, nilpotent Q, BRST pairing, or cancellation ledger proves zero contact beta rows"},
			{Name: "ambiguous branches blocked by firewall", Passed: len(a.FirewallWitnesses) >= 5 && a.BetaCorrectionRowsAllowed == 0, Detail: FormatWitnesses(a.FirewallWitnesses)},
			{Name: "threshold beta tensor remains sealed", Passed: a.BetaCorrectionRowsAllowed == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived && !a.RepresentationOrConstraintDichotomyDerived && !a.AllContactModesResolved, Detail: fmt.Sprintf("beta rows allowed=%d; resolved contact rows=%d/%d", a.BetaCorrectionRowsAllowed, a.ResolvedContactRows, a.ContactRows)},
			{Name: "residual physical-flow nullity unchanged", Passed: a.ResidualNullityAfter == a.ResidualNullityBefore && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken, Detail: fmt.Sprintf("nullity before=%d; after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)},
			{Name: "physical predictions remain sealed", Passed: !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, M*, g_*, or observed thresholds are used"},
		}, Notes: []string{a.TruthStatement, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
