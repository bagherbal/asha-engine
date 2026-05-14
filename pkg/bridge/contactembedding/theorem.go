package contactembedding

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactMatterHyperchargeEmbeddingTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-MATTER-HYPERCHARGE-EMBEDDING"
	const name = "contact-to-matter hypercharge embedding and finite normalization threshold"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact-to-matter embedding", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 101 boundary seed inherited", Passed: a.Boundary.ActionSelectedDimensionlessBoundarySeed, Detail: fmt.Sprintf("K_gen=%s", FormatMatrix(a.Boundary.GeneratorHessian))},
			{Name: "abelian coefficients identified", Passed: a.ContactU1KineticCoefficient > 0 && a.MatterHyperchargeKY > 0, Detail: fmt.Sprintf("K(%s)=%.10f, k_Y(%s)=%.10f", a.ContactGeneratorName, a.ContactU1KineticCoefficient, a.MatterGeneratorName, a.MatterHyperchargeKY)},
			{Name: "embedding equation solved", Passed: a.ContactToMatterEmbeddingSelected, Detail: fmt.Sprintf("%s; lambda²=%.10f, lambda=%.10f", a.EmbeddingEquation, a.EmbeddingScaleSq, a.EmbeddingScale)},
			{Name: "orientation branch selected", Passed: a.OrientationSelected && a.OneDimensionalMapUnique, Detail: FormatSigns(a.SignCandidates)},
			{Name: "embedded matter Hessian selected", Passed: a.EmbeddedMatterHessianSelected, Detail: fmt.Sprintf("K_embedded=%s", FormatMatrix(a.EmbeddedMatterHessian))},
			{Name: "finite boundary diagnostic lifted", Passed: a.EmbeddedMatterBoundaryDiagnostic && !a.PhysicalWeakAngleDerived, Detail: fmt.Sprintf("sin²_embedded=1/(1+k_Y)=%.10f; diagnostic only", a.EmbeddedMatterBoundarySin2)},
			{Name: "contact/matter mismatch removed only by embedding", Passed: a.ContactMatterMismatchBeforeMap > 0 && a.ContactMatterMismatchAfterMap <= 1e-8, Detail: fmt.Sprintf("before=%.10f, after=%.10f", a.ContactMatterMismatchBeforeMap, a.ContactMatterMismatchAfterMap)},
			{Name: "normalization threshold is not a mass threshold", Passed: a.ThresholdNormalizationMapSelected && !a.FiniteNormalizationThreshold.PhysicalMassThreshold && !a.FiniteNormalizationThreshold.PhysicalScaleInserted, Detail: fmt.Sprintf("%s; deficit=%.10f", a.FiniteNormalizationThreshold.Formula, a.FiniteNormalizationThreshold.NormalizationDeficit)},
			{Name: "RG/scale/physical firewall", Passed: !a.RGBoundaryScaleDerived && !a.BetaThresholdFlowDerived && !a.FineStructureDerived && !a.PhysicalCouplingsDerived && !a.PhysicalMassesDerived, Detail: "no M*, beta-flow, thresholds, alpha, thetaW, g2, gY, W/Z masses, Higgs vev, or fermion masses are derived"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed electroweak coupling, weak angle, mass, or scale was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			"threshold map: " + a.FiniteNormalizationThreshold.Detail,
			"rejected claims: " + Join(a.RejectedClaims),
			"remaining unknowns: " + Join(a.RemainingUnknowns),
			"Next: " + a.RecommendedNextGate,
		}}
	}}
}
