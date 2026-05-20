package generation2centralscalarbaselineandupliftonlyresponseisolationaudit

import (
	"math"
	"strings"
	"testing"
)

func nearly(a, b float64) bool { return math.Abs(a-b) < 1e-15 }

func TestCentralBaselineAndUpliftIsolation(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ScalarBaselineK7UpliftInherited || !a.Inherited.KSumScalarBaselineUplift || !a.Inherited.K7SplitUpliftNotPrimitiveGauge || !a.Inherited.NoNativeScalarBaseline || a.Inherited.Verdict != StatusGate705ScalarBaselineK7UpliftInherited {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	c := a.CentralBaseline
	if !strings.Contains(c.Operator, "|lambda|I_H72") || !nearly(c.ScalarBaseline, math.Abs(lambdaLambda12)) || !c.CommutesWithPK7 || !c.CommutesWithPB || !c.CommutesWithPG || !c.CommutesWithPPerp || !c.ProjectorBlind || !c.IdentityShift || !strings.Contains(c.Verdict, StatusScalarBaselineCentralIdentityShift) {
		t.Fatalf("bad central baseline: %+v", c)
	}
	b := a.BaselineExpect
	if !b.ObserverIndependent || !nearly(b.Rho72Expectation, math.Abs(lambdaLambda12)) || !nearly(b.FiniteStateExpectation, math.Abs(lambdaLambda12)) || !nearly(b.KernelStateExpectation, math.Abs(lambdaLambda12)) || !nearly(b.LocalK7StateExpectation, math.Abs(lambdaLambda12)) {
		t.Fatalf("bad baseline expectation: %+v", b)
	}
	u := a.Uplift
	if !strings.Contains(u.UpliftOperator, "S_split P_K7") || !nearly(u.BoundarySplit, lambdaLambda12+r3Minus1) || !nearly(u.DBase, kappaLambda+kappaE+lambdaLambda12) || !nearly(u.KSumMinusScalarBaseline, u.DBase) || !nearly(u.UpliftExpectationRho72, pK7*(lambdaLambda12+r3Minus1)) || math.Abs(u.Residual-8.5258344e-10) > 1e-16 || !u.IsolatesGate700Law || !strings.Contains(u.Verdict, StatusNontrivialBridgeContentK7Uplift) {
		t.Fatalf("bad uplift isolation: %+v", u)
	}
}

func TestObserverSupportRelationAndFirewall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	o := a.Observer
	if !o.BaselineIndependent || !strings.Contains(o.UpliftFormula, "Tr(rho P_K7)") || !nearly(o.Rho72UpliftExpectation, pK7*a.Uplift.BoundarySplit) || !nearly(o.FiniteUpliftExpectation, pK7Finite*a.Uplift.BoundarySplit) || !nearly(o.KernelUpliftExpectation, pK7Kernel*a.Uplift.BoundarySplit) || !nearly(o.LocalK7UpliftExpectation, a.Uplift.BoundarySplit) || !o.OnlyUpliftRequiresRho72 {
		t.Fatalf("bad observer dependence: %+v", o)
	}
	s := a.Support
	if s.BaselineSelectsPK7 || s.BaselineSelectsRho72 || !s.UpliftProjectorSensitive || !s.UpliftSupportSelected || !strings.Contains(s.SelectorStatement, "P_K7") || !s.K7IdentityRequiresSupport || !strings.Contains(s.Verdict, StatusBaselineDoesNotSelectK7OrRho72) {
		t.Fatalf("bad support dependence: %+v", s)
	}
	r := a.Relation
	if !strings.Contains(r.Gate705PositiveDistanceLaw, "K_sum") || !strings.Contains(r.Gate706SubtractedLaw, "D_base") || !nearly(r.KSum, kappaLambda+kappaE) || !nearly(r.DBase, kappaLambda+kappaE+lambdaLambda12) || !nearly(r.ScalarBaseline, math.Abs(lambdaLambda12)) || !r.NoNewNumericalRelation {
		t.Fatalf("bad relation: %+v", r)
	}
	m := a.Missing
	if len(m.Missing) != 6 || !strings.Contains(m.Verdict, StatusBaselineDoesNotSelectK7OrRho72) || !strings.Contains(m.Verdict, StatusNoNativeHistoryResponseTheorem) || !strings.Contains(m.Verdict, StatusNoNativeSevenOver72Theorem) {
		t.Fatalf("bad missing theorem audit: %+v", m)
	}
	f := a.Firewall
	if f.ClaimsBaselineSelectsK7OrRho72 || f.ClaimsScalarWoundFullChamberBaselineNative || f.ClaimsK7ReceivesSplitUpliftNative || f.ClaimsNativeBoundaryWoundUpliftTheorem || f.ClaimsNativeHistoryResponseTheorem || f.ClaimsNativeSevenOver72Theorem || f.ClaimsBoundaryStressDerived || f.ClaimsScalarRGMatching || f.ClaimsHiggsMass || f.ClaimsGaugeUnification || f.ClaimsFlavorDerivation || f.ClaimsCKMPMNS || f.Verdict != StatusGate706CentralBaselineUpliftBoundary {
		t.Fatalf("firewall violated: %+v", f)
	}
	res := Generation2CentralScalarBaselineAndUpliftOnlyResponseIsolationAuditTheorem().Verify()
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
