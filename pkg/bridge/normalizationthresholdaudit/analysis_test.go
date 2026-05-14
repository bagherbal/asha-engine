package normalizationthresholdaudit

import (
	"math"
	"testing"
)

func TestBuildDefaultNormalizationThresholdAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Input.Gate176ConditionalUOneRejected || !a.Input.Gate176StrictUStillOpen || !a.Input.Gate176RatioOnlyCheckFailed {
		t.Fatalf("expected inherited Gate 176 rejection and open strict u: %+v", a.Input)
	}
	if !a.Input.ObservedComparisonQuarantined || a.Input.UsesObservedInputForFiniteTheorem {
		t.Fatalf("observed ledger must remain comparison-only: %+v", a.Input)
	}
}

func TestNormalizationOnlyOverconstrained(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Normalization.ExactTripleFit {
		t.Fatalf("normalization-only branch should not solve all three sectors: %s", FormatNormalization(a.Normalization))
	}
	if a.Normalization.Unknowns != 2 || a.Normalization.Equations != 3 {
		t.Fatalf("expected 2 unknowns and 3 equations: %+v", a.Normalization)
	}
	if !a.Normalization.PositiveU || !a.Normalization.PositiveL {
		t.Fatalf("least-squares normalization witness should be positive: %s", FormatNormalization(a.Normalization))
	}
	if a.Normalization.PairLogIntervalsConsistent {
		t.Fatalf("pairwise L values should remain inconsistent: %s", FormatNormalization(a.Normalization))
	}
	if a.Normalization.MaxRelativeResidual < 0.01 {
		t.Fatalf("expected visible residual for best fit: %s", FormatNormalization(a.Normalization))
	}
}

func TestUniversalThresholdAddsNoRatioFreedom(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Universal.AddsSectorRatioFreedom || !a.Universal.EquivalentToInterceptShift || a.Universal.CanRepairGate176Failure {
		t.Fatalf("universal threshold should not repair relative running: %+v", a.Universal)
	}
	if !a.Universal.PairLogIntervalsStillInconsistent || !a.Universal.RatioOnlyMismatchStillPresent {
		t.Fatalf("universal shift should preserve the ratio obstruction: %+v", a.Universal)
	}
}

func TestNonUniversalThresholdFamilyIsFitButUnderived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	th := a.Thresholds
	if !th.FitsExactlyForAnyChosenPositiveL || !th.CanRepairPhenomenologyByFit {
		t.Fatalf("non-universal threshold family should fit by construction: %s", FormatThresholds(th))
	}
	if !th.UnderdeterminedWithoutFiniteRule || th.FiniteThresholdOperatorDerived || th.CanReduceStrictNullity {
		t.Fatalf("threshold family must remain underived and unable to reduce nullity: %s", FormatThresholds(th))
	}
	if th.MinimumNormForUOne.LogIntervalL <= 0 || th.MinimumNormForUOne.EuclideanNorm <= 0 {
		t.Fatalf("bad minimum-norm threshold witness: %s", FormatThresholdVector(th.MinimumNormForUOne))
	}
	if th.MinimumNormForUOne.SignPatternPreserved {
		t.Fatalf("minimum-norm u=1 threshold witness should alter the SM beta sign pattern: %s", FormatThresholdVector(th.MinimumNormForUOne))
	}
}

func TestAlpha3ThresholdWitnessKeepsDeltaB3Zero(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	w := a.Thresholds.Alpha3FitForUOne
	if math.Abs(w.DeltaB3) > 1e-9 {
		t.Fatalf("alpha3-fit u=1 threshold witness should leave Δb3 approximately zero: %s", FormatThresholdVector(w))
	}
	if w.DeltaB1 < 10 || w.DeltaB2 < 10 {
		t.Fatalf("alpha3-fit u=1 witness should require large EW threshold shifts: %s", FormatThresholdVector(w))
	}
}

func TestFirewallNoNullityReduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	fw := a.Firewall
	if fw.StrictNullityBefore != 3 || fw.StrictNullityAfter != 3 || fw.ConditionalNullityBefore != 2 || fw.ConditionalNullityAfter != 2 {
		t.Fatalf("nullity should not reduce: %+v", fw)
	}
	if fw.PhysicalConstantsDerived || fw.BoundaryScaleDerivedStrict || fw.ThresholdCorrectionsDerived || fw.HiddenObservedInputUsedForDerivation {
		t.Fatalf("firewall violation: %+v", fw)
	}
	if !fw.NonUniversalThresholdCanFitByConstruction || fw.NonUniversalThresholdDerived {
		t.Fatalf("expected fit family but no derived threshold theorem: %+v", fw)
	}
}
