package generation2boundaryweighteddeficitclosureaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate626Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.Verdict != StatusGate625Inherited || !a.Inherited.Gate625ClosureSealDefined {
		t.Fatalf("bad Gate625 inheritance: %+v", a.Inherited)
	}
	if a.Inherited.Gate625NativeClosureTheorem || a.Inherited.Gate625NativeScalarRGMatching || a.Inherited.Gate625NativeFlavorOrientation {
		t.Fatalf("Gate625 native theorem overpromoted: %+v", a.Inherited)
	}
	if math.Abs(a.BoundarySplit.BoundarySplit-0.00129244481881633) > 1e-14 {
		t.Fatalf("bad boundary split %.18g", a.BoundarySplit.BoundarySplit)
	}
	if math.Abs(a.BoundarySplit.ClosureResidualOverSplit-0.0972228818865684) > 1e-12 {
		t.Fatalf("bad residual/split ratio %.18g", a.BoundarySplit.ClosureResidualOverSplit)
	}
	if math.Abs(a.WeightCandidate.Value-7.0/72.0) > 1e-15 {
		t.Fatalf("bad weight %.18g", a.WeightCandidate.Value)
	}
	if math.Abs(a.WeightCandidate.AbsoluteRatioResidual-6.5966e-7) > 5e-10 {
		t.Fatalf("bad ratio residual %.18g", a.WeightCandidate.AbsoluteRatioResidual)
	}
	if math.Abs(a.WeightedClosure.WeightedMixture-0.0498265964350682) > 1e-14 {
		t.Fatalf("bad weighted mixture %.18g", a.WeightedClosure.WeightedMixture)
	}
	if math.Abs(a.WeightedClosure.Residual-8.5258e-10) > 5e-14 {
		t.Fatalf("bad weighted residual %.18g", a.WeightedClosure.Residual)
	}
	if !a.WeightedClosure.ImprovesGate625 || a.WeightedClosure.ImprovementFactor < 100000 {
		t.Fatalf("weighted closure did not improve Gate625: %+v", a.WeightedClosure)
	}
	if math.Abs(a.ScalarFormula.KappaLambdaResidualExact-a.WeightedClosure.Residual) > 1e-15 {
		t.Fatalf("scalar formula residual mismatch: %+v", a.ScalarFormula)
	}
	if len(a.ScalarPrediction.Rows) != 2 || !a.ScalarPrediction.ImprovesGate625Prediction {
		t.Fatalf("bad scalar prediction: %+v", a.ScalarPrediction)
	}
	if math.Abs(a.ScalarPrediction.Rows[0].PredictedLambda-0.129652565054713) > 1e-14 {
		t.Fatalf("bad exact scalar prediction %.18g", a.ScalarPrediction.Rows[0].PredictedLambda)
	}
	if math.Abs(a.ScalarPrediction.Rows[0].Residual) > 5e-12 {
		t.Fatalf("exact weighted scalar prediction residual too large: %+v", a.ScalarPrediction.Rows[0])
	}
	if a.NativeStatus.NativeSevenOverSeventyTwoSource || a.NativeStatus.NativeGaugeScalarFlavorDeficitTransport || a.NativeStatus.NativeBoundaryWeightedClosureTheorem {
		t.Fatalf("native closure theorem incorrectly certified: %+v", a.NativeStatus)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2BoundaryWeightedDeficitClosureAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate625Inherited, StatusBoundarySplitComputed, StatusSevenOverSeventyTwoAudited, StatusBoundaryWeightedClosure, StatusSevenOverSeventyTwoCandidate, StatusScalarFormulaComputed, StatusNoNativeWeightSource, StatusNoNativeTransportTheorem, StatusGate626Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
