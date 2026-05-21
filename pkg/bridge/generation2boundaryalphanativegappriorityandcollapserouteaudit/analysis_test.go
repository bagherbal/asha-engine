package generation2boundaryalphanativegappriorityandcollapserouteaudit

import (
	"strings"
	"testing"
)

func TestGate919GapIndependenceCollapseCandidate(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.GapIndependence.InheritedStatus != Gate918ShortStatus || len(a.GapIndependence.GapNames) != 5 || !a.GapIndependence.MayCollapseToMeasure || a.GapIndependence.NativeCertified {
		t.Fatalf("bad gap independence audit: %s", FormatGapIndependence(a.GapIndependence))
	}
	for _, want := range []string{"boundary degree", "boundary response", "boundary target", "boundary normalization"} {
		if !strings.Contains(FormatGapIndependence(a.GapIndependence), want) {
			t.Fatalf("missing shared structure %s: %s", want, FormatGapIndependence(a.GapIndependence))
		}
	}
}

func TestGate919PriorityRanking(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.PriorityRanking.RankingComplete || len(a.PriorityRanking.OrderedGaps) != 5 || a.PriorityRanking.OrderedGaps[0].Failure != FailureNoNativeSSplitTransportMap || a.PriorityRanking.OrderedGaps[4].Failure != FailureNoNativeZ2CrossLaneExclusionTheorem {
		t.Fatalf("bad priority ranking: %s", FormatPriorityRanking(a.PriorityRanking))
	}
	if !a.PriorityRanking.CrossLaneDependent || a.PriorityRanking.BoundaryActivationRank != 2 {
		t.Fatalf("bad dependent/collapse priorities: %s", FormatPriorityRanking(a.PriorityRanking))
	}
}

func TestGate919CollapseRouteReconstructsAlpha(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.CollapseRoute
	if c.MasterObject != BoundaryActivationMeasureName || c.AlternateName != BoundaryResponseFunctorName || !c.ReassemblesAllFive || c.NativeTheorem {
		t.Fatalf("bad collapse route flags: %s", FormatCollapseRoute(c))
	}
	if c.RankI1 != 3 || c.RankI2 != 7 || c.RankH1 != 10 || c.RankH2 != 72 {
		t.Fatalf("bad ranks: %s", FormatCollapseRoute(c))
	}
	if !near(c.LinearContribution, AlphaLinear) || !near(c.QuadraticContribution, AlphaQuad) || !near(c.Alpha, AlphaB) {
		t.Fatalf("bad alpha reconstruction: %s", FormatCollapseRoute(c))
	}
}

func TestGate919MasterRequirementsOpen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Requirements.AllRequired || a.Requirements.AllCertified || len(a.Requirements.Requirements) != 6 {
		t.Fatalf("bad requirements: %s", FormatRequirements(a.Requirements))
	}
	for _, want := range []string{"source", "parameter", "degree", "target", "normalizer", "exclusion"} {
		if !strings.Contains(FormatRequirements(a.Requirements), want) {
			t.Fatalf("missing requirement %s: %s", want, FormatRequirements(a.Requirements))
		}
	}
}

func TestGate919PromotionAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Promotion.BridgeCandidate || a.Promotion.NativeAlpha || a.Promotion.NativeR3 || a.Promotion.OfficialUpdate {
		t.Fatalf("bad promotion status: %s", FormatPromotion(a.Promotion))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureNoNativeBoundaryActivationMeasureCertified, FailureNoNativeBoundaryResponseMeasure, FailureMuBFormalNotNative, FailureNoNativeSSplitTransportMap, FailureNoNativeReducedB2ResponseFunctional, FailureNoNativeDegreeToZ2FlagClassFunctor, FailureNoNativeZ2CrossLaneExclusionTheorem, FailureNoNativeResponseChamberNormalization, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator} {
		if !containsAll(a.Firewalls.List(), []string{want}) {
			t.Fatalf("missing firewall %s", want)
		}
	}
}

func TestGate919Theorem(t *testing.T) {
	res := Generation2BoundaryAlphaNativeGapPriorityAndCollapseRouteAuditTheorem().Verify()
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
	for _, want := range []string{FinalTruth, Classification, ShortStatus, StrategicConclusion, NextGate, BoundaryActivationMeasureName, BoundaryResponseFunctorName} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}
