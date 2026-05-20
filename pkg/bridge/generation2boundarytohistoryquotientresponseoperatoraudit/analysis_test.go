package generation2boundarytohistoryquotientresponseoperatoraudit

import (
	"math"
	"strings"
	"testing"
)

func nearly(a, b float64) bool { return math.Abs(a-b) < tolerance }

func TestGate698InheritanceAndQuotientCoordinates(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.HistoryReadoutInherited || !strings.Contains(a.Inherited.SigmaHistory, "sigma_history") || !nearly(a.Inherited.DBase, kappaLambda+kappaE+lambdaLambda12) || !strings.Contains(a.Inherited.SigmaBoundary, "sigma_boundary") || !nearly(a.Inherited.BoundaryExpectation, responseCoeff*(lambdaLambda12+r3Minus1)) || !a.Inherited.NoNativeHistoryBoundary || !a.Inherited.NoNativeSevenOver72 {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.BoundaryIn.QuotientLine != "Q_boundary=B_boundary/L_anti" || !strings.Contains(a.BoundaryIn.Coordinate, "lambda+(R_3-1)") || a.BoundaryIn.BoundaryVector != [2]float64{lambdaLambda12, r3Minus1} || !nearly(a.BoundaryIn.SBoundary, lambdaLambda12+r3Minus1) || !a.BoundaryIn.VanishesOnAntiAlignment || a.BoundaryIn.AntiAlignmentWall != "lambda+(R_3-1)=0" {
		t.Fatalf("bad boundary input: %+v", a.BoundaryIn)
	}
	if !strings.Contains(a.HistoryOut.QuotientLine, "Q_history") || !strings.Contains(a.HistoryOut.Coordinate, "kappa_lambda+kappa_e+lambda") || a.HistoryOut.HistoryVector != [3]float64{kappaLambda, kappaE, lambdaLambda12} || !nearly(a.HistoryOut.SHistory, kappaLambda+kappaE+lambdaLambda12) || !a.HistoryOut.VanishesOnClosureWall || a.HistoryOut.ClosureWall != "kappa_lambda+kappa_e+lambda=0" {
		t.Fatalf("bad history output: %+v", a.HistoryOut)
	}
}

func TestResponseOperatorAndBridgeReconstruction(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Operator.Domain != "Q_boundary" || a.Operator.Codomain != "Q_history" || !strings.Contains(a.Operator.Definition, "Tr(rho_72 s P_K7)") || !nearly(a.Operator.Input, a.BoundaryIn.SBoundary) || !nearly(a.Operator.Coefficient, responseCoeff) || !nearly(a.Operator.Output, responseCoeff*a.BoundaryIn.SBoundary) || !a.Operator.LinearityCertified || !a.Operator.CoefficientFromEvent {
		t.Fatalf("bad operator: %+v", a.Operator)
	}
	for _, want := range []string{StatusResponseOperatorRK7Defined, StatusResponseCoefficientComputedAsK7EventWeight, StatusSevenOver72ResponseCoefficientFromK7Weight} {
		if !strings.Contains(a.Operator.Verdict, want) {
			t.Fatalf("missing operator verdict %s in %q", want, a.Operator.Verdict)
		}
	}
	if !strings.Contains(a.Bridge.Equation, "D_base") || !strings.Contains(a.Bridge.ExpandedEquation, "lambda(Lambda_12)+(R_3-1)") || !nearly(a.Bridge.DBase, a.HistoryOut.SHistory) || !nearly(a.Bridge.RK7OfSSplit, a.Operator.Output) || math.Abs(a.Bridge.ResidualE1-a.Inherited.ResidualE1) > 1e-17 || !a.Bridge.MatchesInheritedResidual {
		t.Fatalf("bad bridge: %+v", a.Bridge)
	}
}

func TestSharedLambdaNonTautology(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	n := a.NonTautology
	wantRight := -(65.0/72.0)*lambdaLambda12 + responseCoeff*r3Minus1
	if !n.LeftContainsLambda || !n.RightContainsLambda || !strings.Contains(n.RearrangedEquation, "-(65/72)lambda") || !strings.Contains(n.RearrangedEquation, "(7/72)(R_3-1)") || !nearly(n.KSum, kappaLambda+kappaE) || !nearly(n.WeightedClosureRight, wantRight) || math.Abs(n.Residual-a.Inherited.ResidualE1) > 1e-17 || !n.CoefficientsDiffer || !n.IncludesIndependentGauge || !n.NotIdentity {
		t.Fatalf("bad non-tautology audit: %+v", n)
	}
	for _, want := range []string{StatusSharedLambdaNonTautologyAudited, StatusSharedLambdaDoesNotMakeTautological} {
		if !strings.Contains(n.Verdict, want) {
			t.Fatalf("missing non-tautology verdict %s in %q", want, n.Verdict)
		}
	}
}

func TestAlternativeResponseCoefficientsAndSourceTypes(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	alts := a.Alternatives
	if len(alts.Alternatives) != 6 || !alts.ZeroRejected || !alts.IdentityRejected || !alts.FiniteOnlyRejected || !alts.KernelRejected || !alts.SignedRejected || !alts.SevenOver72Active || !alts.AllAudited {
		t.Fatalf("bad alternatives: %+v", alts)
	}
	byName := map[string]AlternativeResponseCoefficient{}
	for _, x := range alts.Alternatives {
		byName[x.Name] = x
	}
	if byName["zero response"].Active || byName["identity response"].Active || byName["finite-only state response"].Active || byName["kernel-state response"].Active || byName["Hodge-signed event response"].Active {
		t.Fatalf("rejected alternatives active: %+v", byName)
	}
	if !byName["full augmented no-bias K7 event response"].Active || !nearly(byName["full augmented no-bias K7 event response"].Coefficient, responseCoeff) {
		t.Fatalf("active 7/72 alternative missing: %+v", byName["full augmented no-bias K7 event response"])
	}
	if !strings.Contains(a.SourceTypes.BoundaryQuotientRole, "input") || !strings.Contains(a.SourceTypes.HistoryQuotientRole, "output") || !strings.Contains(a.SourceTypes.Rho72Role, "maximum-entropy") || !strings.Contains(a.SourceTypes.PK7Role, "Boolean-octonionic") || !strings.Contains(a.SourceTypes.CoefficientRole, "7/72") || !strings.Contains(a.SourceTypes.BridgeRole, "boundary quotient") {
		t.Fatalf("bad source types: %+v", a.SourceTypes)
	}
}

func TestMissingTheoremFirewallAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Missing.Candidates) != 5 || !strings.Contains(a.Missing.PreciseGap, "boundary") || !strings.Contains(a.Missing.PreciseGap, "history") || !strings.Contains(a.Missing.Verdict, StatusNoNativeBoundaryControlsHistory) || !strings.Contains(a.Missing.Verdict, StatusNoNativeBoundaryHistoryResponseTheorem) || !strings.Contains(a.Missing.Verdict, StatusNoNativeSevenOver72Theorem) {
		t.Fatalf("bad missing theorem audit: %+v", a.Missing)
	}
	if a.Firewall.ClaimsNativeBoundaryControlsHistory || a.Firewall.ClaimsNativeBoundaryHistoryTheorem || a.Firewall.ClaimsNativeSevenOver72Theorem || a.Firewall.ClaimsBoundaryStressDerived || a.Firewall.ClaimsScalarRGMatching || a.Firewall.ClaimsHiggsMass || a.Firewall.ClaimsGaugeUnification || a.Firewall.ClaimsFlavorDerivation || a.Firewall.ClaimsCKMPMNS || a.Firewall.Verdict != StatusGate699BoundaryHistoryQuotientResponseBoundary {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
	res := Generation2BoundaryToHistoryQuotientResponseOperatorAuditTheorem().Verify()
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
