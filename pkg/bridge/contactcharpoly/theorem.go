package contactcharpoly

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ExactContactOverlapCharacteristicPolynomialSymbolicNumberFieldAttemptTheorem() theorem.Theorem {
	const id = "BRIDGE-EXACT-CONTACT-OVERLAP-CHARPOLY-SYMBOLIC-NUMBER-FIELD-ATTEMPT"
	const name = "Exact contact overlap characteristic polynomial / symbolic number-field construction attempt"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact characteristic-polynomial attempt", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 147 algebraic-origin firewall is inherited", Passed: a.Previous.BetaPermissionFirewallClosed && a.Previous.ContactRows == 7 && a.Previous.ExactNumberFieldLifts == 0 && a.Previous.ChargeSemanticRows == 0 && a.ContactBetaRowsAllowed == 0, Detail: fmt.Sprintf("previous beta=%d exactFieldLifts=%d", a.Previous.ContactBetaRowsAllowed, a.Previous.ExactNumberFieldLifts)},
			{Name: "rational characteristic-polynomial candidate covers all seven partial rows", Passed: a.Candidate.CandidateFactorizationRecognized && a.CandidateCoveredRows == 7 && a.Candidate.MaxPartialResidual < 1e-8, Detail: FormatCandidate(a.Candidate)},
			{Name: "three rational factors and one quartic candidate are isolated", Passed: a.RationalFactorRows == 3 && a.QuarticCandidateRows == 4 && a.CandidateNumberFieldDegree == 4 && a.Candidate.MaxQuarticResidual < 1e-8, Detail: FormatRows(a.Rows)},
			{Name: "symbolic candidate is not promoted to an exact determinant certificate", Passed: !a.Candidate.ExactMatrixOverNumberField && !a.Candidate.ExactDeterminantComputed && !a.Candidate.ExactCharacteristicCertified && !a.Candidate.RowMinimalPolynomialsCertified && !a.ExactCharacteristicCertified && !a.ExactNumberFieldLiftCertified && a.ExactCertifiedRows == 0 && a.ExactCharacteristicProofs == 0, Detail: FormatCandidate(a.Candidate)},
			{Name: "exact construction requirements remain unsatisfied", Passed: a.Requirements.RationalReconstructionCandidate && a.Requirements.ResidualCheckPassed && !a.Requirements.ExactOverlapMatrix && !a.Requirements.ExactDeterminant && !a.Requirements.IndependentCertificate && !a.Requirements.RootIsolationCertificate && !a.Requirements.RowwiseRootAssignmentProof && !a.Requirements.AlgebraicRowSemantics && !a.Requirements.ChargeOperatorSelected && !a.Requirements.RepresentationRowsSelected && a.Requirements.ObservedInputFree && !a.Requirements.AllSatisfied, Detail: FormatRequirements(a.Requirements)},
			{Name: "charpoly candidate opens no contact charge, representation, beta, or physical constants", Passed: a.ChargeSemanticRows == 0 && a.T3RRowsDerived == 0 && a.BMinusLRowsDerived == 0 && a.HyperchargeRowsDerived == 0 && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}}
	}}
}
