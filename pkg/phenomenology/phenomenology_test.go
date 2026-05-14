package phenomenology

import (
	"math"
	"testing"
)

func TestVacuumFateRunsAndAppliesThreshold(t *testing.T) {
	res := ComputeVacuumFate()
	if !res.Executed || !res.ThresholdApplied {
		t.Fatalf("expected executed threshold run: %+v", res)
	}
	if !(res.InitialLambda > 0.12 && res.InitialLambda < 0.14) {
		t.Fatalf("unexpected initial lambda: %.12f", res.InitialLambda)
	}
	if !(res.LambdaAfterThreshold < res.LambdaBeforeThreshold) {
		t.Fatalf("threshold did not decrease lambda: before=%v after=%v", res.LambdaBeforeThreshold, res.LambdaAfterThreshold)
	}
	if !res.Metastable || res.InstabilityScaleGeV <= 0 {
		t.Fatalf("expected conditional metastability with crossing: %+v", res)
	}
	if math.IsNaN(res.Log10LifetimeYears) {
		t.Fatal("lifetime log is NaN")
	}
}

func TestDarkMatterYieldConstraint(t *testing.T) {
	res := ComputeDarkMatterConstraint()
	if !res.Executed || res.ExactAbundancePredicted {
		t.Fatalf("unexpected dark matter result: %+v", res)
	}
	if !(res.RequiredYield > 1e-17 && res.RequiredYield < 1e-14) {
		t.Fatalf("required yield out of expected range: %.12e", res.RequiredYield)
	}
	if !(res.OverclosureFactor > 1e10) {
		t.Fatalf("expected severe overclosure: %.12e", res.OverclosureFactor)
	}
}

func TestCosmologicalConstantFineTuning(t *testing.T) {
	res := ComputeCosmologicalConstantSubtraction()
	if !res.Executed || res.OrganicPrediction {
		t.Fatalf("unexpected cosmological result: %+v", res)
	}
	if !(res.DigitsOfCancellation > 120) {
		t.Fatalf("expected >120 digits of cancellation, got %.6f", res.DigitsOfCancellation)
	}
}

func TestReportMarkdown(t *testing.T) {
	rep := ComputeReport()
	if !rep.Executed || len(rep.Statuses) == 0 {
		t.Fatalf("bad report: %+v", rep)
	}
	md := rep.Markdown()
	if len(md) < 1000 {
		t.Fatalf("markdown too short: %d", len(md))
	}
}
