package generation2boundaryexteriortargetselectionsourcesearchaudit

import (
	"strings"
	"testing"
)

func TestGate875InheritsWoundAndFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.OfficialFrozen || a.Ledger.CanUpdateOfficial {
		t.Fatalf("official freeze broken: %s", FormatLedger(a.Ledger))
	}
	if !a.Wound.NeedLambda1ToPiTop || !a.Wound.NeedLambda2ToHRMin || !a.Wound.NeedNoLambda1ToHRMin || !a.Wound.NeedNoLambda2ToPiTop {
		t.Fatalf("wound not inherited: %s", FormatWound(a.Wound))
	}
}

func TestGate875PunctureComplementStrongestRoute(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.PunctureRoute.NativeFunctor || a.PunctureRoute.Strength != "strongest internal source candidate" {
		t.Fatalf("bad puncture route: %s", FormatRoute(a.PunctureRoute))
	}
	if !containsAll(a.PunctureRoute.Supports, []string{SupportPunctureComplementStrongestRoute, SupportExposureVisibleComplementPiTop, SupportEnclosurePuncturedActiveDomain}) {
		t.Fatalf("missing puncture supports: %s", FormatRoute(a.PunctureRoute))
	}
}

func TestGate875AlternativeRoutesRemainCandidatesOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.CodimRoute.NativeFunctor || a.ChamberRoute.NativeFunctor {
		t.Fatalf("alternative route promoted: %s | %s", FormatRoute(a.CodimRoute), FormatRoute(a.ChamberRoute))
	}
	if !containsAll(a.CodimRoute.Failures, []string{FailureCodimRouteNotFunctor}) || !containsAll(a.ChamberRoute.Failures, []string{FailureChamberRouteNotFunctor, FailureResponseChamberTypingNotTheorem}) {
		t.Fatalf("missing route failures: %s | %s", FormatRoute(a.CodimRoute), FormatRoute(a.ChamberRoute))
	}
}

func TestGate875AlphaReconstructionStillSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Candidate.ShapeCoherent || !near(a.Candidate.ReconstructedAlpha, AlphaB) || a.Candidate.TargetSelectionNative {
		t.Fatalf("bad alpha candidate: %s", FormatCandidate(a.Candidate))
	}
	if !containsAll(a.Candidate.Failures, []string{FailureNoNativeTargetSelectionFunctor, FailureAlphaStillSealed, FailureNoNativeAlphaSource}) {
		t.Fatalf("missing alpha failures: %s", FormatCandidate(a.Candidate))
	}
}

func TestGate875R3Blocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.R3.EligibleForR3 || a.R3.EligibleForR4 || a.Impact.CanPromoteToR3 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("promotion/update leaked: %s | %s", FormatR3(a.R3), FormatImpact(a.Impact))
	}
}

func TestGate875Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate875Theorem(t *testing.T) {
	res := Generation2BoundaryExteriorTargetSelectionSourceSearchAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
