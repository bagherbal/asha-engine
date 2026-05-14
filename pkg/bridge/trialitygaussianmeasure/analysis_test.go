package trialitygaussianmeasure

import (
	"math"
	"testing"
)

func close(a, b, tol float64) bool {
	if a > b {
		return a-b <= tol
	}
	return b-a <= tol
}

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || a.Inputs.NGen != 3 || a.Inputs.STop <= 78 {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestFiniteGrassmannMeasure(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Measure.BerezinPfaffianApplies || a.Measure.GenerationDimension != 3 {
		t.Fatalf("bad finite measure: %s", FormatMeasure(a.Measure))
	}
}

func TestZeroModePfaffianFactor(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !close(a.ZeroMode.PfaffianPerGeneration, math.Sqrt2, 1e-15) {
		t.Fatalf("bad per-generation pfaffian: %s", FormatZeroMode(a.ZeroMode))
	}
	if !close(a.ZeroMode.CombinedPfaffian, math.Sqrt(8), 1e-15) {
		t.Fatalf("bad combined pfaffian: %s", FormatZeroMode(a.ZeroMode))
	}
	if !close(a.ZeroMode.CombinedDeterminant, 8, 1e-15) {
		t.Fatalf("bad combined determinant: %s", FormatZeroMode(a.ZeroMode))
	}
}

func TestHierarchySynthesis(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !close(a.Hierarchy.PredictedRatio, 2.024352198454697e-17, 1e-31) {
		t.Fatalf("unexpected predicted ratio: %s", FormatHierarchy(a.Hierarchy))
	}
	if !close(a.Hierarchy.RatioToUnreducedTarget, 1.0037817218631122, 1e-15) {
		t.Fatalf("unexpected unreduced agreement: %s", FormatHierarchy(a.Hierarchy))
	}
	if a.Hierarchy.RatioToReducedTarget > 0.201 || a.Hierarchy.RatioToReducedTarget < 0.200 {
		t.Fatalf("reduced Planck branch should not match: %s", FormatHierarchy(a.Hierarchy))
	}
}

func TestStatusesAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	statuses := Statuses(a)
	required := []string{StatusNativeZeroModeNormalizationDerived, StatusHierarchyScalingConditionallyDerived, StatusFailedF2CutoffMomentStillUnlocked, StatusFailedUnconditionalHierarchyNotClaimed}
	for _, req := range required {
		found := false
		for _, got := range statuses {
			if got == req {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing status %s in %v", req, statuses)
		}
	}
}

func TestTheoremPasses(t *testing.T) {
	res := TrialityGaussianMeasureZeroModeNormalizationAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
