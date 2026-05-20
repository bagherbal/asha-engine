package generation2scalarbaselineandk7boundarysplitupliftobservableaudit

import (
	"math"
	"strings"
	"testing"
)

func nearly(a, b float64) bool { return math.Abs(a-b) < tolerance }

func TestGate704InheritanceAndRewrite(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	inh := a.Inherited
	if !inh.BoundaryWoundMixtureInherited || !inh.EventComplementMeaning || !inh.KSumExpectedBoundaryWound || !nearly(inh.K7EventProbability, pK7) || !nearly(inh.ComplementEventProbability, pComplement) || math.Abs(inh.ExpectedWound-0.0498265964350682) > 5e-16 || math.Abs(inh.Residual-8.5258344e-10) > 1e-16 || !inh.NoNativeBoundaryWoundMixture || !inh.NoNativeHistoryResponse || !inh.NoNativeSevenOver72 || inh.Verdict != StatusGate704BoundaryWoundMixtureInherited {
		t.Fatalf("bad inheritance: %+v", inh)
	}
	r := a.Rewrite
	if !strings.Contains(r.Gate704Observable, "P_perp") || !strings.Contains(r.ProjectorIdentity, "I_H72-P_K7") || !strings.Contains(r.RewrittenObservable, "|lambda|I_H72") || !strings.Contains(r.RewrittenObservable, "S_split P_K7") || !nearly(r.K7PayoffBefore, r3Minus1) || !nearly(r.ComplementPayoffBefore, math.Abs(lambdaLambda12)) || !r.AlgebraicEquivalence || r.Verdict != StatusTwoPayoffObservableRewritten {
		t.Fatalf("bad rewrite: %+v", r)
	}
}

func TestScalarBaselineUpliftAndExpectation(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	d := a.Decomposition
	sSplit := lambdaLambda12 + r3Minus1
	if !nearly(d.ScalarBaseline, math.Abs(lambdaLambda12)) || !nearly(d.BoundarySplitUplift, sSplit) || !nearly(d.SFromSignedSplit, sSplit) || !nearly(d.K7PayoffAfterUplift, r3Minus1) || !nearly(d.ComplementPayoff, math.Abs(lambdaLambda12)) || !d.FullChamberBaseline || !d.K7LocalizedUplift || !d.K7PayoffEqualsGaugeWound || !strings.Contains(d.Verdict, StatusK7ReceivesSplitUpliftNotPrimitiveGaugeWound) {
		t.Fatalf("bad decomposition: %+v", d)
	}
	e := a.Expectation
	if !strings.Contains(e.Formula, "|lambda|") || !strings.Contains(e.Formula, "p_K7 S_split") || !nearly(e.BaselineExpectation, math.Abs(lambdaLambda12)) || !nearly(e.UpliftExpectation, pK7*sSplit) || math.Abs(e.TotalExpectation-0.0498265964350682) > 5e-16 || math.Abs(e.KSum-0.0498265972876517) > 5e-16 || math.Abs(e.Residual-8.5258344e-10) > 1e-16 || !e.ReproducesKSumClosure || !e.EquivalentToGate704 || !strings.Contains(e.Verdict, StatusKSumScalarBaselineExpectedK7SplitUplift) {
		t.Fatalf("bad expectation: %+v", e)
	}
}

func TestRelationSourceUpgradeAndAlternatives(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Relation
	if !strings.Contains(r.Gate700ResponseLaw, "D_base") || !strings.Contains(r.PositiveDistanceLaw, "K_sum") || !nearly(r.DBase, kappaLambda+kappaE+lambdaLambda12) || !nearly(r.KSum, kappaLambda+kappaE) || !nearly(r.ScalarDepth, math.Abs(lambdaLambda12)) || !nearly(r.ExpectedSplitUplift, pK7*(lambdaLambda12+r3Minus1)) || !strings.Contains(r.BaselineIdentity, "K_sum=|lambda|+D_base") || !r.NotNewNumericalRelation || r.Verdict != StatusRelationToGate700ResponseLawAudited {
		t.Fatalf("bad Gate700 relation: %+v", r)
	}
	s := a.SourceUpgrade
	if !strings.Contains(s.Gate704Reading, "gauge wound") || !strings.Contains(s.Gate705Reading, "scalar wound baseline") || !s.RemovesPrimitiveGaugeAssignment || !s.K7ReceivesBoundarySplitUplift || !s.ScalarWallAirlockSupportsBaseline || !strings.Contains(s.Verdict, StatusSourceTypeUpgradeAudited) {
		t.Fatalf("bad source upgrade: %+v", s)
	}
	alts := a.Alternatives
	if len(alts.Alternatives) != 4 || !alts.GaugeBaselineRejected || !alts.MidpointBaselineRejected || !alts.HodgeSignedUpliftRejected || !alts.ActiveBaselineAccepted || alts.Verdict != StatusAlternativeBaselineDecompositionsAudited {
		t.Fatalf("bad alternatives: %+v", alts)
	}
	if !alts.Alternatives[0].Equivalent || !alts.Alternatives[0].Rejected || !strings.Contains(alts.Alternatives[0].Reason, "scalar-wall airlock") {
		t.Fatalf("bad gauge baseline alternative: %+v", alts.Alternatives[0])
	}
	if alts.Alternatives[1].Equivalent || !alts.Alternatives[1].Rejected || !strings.Contains(alts.Alternatives[1].Reason, "less minimal") {
		t.Fatalf("bad midpoint alternative: %+v", alts.Alternatives[1])
	}
	if alts.Alternatives[2].Equivalent || !alts.Alternatives[2].Rejected || !strings.Contains(alts.Alternatives[2].Reason, "1/72") {
		t.Fatalf("bad Hodge alternative: %+v", alts.Alternatives[2])
	}
	if !alts.Alternatives[3].Equivalent || !alts.Alternatives[3].Active || alts.Alternatives[3].Rejected || !nearly(alts.Alternatives[3].Expectation, a.Expectation.TotalExpectation) {
		t.Fatalf("bad active alternative: %+v", alts.Alternatives[3])
	}
}

func TestMissingFirewallAndTheoremNotes(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	m := a.Missing
	if len(m.Missing) != 5 || !strings.Contains(m.Verdict, StatusNoNativeScalarWoundFullChamberBaseline) || !strings.Contains(m.Verdict, StatusNoNativeK7ReceivesSplitUplift) || !strings.Contains(m.Verdict, StatusNoNativeBoundaryWoundUpliftTheorem) || !strings.Contains(m.Verdict, StatusNoNativeHistoryResponseTheorem) || !strings.Contains(m.Verdict, StatusNoNativeSevenOver72Theorem) {
		t.Fatalf("bad missing theorem audit: %+v", m)
	}
	f := a.Firewall
	if f.ClaimsScalarWoundFullChamberBaselineNative || f.ClaimsK7ReceivesSplitUpliftNative || f.ClaimsNativeBoundaryWoundUpliftTheorem || f.ClaimsNativeHistoryResponseTheorem || f.ClaimsNativeSevenOver72Theorem || f.ClaimsBoundaryStressDerived || f.ClaimsScalarRGMatching || f.ClaimsHiggsMass || f.ClaimsGaugeUnification || f.ClaimsFlavorDerivation || f.ClaimsCKMPMNS || f.Verdict != StatusGate705ScalarBaselineK7UpliftBoundary {
		t.Fatalf("firewall violated: %+v", f)
	}
	res := Generation2ScalarBaselineAndK7BoundarySplitUpliftObservableAuditTheorem().Verify()
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
