package generation2z2boundaryalphafunctorsourcedecompositionaudit

import (
	"strings"
	"testing"
)

func TestGate912InheritedRailAndFormula(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.ReopensPhaseSign || a.Inherited.ReopensSocketOrder || a.Inherited.ReopensRepAlpha || a.Inherited.DerivesAlpha || a.Inherited.UpdatesOfficialLedger {
		t.Fatalf("bad inherited rail: %s", FormatInherited(a.Inherited))
	}
	if !a.Formula.RepresentativeFree || a.Formula.Native || !near(a.Formula.Alpha, AlphaB) {
		t.Fatalf("bad formula status: %s", FormatFormula(a.Formula))
	}
	if a.Formula.RankPair != [2]int{RankF1OverF0, RankF2OverF0} || a.Formula.Denominators != [2]int{LinearDenom, QuadDenom} {
		t.Fatalf("bad rank/denominator pair: %s", FormatFormula(a.Formula))
	}
}

func TestGate912DecompositionHasFiveMissingSubobjects(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	d := a.Decomposition
	if d.RequiredCount != 5 || d.CertifiedCount != 0 || d.MissingCount != 5 || d.NativeFunctor || d.AlphaNative || d.R3Native {
		t.Fatalf("bad decomposition accounting: %s", FormatDecomposition(d))
	}
	if !allRequiredSubobjectsPresent(d.Subobjects) {
		t.Fatalf("required subobjects not complete: %s", FormatDecomposition(d))
	}
	wantTheorems := []string{NativeReducedB2Theorem, DegreeSelectorTheorem, CrossLaneTheorem, SsplitTransportTheorem, DenominatorTypingTheorem}
	for i, want := range wantTheorems {
		if d.Subobjects[i].RequiredTheorem != want || d.Subobjects[i].CertifiedNative {
			t.Fatalf("bad subobject %d: %s", i+1, FormatSubobject(d.Subobjects[i]))
		}
	}
}

func TestGate912SubobjectSpecificFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	subs := a.Decomposition.Subobjects
	checks := []struct {
		idx      int
		supports []string
		failures []string
	}{
		{0, []string{SupportReducedB2Required, SupportReducedB2CorrectShape, SupportZeroOrderSuppressed, SupportCubicAbsent}, []string{FailureReducedB2NotNativeFunctional, FailureNoNativeReasonEBMinusOne, FailureNoNativeTransportSInB2}},
		{1, []string{SupportDegreeSelectorRequired, SupportDegreeOneExposed, SupportDegreeTwoFull, SupportRankPairRepresentativeFree}, []string{FailureNoNativeDegreeToZ2FlagFunctor, FailureNoNativeLambda1ExposedMap, FailureNoNativeLambda2FullMap}},
		{2, []string{SupportCrossLaneRequired, SupportCrossLanesExcludedIfFunctor}, []string{FailureNoNativeZ2CrossLane, FailureNoNativeLinearDomainExclusion, FailureNoNativeQuadraticFaceExclusion}},
		{3, []string{SupportSsplitTransportRequired, SupportSsplitFeedsDegreeShape}, []string{FailureNoNativeTransportS, FailureNoTypedSToLambda1, FailureNoTypedS2ToLambda2}},
		{4, []string{SupportDenominatorChambersRequired, SupportDenominatorsTyped}, []string{FailureDenominatorNotActivation}},
	}
	for _, c := range checks {
		s := subs[c.idx]
		if !s.Required || s.CertifiedNative || !containsAll(s.Supports, c.supports) || !containsAll(s.Failures, c.failures) {
			t.Fatalf("bad subobject %d: %s", c.idx+1, FormatSubobject(s))
		}
	}
	if !containsAll(subs[2].ForbiddenTargets, []string{ForbiddenLinearFull, ForbiddenQuadraticExposed}) || !containsAll(subs[2].WrongTerms, []string{WrongLinearFullTerm, WrongQuadraticExposedTerm}) {
		t.Fatalf("cross-lane object failed to record forbidden targets/terms: %s", FormatSubobject(subs[2]))
	}
}

func TestGate912GlobalFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureNoNativeZ2BoundaryAlphaFunctor, FailureReducedB2NotNativeFunctional, FailureNoNativeDegreeToZ2FlagFunctor, FailureNoNativeZ2CrossLane, FailureNoNativeTransportS, FailureDenominatorNotActivation, FailureAlphaStillSealed, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator} {
		if !containsAll(a.FirewallsList(), []string{want}) {
			t.Fatalf("missing firewall %s from %s", want, FormatFirewalls(a.Firewalls))
		}
	}
}

func TestGate912Theorem(t *testing.T) {
	res := Generation2Z2BoundaryAlphaFunctorSourceDecompositionAuditTheorem().Verify()
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
	for _, want := range []string{FinalTruth, Classification, ShortStatus, StrategicConclusion, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s in notes: %s", want, joined)
		}
	}
}
