package sourcemap

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func SourceTensorSelectionTheorem() theorem.Theorem {
	const id = "MATTER-SOURCE-TENSOR-SELECTION"
	const name = "source tensor selection and active-generation map search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.OpenTest, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct source tensor search", Passed: false, Detail: err.Error()}}}
		}
		status := theorem.OpenTest
		if a.CanonicalSourceTensorFound {
			status = theorem.BridgeRequired
		}
		candDetails := make([]string, 0, len(a.Candidates))
		for _, c := range a.Candidates {
			candDetails = append(candDetails, FormatCandidate(c))
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: status, Checks: []theorem.Check{
			{Name: "source tensor domain", Passed: a.GenerationDimension == 3 && a.ActiveDimension == 4, Detail: fmt.Sprintf("Hom(H_active,H_generation): %d→%d, dimension=%d", a.ActiveDimension, a.GenerationDimension, a.MapSpaceDimension)},
			{Name: "existing connection cross-map", Passed: a.ExistingConnectionMapFound, Detail: fmt.Sprintf("rank=%d, max ||GᵀAH||=%.6e", a.Bridge.CrossMapSpanRank, a.Bridge.MaxCrossMapNorm)},
			{Name: "BF curvature cross-map", Passed: a.BFCurvatureMapFound, Detail: fmt.Sprintf("rank=%d, max ||GᵀFH||=%.6e", a.BFSource.MixedBFResponseRank, a.BFSource.MixedBFMaxNorm)},
			{Name: "BF action mixed source", Passed: a.BFSourceMapFound, Detail: fmt.Sprintf("rank=%d, norm=%.6e", a.BFSource.MixedQuadratic.Rank, a.BFSource.MixedQuadratic.Norm)},
			{Name: "diagonal spurion is not a source tensor", Passed: a.GenerationBreak.DiagonalSpurionFound && !a.GenerationBreak.MixingOperatorFound, Detail: fmt.Sprintf("diagonal split exists=%t, mixing=%t", a.GenerationBreak.DiagonalSpurionFound, a.GenerationBreak.MixingOperatorFound)},
			{Name: "arbitrary map space rejected as derivation", Passed: a.ArbitraryMapsExist && !a.CanonicalSourceTensorFound, Detail: fmt.Sprintf("%d arbitrary 3x4 entries exist, but no finite theorem selects them", a.MapSpaceDimension)},
			{Name: "canonical source tensor selected", Passed: a.CanonicalSourceTensorFound, Detail: fmt.Sprintf("best=%s", FormatCandidate(a.BestCandidate))},
			{Name: "CKM/PMNS discipline", Passed: !a.CKMDerived && !a.PMNSDerived, Detail: "no non-commuting generation texture pair has been derived"},
		}, Notes: []string{
			a.TruthStatement,
			"The search distinguishes existence of arbitrary maps from derivation of a canonical finite source tensor.",
			fmt.Sprintf("candidate audit: %s", strings.Join(candDetails, " | ")),
			fmt.Sprintf("remaining unknowns: %v", a.RemainingUnknowns),
		}}
	}}
}
