package innerfluctuationedgemeasure

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func InnerFluctuationOneFormSupportCCMEdgeMeasureSelectionSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-INNER-FLUCTUATION-ONE-FORM-CCM-EDGE-MEASURE-SELECTION-SIEVE"
	const name = "Inner Fluctuation 1-Form Support / CCM Edge Measure Selection Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build inner fluctuation edge-measure audit", Passed: false, Detail: err.Error()}}}
		}
		c := a.Calculation
		checks := []theorem.Check{
			{Name: "Higgs field formalized as finite inner fluctuation", Passed: c.InnerFluctuation.UsesCommutatorDF && c.InnerFluctuation.IsFiniteOneForm, Detail: c.InnerFluctuation.Verdict},
			{Name: "one-form support has ten J-doubled edge slots", Passed: len(c.Support.Edges) == int(JDoubledEdgeCount) && math.Abs(c.Support.NodeCount-7) < 1e-12 && math.Abs(c.Support.EdgeCount-10) < 1e-12, Detail: c.Support.SupportProjectionFormula},
			{Name: "canonical one-form trace restricts to edge support", Passed: c.Support.AEqualsPEAPE && c.Support.EdgeMeasureMandated && !c.Support.NodeMeasureAdmissible, Detail: c.Support.Verdict},
			{Name: "CCM edge measure selection theorem derived", Passed: c.Theorem.Proven && c.Theorem.AvoidsDoubleCount && strings.Contains(StatusLine(c), StatusCCMEdgeMeasureSelectionDerived), Detail: c.Theorem.Verdict},
			{Name: "Gate 384 raw trace recomputation inherited without post-hoc multiplier", Passed: strings.Contains(StatusLine(c), StatusRawTraceGate384Inherited) && strings.Contains(StatusLine(c), StatusNoPostHocMultiplierPreserved), Detail: "The selected lane uses R_edge=(7/10)R_node inside e/a²."},
			{Name: "Higgs tree-level CCM+Pfaffian proxy is sealed", Passed: c.HiggsTreeProxySealed && math.Abs(c.Higgs.MassPfaffianGeV-HiggsTargetGeV) < 0.3 && strings.Contains(StatusLine(c), StatusHiggsProxyGeometricallySealed), Detail: c.Higgs.Verdict},
			{Name: "CCM f0 remains separate from edge count", Passed: strings.Contains(StatusLine(c), StatusTensionContinuumMomentSeparate), Detail: "Gate 385 does not redefine f0; it selects the finite one-form trace measure."},
			{Name: "physical pole mass and full numerical ToE closure are not overclaimed", Passed: !c.PhysicalPoleMassDerived && !c.FullNumericalTOEClosed && strings.Contains(StatusLine(c), StatusFailedPhysicalPoleMassNotDerived), Detail: c.Boundary.FinalConclusion},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{c.Truth}}
	}}
}
