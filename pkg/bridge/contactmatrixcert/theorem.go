package contactmatrixcert

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ExactRationalContactOverlapMatrixLiftDeterminantCertificateSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-EXACT-RATIONAL-CONTACT-OVERLAP-MATRIX-DETERMINANT-CERTIFICATE"
	const name = "Exact rational contact-overlap matrix lift / determinant certificate search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build exact rational contact-overlap certificate", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 148 candidate is inherited but not treated as physical input", Passed: a.Previous.BetaPermissionFirewallClosed && a.Previous.CandidateCoveredRows == 7 && a.Previous.ExactCharacteristicProofs == 0 && !a.Previous.ExactCharacteristicCertified && !a.Previous.ExactNumberFieldLiftCertified && a.Previous.ContactBetaRowsAllowed == 0, Detail: fmt.Sprintf("prevCovered=%d prevExact=%t prevBeta=%d", a.Previous.CandidateCoveredRows, a.Previous.ExactCharacteristicCertified, a.Previous.ContactBetaRowsAllowed)},
			{Name: "Boolean Gram inverse and G2 raw calibration are exact finite inputs", Passed: a.MatrixLift.RawG2GramIsFourIdentity && a.MatrixLift.BooleanGramInverseClosed, Detail: FormatGramInverse(a.GramInverse)},
			{Name: "exact rational contact-overlap matrix is lifted", Passed: a.ExactRationalOverlapMatrix && a.MatrixLift.RationalMatrixBuilt && a.MatrixLift.Symmetric && a.MatrixLift.Rows == 14 && a.MatrixLift.Cols == 14 && a.MatrixLift.MaxFloatResidual < 1e-10, Detail: FormatMatrixLift(a.MatrixLift)},
			{Name: "determinant and characteristic polynomial are certified exactly", Passed: a.ExactDeterminantComputed && a.ExactCharacteristicCertified && a.ExactAnnihilationCertified && a.Certificate.ExactCharpolyComputed && a.Certificate.CandidateCharpolyMatches && a.Certificate.CandidatePolynomialAnnihilatesMatrix && a.Certificate.TraceMatches && a.Certificate.DeterminantMatches && a.Certificate.Degree == 14 && a.Certificate.UnitEigenMultiplicity == 7 && a.Certificate.PartialDegree == 7 && a.Certificate.QuarticFactorCertified, Detail: FormatCertificate(a.Certificate)},
			{Name: "exact certificate still does not provide root isolation or row semantics", Passed: a.Requirements.ExactRationalOverlapMatrix && a.Requirements.ExactCharacteristicPolynomial && a.Requirements.ExactDeterminantCertificate && a.Requirements.ExactAnnihilationCertificate && !a.Requirements.RootIsolationCertificate && !a.Requirements.RowwiseRootAssignmentProof && !a.Requirements.ChargeOperatorSelected && !a.Requirements.RepresentationRowsSelected && a.Requirements.ObservedInputFree && !a.Requirements.AllSatisfiedForPhysics, Detail: FormatRequirements(a.Requirements)},
			{Name: "contact beta and physical constants remain sealed", Passed: a.ContactRows == 7 && a.RootIsolationCertificates == 0 && a.RowwiseRootAssignmentProofs == 0 && a.ChargeSemanticRows == 0 && a.T3RRowsDerived == 0 && a.BMinusLRowsDerived == 0 && a.HyperchargeRowsDerived == 0 && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && a.BetaPermissionFirewallClosed && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}}
	}}
}
