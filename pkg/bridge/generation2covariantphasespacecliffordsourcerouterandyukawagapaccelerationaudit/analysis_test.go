package generation2covariantphasespacecliffordsourcerouterandyukawagapaccelerationaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate824RoutesClaimsWithoutPromotion(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Router.Claims) != 12 {
		t.Fatalf("expected 12 routed claims, got %d", len(a.Router.Claims))
	}
	claims := FormatClaims(a.Router.Claims)
	for _, fragment := range []string{"Cl(1,7) phase-space board", "SO(8)/D4 triality", "Higgs mass bridge", "1+3 rest simplex"} {
		if !strings.Contains(claims, fragment) {
			t.Fatalf("claim table missing %q:\n%s", fragment, claims)
		}
	}
	if a.BoundaryFN.CanMoveBeyondPartialR2 {
		t.Fatalf("boundary-FN should not move beyond partial R2 without map")
	}
}

func TestGate824RealFormChiralityFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.RealForm.OmegaSquared != "-1" || a.RealForm.NaiveProjectorsIdempotent || !a.RealForm.RequiresAirlock {
		t.Fatalf("bad real form precheck: %+v", a.RealForm)
	}
	if !containsAll(a.RealForm.Failures, []string{FailureNaiveRealChirality, FailureChiralityAirlockNotYukawa}) {
		t.Fatalf("missing real form failures: %+v", a.RealForm.Failures)
	}
}

func TestGate824OutcomeAndImpact(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Outcome.Outcome != OutcomePartialContainer {
		t.Fatalf("unexpected outcome: %+v", a.Outcome)
	}
	if a.Impact.CanUpdate {
		t.Fatalf("must not update C_Yukawa")
	}
	if math.Abs(a.Impact.OfficialCYukawa-CYukawa) > 1e-18 || math.Abs(a.Impact.OfficialCHiggs-CHiggs) > 1e-18 {
		t.Fatalf("official impact changed: %+v", a.Impact)
	}
}

func TestGate824TheoremAndFirewalls(t *testing.T) {
	res := Generation2CovariantPhaseSpaceCliffordSourceRouterAndYukawaGapAccelerationAuditTheorem().Verify()
	if string(res.Status) == "FAILED_ROUTE" {
		t.Fatalf("theorem failed: %+v", res)
	}
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewalls.Enforced || !a.Firewalls.ContainerNotSpectrum || !a.Firewalls.NoCYukawaUpdate || a.Firewalls.Verdict != StatusFirewallGate824 {
		t.Fatalf("firewalls not enforced: %+v", a.Firewalls)
	}
}
