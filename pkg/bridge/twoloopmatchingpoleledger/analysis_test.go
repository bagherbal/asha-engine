package twoloopmatchingpoleledger

import (
	"math"
	"testing"
)

func TestGate309Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Inheritance.InheritedAsDiagnostic || !a.Inheritance.OneLoopOnly || !a.Inheritance.PureSMHighScaleRejected || !a.Inheritance.ThresholdMatchingOmitted || !a.Inheritance.PoleMassMatchingOmitted || a.Inheritance.FinalColliderMassClaimed {
		t.Fatalf("bad inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if math.Abs(a.Inheritance.UVLambda-1197.0/4624.0) > 1e-15 {
		t.Fatalf("bad UV lambda: %s", FormatInheritance(a.Inheritance))
	}
	if a.Inheritance.PrimaryMassGeV < 330 || a.Inheritance.PrimaryMassGeV > 333 {
		t.Fatalf("unexpected Gate 309 mass: %s", FormatInheritance(a.Inheritance))
	}
}

func TestTwoLoopLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.TwoLoop.Formalized || !a.TwoLoop.PositiveBetaSoftensDownward || !a.TwoLoop.NegativeBetaAmplifiesDownward || len(a.TwoLoop.RepresentativeTerms) < 4 {
		t.Fatalf("bad two-loop ledger: %s", FormatTwoLoop(a.TwoLoop))
	}
	if a.TwoLoop.ExactFullSystemInstalled || a.TwoLoop.TwoLoopIntegrationExecuted || a.TwoLoop.CanResolveAlone {
		t.Fatalf("two-loop ledger overclaimed: %s", FormatTwoLoop(a.TwoLoop))
	}
}

func TestThresholdAndPoleLedgers(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Thresholds.Formalized || !a.Thresholds.HasCapacityToResolve || a.Thresholds.ValuesDerived || a.Thresholds.ThresholdsExecuted || len(a.Thresholds.Sources) < 3 {
		t.Fatalf("bad threshold ledger: %s", FormatThresholds(a.Thresholds))
	}
	if !a.PoleMass.Formalized || a.PoleMass.CanResolveAlone || a.PoleMass.SelfEnergiesComputed || a.PoleMass.UsesMeasuredMassForFit {
		t.Fatalf("bad pole ledger: %s", FormatPole(a.PoleMass))
	}
	if a.PoleMass.RequiredLambdaAtV < 0.12 || a.PoleMass.RequiredLambdaAtV > 0.14 {
		t.Fatalf("unexpected reference lambda: %s", FormatPole(a.PoleMass))
	}
	if a.PoleMass.RequiredLambdaShiftAtV > -0.7 {
		t.Fatalf("expected large negative lambda shift: %s", FormatPole(a.PoleMass))
	}
}

func TestTensionCapacityAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Tension.Formalized || a.Tension.TwoLoopCanResolveAlone || a.Tension.PoleMassCanResolveAlone || !a.Tension.ThresholdsCanResolveInPrinciple || !a.Tension.ModifiedTopSectorMayBeRequired || !a.Tension.NeedsFullPrecisionRun || a.Tension.FinalMassResolved {
		t.Fatalf("bad tension audit: %s", FormatTension(a.Tension))
	}
	if !a.Firewalls.NoTwoLoopNumericalTransportRun || !a.Firewalls.NoThresholdJumpInserted || !a.Firewalls.NoPoleSelfEnergyInserted || !a.Firewalls.NoObservedHiggsUsedAsFit || !a.Firewalls.NoObservedTopUsedAsFit || !a.Firewalls.NoFinalMassClaimed || !a.Firewalls.NoFiniteCorePolluted {
		t.Fatalf("firewall failed: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := TwoLoopMatchingPoleMassConversionLedgerAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
