package rawfinitetracerecomputation

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func RawFiniteTraceRecomputationEdgeMeasureSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-RAW-FINITE-TRACE-RECOMPUTATION-EDGE-MEASURE-SIEVE"
	const name = "Raw Finite Trace Recomputation / Edge Measure Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build raw finite trace recomputation audit", Passed: false, Detail: err.Error()}}}
		}
		c := a.Calculation
		lanes := map[string]HiggsLane{}
		for _, l := range c.Lanes {
			lanes[l.Name] = l
		}
		edgeRatio := TraceRatioNode * ContactNodeCount / JDoubledEdgeCount
		checks := []theorem.Check{
			{Name: "raw node and edge trace symbolics reconstructed", Passed: c.Executed && c.Symbolics.UniformMeasureLift && math.Abs(c.Symbolics.MeasureScale-10.0/7.0) < 1e-12, Detail: c.Symbolics.Verdict},
			{Name: "edge measure reduces e/a² by 7/10", Passed: math.Abs(c.Measures[1].RatioEOverA2-edgeRatio) < 1e-12 && math.Abs(c.Measures[1].RatioRelativeToNode-0.7) < 1e-12, Detail: c.Measures[1].Verdict},
			{Name: "10/7 appears inside the raw ratio rather than post-hoc", Passed: c.EdgeMeasureTraceComputed && c.TenOverSevenDerivedInsideRatio && strings.Contains(StatusLine(c), StatusNoPostHocMultiplier), Detail: c.DoubleCount.Verdict},
			{Name: "edge-measure raw-ratio lane near-closes Higgs", Passed: math.Abs(lanes["edge-measure raw ratio with inherited node normalization"].MassPfaffianGeV-HiggsTargetGeV) < 0.3 && !lanes["edge-measure raw ratio with inherited node normalization"].DoubleCounts, Detail: lanes["edge-measure raw ratio with inherited node normalization"].Verdict},
			{Name: "literal f0=1 alone does not close", Passed: lanes["literal CCM f0=1 with edge-measure ratio only"].MassPfaffianGeV > 320 && strings.Contains(StatusLine(c), StatusTensionLiteralF0UnitDoesNotClose), Detail: lanes["literal CCM f0=1 with edge-measure ratio only"].Verdict},
			{Name: "double-counted edge lane is rejected", Passed: lanes["double-counted edge ratio plus edge denominator"].DoubleCounts && lanes["double-counted edge ratio plus edge denominator"].MassPfaffianGeV < HiggsTargetGeV-15, Detail: lanes["double-counted edge ratio plus edge denominator"].Verdict},
			{Name: "edge measure selection theorem still missing", Passed: !c.EdgeMeasureSelectedNatively && strings.Contains(StatusLine(c), StatusFailedNativeEdgeMeasureNotDerived), Detail: c.Closure.Conclusion},
			{Name: "Higgs mass geometric sealing is not claimed", Passed: !c.HiggsMassSealed && !c.FullNumericalTOEClosed && strings.Contains(StatusLine(c), StatusFailedHiggsMassNotGeometricallySealed), Detail: c.Truth},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{c.Truth}}
	}}
}
