package contactpropagator

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactOverlapKineticSignLocalityPropagatorClassifierTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-OVERLAP-KINETIC-SIGN-LOCALITY-PROPAGATOR-CLASSIFIER"
	const name = "contact-overlap kinetic-sign / locality / propagator classifier search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact-overlap propagator classifier", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 109 class-open contact obstruction inherited", Passed: a.ModeClass.ContactOpenRows == 7 && !a.ModeClass.ContactOverlapClassDerived && a.ModeClass.ResidualNullityAfter == 3, Detail: fmt.Sprintf("open contact=%d; nullity=%d", a.ModeClass.ContactOpenRows, a.ModeClass.ResidualNullityAfter)},
			{Name: "contact finite-overlap spectrum is positive", Passed: a.ContactRows == 7 && a.PositiveFiniteContactRows == 7 && a.PositiveFiniteOverlapSpectrumDerived, Detail: fmt.Sprintf("contact rows=%d; positive=%d; %s", a.ContactRows, a.PositiveFiniteContactRows, FormatRows(a.Rows, 10))},
			{Name: "Lorentz kinetic sign is not derived from overlap eigenvalues", Passed: !a.LorentzKineticSignDerived, Detail: "finite overlap positivity is not a spacetime kinetic operator or p²+m² law"},
			{Name: "locality and finite-to-continuum field map absent", Passed: !a.LocalityDerived, Detail: "no local continuum variable is assigned to the seven contact partial-overlap modes"},
			{Name: "pole denominator and residue sign remain unselected", Passed: !a.PoleDenominatorDerived && !a.ResidueSignDerived && a.DenominatorAmbiguityWitnessed, Detail: FormatDenominators(a.DenominatorWitnesses)},
			{Name: "no physical/ghost/constrained/vacuum contact class selected", Passed: !a.PhysicalPositiveNormContactPropagatorDerived && !a.RegulatorGhostContactClassDerived && !a.ConstrainedContactClassDerived && !a.VacuumFrustrationContactClassDerived && !a.ContactPropagatorClassDerived, Detail: FormatBranches(a.BranchWitnesses)},
			{Name: "propagator classifier attempts audited", Passed: len(a.TestAttempts) >= 6, Detail: FormatAttempts(a.TestAttempts)},
			{Name: "threshold beta tensor remains sealed", Passed: a.BetaCorrectionRowsAllowed == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived, Detail: fmt.Sprintf("beta rows allowed=%d", a.BetaCorrectionRowsAllowed)},
			{Name: "residual nullity unchanged", Passed: a.ResidualNullityAfter == a.ResidualNullityBefore && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken, Detail: fmt.Sprintf("nullity before=%d; after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)},
			{Name: "physical predictions remain sealed", Passed: !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, observed scales, or fitted thresholds are used"},
		}, Notes: []string{a.TruthStatement, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
