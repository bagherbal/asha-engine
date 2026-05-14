package contactorientation

import "testing"

func TestGate141OrientationEnumeration(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.ContactRows != 7 || a.LargestGapHighRows != 3 || a.LargestGapLowRows != 4 || a.SplitPattern != "3|4" {
		t.Fatalf("expected inherited 3|4 contact split: %+v", a.Summary)
	}
	if a.OrientationCandidates != 2 || a.SpectrallyMonotoneOrientations != 2 {
		t.Fatalf("expected two compatible orientations: %+v candidates=%+v", a.Summary, a.Candidates)
	}
	if a.SelectedOrientations != 0 || a.T3RSemanticOrientations != 0 {
		t.Fatalf("orientation must not be selected or promoted to T3R: %+v", a.Summary)
	}
}

func TestGate141TraceAndFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.TracelessOrientations != 0 || a.PureHalfSignTraceMagnitudeNumerator != 1 || a.PureHalfSignTraceMagnitudeDenom != 2 {
		t.Fatalf("expected nontraceless +/- half-sign orientations: %+v candidates=%+v", a.Summary, a.Candidates)
	}
	if a.FockContactIntertwiners != 0 || a.T3RPullbackRowsDerived != 0 || a.ChiralityPullbackRowsDerived != 0 || a.BMinusLPullbackRowsDerived != 0 || a.SU2LPullbackRowsDerived != 0 || a.HyperchargeRowsDerived != 0 {
		t.Fatalf("unexpected contact operator pullback/hypercharge rows: %+v", a.Summary)
	}
	if !a.BetaPermissionFirewallClosed || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("contact beta firewall should remain closed: %+v", a.Summary)
	}
	if a.ResidualS6Choices != 720 || a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 {
		t.Fatalf("residual obstruction changed unexpectedly: %+v", a.Summary)
	}
}
