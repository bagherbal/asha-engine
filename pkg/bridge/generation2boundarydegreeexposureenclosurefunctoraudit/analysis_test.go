package generation2boundarydegreeexposureenclosurefunctoraudit

import (
	"strings"
	"testing"
)

func TestGate924DegreeExteriorTyping(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.InheritedStatus != Gate923ShortStatus {
		t.Fatalf("bad inherited status: %s", a.InheritedStatus)
	}
	if a.DegreeOne.Degree != 1 || a.DegreeOne.BoundaryFactors != 1 || !a.DegreeOne.NativeShape || a.DegreeOne.NativeTargetMap || a.DegreeOne.TopDegree {
		t.Fatalf("bad degree one: %s", FormatDegree(a.DegreeOne))
	}
	if a.DegreeTwo.Degree != 2 || a.DegreeTwo.BoundaryFactors != 2 || !a.DegreeTwo.TopDegree || !a.DegreeTwo.NativeShape || a.DegreeTwo.NativeTargetMap {
		t.Fatalf("bad degree two: %s", FormatDegree(a.DegreeTwo))
	}
}

func TestGate924ContrastCumulativeSelector(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Contrast.GroundedInDegree || a.Contrast.ArbitraryLabels || a.Contrast.SelectsZ2Targets {
		t.Fatalf("bad contrast: %s", FormatContrast(a.Contrast))
	}
	if !a.Cumulative.TopDegreeSourceCandidate || a.Cumulative.NativeTargetFunctor || a.Cumulative.RankCumulative != RankF2OverF0 || a.Cumulative.RankAssociatedGraded != RankF2OverF1 {
		t.Fatalf("bad cumulative: %s", FormatCumulative(a.Cumulative))
	}
	if !a.Selector.StrengthenedByDegreeType || a.Selector.TargetFunctorNative {
		t.Fatalf("bad selector: %s", FormatSelector(a.Selector))
	}
}

func TestGate924MuBAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.MuB.SelectorInputExteriorTyped || !a.MuB.FunctionhoodGapWeakened || a.MuB.NativeMuB {
		t.Fatalf("bad muB: %s", FormatMuB(a.MuB))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureExposureLanguageNotTargetFunctor, FailureEnclosureLanguageNotTargetFunctor, FailureExteriorContrastDoesNotSelectZ2, FailureNoNativeTopDegreeToF2OverF0, FailureSelectorFunctionhoodTargetFunctor, FailureMuBStillNotNativeTargetFunctor, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3} {
		if !containsAll(a.Firewalls.List(), []string{want}) {
			t.Fatalf("missing firewall %s", want)
		}
	}
}

func TestGate924Theorem(t *testing.T) {
	res := Generation2BoundaryDegreeExposureEnclosureFunctorAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range append(append(Statuses(), Supports()...), Failures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
	for _, want := range []string{FinalTruth, Classification, ShortStatus, BoundaryPair, ExteriorLedger, ReducedResponse, DegreeOneTerm, DegreeTwoTerm, TargetFunctorGap, NextGate, SourceNativeExteriorShape, SourceBridgeTargetBlocked} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}
