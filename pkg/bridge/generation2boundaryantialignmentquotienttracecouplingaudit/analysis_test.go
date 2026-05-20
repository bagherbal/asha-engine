package generation2boundaryantialignmentquotienttracecouplingaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate676Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.TraceResponseCandidateInherited || !a.Inherited.MissingReasonTraceActsOnLine || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.BoundaryPlane.Dimension != 2 || a.BoundaryPlane.Vector != [2]float64{lambda12, r3Minus1} {
		t.Fatalf("bad boundary plane: %+v", a.BoundaryPlane)
	}
	if !a.AntiAlignment.IsInKernelOfSigma || a.AntiAlignment.AntiAlignmentGenerator != [2]float64{-1, 1} {
		t.Fatalf("bad anti-alignment: %+v", a.AntiAlignment)
	}
	if !a.Quotient.CanonicalCokernelDefect || math.Abs(a.Quotient.SSplit-0.0012924448188163) > 1e-15 {
		t.Fatalf("bad quotient: %+v", a.Quotient)
	}
	if math.Abs(a.Coupling.TauDefect-7.0/72.0) > 1e-15 || math.Abs(a.Coupling.Residual-8.52583439801e-10) > 1e-14 {
		t.Fatalf("bad coupling: %+v", a.Coupling)
	}
	if !a.Upgrade.LessTautological || a.Upgrade.PromotableToTheorem {
		t.Fatalf("bad upgrade: %+v", a.Upgrade)
	}
	if a.Discipline.ClaimsNativeTraceBoundaryQuotient || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.ClaimsNativeWallAirlock || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsFullK7BoundaryMap || a.Discipline.Verdict != StatusGate676Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestQuotientFunctionalAnnihilatesAntiAlignment(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	sigma := a.Quotient.FunctionalVector[0]*a.AntiAlignment.AntiAlignmentGenerator[0] + a.Quotient.FunctionalVector[1]*a.AntiAlignment.AntiAlignmentGenerator[1]
	if math.Abs(sigma) > 1e-15 {
		t.Fatalf("sigma should annihilate anti-alignment generator, got %.17g", sigma)
	}
	if math.Abs(a.Quotient.SSplit-(lambda12+r3Minus1)) > 1e-15 {
		t.Fatalf("split mismatch")
	}
}

func TestTraceCouplingIdentity(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	pred := (7.0 / 72.0) * a.Quotient.SSplit
	if math.Abs(pred-a.Coupling.PredictedDBase) > 1e-15 {
		t.Fatalf("pred mismatch")
	}
	if math.Abs((a.Coupling.DBase-pred)-a.Coupling.Residual) > 1e-15 {
		t.Fatalf("residual mismatch")
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2BoundaryAntiAlignmentQuotientLineTraceCouplingAuditTheorem().Verify()
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
