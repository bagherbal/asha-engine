package generation2halftraceboundarycoordinateweightaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate656Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.FanoSealDefined || !a.Inherited.FanoSealInternalOnly || !a.Inherited.FanoStructuresNumerator || !a.Inherited.NoBoundaryInterface || !a.Inherited.NoSevenOver72Theorem || !a.Inherited.NoBoundaryStress || !a.Inherited.NoScalarFlavorMap || !a.Inherited.NoHistoryLoopSource || a.Inherited.ClaimsBoundaryStress || a.Inherited.ClaimsSevenOver72 || a.Inherited.ClaimsScalarFlavor || a.Inherited.ClaimsHistoryLoopUnit || !a.Inherited.Gate655Firewall {
		t.Fatalf("bad Gate655 inheritance: %+v", a.Inherited)
	}
	if !near(a.SourceType.FullWeight, 7.0/72.0) || !near(a.SourceType.HalfWeight, 7.0/144.0) || !a.SourceType.SevenTyped || !a.SourceType.SeventyTwoTyped || !a.SourceType.HalfTyped || a.SourceType.HalfNative || !a.SourceType.AllFactorsTyped || a.SourceType.CertifiedHalfTraceMap {
		t.Fatalf("bad source type audit: %+v", a.SourceType)
	}
	if len(a.Boundary.Rows) != 3 || a.Boundary.ClosestTarget != "|lambda(Lambda_12)|" || !near(a.Boundary.ClosestResidual, math.Abs(absLambda-wHalf)) || a.Boundary.CertifiedMatch || !a.Boundary.NoProximityCertification {
		t.Fatalf("bad boundary comparison: %+v", a.Boundary)
	}
	if !near(a.MeanStress.HalfWeight, wHalf) || !near(a.MeanStress.XiBoundary, xiBoundary) || a.MeanStress.SignedResidual <= 0 || !a.MeanStress.ExistingMeanStressBetter || !a.MeanStress.AntiAlignmentSealStronger {
		t.Fatalf("bad mean stress audit: %+v", a.MeanStress)
	}
	if !near(a.Split.FullWeight, wFull) || !near(a.Split.HalfWeight, wHalf) || !near(a.Split.SignedPair[0], wHalf) || !near(a.Split.SignedPair[1], -wHalf) || !a.Split.FullWeightTyped || !a.Split.PerCoordinateTyped || !a.Split.SignedPairTyped || !a.Split.MeanStressTyped || a.Split.SuppliesBoundaryMap || a.Split.SuppliesTraceTheorem {
		t.Fatalf("bad split audit: %+v", a.Split)
	}
	if len(a.Relations.Rows) != 4 || !a.Relations.FanoHitchinSource || a.Relations.HistoryLoopSource || a.Relations.BoundaryStressSource || a.Relations.OrientationBalanceSource {
		t.Fatalf("bad relations audit: %+v", a.Relations)
	}
	if a.BoundaryMap.HasHalfTraceMap || a.BoundaryMap.HasSevenOver72Map || a.BoundaryMap.HasBoundaryStressMap || a.BoundaryMap.CanDeriveBoundaryStress || a.BoundaryMap.CanDeriveLambdaOrR3 {
		t.Fatalf("bad boundary map audit: %+v", a.BoundaryMap)
	}
	if a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsLambdaR3 || a.Firewalls.ClaimsSevenOver144 || a.Firewalls.ClaimsSevenOver72 || a.Firewalls.ClaimsHistoryLoopUnit || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.Verdict != StatusGate656Boundary {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2HalfTraceBoundaryCoordinateWeightAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate655FanoHitchinSealInherited, StatusSourceTypeAudited, StatusBoundaryComparisonAudited, StatusMeanStressAudited, StatusTwoCoordinateSplitAudited, StatusRelationToPreviousSealsAudited, StatusHalfTraceCandidate, StatusFanoNumeratorStrengthensClue, StatusBoundaryClueOnly, StatusNoNativeHalfTraceMap, StatusNoSevenOver144Theorem, StatusNoSevenOver72Theorem, StatusNoBoundaryStressFromK7, StatusNoBoundaryStressDerived, StatusNoHistoryLoopSource, StatusNoScalarFlavorMap, StatusGate656Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
