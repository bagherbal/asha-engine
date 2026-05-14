package canonicalboundary

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CanonicalFiniteRGBoundarySeedTheorem() theorem.Theorem {
	const id = "BRIDGE-CANONICAL-FINITE-RG-BOUNDARY-SEED"
	const name = "canonical finite RG boundary seed and scale firewall"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build canonical RG boundary seed", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 100 action-selected Hessian inherited", Passed: a.CanonicalAction.FullGaugeHessianSelected, Detail: fmt.Sprintf("K_EW in closed basis=%s", FormatMatrix(a.CanonicalAction.FullGaugeHessian))},
			{Name: "basis transform typed", Passed: a.BasisTransform.Matrix.Rows() == 4 && a.BasisTransform.Matrix.Cols() == 4, Detail: fmt.Sprintf("from=%v to=%v", a.ClosedBasis, a.GeneratorBasis)},
			{Name: "generator-basis Hessian computed", Passed: a.GeneratorHessian.IsSymmetric(1e-8), Detail: fmt.Sprintf("K_gen=%s", FormatMatrix(a.GeneratorHessian))},
			{Name: "SU(2)L kinetic isotropy selected", Passed: a.SU2KineticIsotropic, Detail: fmt.Sprintf("entries=%s, coefficient=%.10f", FormatFloatSlice(a.SU2KineticEntries), a.SU2KineticCoefficient)},
			{Name: "scalar/contact U(1) kinetic seed selected", Passed: a.ActionSelectedDimensionlessBoundarySeed, Detail: fmt.Sprintf("K(Y_phi)=%.10f, K(Y_phi)/K_SU2=%.10f", a.ScalarContactU1KineticCoefficient, a.ScalarContactU1ToSU2Ratio)},
			{Name: "contact-sector boundary diagnostic exposed", Passed: a.ScalarContactBoundarySin2 > 0 && !a.PhysicalWeakAngleDerived, Detail: fmt.Sprintf("inverse-kinetic ratio=%.10f gives sin²_contact=%.10f; diagnostic only", a.InverseKineticCouplingRatio, a.ScalarContactBoundarySin2)},
			{Name: "matter hypercharge normalization kept separate", Passed: a.MatterHyperchargeKY > 0 && a.MatterBoundarySin2 > 0, Detail: fmt.Sprintf("k_Y=%.10f gives sin²_matter=%.10f under equal normalized matter-table coupling", a.MatterHyperchargeKY, a.MatterBoundarySin2)},
			{Name: "contact-to-matter embedding remains missing", Passed: !a.EmbeddingMapSelected, Detail: fmt.Sprintf("K(Y_phi)=%.10f versus k_Y=%.10f; required lambda²=%.10f, lambda=%.10f is exposed but not selected", a.ScalarContactU1KineticCoefficient, a.MatterHyperchargeKY, a.RequiredEmbeddingScaleSq, a.RequiredEmbeddingScale)},
			{Name: "RG and scale firewall", Passed: !a.RGFlowDetermined && !a.BoundaryScaleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived, Detail: "no M*, beta coefficients, thresholds, alpha, thetaW, W/Z masses, or Higgs scale are derived"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed alpha, thetaW, electroweak vev, Higgs mass, or measured coupling was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			"boundary candidates: " + FormatCandidates(a.BoundaryCandidates),
			"rejected claims: " + Join(a.RejectedClaims),
			"remaining unknowns: " + Join(a.RemainingUnknowns),
			"Next: " + a.RecommendedNextGate,
		}}
	}}
}
