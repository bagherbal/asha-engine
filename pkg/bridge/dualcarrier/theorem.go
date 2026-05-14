package dualcarrier

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func DualCarrierGaugeArchitectureSplitTheorem() theorem.Theorem {
	const id = "BRIDGE-DUAL-CARRIER-GAUGE-ARCHITECTURE"
	const name = "dual-carrier gauge architecture split"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build dual-carrier architecture", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 71 obstruction inherited", Passed: a.ForcedEmbeddingRejected, Detail: "the direct u(4)→contact-block embedding is rejected rather than patched with arbitrary coefficients"},
			{Name: "Pati-Salam/Fock current carrier", Passed: a.PatiSalamCarrierPreservesU4, Detail: FormatCarrier(a.PatiSalamCarrier)},
			{Name: "Boolean/contact electroweak block carrier", Passed: a.ContactCarrierPreservesEWSeed, Detail: FormatCarrier(a.ContactCarrier)},
			{Name: "color and leptoquark sectors preserved", Passed: a.ColorCarrierPreserved && a.LeptoquarkCarrierPreserved, Detail: "SU(3)c and leptoquark currents remain on the u(4) carrier instead of being compressed into a 4D contact target"},
			{Name: "direct embedding replaced by coupling problem", Passed: a.DualCarrierSplitDefined, Detail: fmt.Sprintf("former Hom(R^16,R^4) embedding space dim=%d with kernel≥%d; new bridge is a coupling/action between carriers", a.DirectEmbeddingDimension, a.DirectEmbeddingKernelMin)},
			{Name: "coupling tensor domain exposed", Passed: a.CouplingTensorDimension == 64, Detail: fmt.Sprintf("formal current-contact coupling tensor dimension=%d; this is a bridge domain, not a selected tensor", a.CouplingTensorDimension)},
			{Name: "coupling problem audit", Passed: len(a.CouplingProblems) == 4, Detail: FormatCouplingProblems(a.CouplingProblems)},
			{Name: "abelian separation", Passed: !a.AbelianSeparationStillOpen, Detail: "open; central u(1), B-L, and contact-u1/hypercharge still need a kinetic/mixing separation theorem"},
			{Name: "coupling tensor selected", Passed: a.CouplingTensorSelected, Detail: "open; arbitrary 64 coefficients would be fitting"},
			{Name: "dual-carrier action derived", Passed: a.CouplingActionDerived, Detail: "open; S_coupling[j,A,Φ] is typed but not constructed"},
			{Name: "current Hessian computable", Passed: a.CurrentHessianComputable, Detail: "open; K_current cannot be computed until the dual-carrier coupling action is derived"},
			{Name: "exchange kernel update", Passed: a.ExchangeKernelUpdated, Detail: "open; G_hat remains unchanged"},
			{Name: "attractive scalar-channel theorem", Passed: a.AttractiveScalarDerived, Detail: "open; no NJL attraction or condensation follows from architecture split alone"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; no top-like versus bottom-like distinction is selected"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed coupling, mass, scale, Higgs parameter, or physical threshold was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
