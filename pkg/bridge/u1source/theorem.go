package u1source

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactU1BLKineticSourceSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-U1-BL-KINETIC-SOURCE-SEARCH"
	const name = "contact-u1 / B-L kinetic Hessian source search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact-u1/B-L source search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 75 U(1) kinetic input", Passed: a.U1Kinetic.MatterGramDerived && a.U1Kinetic.ContactU1NormDerived, Detail: "central/B-L matter Gram and contact-u1 scalar norm are available"},
			{Name: "abelian fields under audit", Passed: len(a.Fields) == 3, Detail: fmt.Sprintf("fields=%v", a.Fields)},
			{Name: "trace-zero ingredients", Passed: a.ContactU1TraceZero && a.BMinusLTraceZero, Detail: fmt.Sprintf("Tr(T_phi)=%.1e, Tr(B-L)=%.1e", a.U1Kinetic.ContactU1.Trace, a.U1Kinetic.BMinusL.Trace)},
			{Name: "central/contact factorized source", Passed: a.FactorizedTraceSourceDerived && math.Abs(a.FactorizedCentralContact) < 1e-10, Detail: fmt.Sprintf("Tr(I)Tr(T_phi)=%.3e; central u(1) is not part of hypercharge", a.FactorizedCentralContact)},
			{Name: "B-L/contact factorized source", Passed: a.FactorizedTraceSourceDerived && math.Abs(a.FactorizedBLContact) < 1e-10, Detail: fmt.Sprintf("Tr(B-L)Tr(T_phi)=%.3e; factorized trace gives no kinetic mixing source", a.FactorizedBLContact)},
			{Name: "tensor-product trace source", Passed: math.Abs(a.TensorTraceBLContact) < 1e-10 && math.Abs(a.TensorTraceCentralContact) < 1e-10, Detail: fmt.Sprintf("tensor traces central/contact=%.3e, B-L/contact=%.3e", a.TensorTraceCentralContact, a.TensorTraceBLContact)},
			{Name: "non-factorized abelian action", Passed: a.NonFactorizedActionDerived, Detail: "not derived; a finite dual-carrier action must produce this if kinetic mixing is physical"},
			{Name: "cross-carrier kinetic source", Passed: a.CrossCarrierSourceDerived, Detail: "not derived; no finite source for K(B-L, contact-u1) is selected"},
			{Name: "full U(1) Hessian", Passed: a.FullHessianDerived, Detail: "not derived; physical U(1)_Y coupling remains open"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed alpha, theta_W, v, or measured coupling was inserted"},
		}, Notes: []string{a.TruthStatement, "Next: " + a.RecommendedNextGate}}
	}}
}
