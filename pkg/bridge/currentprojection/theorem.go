package currentprojection

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CurrentActionScalarProjectionTheorem() theorem.Theorem {
	const id = "BRIDGE-CURRENT-ACTION-SCALAR-LR-PROJECTION"
	const name = "current action on scalar LR projector / coefficient audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct current action scalar projection", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "scalar LR projector input", Passed: a.Chiral.ScalarLRProjectorConstructed, Detail: fmt.Sprintf("rank(P_LR)=%d on domain dim=%d; right dim=%d", a.ScalarLRRank, a.DomainDimension, a.RightDimension)},
			{Name: "u(4)-shaped current generator inventory", Passed: a.CurrentActionConstructed, Detail: fmt.Sprintf("flavor dim=%d; generators=%d = central 1 + color 8 + B-L 1 + leptoquark 6", a.FlavorDimension, len(a.Generators))},
			{Name: "induced current action on scalar LR image", Passed: a.InducedCurrentActionConstructed, Detail: fmt.Sprintf("max ||UᵀT_DU−T_R||_F across generators/sectors = %.3e", a.MaxIntertwinerResidual)},
			{Name: "finite scalar-projection overlap coefficients", Passed: a.UnsignedScalarProjectionCoefficientsKnown, Detail: FormatSectorCoefficients(a.SectorCoefficients)},
			{Name: "lepton/color split visible", Passed: a.LeptonQuarkSplitVisible, Detail: fmt.Sprintf("central=%.10f, color=%.10f, B-L=%.10f, leptoquark=%.10f", a.CentralOverlap, a.ColorOverlap, a.BLOverlap, a.LeptoquarkOverlap)},
			{Name: "signed Fierz coefficients", Passed: a.SignedScalarProjectionCoefficientsKnown, Detail: "open; overlaps are finite representation diagnostics, not Lorentz/Fierz-signed coefficients"},
			{Name: "generator kinetic normalization", Passed: a.GeneratorKineticNormalizationDerived, Detail: "open; relative current kinetic weights are still not derived"},
			{Name: "attractive scalar-channel sign", Passed: a.AttractiveSignDerived, Detail: "open; finite action/propagator sign is still missing"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; u(4) lepton/color currents do not distinguish top-like up from bottom-like down"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed Yukawa, v, Higgs mass, top mass, or fitted coupling was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
