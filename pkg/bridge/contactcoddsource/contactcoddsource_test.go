package contactcoddsource

import "testing"

func TestGate144ContactCoddSignedCurrentConstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.ContactRows != 7 || a.LargestGapHighRows != 3 || a.LargestGapLowRows != 4 || a.OrientationCandidates != 2 {
		t.Fatalf("unexpected inherited split: %+v", a.Summary)
	}
	if !a.CenteredFunctional.CanonicalAsDiagnostic || a.CenteredFunctional.PhysicalCoddSource {
		t.Fatalf("centered spectral current should be diagnostic only: %+v", a.CenteredFunctional)
	}
	if a.CenteredPositiveRows != 3 || a.CenteredNegativeRows != 4 || a.CenteredZeroRows != 0 || !a.CenteredFunctional.MatchesLargestGap {
		t.Fatalf("unexpected centered-sign split: %+v", a.CenteredFunctional)
	}
	if a.CenteredFunctional.Trace > 1e-9 || a.CenteredFunctional.Trace < -1e-9 {
		t.Fatalf("centered functional should be trace-zero, trace=%g", a.CenteredFunctional.Trace)
	}
	if a.SignedSourcesAudited != 7 || a.AvailableSignedDiagnostics != 3 || a.TraceZeroDiagnostics != 2 || a.CanonicalSignedDiagnostics != 1 {
		t.Fatalf("unexpected signed source audit counts: %+v", a.Summary)
	}
	if a.CoddContactFunctionals != 0 || a.CBreakingSources != 0 || a.SourcesSelectingPhysicalSign != 0 {
		t.Fatalf("no physical C-odd source should be derived: %+v", a.Summary)
	}
	if a.BinaryT3RRowsDerived != 0 || a.T3RPullbackRowsDerived != 0 || a.BMinusLPullbackRowsDerived != 0 || a.HyperchargeRowsDerived != 0 {
		t.Fatalf("no contact charge rows should be derived: %+v", a.Summary)
	}
	if !a.BetaPermissionFirewallClosed || a.ContactBetaRowsAllowed != 0 || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 {
		t.Fatalf("firewall should remain closed: %+v", a.Summary)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("hidden observed physics leaked into Gate 144")
	}
}
