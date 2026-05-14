package threshold

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ThresholdSpectrumMatchingAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-THRESHOLD-SPECTRUM-MATCHING-AUDIT"
	const name = "finite threshold spectrum and matching audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct threshold audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "finite spectral anchors available", Passed: a.DimensionlessSpectralAnchorsAvailable, Detail: fmt.Sprintf("B-positive modes=%d, contact partial-overlap modes=%d, scalar active modes=%d", len(a.BPositiveEigenvalues), len(a.ContactPartialOverlap), len(a.ScalarActiveSpectrum))},
			{Name: "B-sector gap anchor", Passed: a.BGap > 0, Detail: fmt.Sprintf("first positive O_B eigenvalue=%.10f; positive Boolean modes=%d", a.BGap, len(a.BPositiveEigenvalues))},
			{Name: "contact partial-overlap spectrum", Passed: len(a.ContactPartialOverlap) == 7, Detail: fmt.Sprintf("partial modes=%s", FormatFloats(a.ContactPartialOverlap, 7))},
			{Name: "scalar/contact active spectrum", Passed: len(a.ScalarActiveSpectrum) == 4, Detail: fmt.Sprintf("active=%s; clusters=%s", FormatFloats(a.ScalarActiveSpectrum, 4), FormatClusters(a.ScalarClusters))},
			{Name: "dimensionless threshold candidates", Passed: len(a.Candidates) > 0, Detail: fmt.Sprintf("candidates=%d; %s", len(a.Candidates), FormatCandidates(a.Candidates, 6))},
			{Name: "dimensional-analysis firewall", Passed: !a.PhysicalMassUnitDerived, Detail: a.ThresholdMassFamily},
			{Name: "threshold activation rule", Passed: a.ThresholdActivationRuleDerived, Detail: "not derived; the engine does not yet know which finite modes are continuum-active heavy thresholds"},
			{Name: "finite-to-continuum matching", Passed: a.FiniteToContinuumMatchingDerived, Detail: "not derived; mode-to-gauge-factor matching and decoupling rules remain open"},
			{Name: "threshold-corrected beta coefficients", Passed: a.ThresholdCorrectedBetaDerived, Detail: "not derived; Gate 43 coefficients remain unthresholded continuum one-loop candidates"},
			{Name: "hidden threshold insertion", Passed: !a.HiddenThresholdScaleInserted && !a.ObservedMassesUsed, Detail: "no observed masses, GUT scale, Planck scale, or arbitrary threshold scale was inserted"},
		}, Notes: []string{a.TruthStatement, fmt.Sprintf("minimum missing data: %v", a.MinimumMissingData)}}
	}}
}
