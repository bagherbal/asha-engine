package normalizationthresholdaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func NormalizationPrefactorThresholdDeformationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-NORMALIZATION-PREFACTOR-THRESHOLD-DEFORMATION-AUDIT"
	const name = "normalization-prefactor or threshold-deformation branch audit after u=1 rejection"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build normalization/threshold audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 176 failure is inherited only as quarantined comparison data", Passed: a.Input.Gate176ConditionalUOneRejected && a.Input.Gate176StrictUStillOpen && a.Input.Gate176RatioOnlyCheckFailed && a.Input.ObservedComparisonQuarantined && !a.Input.UsesObservedInputForFiniteTheorem, Detail: FormatInput(a.Input)},
			{Name: "normalization prefactor alone is overconstrained", Passed: a.Normalization.Unknowns == 2 && a.Normalization.Equations == 3 && !a.Normalization.ExactTripleFit && !a.Normalization.PairLogIntervalsConsistent && a.Normalization.PositiveU && a.Normalization.PositiveL, Detail: FormatNormalization(a.Normalization)},
			{Name: "universal threshold shift cannot repair relative running", Passed: !a.Universal.AddsSectorRatioFreedom && a.Universal.EquivalentToInterceptShift && a.Universal.PairLogIntervalsStillInconsistent && a.Universal.RatioOnlyMismatchStillPresent && !a.Universal.CanRepairGate176Failure, Detail: FormatUniversal(a.Universal)},
			{Name: "non-universal thresholds can fit only as underived deformation family", Passed: a.Thresholds.FitsExactlyForAnyChosenPositiveL && a.Thresholds.UnderdeterminedWithoutFiniteRule && a.Thresholds.CanRepairPhenomenologyByFit && !a.Thresholds.FiniteThresholdOperatorDerived && !a.Thresholds.CanReduceStrictNullity, Detail: FormatThresholds(a.Thresholds)},
			{Name: "minimum-norm u=1 threshold witness is not finite-derived", Passed: a.Thresholds.MinimumNormForUOne.UInverseGStar == 1 && a.Thresholds.MinimumNormForUOne.LogIntervalL > 0 && !a.Thresholds.MinimumNormForUOne.FiniteDerived && !a.Thresholds.MinimumNormForUOne.SignPatternPreserved, Detail: FormatThresholdVector(a.Thresholds.MinimumNormForUOne)},
			{Name: "strict nullity and beta firewall remain closed", Passed: !a.Firewall.NormalizationPrefactorAloneSufficient && !a.Firewall.UniversalThresholdAloneSufficient && a.Firewall.NonUniversalThresholdCanFitByConstruction && !a.Firewall.NonUniversalThresholdDerived && a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3 && !a.Firewall.PhysicalConstantsDerived && !a.Firewall.ThresholdCorrectionsDerived && !a.Firewall.HiddenObservedInputUsedForDerivation, Detail: FormatFirewall(a.Firewall)},
		}, Notes: []string{
			a.TruthStatement,
			"Gate 177 does not fit physical constants into the finite core. It only classifies which deformation type would be mathematically capable of repairing the Gate 176 comparison failure.",
			"The next required object is a finite threshold/decoupling operator; without it, arbitrary Δb_i vectors are phenomenological parameters, not theorem data.",
		}}
	}}
}
