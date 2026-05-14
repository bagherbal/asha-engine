package currentcontact

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CurrentToContactEmbeddingMapSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-CURRENT-TO-CONTACT-EMBEDDING-MAP-SEARCH"
	const name = "current-to-contact embedding map search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build current-to-contact embedding search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 70 typed-current input", Passed: a.Previous.FieldSlotsDefined && !a.Previous.CurrentToContactEmbeddingDerived, Detail: "current fields and action slots are typed, but E_current_to_block was left open"},
			{Name: "source current inventory", Passed: a.SourceSectorCount == 4 && a.SourceGeneratorCount == 16, Detail: fmt.Sprintf("u(4)-shaped source has %d sectors and %d generator fields", a.SourceSectorCount, a.SourceGeneratorCount)},
			{Name: "Boolean/contact block target", Passed: a.TargetHasSU2U1Shape, Detail: fmt.Sprintf("target block seed has %d generators with span rank %d; Boolean dim=%d, K=%d, K⊥=%d", a.TargetBlockSeedCount, a.TargetBlockSpanRank, a.BooleanDimension, a.ContactDimension, a.ComplementDimension)},
			{Name: "abstract linear map space", Passed: a.AbstractMapSpaceExists, Detail: fmt.Sprintf("Hom(R^%d,R^%d) has dimension %d; any such map would have kernel dimension at least %d if restricted to the current target rank", a.SourceGeneratorCount, a.TargetBlockSeedCount, a.AbstractMapDimension, a.MinimalKernelDimension)},
			{Name: "sector embedding audit", Passed: len(a.SectorAudits) == 4, Detail: FormatSectorAudits(a.SectorAudits)},
			{Name: "abelian ambiguity", Passed: a.AbelianAmbiguity, Detail: "central and B-L each have one source generator, but the contact target supplies only one abelian slot without a derived separation rule"},
			{Name: "color-sector carrier", Passed: a.ColorSectorCarrierDerived, Detail: "open; no SU(3)c adjoint 8D carrier is present in the contact block target"},
			{Name: "leptoquark-sector carrier", Passed: a.LeptoquarkCarrierDerived, Detail: "open; no 6D off-diagonal lepton-color carrier is derived in the contact block target"},
			{Name: "full u(4) target capacity", Passed: a.TargetCanHostFullU4, Detail: "open; current Boolean/contact target is su(2)+u(1)-shaped, not u(4)-shaped"},
			{Name: "current-to-contact embedding map", Passed: a.CurrentToContactMapDerived, Detail: "open; E_current_to_block is not derived and arbitrary 16→4 coefficients would be fitting"},
			{Name: "source functional", Passed: a.SourceFunctionalDerived, Detail: "open; J_source[B,A] remains unconstructed"},
			{Name: "Hessian computable", Passed: a.HessianComputable, Detail: "open; K_current cannot be computed without E_current_to_block or a dual-carrier action"},
			{Name: "exchange kernel update", Passed: a.ExchangeKernelUpdated, Detail: "open; G_hat cannot be updated"},
			{Name: "attractive scalar-channel theorem", Passed: a.AttractiveScalarDerived, Detail: "open; no NJL attraction or condensation follows from this gate"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; no top-like versus bottom-like distinction is selected"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed coupling, mass, scale, Higgs parameter, or physical threshold was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
