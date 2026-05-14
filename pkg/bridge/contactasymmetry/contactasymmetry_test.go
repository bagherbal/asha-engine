package contactasymmetry

import "testing"

func TestGate143ContactCBreakingAsymmetrySearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.ContactRows != 7 || a.LargestGapHighRows != 3 || a.LargestGapLowRows != 4 || a.OrientationCandidates != 2 {
		t.Fatalf("unexpected inherited split: %+v", a.Summary)
	}
	if !a.Z2OrientationDegeneracy || a.ChargeConjugationBroken {
		t.Fatalf("charge-conjugation degeneracy should remain unbroken: %+v", a.Summary)
	}
	if a.CardinalityImbalance != 1 || a.AsymmetryDiagnostics != 2 || a.CInvariantDiagnostics != 2 || a.CBreakingDiagnostics != 0 {
		t.Fatalf("unexpected asymmetry diagnostic audit: %+v", a.SplitAudit)
	}
	if a.AsymmetrySourcesAudited != 7 || a.AsymmetrySourcesAvailable != 2 || a.CBreakingSources != 0 || a.SourcesSelectingOrientation != 0 || a.CoddContactFunctionals != 0 {
		t.Fatalf("unexpected source audit counts: %+v", a.Summary)
	}
	if a.T3RPullbackRowsDerived != 0 || a.ChiralityPullbackRowsDerived != 0 || a.BMinusLPullbackRowsDerived != 0 || a.SU2LPullbackRowsDerived != 0 || a.HyperchargeRowsDerived != 0 {
		t.Fatalf("no contact charge rows should be derived: %+v", a.Summary)
	}
	if !a.BetaPermissionFirewallClosed || a.ContactBetaRowsAllowed != 0 || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 {
		t.Fatalf("firewall should remain closed: %+v", a.Summary)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("hidden observed physics leaked into Gate 143")
	}
}
