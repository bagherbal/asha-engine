package contactsignsplit

import "testing"

func TestGate140SpectralCutAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.ContactRows != 7 || a.DistinctSpectralRows != 7 {
		t.Fatalf("unexpected contact rows: %+v", a.Summary)
	}
	if a.ProperSpectralCuts != 6 || a.SpectralCutSignAssignments != 12 || a.AbstractSignAssignments != 128 {
		t.Fatalf("unexpected cut/sign count: %+v", a.Summary)
	}
	if a.UniqueLargestGapCuts != 1 || a.LargestGapHighRows != 3 || a.LargestGapLowRows != 4 || a.CanonicalDiagnosticSplits != 1 {
		t.Fatalf("expected unique diagnostic 3|4 largest-gap split: %+v largest=%+v", a.Summary, a.LargestGap)
	}
}

func TestGate140Firewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.OrientationSelectedSplits != 0 || a.T3RSemanticSplits != 0 || a.T3RTargetOperatorsDerived != 0 || a.QuotientInducedT3RTargetOps != 0 {
		t.Fatalf("spectral cut must not become T3R semantics: %+v", a.Summary)
	}
	if a.FullOperatorIntertwiners != 0 || a.T3RPullbackRowsDerived != 0 || a.ChiralityPullbackRowsDerived != 0 || a.HyperchargeRowsDerived != 0 {
		t.Fatalf("unexpected contact pullback or hypercharge rows: %+v", a.Summary)
	}
	if !a.BetaPermissionFirewallClosed || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("contact beta firewall should remain closed: %+v", a.Summary)
	}
	if a.ResidualS6Choices != 720 || a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 {
		t.Fatalf("residual obstruction changed unexpectedly: %+v", a.Summary)
	}
}
