package fockcontactkernel

import "testing"

func TestGate138KernelSelectionObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.MatterDimension != 16 || a.ContactRows != 7 || a.RequiredKernelDim != 9 || a.UnconstrainedGrassmannDimension != 63 {
		t.Fatalf("unexpected quotient/kernel data: %+v", a.Summary)
	}
	if a.T3RInvariantSplitPatterns != 8 || a.Summary.T3RInvariantResidualDimMin != 7 || a.Summary.T3RInvariantResidualDimMax != 31 {
		t.Fatalf("unexpected T3R split audit: %+v", a.Summary)
	}
	if a.T3RChiralitySplitPatterns != 80 || a.Summary.T3RChiralityResidualDimMin != 3 || a.Summary.T3RChiralityResidualDimMax != 15 {
		t.Fatalf("unexpected joint split audit: %+v", a.Summary)
	}
	if a.FullOperatorIntertwinersDerived != 0 || a.TargetContactOperatorsDerived != 0 || a.CanonicalKernelCandidates != 1 {
		t.Fatalf("operator-intertwiner obstruction should remain: %+v", a.Summary)
	}
}

func TestGate138FirewallRemainsClosed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.T3RPullbackRowsDerived != 0 || a.ChiralityPullbackRowsDerived != 0 || a.BMinusLPullbackRowsDerived != 0 || a.SU2LPullbackRowsDerived != 0 || a.HyperchargeRowsDerived != 0 {
		t.Fatalf("unexpected operator pullback rows: %+v", a.Summary)
	}
	if !a.BetaPermissionFirewallClosed || a.RepresentationCompleteRows != 0 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("contact beta firewall should remain closed: %+v", a.Summary)
	}
	if a.ResidualS6Choices != 720 || a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 {
		t.Fatalf("residual obstruction changed unexpectedly: %+v", a.Summary)
	}
}
