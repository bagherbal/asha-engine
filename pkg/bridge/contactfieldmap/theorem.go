package contactfieldmap

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactOverlapLocalFieldMapConstraintBRSTClassifierTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-OVERLAP-LOCAL-FIELD-MAP-CONSTRAINT-BRST-CLASSIFIER"
	const name = "contact-overlap local field map / constraint-BRST classifier search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact local-field/BRST classifier", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 110 propagator obstruction inherited", Passed: a.ContactPropagator.ContactRows == 7 && a.ContactPropagator.PositiveFiniteContactRows == 7 && !a.ContactPropagator.ContactPropagatorClassDerived && a.ContactPropagator.ResidualNullityAfter == 3, Detail: fmt.Sprintf("contact=%d; positive=%d; nullity=%d", a.ContactPropagator.ContactRows, a.ContactPropagator.PositiveFiniteContactRows, a.ContactPropagator.ResidualNullityAfter)},
			{Name: "seven contact modes remain positive finite-overlap rows", Passed: a.ContactRows == 7 && a.PositiveFiniteContactRows == 7, Detail: FormatRows(a.Rows, 12)},
			{Name: "local field map candidate constructed but not derived", Passed: a.LocalFieldMapCandidateConstructed && !a.LocalCoordinateDerived && !a.SpacetimeSupportDerived && !a.InvertibleFieldMapDerived, Detail: FormatLocality(a.LocalityCriteria)},
			{Name: "contact Lorentz kinetic action and pole/residue absent", Passed: !a.LorentzKineticOperatorDerived && !a.PoleResidueForContactDerived && !a.CanonicalNormalizationForContactDerived, Detail: "positive finite overlap is still not a Lorentzian local quadratic action or pole/residue theorem"},
			{Name: "contact gauge representation rows absent", Passed: !a.GaugeRepresentationForContactDerived, Detail: "no SU(3)c×SU(2)L×U(1)Y representation row is selected for partial-overlap modes"},
			{Name: "constraint/BRST classifier constructed but not derived", Passed: a.ConstraintClassifierConstructed && !a.ConstraintGeneratorDerived && !a.GhostGradingDerived && !a.NilpotentBRSTDerived && !a.BRSTPairingDerived && !a.SupertraceCancellationDerived, Detail: FormatBRST(a.BRSTCriteria)},
			{Name: "no physical/local, constrained, regulator, or vacuum class selected", Passed: !a.PhysicalLocalContactFieldsDerived && !a.ConstrainedContactClassDerived && !a.RegulatorGhostContactClassDerived && !a.VacuumFrustrationContactClassDerived && !a.ContactFieldClassDerived, Detail: FormatBranches(a.BranchWitnesses)},
			{Name: "threshold beta permission remains zero", Passed: a.BetaCorrectionRowsAllowed == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived, Detail: fmt.Sprintf("beta rows allowed=%d", a.BetaCorrectionRowsAllowed)},
			{Name: "residual nullity unchanged", Passed: a.ResidualNullityAfter == a.ResidualNullityBefore && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken, Detail: fmt.Sprintf("nullity before=%d; after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)},
			{Name: "physical predictions remain sealed", Passed: !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, boundary scale, or observed thresholds are used"},
		}, Notes: []string{a.TruthStatement, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
