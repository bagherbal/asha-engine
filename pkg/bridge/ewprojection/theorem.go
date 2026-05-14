package ewprojection

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ElectroweakProjectionTheorem() theorem.Theorem {
	const id = "BRIDGE-ELECTROWEAK-PROJECTION-MIXING"
	const name = "electroweak projection and mixing-angle search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct electroweak projection audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "finite electromagnetic charge identity", Passed: a.ElectromagneticGeneratorDerived, Detail: fmt.Sprintf("Q=T3_L+Y on the derived left-doublet space; residual ||Q-(T3+Y)||_F=%.3e", a.ChargeIdentityResidual)},
			{Name: "hypercharge commutes with SU(2)_L ladders", Passed: a.HyperchargeCommutesWithSU2LNorm < 1e-8, Detail: fmt.Sprintf("||[Y,T+]+[Y,T-]||_F=%.3e", a.HyperchargeCommutesWithSU2LNorm)},
			{Name: "left-doublet trace metric", Passed: a.LeftDimension == 8 && a.T3Trace2Left > 0 && a.YTrace2Left > 0, Detail: fmt.Sprintf("dim=%d, Tr(T3²)=%.10f, Tr(Y²)=%.10f, Tr(T3Y)=%.3e, Tr(Q²)=%.10f", a.LeftDimension, a.T3Trace2Left, a.YTrace2Left, a.T3YTraceLeft, a.QTrace2Left)},
			{Name: "charge-direction diagnostic", Passed: a.TraceDirectionSin2 > 0 && a.TraceDirectionCos2 > 0, Detail: fmt.Sprintf("left trace direction: sin²_trace=Tr(Y²)/Tr(Q²)=%.10f, cos²_trace=%.10f; this is not θ_W", a.TraceDirectionSin2, a.TraceDirectionCos2)},
			{Name: "full one-generation hypercharge normalization", Passed: fmt.Sprintf("%.10f", a.HyperchargeNormalizationKY) == "1.6666666667", Detail: fmt.Sprintf("Tr(Y²)_full=%.10f, Tr(T3²)_full=%.10f, k_Y=Tr(Y²)/Tr(T3²)=%.10f", a.FullYTrace2OneGeneration, a.FullT3Trace2OneGeneration, a.HyperchargeNormalizationKY)},
			{Name: "normalized hypercharge factor", Passed: a.NormalizedHyperchargeFactor > 0, Detail: fmt.Sprintf("Y_N=sqrt(1/k_Y)Y gives factor %.10f and Tr(Y_N²)=%.10f", a.NormalizedHyperchargeFactor, a.NormalizedYTrace2)},
			{Name: "equal-normalized-coupling boundary candidate", Passed: a.EqualNormalizedCouplingBoundaryCandidate, Detail: fmt.Sprintf("if g_2=g_1 at a boundary and g_1²=k_Y g_Y², then sin²θ_boundary=1/(1+k_Y)=%.10f", a.EqualNormalizedCouplingBoundarySin2)},
			{Name: "gauge kinetic normalization", Passed: a.GaugeKineticNormalizationDerived, Detail: "not derived; trace metrics do not yet prove the finite kinetic terms for W3 and B"},
			{Name: "RG boundary scale", Passed: a.RGBoundaryScaleDerived, Detail: "not derived; no physical scale has been assigned to the boundary candidate"},
			{Name: "physical weak mixing angle", Passed: a.WeakMixingAngleDerived, Detail: "not derived; θ_W requires g and g' after kinetic normalization and RG running"},
			{Name: "fine-structure constant", Passed: a.FineStructureDerived, Detail: "not derived; α_em requires e=g sinθ plus scale and normalization"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no measured weak angle or electromagnetic coupling was inserted"},
		}, Notes: []string{a.TruthStatement, fmt.Sprintf("remaining unknowns: %v", a.RemainingUnknowns)}}
	}}
}
