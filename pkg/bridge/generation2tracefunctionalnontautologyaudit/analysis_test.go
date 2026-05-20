package generation2tracefunctionalnontautologyaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate675Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.TraceCandidateInherited || !a.Inherited.FirewallPreserved || !a.Inherited.FullK7BoundaryMapFailed {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.Projector.TotalDimension != 72 || a.Projector.RankPDefect != 7 || a.Projector.BoundaryActionRank != 0 || a.Projector.BoundaryVectorMapNeeded {
		t.Fatalf("bad projector: %+v", a.Projector)
	}
	if math.Abs(a.Trace.TauDefect-7.0/72.0) > 1e-15 {
		t.Fatalf("bad tau defect: %+v", a.Trace)
	}
	if a.BoundaryLine.ChosenLine.Name != "split line" || len(a.BoundaryLine.Candidates) != 5 {
		t.Fatalf("bad boundary line: %+v", a.BoundaryLine)
	}
	if math.Abs(a.Ansatz.Residual-8.52583439801e-10) > 1e-14 || !a.Ansatz.RequiresScalarFunctional || a.Ansatz.RequiresVectorBoundaryMap {
		t.Fatalf("bad ansatz: %+v", a.Ansatz)
	}
	if a.NonTautology.CertifiedCriteriaCount != 4 || a.NonTautology.RequiredCriteriaCount != 5 || a.NonTautology.PromotableToNativeTheorem {
		t.Fatalf("bad non-tautology audit: %+v", a.NonTautology)
	}
	if len(a.Sources) != 5 || a.Sources[2].Status != "sealed" || a.Sources[4].Status != "missing theorem" {
		t.Fatalf("bad sources: %+v", a.Sources)
	}
	if a.Discipline.ClaimsNativeTraceResponse || a.Discipline.ClaimsTraceActsOnSplitLine || a.Discipline.ClaimsNativeWallDistanceAirlock || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.ClaimsNativeStressSplitPullback || a.Discipline.ClaimsFullK7BoundaryMap || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsHiggsMassPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.Verdict != StatusGate675Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestTraceFunctionalIdentity(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	tau := float64(a.Projector.RankPDefect) / float64(a.Projector.TotalDimension)
	got := a.Ansatz.DBase - tau*a.Ansatz.SSplit
	if math.Abs(tau-a.Trace.TauDefect) > 1e-15 {
		t.Fatalf("tau mismatch %.17g %.17g", tau, a.Trace.TauDefect)
	}
	if math.Abs(got-a.Ansatz.Residual) > 1e-15 {
		t.Fatalf("residual mismatch got %.17g stored %.17g", got, a.Ansatz.Residual)
	}
}

func TestNonTautologyMissingCriterion(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, c := range a.NonTautology.Criteria {
		if c.Criterion == "typed reason trace acts on S_split" {
			found = true
			if c.Certified || c.Status != "missing" {
				t.Fatalf("expected missing action criterion: %+v", c)
			}
		}
	}
	if !found {
		t.Fatal("missing criterion not found")
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2AugmentedChamberTraceResponseFunctionalNonTautologyAuditTheorem().Verify()
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
