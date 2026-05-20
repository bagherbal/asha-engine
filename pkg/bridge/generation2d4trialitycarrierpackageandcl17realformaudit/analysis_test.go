package generation2d4trialitycarrierpackageandcl17realformaudit

import (
	"strings"
	"testing"
)

func TestGate800BoardArithmetic(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Board.N != 8 || a.Board.DifferenceMod8 != 2 || a.Board.VolumeSquare != -1 {
		t.Fatalf("bad Cl(1,7) arithmetic: %s", FormatBoard(a.Board))
	}
	if a.Board.RealChiralityProjectorsExist || a.Board.MinimalRealSpinorDim != 16 || a.Board.ComplexHalfSpinorDimC != 8 {
		t.Fatalf("bad chirality audit: %s", FormatBoard(a.Board))
	}
}

func TestGate800RealFormAndOutcome(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ComplexD4.Recorded || !a.ComplexD4.LawfulComplexCandidate || a.ComplexD4.RealNative {
		t.Fatalf("bad complex D4 candidate: %+v", a.ComplexD4)
	}
	if !a.RealForm.Defined || !strings.Contains(a.RealForm.PreservesRealForm, "not certified") || !containsAll(a.RealForm.Failures, []string{StatusNoNativeUnlessPreserves, StatusComplexNotNative}) {
		t.Fatalf("bad real form test: %s", FormatRealForm(a.RealForm))
	}
	if a.Outcome.FullNativeFound || !strings.Contains(a.Outcome.Selected, "Outcome C") || !containsAll(a.Outcome.Supports, []string{StatusComplexOnlyTriality, StatusTrialityNeedsAirlock}) {
		t.Fatalf("bad outcome: %+v", a.Outcome)
	}
}

func TestGate800FirewallsAndExistingObjects(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	existing := FormatExisting(a.Existing)
	for _, want := range []string{StatusK7Not8, StatusK7HodgeNotCarrier, StatusLambda4NotModule, StatusAggregateTracesNotCarrier} {
		if !strings.Contains(existing, want) {
			t.Fatalf("missing existing-object failure %s in %s", want, existing)
		}
	}
	if !a.YukawaFirewall.Defined || !a.YukawaFirewall.NotYukawaTheorem || !containsAll(a.YukawaFirewall.Failures, []string{StatusTrialityCarrierNotYukawa, StatusNoTraceAtomReadout, StatusNoPMNSCKMReadout}) {
		t.Fatalf("bad Yukawa firewall: %+v", a.YukawaFirewall)
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NoYukawa || !a.Firewalls.NoNEff || !a.Firewalls.NoPoleMass {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate800TheoremStatusesAndFinal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Final, "does not use triality") || !strings.Contains(a.Final, "Outcome C") || !strings.Contains(a.Final, "RealFormAirlock") {
		t.Fatalf("bad final: %s", a.Final)
	}
	res := Generation2D4TrialityCarrierPackageAndCL17RealFormAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
