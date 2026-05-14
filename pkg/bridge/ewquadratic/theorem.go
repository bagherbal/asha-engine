package ewquadratic

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FullElectroweakQuadraticActionAbelianCompletionTheorem() theorem.Theorem {
	const id = "BRIDGE-FULL-ELECTROWEAK-QUADRATIC-ABELIAN-COMPLETION"
	const name = "full electroweak quadratic action / abelian completion search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 98 full EW quadratic audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 97 full-curvature input", Passed: a.Gate97.Closed && a.Gate97.FullFieldStrengthTyped, Detail: "closed carrier {T1,T2,Z,Q} inherited from Gate 97"},
			{Name: "semisimple curvature metric", Passed: a.SemisimplePositiveSemidefinite && a.SemisimpleRank == 3, Detail: fmt.Sprintf("rank=%d, K_SU2=%s", a.SemisimpleRank, MatrixString(a.SemisimpleMetric))},
			{Name: "abelian null direction identified", Passed: len(a.SemisimpleNull) == 4, Detail: "null direction [0,0,-1,1] = Q-Z = 2Y_phi requires abelian completion"},
			{Name: "abelian completion family typed", Passed: a.AbelianCompletionTyped && a.FullQuadraticActionFamilyTyped, Detail: fmt.Sprintf("%s; positive for %s", a.AbelianCompletion.Convention, a.AbelianCompletion.PositiveFor)},
			{Name: "broken-coordinate diagnostic family", Passed: a.BrokenDiagnostic.Diag114Reachable, Detail: fmt.Sprintf("with Q=0 and charged normalization: %s; diag(1,1,4) at kappa=%.10f", a.BrokenDiagnostic.NormalizedFormula, a.BrokenDiagnostic.KappaForDiag114)},
			{Name: "positive quadratic family exists", Passed: a.PositiveQuadraticFamilyExists, Detail: "a positive family exists for kappa_U1>0, but a family is not a selected physical action"},
			{Name: "diag(1,1,4) selected by action", Passed: a.Diag114SelectedByAction, Detail: "not derived; kappa_U1=6 is the value required by the whitening candidate, not selected by finite second variation"},
			{Name: "abelian coefficient selected", Passed: a.AbelianCoefficientSelected, Detail: "not derived; U(1) completion coefficient remains free"},
			{Name: "gauge kinetic Hessian selected", Passed: a.GaugeKineticHessianSelected, Detail: "not selected; the gate exposes an action family, not a physical Hessian"},
			{Name: "physical couplings or masses", Passed: a.PhysicalCouplingsOrMasses, Detail: "not derived; g2/gY/thetaW/alpha and physical W/Z masses remain bridge-gated"},
		}, Notes: []string{a.TruthStatement, "variables: " + strings.Join(a.Variables, ", "), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
