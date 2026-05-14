package pfaffianhierarchy

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

func TestPfaffianHalfAction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	want := math.Exp(-4 * math.Pi * math.Pi)
	if !a.Pfaffian.HalfActionAuthorized || !close(a.Pfaffian.HalfExponential, want, 1e-30) {
		t.Fatalf("bad pfaffian ledger: %s", FormatPfaffian(a.Pfaffian))
	}
	if a.Pfaffian.FiniteCoreDerived {
		t.Fatalf("pfaffian rule should remain continuum-measure, not finite-core: %s", FormatPfaffian(a.Pfaffian))
	}
}

func TestGenerationFactor(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !close(a.Generation.CombinedFactor, math.Sqrt(8), 1e-15) {
		t.Fatalf("bad generation factor: %s", FormatGeneration(a.Generation))
	}
	if a.Generation.FiniteCoreDerived {
		t.Fatalf("zero-mode normalization should remain firewalled: %s", FormatGeneration(a.Generation))
	}
}

func TestCombinedHierarchy(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !close(a.Prediction.PredictedRatio, 2.024352198454697e-17, 1e-31) {
		t.Fatalf("unexpected predicted hierarchy: %s", FormatPrediction(a.Prediction))
	}
	if !close(a.Prediction.RatioToUnreducedTarget, 1.0037817218631122, 1e-15) {
		t.Fatalf("unexpected unreduced agreement: %s", FormatPrediction(a.Prediction))
	}
	if a.Prediction.RatioToReducedTarget > 0.201 || a.Prediction.RatioToReducedTarget < 0.200 {
		t.Fatalf("reduced Planck branch should not match: %s", FormatPrediction(a.Prediction))
	}
}

func TestStatusesAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	statuses := Statuses(a)
	required := []string{StatusCombinedHierarchyComputed, StatusTensionPfaffianMechanismExternal, StatusFailedUnconditionalHierarchyNotClaimed, StatusFailedF2MomentStillNotLocked}
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
	res := PfaffianHalfActionHierarchyFermionicFluctuationDeterminantTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
