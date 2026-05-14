package u1nonfactor

import (
	"fmt"
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func NonFactorizedAbelianActionSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-NONFACTORIZED-ABELIAN-ACTION-SEARCH"
	const name = "non-factorized abelian action / kinetic-mixing search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build non-factorized abelian source search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 76 factorized no-source inherited", Passed: !a.Previous.CrossCarrierSourceDerived && math.Abs(a.Previous.FactorizedBLContact) < 1e-10, Detail: "factorized B-L/contact trace source vanishes; testing non-factorized incidence source"},
			{Name: "Yukawa-incidence support", Passed: a.NonFactorizedSupportDerived, Detail: fmt.Sprintf("support entries=%d inside full left⊗scalar tensor dim=%d, fraction=%.6f", a.YukawaSupportEntries, a.FullTensorDimension, a.SupportFraction)},
			{Name: "local B-L/contact correlation", Passed: a.LocalNonzeroCorrelation, Detail: fmt.Sprintf("Σ|B-L·T_phi|=%.10f, Σ(B-L·T_phi)^2=%.10f, RMS=%.10f", a.AbsoluteBLContactMoment, a.QuadraticBLContactMoment, a.BLContactRMS)},
			{Name: "kind-level signed contributions", Passed: len(a.KindContributions) == 4, Detail: formatKindContributions(a.KindContributions)},
			{Name: "up/down cancellation", Passed: a.UpDownCancellation, Detail: "up and down quark branches carry opposite scalar T_phi and cancel in signed B-L/contact moment"},
			{Name: "lepton-pair cancellation", Passed: a.LeptonPairCancellation, Detail: "neutrino and electron branches carry opposite scalar T_phi and cancel in signed B-L/contact moment"},
			{Name: "total signed non-factorized source", Passed: !a.TotalSignedCancellation, Detail: fmt.Sprintf("signed B-L/contact=%.3e, signed central/contact=%.3e; total cancels, so no net kinetic source is derived", a.SignedBLContactMoment, a.SignedCentralContactMoment)},
			{Name: "non-factorized action object", Passed: a.NonFactorizedActionDerived, Detail: "Yukawa-incidence correlation is a genuine non-factorized finite object, but not a nonzero Hessian source"},
			{Name: "cross-carrier kinetic source", Passed: a.CrossCarrierSourceDerived, Detail: "not derived; nonzero local correlation cancels in the signed total"},
			{Name: "full U(1) Hessian", Passed: a.FullU1HessianDerived, Detail: "not derived; physical U(1)_Y coupling remains open"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed alpha, theta_W, v, or measured coupling was inserted"},
		}, Notes: []string{a.TruthStatement, "Next: " + a.RecommendedNextGate}}
	}}
}

func formatKindContributions(xs []KindContribution) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(ch=%d,fiber=%d,signed=%.10f,abs=%.10f,quad=%.10f)", x.Kind, x.Channels, x.FiberEntries, x.SignedBLContact, x.AbsoluteBLContact, x.QuadraticBLContact))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}
