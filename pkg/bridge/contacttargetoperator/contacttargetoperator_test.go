package contacttargetoperator

import "testing"

func TestGate139ContactTargetOperatorObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.ContactRows != 7 || a.ContactSpectralRows != 7 || a.ContactSpectralDistinctRows != 7 {
		t.Fatalf("unexpected contact spectrum data: %+v", a.Summary)
	}
	if a.DiagnosticContactOperators < 1 {
		t.Fatalf("expected at least one canonical diagnostic contact operator: %+v", a.Candidates)
	}
	if a.CanonicalT3RTargetOperators != 0 || a.QuotientInducedT3RTargetOperators != 0 {
		t.Fatalf("no contact T3R target operator should be derived: %+v", a.Summary)
	}
}

func TestGate139T3RSplitAndFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.AbstractT3RMultiplicitySplits != 8 || a.T3RRowSignAssignments != 128 || a.NonScalarT3RRowSignAssignments != 126 {
		t.Fatalf("unexpected T3R split/sign audit: %+v", a.Summary)
	}
	if a.FullOperatorIntertwinersDerived != 0 || a.T3RPullbackRowsDerived != 0 || a.ChiralityPullbackRowsDerived != 0 || a.HyperchargeRowsDerived != 0 {
		t.Fatalf("unexpected quotient-side pullback rows: %+v", a.Summary)
	}
	if !a.BetaPermissionFirewallClosed || a.RepresentationCompleteRows != 0 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("contact beta firewall should remain closed: %+v", a.Summary)
	}
	if a.ResidualS6Choices != 720 || a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 {
		t.Fatalf("residual obstruction changed unexpectedly: %+v", a.Summary)
	}
}
