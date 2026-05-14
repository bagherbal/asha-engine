package bfcurvature

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteMaurerCartanCurvatureTheorem() theorem.Theorem {
	const id = "MATTER-FINITE-BF-MAURER-CARTAN-CURVATURE"
	const name = "finite BF/Maurer-Cartan curvature on the Boolean block connection"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.OpenTest, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct finite BF curvature", Passed: false, Detail: err.Error()}}}
		}
		canonical := a.CanonicalTextureFound
		status := theorem.OpenTest
		if canonical {
			status = theorem.BridgeRequired
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: status, Checks: []theorem.Check{
			{Name: "Boolean connection seed", Passed: a.GeneratorCount > 0 && a.SeedSpanRank > 0, Detail: fmt.Sprintf("generators=%d, seed span rank=%d on %dD Boolean support", a.GeneratorCount, a.SeedSpanRank, a.FullDimension)},
			{Name: "finite Maurer-Cartan residual", Passed: a.FullMaxNorm > 1e-8 && a.FullCurvatureSpanRank > 0, Detail: fmt.Sprintf("max ||[A_i,A_j]-Π_seed[A_i,A_j]||_F=%.6e, span rank=%d", a.FullMaxNorm, a.FullCurvatureSpanRank)},
			{Name: "protected generation restriction", Passed: a.ProtectedMaxNorm > 1e-8 && a.ProtectedSpanRank > 0, Detail: fmt.Sprintf("dim=%d, max ||GᵀFG||=%.6e, span rank=%d, offdiag=%.6e", a.ProtectedDimension, a.ProtectedMaxNorm, a.ProtectedSpanRank, a.ProtectedMaxOffDiagonalNorm)},
			{Name: "active Higgs/contact restriction", Passed: a.ActiveMaxNorm > 1e-8 && a.ActiveSpanRank > 0, Detail: fmt.Sprintf("dim=%d, max ||HᵀFH||=%.6e, span rank=%d", a.ActiveDimension, a.ActiveMaxNorm, a.ActiveSpanRank)},
			{Name: "active-to-generation curvature bridge", Passed: a.ActiveToGenerationBridgeFound, Detail: fmt.Sprintf("max ||GᵀFH||=%.6e, span rank=%d", a.CrossMaxNorm, a.CrossSpanRank)},
			{Name: "non-diagonal generation texture candidate", Passed: a.NonDiagonalGenerationMixingFound, Detail: fmt.Sprintf("nonDiagonal=%t; canonicalTexture=%t", a.NonDiagonalGenerationMixingFound, a.CanonicalTextureFound)},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("protected operators: %s", FormatOperators(a.ProtectedCurvatures)),
			fmt.Sprintf("cross operators: %s", FormatOperators(a.CrossCurvatures)),
			fmt.Sprintf("remaining unknowns: %v", a.RemainingUnknowns),
		}}
	}}
}
