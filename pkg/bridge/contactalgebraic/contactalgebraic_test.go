package contactalgebraic

import "testing"

func TestGate147ContactAlgebraicOriginObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.ContactRows != 7 || len(a.Rows) != 7 || len(a.PartialValues) != 7 {
		t.Fatalf("expected seven contact rows, got summary=%+v", a.Summary)
	}
	if !a.MatrixOrigin.FiniteSymmetricOperator || !a.MatrixOrigin.NumericEigenDecomposition || a.MatrixOrigin.AmbientOverlapDimension != 14 || a.MatrixOrigin.ContactMultiplicity != 7 || a.MatrixOrigin.PartialRows != 7 {
		t.Fatalf("unexpected matrix origin audit: %+v", a.MatrixOrigin)
	}
	if a.MatrixOrigin.ExactMatrixOverNumberField || a.MatrixOrigin.CharacteristicPolynomialExact || a.MatrixOrigin.RowMinimalPolynomialsExact {
		t.Fatalf("exact symbolic matrix/charpoly/minpoly should not be certified: %+v", a.MatrixOrigin)
	}
	if a.RationalRows != 3 || a.DegreeOneCertifiedRows != 3 || a.ExactMinimalPolynomialRows != 3 {
		t.Fatalf("expected exactly three rational degree-one diagnostics: %+v", a.Summary)
	}
	if a.NonRationalCandidateRows != 4 || a.NonDegreeOneCertifiedRows != 4 || a.NumericOnlyRows != 4 || a.ExactNumberFieldLifts != 0 {
		t.Fatalf("expected four numerical algebraic candidates and no exact field lift: %+v", a.Summary)
	}
	if a.Requirements.AllSatisfied || a.Requirements.ExactNumberFieldBasis || a.Requirements.ExactCharacteristicPolynomial || a.Requirements.RowwiseMinimalPolynomials || a.Requirements.AlgebraicRowSemantics || a.Requirements.ChargeLatticeSelected || a.Requirements.OperatorPullback || a.Requirements.LocalFieldMap {
		t.Fatalf("minimal-polynomial requirements should remain unsatisfied: %+v", a.Requirements)
	}
	if !a.BetaPermissionFirewallClosed || a.ContactBetaRowsAllowed != 0 || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ChargeSemanticRows != 0 {
		t.Fatalf("firewall should remain closed: %+v", a.Summary)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("hidden observed physics leaked into Gate 147")
	}
}
