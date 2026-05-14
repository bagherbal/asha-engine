package bfbridge

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ActiveGenerationProjectionBridgeTheorem() theorem.Theorem {
	const id = "MATTER-ACTIVE-GENERATION-PROJECTION-BRIDGE"
	const name = "active Higgs-curvature to generation-carrier projection bridge"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.OpenTest, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct active-generation bridge", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.OpenTest, Checks: []theorem.Check{
			{Name: "carrier dimensions", Passed: a.ProtectedDimension == 3 && a.ActiveDimension == 4, Detail: fmt.Sprintf("protected generation carrier dim=%d, active Higgs/contact carrier dim=%d", a.ProtectedDimension, a.ActiveDimension)},
			{Name: "active curvature source present", Passed: a.ActiveCurvatureNorm > 1e-8 && a.ActiveCurvatureRank > 0, Detail: fmt.Sprintf("strongest active curvature pair=%s, norm=%.6e, rank=%d", a.ActiveCurvaturePair, a.ActiveCurvatureNorm, a.ActiveCurvatureRank)},
			{Name: "existing active-to-protected connection maps", Passed: a.ExistingConnectionBridgeFound, Detail: fmt.Sprintf("max ||GᵀA_iH||_F=%.6e, span rank=%d across %d generators", a.MaxCrossMapNorm, a.CrossMapSpanRank, a.GeneratorCount)},
			{Name: "induced skew generation operators", Passed: a.InducedSkewSpanRank > 0 && a.MaxInducedSkewNorm > 1e-8, Detail: fmt.Sprintf("max ||B_i F B_jᵀ||_F=%.6e, span rank=%d", a.MaxInducedSkewNorm, a.InducedSkewSpanRank)},
			{Name: "induced symmetric texture operators", Passed: a.InducedSymmetricSpanRank > 0 && a.MaxInducedSymmetricNorm > 1e-8, Detail: fmt.Sprintf("max ||B_i FᵀF B_jᵀ||_F=%.6e, span rank=%d", a.MaxInducedSymmetricNorm, a.InducedSymmetricSpanRank)},
			{Name: "non-diagonal generation mixing selected", Passed: a.NonDiagonalGenerationMixingFound, Detail: fmt.Sprintf("nonDiagonal=%t; current bridge found=%t", a.NonDiagonalGenerationMixingFound, a.ExistingConnectionBridgeFound)},
			{Name: "finite BF curvature still required", Passed: !a.BFCurvatureImplemented && !a.CanonicalGenerationMixingFound, Detail: "existing second-fundamental data do not induce generation mixing; implement genuine BF/Maurer-Cartan curvature before claiming CKM/PMNS"},
		}, Notes: []string{
			a.TruthStatement,
			"Gate 31 is a no-go/bridge gate: it prevents forcing active Higgs curvature into generations without a canonical map.",
			fmt.Sprintf("remaining unknowns: %v", a.RemainingUnknowns),
		}}
	}}
}
