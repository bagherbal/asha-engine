package contactcharpoly

import "testing"

func TestGate148ContactCharpolyCandidateObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.ContactRows != 7 || len(a.Rows) != 7 || a.CandidateCoveredRows != 7 {
		t.Fatalf("expected seven covered contact rows, got %+v", a.Summary)
	}
	if !a.Candidate.CandidateFactorizationRecognized || a.Candidate.PartialDegree != 7 || a.Candidate.FullDegree != 14 {
		t.Fatalf("candidate factorization should be recognized: %+v", a.Candidate)
	}
	if a.RationalFactorRows != 3 || a.QuarticCandidateRows != 4 || a.CandidateNumberFieldDegree != 4 {
		t.Fatalf("expected 3 rational factors and 4 quartic rows: %+v", a.Summary)
	}
	if a.Candidate.MaxPartialResidual > 1e-8 || a.Candidate.MaxQuarticResidual > 1e-8 {
		t.Fatalf("candidate residuals too large: %+v", a.Candidate)
	}
	if a.Candidate.ExactMatrixOverNumberField || a.Candidate.ExactDeterminantComputed || a.Candidate.ExactCharacteristicCertified || a.Candidate.RowMinimalPolynomialsCertified {
		t.Fatalf("candidate should not be exact-certified: %+v", a.Candidate)
	}
	if a.Requirements.AllSatisfied || a.Requirements.ExactOverlapMatrix || a.Requirements.ExactDeterminant || a.Requirements.IndependentCertificate || a.Requirements.RootIsolationCertificate || a.Requirements.RowwiseRootAssignmentProof || a.Requirements.AlgebraicRowSemantics {
		t.Fatalf("exact construction requirements should remain unsatisfied: %+v", a.Requirements)
	}
	if !a.BetaPermissionFirewallClosed || a.ContactBetaRowsAllowed != 0 || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ChargeSemanticRows != 0 {
		t.Fatalf("firewall should remain closed: %+v", a.Summary)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("hidden observed physics leaked into Gate 148")
	}
}
