package generation2conditionalashahistoryresponselawclosureaudit

import (
	"math"
	"strings"
	"testing"
)

func nearly(a, b float64) bool { return math.Abs(a-b) < 1e-15 }

func TestGate699InheritanceAndFunctional(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.InheritedBoundaryHistoryResponse || !strings.Contains(a.Inherited.ResponseOperator, "R_K7") || !nearly(a.Inherited.SBoundary, lambdaLambda12+r3Minus1) || !nearly(a.Inherited.SHistory, kappaLambda+kappaE+lambdaLambda12) || !nearly(a.Inherited.RK7OfSSplit, responseCoeff*a.Inherited.SBoundary) || math.Abs(a.Inherited.ResidualE1-(a.Inherited.SHistory-a.Inherited.RK7OfSSplit)) > 1e-17 || !a.Inherited.SharedLambdaNonTautology || !a.Inherited.NoNativeBoundaryHistoryTheorem || !a.Inherited.NoNativeSevenOver72 || a.Inherited.Verdict != StatusGate699BoundaryHistoryResponseInherited {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	f := a.Functional
	if f.FunctionalName != "A_history(b,h)" || !strings.Contains(f.Equation, "sigma_history") || !strings.Contains(f.Equation, "rho_72") || !nearly(f.SigmaHistory, a.Inherited.SHistory) || !nearly(f.ExpectedBoundaryK7, a.Inherited.RK7OfSSplit) || math.Abs(f.AHistory-a.Inherited.ResidualE1) > 1e-17 || f.AbsoluteResidual > 1e-8 || !f.ApproxLawCertified || !f.UsesFirstExpectation || !strings.Contains(f.Verdict, StatusConditionalHistoryResponseFunctionalDefined) || !strings.Contains(f.Verdict, StatusCompleteConditionalResponseLaw) {
		t.Fatalf("bad conditional functional: %+v", f)
	}
}

func TestPremiseLadderAndRemovalAudit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	p := a.Premises
	if len(p.Premises) != 7 || !p.Complete || !p.AllStructurallyUsed || p.Verdict != StatusPremiseLadderConstructed {
		t.Fatalf("bad premise ladder: %+v", p)
	}
	for i, premise := range p.Premises {
		if premise.Index != i+1 || premise.Name == "" || premise.Object == "" || premise.Role == "" || !premise.StructurallyUsed {
			t.Fatalf("bad premise %d: %+v", i, premise)
		}
	}
	r := a.Removal
	if len(r.Removals) != 7 || !r.RemoveRho72ChangesCoeff || !r.RemovePK7RestoresDegeneracy || !r.RemoveSupportBreaksIdentity || !r.RemoveSupportLocalityRestoresPayoffDegeneracy || !r.RemoveSigmaBoundaryBreaksPayoffRole || !r.RemoveSigmaHistoryBreaksReadoutRole || !r.RemoveFirstExpectationBreaksLeadingOrder || !r.EachPremiseNonredundant {
		t.Fatalf("bad removal audit: %+v", r)
	}
	for _, removal := range r.Removals {
		if removal.RemovedPremise == "" || removal.FailureMode == "" || removal.ExpectedFailure == "" || !removal.Nonredundant {
			t.Fatalf("bad removal item: %+v", removal)
		}
	}
	if !strings.Contains(r.Verdict, StatusPremiseRemovalAuditComputed) || !strings.Contains(r.Verdict, StatusEachPremiseNonredundantStructuralRole) {
		t.Fatalf("bad removal verdict: %q", r.Verdict)
	}
}

func TestMasterBridgeResidualAndMissingTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	m := a.Master
	if !strings.Contains(m.CompactEquation, "sigma_history") || !strings.Contains(m.CompactEquation, "rho_72") || !strings.Contains(m.ExpandedEquation, "I_H72/72") || !nearly(m.DBase, a.Inherited.SHistory) || !nearly(m.Expectation, a.Inherited.RK7OfSSplit) || math.Abs(m.ResidualE1-a.Inherited.ResidualE1) > 1e-17 || !m.Reconstructed || !strings.Contains(m.Verdict, StatusMasterBridgeEquationReconstructed) || !strings.Contains(m.Verdict, StatusCompleteConditionalResponseLaw) {
		t.Fatalf("bad master bridge: %+v", m)
	}
	res := a.Residual
	if math.Abs(res.ResidualE1-a.Inherited.ResidualE1) > 1e-17 || !strings.Contains(res.QuadraticCorrectionCandidate, "Gate690") || !res.QuadraticSubleading || !res.QuadraticNotIndependent || res.ResidualAbsorbed || res.Verdict != StatusResidualStatusRecorded {
		t.Fatalf("bad residual audit: %+v", res)
	}
	missing := a.Missing
	if len(missing.Candidates) != 7 || !strings.Contains(missing.PreciseGap, "full augmented no-bias state") || !strings.Contains(missing.PreciseGap, "K7 event support") || !strings.Contains(missing.PreciseGap, "first ordinary expectation") {
		t.Fatalf("bad missing theorem gap: %+v", missing)
	}
	for _, want := range []string{StatusAshaHistoryResponseLawTargetSharpened, StatusPremisesNotNativelyDerived, StatusNoNativeBoundaryHistoryResponsePrinciple, StatusNoNativeStateSelectionTheorem, StatusNoNativeK7EventPayoffTheorem, StatusNoNativeSevenOver72Theorem} {
		if !strings.Contains(missing.Verdict, want) {
			t.Fatalf("missing verdict %s in %q", want, missing.Verdict)
		}
	}
}

func TestFirewallAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if f.ClaimsPremisesNativelyDerived || f.ClaimsNativeBoundaryHistoryPrinciple || f.ClaimsNativeStateSelectionTheorem || f.ClaimsNativeK7EventPayoffTheorem || f.ClaimsNativeSevenOver72Theorem || f.ClaimsBoundaryStressDerived || f.ClaimsScalarRGMatching || f.ClaimsHiggsMass || f.ClaimsGaugeUnification || f.ClaimsFlavorDerivation || f.ClaimsCKMPMNS || f.Verdict != StatusGate700ConditionalHistoryResponseLawBoundary {
		t.Fatalf("firewall violated: %+v", f)
	}
	res := Generation2ConditionalAshaHistoryResponseLawClosureAuditTheorem().Verify()
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
