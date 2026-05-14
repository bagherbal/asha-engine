package currentembedding

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CurrentFieldEmbeddingTheorem() theorem.Theorem {
	const id = "BRIDGE-CURRENT-FIELD-EMBEDDING-BF-CONTACT"
	const name = "current-field embedding into finite BF/contact action"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build current-field embedding audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 69 current-Hessian input", Passed: !a.Previous.CurrentHessianDerived && a.Previous.SectorCount == 4, Detail: "previous gate exposed K_current as missing and left propagator selection open"},
			{Name: "typed current-sector fields", Passed: a.FieldSlotsDefined && a.SectorFieldCount == 4 && a.GeneratorFieldCount == 16, Detail: fmt.Sprintf("sector fields=%d, generator fields=%d; %s", a.SectorFieldCount, a.GeneratorFieldCount, FormatFields(a.Fields))},
			{Name: "finite action slots exposed", Passed: a.ActionSlotCount == 5 && a.BooleanContactActionAvailable, Detail: FormatSlots(a.ActionSlots)},
			{Name: "minimal current-field action template", Passed: a.MinimalActionTemplate != "", Detail: a.MinimalActionTemplate},
			{Name: "source coupling template", Passed: a.SourceCouplingTemplate != "", Detail: a.SourceCouplingTemplate},
			{Name: "Fock/current inventory available", Passed: a.FockCurrentInventoryAvailable, Detail: "u(4)-shaped current inventory central/color/B-L/leptoquark is available from previous gates"},
			{Name: "Boolean/contact action carrier available", Passed: a.ContactBlockActionAvailable, Detail: "K ⊕ K⊥ block-connection and B-sector action carriers are available as finite action substrates"},
			{Name: "current-to-contact embedding map", Passed: a.CurrentToContactEmbeddingDerived, Detail: "open; E_current_to_block is typed as an action slot but not derived"},
			{Name: "source functional", Passed: a.SourceFunctionalDerived, Detail: "open; J_source[B,A] is not constructed"},
			{Name: "Hessian computable", Passed: a.HessianComputable, Detail: "open; K_current cannot be computed until E_current_to_block and J_source exist"},
			{Name: "current Hessian derived", Passed: a.CurrentHessianDerived, Detail: "open; no δ²S/δjδj is computed"},
			{Name: "propagator rule", Passed: a.PropagatorRuleDerived, Detail: "open; no exchange propagator follows from typed slots alone"},
			{Name: "exchange kernel update", Passed: a.ExchangeKernelUpdated, Detail: "open; G_hat cannot be updated before K_current is derived"},
			{Name: "attractive scalar-channel theorem", Passed: a.AttractiveScalarDerived, Detail: "open; no NJL attraction or condensation follows from this gate"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; current-field slots still do not distinguish top-like up from bottom-like down"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed coupling, mass, scale, or Higgs parameter was inserted"},
		}, Notes: []string{
			a.EmbeddingObstruction,
			a.TruthStatement,
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
