package contactalgebraic

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactIrrationalSpectrumAlgebraicOriginMinimalPolynomialObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-IRRATIONAL-SPECTRUM-ALGEBRAIC-ORIGIN-MINIMAL-POLYNOMIAL-OBSTRUCTION"
	const name = "Contact irrational-spectrum algebraic-origin / minimal-polynomial obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact algebraic-origin obstruction", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 146 charge-lattice firewall is inherited", Passed: a.Previous.BetaPermissionFirewallClosed && a.Previous.ContactRows == 7 && a.Previous.ChargeSemanticEmbeddings == 0 && a.ContactBetaRowsAllowed == 0, Detail: fmt.Sprintf("previous beta=%d semanticEmbeddings=%d", a.Previous.ContactBetaRowsAllowed, a.Previous.ChargeSemanticEmbeddings)},
			{Name: "finite contact overlap operator is real but only numerically diagonalized", Passed: a.MatrixOrigin.FiniteSymmetricOperator && a.MatrixOrigin.NumericEigenDecomposition && a.MatrixOrigin.AmbientOverlapDimension == 14 && a.MatrixOrigin.ContactMultiplicity == 7 && a.MatrixOrigin.PartialRows == 7 && !a.MatrixOrigin.ExactMatrixOverNumberField && !a.MatrixOrigin.CharacteristicPolynomialExact && !a.MatrixOrigin.RowMinimalPolynomialsExact, Detail: FormatMatrixOrigin(a.MatrixOrigin)},
			{Name: "three degree-one rational diagnostics are recognized", Passed: a.RationalRows == 3 && a.DegreeOneCertifiedRows == 3 && a.ExactMinimalPolynomialRows == 3, Detail: FormatRows(a.Rows)},
			{Name: "four non-rational-looking rows remain numerical algebraic candidates", Passed: a.NonRationalCandidateRows == 4 && a.NonDegreeOneCertifiedRows == 4 && a.NumericOnlyRows == 4 && a.ExactNumberFieldLifts == 0, Detail: FormatRows(a.Rows)},
			{Name: "minimal-polynomial requirements are not satisfied", Passed: !a.Requirements.ExactNumberFieldBasis && !a.Requirements.ExactCharacteristicPolynomial && !a.Requirements.RowwiseMinimalPolynomials && !a.Requirements.AlgebraicRowSemantics && !a.Requirements.ChargeLatticeSelected && !a.Requirements.OperatorPullback && !a.Requirements.LocalFieldMap && !a.Requirements.MassActivation && !a.Requirements.DecouplingRule && a.Requirements.ObservedInputFree && !a.Requirements.AllSatisfied, Detail: FormatRequirements(a.Requirements)},
			{Name: "algebraic-origin audit opens no contact charge, representation, beta, or physical constants", Passed: a.ChargeSemanticRows == 0 && a.T3RRowsDerived == 0 && a.BMinusLRowsDerived == 0 && a.HyperchargeRowsDerived == 0 && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}}
	}}
}
