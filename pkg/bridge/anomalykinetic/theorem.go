package anomalykinetic

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func AnomalyConstrainedU1KineticHessianSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-ANOMALY-CONSTRAINED-U1-KINETIC-HESSIAN"
	const name = "anomaly-constrained U(1) kinetic Hessian search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build anomaly-constrained U(1) Hessian", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 79 anomaly input", Passed: a.Anomaly.AnomalyShadowSupported, Detail: "Y, B-L, and mixed abelian anomaly ledgers cancel; Yukawa-incidence cancellation is anomaly-shadow supported"},
			{Name: "Gate 75 kinetic input", Passed: a.Kinetic.MatterGramDerived && a.Kinetic.ContactU1NormDerived, Detail: "central/B-L matter trace Gram and contact-u1 scalar norm are available"},
			{Name: "general symmetric Hessian domain", Passed: a.SymmetricHessianDimension == 6, Detail: fmt.Sprintf("three abelian fields give symmetric Hessian dim=%d = 3 diagonal + 3 off-diagonal", a.SymmetricHessianDimension)},
			{Name: "diagonal trace-Gram sector survives", Passed: a.DiagonalPositive, Detail: fmt.Sprintf("diag=[%.10f, %.10f, %.10f], determinant=%.10f", a.Kinetic.Central.Trace2, a.Kinetic.BMinusL.Trace2, a.Kinetic.ContactU1.Trace2, a.TraceGramDeterminant)},
			{Name: "known off-diagonal sources cancel", Passed: a.AllKnownOffDiagonalSourcesCancel, Detail: formatConstraints(a.OffDiagonalConstraints)},
			{Name: "nonzero off-diagonal survives", Passed: a.NonzeroOffDiagonalSurvives, Detail: "no; all known central/B-L/contact cross sources are zero in the current finite data"},
			{Name: "stricter no-mixing theorem", Passed: a.StricterNoMixingTheoremDerived, Detail: "not derived; the result is a no-source diagnostic, not a proof that every possible finite action forbids kinetic mixing"},
			{Name: "full U(1) kinetic Hessian", Passed: a.FullU1KineticHessianDerived, Detail: "not derived; diagonal trace-Gram entries are diagnostics until an action selects them as kinetic coefficients"},
			{Name: "physical U(1) coupling and alpha", Passed: a.PhysicalU1CouplingDerived || a.FineStructureDerived, Detail: "not derived; no g_Y, alpha_em, or low-energy alpha follows from this gate"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed coupling, alpha, theta_W, v, or mass scale was inserted"},
		}, Notes: []string{a.TruthStatement, "Next: " + a.RecommendedNextGate}}
	}}
}

func formatConstraints(xs []OffDiagonalConstraint) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s via %s = %.3e (zero=%v)", x.Pair, x.Source, x.Value, x.ForcedZero))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}
