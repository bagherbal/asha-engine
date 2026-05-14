package scalarcovariant

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteScalarCovariantDerivativeMassMatrixTheorem() theorem.Theorem {
	const id = "BRIDGE-SCALAR-COVARIANT-DERIVATIVE-MASS-MATRIX"
	const name = "finite scalar covariant derivative and gauge-boson mass matrix search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build scalar covariant derivative audit", Passed: false, Detail: err.Error()}}}
		}
		eps := 1e-9
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "scalar/contact active frame", Passed: a.ActiveRealDimension == 4, Detail: fmt.Sprintf("active real directions=%d, r0²=%.10f, r0=%.10f", a.ActiveRealDimension, a.VacuumRadiusSquared, a.VacuumRadius)},
			{Name: "abstract covariant-derivative template", Passed: a.AbstractCovariantDerivativeTemplate && a.GeneratorSkewResidual < eps, Detail: fmt.Sprintf("DΦ=(d+g2 T_a W_a+gY YΦ B)Φ; generator skew residual=%.3e", a.GeneratorSkewResidual)},
			{Name: "vacuum orientation discipline", Passed: !a.VacuumOrientationChosen, Detail: fmt.Sprintf("using %s with φ0=%s; orientation is diagnostic, not derived", a.VacuumOrientationName, FormatVector(a.VacuumVector))},
			{Name: "electromagnetic null direction", Passed: a.EMAnnihilatesVacuumNorm < eps, Detail: fmt.Sprintf("||(T3+YΦ)φ0||=%.3e", a.EMAnnihilatesVacuumNorm)},
			{Name: "dimensionless mass matrix diagnostic", Passed: a.DimensionlessWZPhotonSignature, Detail: fmt.Sprintf("M_hat=%s; rank=%d", FormatMassMatrix(a.MassMatrix), a.MassMatrixRank)},
			{Name: "charged W degeneracy", Passed: a.ChargedDegeneracyResidual < eps, Detail: fmt.Sprintf("mW_hat²=%.10f, degeneracy residual=%.3e", a.ChargedMassSquaredHat, a.ChargedDegeneracyResidual)},
			{Name: "neutral Z/photon split", Passed: a.PhotonNullResidual < eps && a.NeutralMassiveSquaredHat > eps, Detail: fmt.Sprintf("mZ_hat²=%.10f, mγ_hat²=%.10f, photon residual=%.3e", a.NeutralMassiveSquaredHat, a.PhotonMassSquaredHat, a.PhotonNullResidual)},
			{Name: "scalar kinetic normalization", Passed: a.FiniteScalarKineticNormalizationDerived, Detail: "not derived; mass matrix is dimensionless and unnormalized"},
			{Name: "gauge couplings", Passed: a.GaugeCouplingsDerived, Detail: "not derived; g2 and gY remain boundary-family data"},
			{Name: "gauge action Hessian", Passed: a.GaugeActionHessianDerived, Detail: "not derived by Gate 83; representation generators are not yet physical gauge fields"},
			{Name: "physical W/Z/Higgs masses", Passed: a.PhysicalMassesDerived, Detail: "not derived; no comparison to observed masses is allowed"},
		}, Notes: []string{a.TruthStatement, fmt.Sprintf("remaining unknowns: %s", formatUnknowns(a.RemainingUnknowns)), "Next: " + a.RecommendedNextGate}}
	}}
}

func formatUnknowns(values []string) string {
	out := ""
	for i, v := range values {
		if i > 0 {
			out += "; "
		}
		out += v
	}
	if out == "" {
		return "none"
	}
	return out
}
