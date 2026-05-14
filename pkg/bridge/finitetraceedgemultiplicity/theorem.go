package finitetraceedgemultiplicity

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteTraceEdgeMultiplicityEffectiveCoefficientSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-TRACE-EDGE-MULTIPLICITY-EFFECTIVE-COEFFICIENT-SIEVE"
	const name = "Finite Trace Edge Multiplicity / Effective Coefficient Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build finite trace edge multiplicity ledger", Passed: false, Detail: err.Error()}}}
		}
		c := a.Calculation
		var unit, numerator, denom, contact CoefficientLane
		for _, l := range c.Lanes {
			switch l.Name {
			case "unit f₀, no extra multiplicity":
				unit = l
			case "wrong numerator multiplication":
				numerator = l
			case "denominator edge normalization witness":
				denom = l
			case "contact f₀=7 ledger":
				contact = l
			}
		}
		checks := []theorem.Check{
			{Name: "continuous CCM moment is locked to f0=1", Passed: c.Executed && math.Abs(c.Moment.LockedValue-1) < 1e-12 && strings.Contains(c.Moment.CCMF0Definition, "f(0)"), Detail: c.Moment.Verdict},
			{Name: "J-doubled finite edge multiplicity ten is inherited", Passed: c.EdgeMultiplicity.InheritedFromGate381 && c.EdgeMultiplicity.JDoubledEdgeCount == 10 && math.Abs(c.EdgeMultiplicity.ProjectionTrace-10) < 1e-12 && !c.EdgeMultiplicity.CanReplaceF0, Detail: c.EdgeMultiplicity.Verdict},
			{Name: "finite Higgs trace ratio is treated as already traced", Passed: c.TraceDecomposition.IsAlreadyTraceRatio && !c.TraceDecomposition.CanPullExtraTenFromRatio && math.Abs(c.TraceDecomposition.FiniteRatioValue-TraceRatioEOverA2) < 1e-15, Detail: c.TraceDecomposition.Verdict},
			{Name: "unit f0 alone overpredicts Higgs", Passed: unit.MassPfaffianGeV > 350 && unit.Native && !unit.CircularRisk, Detail: unit.Verdict},
			{Name: "edge multiplicity in numerator is rejected", Passed: numerator.MassPfaffianGeV > 1000 && !numerator.Native && numerator.CircularRisk, Detail: numerator.Verdict},
			{Name: "edge multiplicity in denominator near-closes but requires theorem", Passed: math.Abs(denom.EffectiveDenom-10) < 1e-12 && math.Abs(denom.MassPfaffianGeV-HiggsBoundaryGeV) < 0.3 && !denom.Native && denom.CircularRisk, Detail: denom.Verdict},
			{Name: "contact f0=7 overpredicts relative to target", Passed: contact.MassPfaffianGeV > 145 && contact.Native, Detail: contact.Verdict},
			{Name: "ten-over-seven gap is isolated, not derived", Passed: math.Abs(c.Gap.RatioTenOverSeven-(10.0/7.0)) < 1e-12 && !c.Gap.RecognizedNative && strings.Contains(StatusLine(c), StatusFailedTenOverSevenNotDerived), Detail: c.Gap.Verdict},
			{Name: "Higgs mass geometric sealing is not claimed", Passed: !c.EdgeMultiplicityExtracted && !c.HiggsMassSealed && !c.FullNumericalTOEClosure && strings.Contains(StatusLine(c), StatusFailedHiggsMassNotGeometricallySealed), Detail: c.Truth},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{c.Truth}}
	}}
}
