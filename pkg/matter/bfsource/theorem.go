package bfsource

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func BFActionSourceTextureTheorem() theorem.Theorem {
	const id = "MATTER-BF-ACTION-SOURCE-TEXTURE"
	const name = "BF action source contraction search for generation texture"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.OpenTest, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct BF action source audit", Passed: false, Detail: err.Error()}}}
		}
		status := theorem.OpenTest
		if a.CanonicalTextureFound {
			status = theorem.BridgeRequired
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: status, Checks: []theorem.Check{
			{Name: "finite curvature input", Passed: a.FullCurvatureRank > 0 && a.FullCurvatureNorm > 1e-8, Detail: fmt.Sprintf("full BF/Maurer-Cartan residual rank=%d, max norm=%.6e", a.FullCurvatureRank, a.FullCurvatureNorm)},
			{Name: "protected BF source response", Passed: a.ProtectedBFTextureFound, Detail: fmt.Sprintf("rank span{GᵀFG}=%d, max ||GᵀFG||=%.6e, offdiag=%.6e", a.ProtectedBFResponseRank, a.ProtectedBFMaxNorm, a.ProtectedBFMaxOffDiag)},
			{Name: "mixed active-generation BF source response", Passed: a.MixedBFBridgeFound, Detail: fmt.Sprintf("rank span{GᵀFH}=%d, max ||GᵀFH||=%.6e", a.MixedBFResponseRank, a.MixedBFMaxNorm)},
			{Name: "active-only BF action response", Passed: a.ActiveOnlySourceFound, Detail: fmt.Sprintf("rank span{HᵀFH}=%d, max ||HᵀFH||=%.6e", a.ActiveBFResponseRank, a.ActiveBFMaxNorm)},
			{Name: "protected quadratic texture", Passed: a.ProtectedQuadratic.Rank > 0, Detail: FormatTexture(a.ProtectedQuadratic)},
			{Name: "mixed quadratic texture", Passed: a.MixedQuadratic.Rank > 0, Detail: FormatTexture(a.MixedQuadratic)},
			{Name: "active scalar quadratic", Passed: a.ActiveQuadratic.Rank > 0, Detail: FormatTexture(a.ActiveQuadratic)},
			{Name: "canonical generation texture selected", Passed: a.CanonicalTextureFound, Detail: fmt.Sprintf("protected=%t, mixed=%t, activeOnly=%t", a.ProtectedBFTextureFound, a.MixedBFBridgeFound, a.ActiveOnlySourceFound)},
		}, Notes: []string{
			a.TruthStatement,
			"A BF source supported only on generations sees GᵀFG; a mixed source sees GᵀFH. Both vanish in the current finite connection, so active curvature cannot be rebranded as generation texture.",
			fmt.Sprintf("remaining unknowns: %v", a.RemainingUnknowns),
		}}
	}}
}
