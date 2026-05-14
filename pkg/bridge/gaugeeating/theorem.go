package gaugeeating

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteScalarKineticGaugeEatingSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-SCALAR-KINETIC-GAUGE-EATING-SEARCH"
	const name = "finite scalar kinetic normalization and gauge-eating theorem search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build scalar kinetic/gauge-eating audit", Passed: false, Detail: err.Error()}}}
		}
		eps := 1e-9
		condText := "infinite"
		if !math.IsInf(a.BrokenImageCondition, 1) {
			condText = fmt.Sprintf("%.10f", a.BrokenImageCondition)
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 84 mass-matrix input", Passed: a.ScalarCovariant.DimensionlessWZPhotonSignature, Detail: fmt.Sprintf("rank=%d, mW_hat²=%.10f, mZ_hat²=%.10f, mγ_hat²=%.10f", a.ScalarCovariant.MassMatrixRank, a.ScalarCovariant.ChargedMassSquaredHat, a.ScalarCovariant.NeutralMassiveSquaredHat, a.ScalarCovariant.PhotonMassSquaredHat)},
			{Name: "canonical Euclidean scalar metric diagnostic", Passed: a.KineticMetricPositive, Detail: fmt.Sprintf("K_Φ=I₄ diagnostic, Tr(K_Φ)=%.10f, det(K_Φ)=%.10f", a.KineticTrace, a.KineticDeterminant)},
			{Name: "scalar kinetic normalization selected", Passed: a.KineticNormalizationSelected, Detail: "not derived; K_Φ=I₄ is the active-frame metric used diagnostically, not an action-selected kinetic normalization"},
			{Name: "radial direction", Passed: a.RadialNormResidual < eps, Detail: fmt.Sprintf("r̂=φ0/||φ0||=%s, norm residual=%.3e", FormatVector(a.RadialDirection), a.RadialNormResidual)},
			{Name: "broken-generator image rank", Passed: a.BrokenImagesIndependent, Detail: fmt.Sprintf("images of {T1,T2,Z=T3−YΦ}φ0 have Gram=%s, rank=%d, eig range=[%.10f, %.10f], condition=%s", FormatGram(a.BrokenImageGram), a.BrokenImageRank, a.BrokenImageMinEigen, a.BrokenImageMaxEigen, condText)},
			{Name: "electromagnetic generator remains unbroken", Passed: a.EMNullNorm < eps, Detail: fmt.Sprintf("||(T3+YΦ)φ0||=%.3e", a.EMNullNorm)},
			{Name: "Goldstone/gauge-eating image diagnostic", Passed: a.GoldstoneImageTheoremDiagnostic, Detail: fmt.Sprintf("%d broken image directions + %d electromagnetic null direction", a.BrokenGeneratorCount, a.UnbrokenGeneratorCount)},
			{Name: "finite gauge-eating theorem", Passed: a.FiniteGaugeEatingTheoremDerived, Detail: "open; requires scalar kinetic action, gauge Hessian, vacuum orientation, and protected-contact/broken-generator intertwiner"},
			{Name: "physical gauge-boson masses", Passed: a.PhysicalMassesDerived, Detail: "not derived; no physical W/Z/photon mass comparison is allowed"},
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
