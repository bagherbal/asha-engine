package generation2boundarystresssplitlinepullbacksourceaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate673Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.InheritedStressSplitPullback || !a.Inherited.BaseClosureComputed || !a.Inherited.StressSplitComputed || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if math.Abs(a.BoundaryLine.SSplit-0.0012924448188163) > 1e-14 {
		t.Fatalf("bad boundary split: %+v", a.BoundaryLine)
	}
	if math.Abs(a.BaseLine.DBase-0.00012565520996836) > 1e-14 {
		t.Fatalf("bad base line: %+v", a.BaseLine)
	}
	if math.Abs(a.Coefficient.QPull-0.0972228818894104) > 1e-12 || a.Coefficient.BestTypedCandidate != "7/72" || len(a.Coefficient.Candidates) != 6 {
		t.Fatalf("bad coefficient: %+v", a.Coefficient)
	}
	if math.Abs(a.Coefficient.SevenOver72Residual-8.52583727234e-10) > 1e-14 {
		t.Fatalf("bad 7/72 residual: %+v", a.Coefficient)
	}
	if len(a.Source.CandidateSupport) != 3 || len(a.Source.MissingTheorems) != 4 || !strings.Contains(a.Source.Verdict, StatusLinePullbackSharperThanFullMap) {
		t.Fatalf("bad source audit: %+v", a.Source)
	}
	if !a.Firewall.FullK7ToBoundaryMapFailed || !a.Firewall.FanoHitchinRouteRemainsSealed || !a.Firewall.LinePullbackStillPossible {
		t.Fatalf("bad firewall: %+v", a.Firewall)
	}
	if !a.ScaleLocal.Lambda12Local || !a.ScaleLocal.CrossingBased || !a.ScaleLocal.StationarityRejected || !a.ScaleLocal.QPullNearSevenOver72OnlyAtLambda12 {
		t.Fatalf("bad scale locality: %+v", a.ScaleLocal)
	}
	if a.Discipline.ClaimsNativeStressSplitPullback || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.ClaimsFullK7BoundaryMap || a.Discipline.ClaimsWallDistanceAirlock || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsHiggsMassPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.Verdict != StatusGate673Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestTypedCandidateOrdering(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	best := a.Coefficient.Candidates[0]
	for _, c := range a.Coefficient.Candidates[1:] {
		if c.AbsResidual < best.AbsResidual {
			best = c
		}
	}
	if best.Name != "7/72" {
		t.Fatalf("expected 7/72 best, got %+v", best)
	}
}

func TestLinePullbackIdentity(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	got := a.BaseLine.DBase - sevenOver72*a.BoundaryLine.SSplit
	if math.Abs(got-a.Coefficient.SevenOver72Residual) > 1e-15 {
		t.Fatalf("identity mismatch: got %.17g residual %.17g", got, a.Coefficient.SevenOver72Residual)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2BoundaryStressSplitLinePullbackSourceAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
