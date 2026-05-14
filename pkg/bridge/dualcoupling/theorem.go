package dualcoupling

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func DualCarrierCouplingTensorActionSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-DUAL-CARRIER-COUPLING-ACTION-SEARCH"
	const name = "dual-carrier coupling tensor / action search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build dual-carrier coupling audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 72 dual-carrier input", Passed: a.Previous.DualCarrierSplitDefined && a.Previous.CouplingTensorDimension == 64, Detail: "previous gate split the Pati-Salam/u(4) and contact/electroweak carriers and exposed a formal 64D coupling tensor"},
			{Name: "raw 64-parameter tensor rejected", Passed: a.Direct64TensorRejected, Detail: fmt.Sprintf("naive tensor dim=%d; arbitrary coefficients would be fitting, not derivation", a.NaiveTensorDimension)},
			{Name: "sector-compatible coupling audit", Passed: len(a.SectorCouplings) == 4, Detail: FormatSectorCouplings(a.SectorCouplings)},
			{Name: "color-contact direct coupling rejected", Passed: a.ColorContactCouplingRejected, Detail: "SU(3)c remains on the Pati-Salam carrier; no contact color target exists"},
			{Name: "leptoquark-contact direct coupling rejected", Passed: a.LeptoquarkContactCouplingRejected, Detail: "leptoquark exchange requires a separate Pati-Salam-sector action, not contact compression"},
			{Name: "abelian bridge domain exposed", Passed: a.AbelianBridgeDomainExposed && a.AbelianBridgeDimension == 2, Detail: fmt.Sprintf("central/contact-u1 and B-L/contact-u1 slots give a %dD abelian bridge domain", a.AbelianBridgeDimension)},
			{Name: "scalar-current bridge domain exposed", Passed: a.ScalarCurrentCouplingDomainExposed, Detail: "possible scalar/contact current terms are typed, but their current and normalization are not derived"},
			{Name: "action-term audit", Passed: len(a.ActionTerms) == 4, Detail: FormatActionTerms(a.ActionTerms)},
			{Name: "coupling coefficients selected", Passed: a.CouplingTensorSelected || a.AbelianCoefficientsSelected || a.ScalarCurrentCoefficientsSelected, Detail: "open; κ coefficients and scalar-current normalizations are not selected by finite action data"},
			{Name: "coupling action derived", Passed: a.CouplingActionDerived, Detail: "open; S_coupling[j,A,Φ] is reduced to typed domains but not constructed"},
			{Name: "dual-carrier Hessian computable", Passed: a.DualCarrierHessianComputable, Detail: "open; coupled Hessian requires selected abelian/scalar-current action"},
			{Name: "exchange kernel update", Passed: a.ExchangeKernelUpdated, Detail: "open; G_hat remains unchanged"},
			{Name: "attractive scalar-channel theorem", Passed: a.AttractiveScalarDerived, Detail: "open; no NJL attraction or condensation follows from coupling-domain exposure"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; no top-like versus bottom-like distinction is selected"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed coupling, mass, scale, Higgs parameter, or threshold was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
