package generation2boundaryweighteddeficitclosurerobustnessaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate661Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ActiveWeightInherited || math.Abs(a.Inherited.W72-0.04982659643506822) > 5e-15 || math.Abs(a.Inherited.KSum-0.04982659728765166) > 5e-15 || math.Abs(a.Inherited.WeightedResidual-8.525834413464217e-10) > 5e-18 || !a.Inherited.FormulaLiftBridgeLayerOnly || !a.Inherited.NoNativeSevenOver72Theorem || !a.Inherited.NoNativeK7BoundaryMap || !a.Inherited.NoNativeTransportTheorem || !a.Inherited.NoFanoHitchinBoundaryRevival || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if len(a.Dependency.Nodes) != 9 || !a.Dependency.KappaLambdaDefinedFromRuntime || !a.Dependency.LambdaLambda12DependsOnRuntime || !a.Dependency.W72DependsOnBoundaryEndpoints || !a.Dependency.FormulaLiftPartlyTautological || !strings.Contains(a.Dependency.NontrivialStatement, "lambda_runtime") {
		t.Fatalf("bad dependency graph: %+v", a.Dependency)
	}
	if math.Abs(a.Closure.ClosureResidualExact-8.525834413464217e-10) > 5e-18 || a.Closure.RelativeToBoundarySplit > 1e-6 || a.Closure.FormulaLiftIndependent || !strings.Contains(a.Closure.NontrivialBridgeStatement, "7/72") {
		t.Fatalf("bad nontrivial closure: %+v", a.Closure)
	}
	if math.Abs(a.Orientation.KappaEDifference-2.775873137889728e-06) > 5e-18 || math.Abs(a.Orientation.ClosureResidualOrientation-2.7767257213331953e-06) > 5e-18 || a.Orientation.RelativeResidualOrientationToW72 >= 6e-5 || a.Orientation.RelativeResidualOrientationToSplit >= 0.003 || a.Orientation.ExactToOrientationResidualRatio <= 3000 {
		t.Fatalf("bad orientation audit: %+v", a.Orientation)
	}
	if len(a.Uncertainty.Slots) != 5 || a.Uncertainty.FullPropagationAvailable || a.Uncertainty.InventedUncertainties || a.Uncertainty.ClosureSignificanceCertified {
		t.Fatalf("bad uncertainty audit: %+v", a.Uncertainty)
	}
	if len(a.Scale.Rows) != 5 || !a.Scale.Lambda12OnlyComputed || a.Scale.NearbyScaleSweepAvailable || a.Scale.EndpointIndependenceCertified {
		t.Fatalf("bad scale audit: %+v", a.Scale)
	}
	if len(a.Weights.Rows) != 6 || a.Weights.BestName != "7/72" || a.Weights.BestResidual > 1e-9 || a.Weights.SecondBestName != "1/10" || a.Weights.ImprovementOverSecond < 4000 || !a.Weights.NoArbitrarySearch {
		t.Fatalf("bad weight audit: %+v", a.Weights)
	}
	if !a.Discipline.ClassifiesRobustV1ExactLedger || !a.Discipline.ClassifiesPendingUncertaintySweep || a.Discipline.ClaimsNativeSevenOver72Theorem || a.Discipline.ClaimsNativeTransportTheorem || a.Discipline.ClaimsIndependentEndpointDerivation || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsHiggsPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsFanoHitchinBoundaryMap || a.Discipline.Verdict != StatusGate661Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2BoundaryWeightedDeficitClosureRobustnessAndNoncircularityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate660ActiveWeightInherited, StatusDependencyGraphAudited, StatusNontrivialClosureIsolated, StatusScalarFormulaLiftCircularityAudited, StatusOrientationApproximationAudited, StatusUncertaintySlotsDefined, StatusScaleSensitivitySlotsDefined, StatusTypedWeightUniquenessAudited, StatusClosureRobustInV1, StatusV1PrecisionCluePendingUncertaintySweep, StatusOrientationApproxStillSmall, StatusSevenOver72TypedWeightBest, StatusFormulaLiftNotIndependentEvidence, StatusNoNativeSevenOver72Theorem, StatusNoNativeScalarFlavorBoundaryTheorem, StatusNoIndependentEndpointDerivation, StatusNoFullUncertaintyLedger, StatusNoScaleSweepData, StatusNoBoundaryStressDerivation, StatusNoHiggsFlavorGaugeClaim, StatusGate661Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
