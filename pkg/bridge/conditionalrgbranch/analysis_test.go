package conditionalrgbranch

import (
	"math"
	"testing"
)

func TestBuildDefaultConditionalRGBranch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Input.Gate175ConditionalBranchAvailable || a.Input.Gate175StrictAbsoluteUDerived {
		t.Fatalf("expected inherited conditional but non-strict branch: %+v", a.Input)
	}
	if math.Abs(a.Input.ConditionalUInverseGStar-1) > 1e-12 {
		t.Fatalf("expected conditional u=1, got %.12g", a.Input.ConditionalUInverseGStar)
	}
}

func TestObservedLedgerQuarantined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Observed.QuarantinedComparison || a.Observed.UsedForDerivation || a.Input.UsesObservedInputForDerivation {
		t.Fatalf("observed ledger must be comparison-only: obs=%+v input=%+v", a.Observed, a.Input)
	}
	if a.Observed.Alpha2Inverse <= 0 || a.Observed.Alpha1GUTInverse <= 0 || a.Observed.Alpha3Inverse <= 0 {
		t.Fatalf("bad observed comparison ledger: %+v", a.Observed)
	}
}

func TestSingleObservableFitsRejectUOneBranch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Fits) != 4 {
		t.Fatalf("expected four fits, got %d", len(a.Fits))
	}
	for _, f := range a.Fits {
		if f.PassesLooseRange {
			t.Fatalf("no single-observable fit should pass simultaneous loose range: %s", FormatFit(f))
		}
	}
	if !a.Firewall.ConditionalBranchRejectedByMZ || a.Firewall.AnyObservedFitViable {
		t.Fatalf("expected conditional branch rejection at MZ: %+v", a.Firewall)
	}
}

func TestAlpha3FitPositiveButNotElectroweakViable(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Fits[0]
	if f.TargetObservable != "α3(M_Z)" || !f.FeasiblePositive || !f.FeasibleLPositive {
		t.Fatalf("alpha3 fit should be positive and L>=0: %s", FormatFit(f))
	}
	if math.Abs(f.Alpha3Residual) > 1e-10 {
		t.Fatalf("alpha3 fit should hit alpha3 by construction: %s", FormatFit(f))
	}
	if math.Abs(f.AlphaEMResidual) < 0.5 || math.Abs(f.Alpha2Residual) < 0.5 {
		t.Fatalf("alpha3 fit should badly miss electroweak inverses: %s", FormatFit(f))
	}
}

func TestRatioOnlyCheckFailsWithoutThresholds(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ratio.IndependentOfU || !a.Ratio.UsesGUTNormalizedAlpha1 || !a.Ratio.ObservedComparisonOnly {
		t.Fatalf("bad ratio audit flags: %+v", a.Ratio)
	}
	if a.Ratio.LIntervalsAgree || a.Firewall.RatioCheckPasses {
		t.Fatalf("ratio-only check should fail under unthresholded one-loop branch: %+v", a.Ratio)
	}
	if math.Abs(a.Ratio.TheoryRatio-(23.0/6.0)/(109.0/15.0)) > 1e-12 {
		t.Fatalf("unexpected theory ratio: %+v", a.Ratio)
	}
}

func TestNullityNotReducedByPhenomenology(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Firewall.StrictNullityBefore != 3 || a.Firewall.StrictNullityAfter != 3 {
		t.Fatalf("strict nullity should remain 3: %+v", a.Firewall)
	}
	if a.Firewall.ConditionalNullityBefore != 2 || a.Firewall.ConditionalNullityAfter != 2 {
		t.Fatalf("conditional nullity ledger should remain 2 before and after rejection: %+v", a.Firewall)
	}
	if a.Firewall.PhysicalConstantsDerived || a.Firewall.HiddenObservedInputUsedForDerivation {
		t.Fatalf("no physical constants should be derived: %+v", a.Firewall)
	}
}
