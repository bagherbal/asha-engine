package normalizationfactoraudit

import (
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CompleteNormalizationFactorAuditProductSpectralActionConventionSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-COMPLETE-NORMALIZATION-FACTOR-AUDIT"
	const name = "Complete Normalization Factor Audit / Product Spectral Action Convention Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build normalization audit", Passed: false, Detail: err.Error()}}}
		}
		au := a.Audit
		checks := []theorem.Check{
			{Name: "six normalization targets are audited", Passed: au.Executed && len(au.Factors) == 6, Detail: FormatAudit(au)},
			{Name: "4D heat-kernel volume factor is fixed", Passed: math.Abs(au.Factors[0].Numeric-1/(16*math.Pi*math.Pi)) < 1e-15, Detail: FormatFactor(au.Factors[0])},
			{Name: "Lichnerowicz term is corrected into the a2 Dirac curvature coefficient", Passed: math.Abs(au.Factors[1].Numeric-1.0/12.0) < 1e-15, Detail: FormatFactor(au.Factors[1])},
			{Name: "reality trace alternatives are computed rather than assumed universal", Passed: au.EHWithDoubledTrace.TraceDimension == 96 && au.EHWithRealityHalfTrace.TraceDimension == 48 && au.EHWithDoubledTrace.CoefficientPerMP2 > au.EHWithRealityHalfTrace.CoefficientPerMP2, Detail: FormatEH(au.EHWithDoubledTrace) + "\n" + FormatEH(au.EHWithRealityHalfTrace)},
			{Name: "current f2 cutoff moment does not canonically normalize Einstein-Hilbert", Passed: au.EHWithDoubledTrace.CurrentF2ShortBy > 200 && au.EHWithRealityHalfTrace.CurrentF2ShortBy > 400 && !au.ClosesEH, Detail: FormatEH(au.EHWithDoubledTrace) + "\n" + FormatEH(au.EHWithRealityHalfTrace)},
			{Name: "f0=7 is kept out of the Einstein-Hilbert a2 channel", Passed: au.Factors[3].Channel != au.Factors[4].Channel && au.Factors[3].Numeric == 7, Detail: FormatFactor(au.Factors[3])},
			{Name: "absolute gauge branch requires representation trace normalization", Passed: au.Gauge.RequiredRepresentationTraceK > 45 && au.Gauge.RequiredRepresentationTraceK < 46, Detail: au.Gauge.Formula},
			{Name: "single six-factor product is rejected as channel mixing", Passed: au.ChannelSeparated && !au.FullNumericalClosure, Detail: au.Verdict},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{au.Truth}}
	}}
}
