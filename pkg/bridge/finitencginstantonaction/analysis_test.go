package finitencginstantonaction

import "testing"

func TestGate286FiniteNCGInstantonAudit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	if !a.Gate285.ContinuumRouteAudited || !a.Gate285.FiniteConnectionMissing || !a.Gate285.IntermediateSealRequired {
		t.Fatalf("bad Gate285 snapshot: %s", FormatGate285(a.Gate285))
	}
	if !a.Calculus.AlgebraicMatrixRouteDefined || a.Calculus.RequiresContinuumForms || a.Calculus.RequiresIntegrationMeasure {
		t.Fatalf("NCG calculus should be finite-matrix route: %s", FormatCalculus(a.Calculus))
	}
	if a.Calculus.PhysicalDFDerived || a.Calculus.FullAlgebraRepresentation {
		t.Fatalf("physical D_F/full representation should not be claimed: %s", FormatCalculus(a.Calculus))
	}
	if !a.Diagnostic.NonVacuousOneForm || a.Diagnostic.OneFormNorm2 <= 0 || !a.Diagnostic.FiniteTraceActionComputed {
		t.Fatalf("expected non-vacuous local diagnostic: %s", FormatDiagnostic(a.Diagnostic))
	}
	if a.Saddle.NontrivialRealSaddleExists || a.Saddle.NontrivialActionGapDerived {
		t.Fatalf("nontrivial saddle should not be derived: %s", FormatSaddle(a.Saddle))
	}
	if a.BGapAudit.TreatingBGapAsMajoranaDerived || a.BGapAudit.TreatingBGapAsInverseDerived || a.BGapAudit.ProducesInverseBGap || a.BGapAudit.ProducesFourOverPi {
		t.Fatalf("B_gap route should remain sealed: %s", FormatBGapAudit(a.BGapAudit))
	}
	if !a.Firewalls.DoesNotUseContinuumForms || !a.Firewalls.DoesNotPromoteBGapToCoupling || !a.Firewalls.DoesNotClaimFourOverPiSaddle || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
	if a.Summary.FiniteInstantonDerived || a.Summary.IntermediateSealGranted || a.Summary.FourOverPiGenerated {
		t.Fatalf("summary should not promote theorem: %s", FormatSummary(a.Summary))
	}
}

func TestGate286TheoremPassesChecks(t *testing.T) {
	res := FiniteSpectralActionSaddlePointBGapInstantonActionAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem checks failed:\n%s", res.Details())
	}
	if res.Status != "BRIDGE_REQUIRED" {
		t.Fatalf("Gate 286 should remain BridgeRequired, got %s", res.Status)
	}
}
