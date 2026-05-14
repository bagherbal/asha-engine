package u1completion

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func AbelianCoefficientU1CompletionSelectionSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-ABELIAN-COEFFICIENT-U1-COMPLETION-SELECTION"
	const name = "abelian coefficient / U(1) completion selection search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 99 U(1) completion selection audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 98 abelian-completion input", Passed: a.CompletionFamilyTyped, Detail: fmt.Sprintf("K(kappa)=K_SU2+kappa(Q-Z)(Q-Z)^T; target kappa=%.10f", a.TargetKappa)},
			{Name: "abelian null/completion direction", Passed: len(a.AbelianDirection) == 4 && a.AbelianDirectionNormSq > 0, Detail: fmt.Sprintf("%s; ||Q-Z||^2=%.10f", a.NullDirectionBasis, a.AbelianDirectionNormSq)},
			{Name: "whitening condition selects numerical value", Passed: a.WhiteningSelectsKappa && a.TargetKappa == 6, Detail: fmt.Sprintf("diag(1,1,4) requires kappa_U1=%.10f in the Gate 98 convention", a.TargetKappa)},
			{Name: "finite count resonances audited", Passed: a.CandidateHitCount >= 2, Detail: fmt.Sprintf("%d independent finite count resonances hit kappa=6: %s", a.CandidateHitCount, CandidateSummary(a.CandidateResonances))},
			{Name: "unique derivation of kappa", Passed: a.UniqueDerivation, Detail: "not derived; multiple count resonances hit 6, so count-matching is not a theorem"},
			{Name: "count resonance selected as action coefficient", Passed: a.CountResonanceSelected, Detail: "rejected; no finite variation attaches any count product to the abelian quadratic term"},
			{Name: "finite second variation", Passed: a.FiniteSecondVariation, Detail: "not computed; the U(1) completion coefficient is still not action-selected"},
			{Name: "kappa_U1 physical", Passed: a.KappaPhysical, Detail: "not physical yet; kappa=6 is only the metric-whitening value until the finite action selects it"},
			{Name: "gauge kinetic Hessian fixed", Passed: a.GaugeKineticHessianFixed, Detail: "not fixed; abelian coefficient and scalar/gauge kinetic normalizations remain open"},
			{Name: "physical couplings or masses", Passed: a.PhysicalCouplingsOrMasses, Detail: "not derived; g2/gY/thetaW/alpha and physical W/Z masses remain bridge-gated"},
		}, Notes: []string{a.TruthStatement, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
