package contactsignsource

import "testing"

func TestGate142ContactSignOrientationSource(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.ContactRows != 7 || a.LargestGapHighRows != 3 || a.LargestGapLowRows != 4 || a.OrientationCandidates != 2 {
		t.Fatalf("unexpected inherited split: %+v", a.Summary)
	}
	if a.OrientationSourcesAudited != 7 || a.SourcesAvailable != 3 {
		t.Fatalf("unexpected source audit counts: %+v", a.Summary)
	}
	if !a.Z2OrientationDegeneracy || a.ChargeConjugationInvolutions != 1 || a.ChargeConjugationSelectedBranches != 0 || a.SourcesSelectingOrientation != 0 {
		t.Fatalf("charge conjugation should prove degeneracy, not selection: %+v", a.ChargeConjugation)
	}
	if a.T3RPullbackRowsDerived != 0 || a.ChiralityPullbackRowsDerived != 0 || a.BMinusLPullbackRowsDerived != 0 || a.SU2LPullbackRowsDerived != 0 || a.HyperchargeRowsDerived != 0 {
		t.Fatalf("no contact charge rows should be derived: %+v", a.Summary)
	}
	if !a.BetaPermissionFirewallClosed || a.ContactBetaRowsAllowed != 0 || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 {
		t.Fatalf("firewall should remain closed: %+v", a.Summary)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("hidden observed physics leaked into Gate 142")
	}
}
