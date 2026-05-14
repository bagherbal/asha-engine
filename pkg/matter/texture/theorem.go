package texture

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func GenerationBreakingTextureSearchTheorem() theorem.Theorem {
	const id = "MATTER-GENERATION-BREAKING-TEXTURE-SEARCH"
	const name = "generation-breaking Yukawa texture operator search"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerMatter,
		Status: theorem.OpenTest,
		Verify: func() theorem.Result {
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct texture search", Passed: false, Detail: err.Error()}},
				}
			}
			eig := TrialityInvariantEigenvalues(1, 0.25)
			residual := TrialityDegeneracyResidual(1, 0.25)
			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerMatter,
				Status: theorem.OpenTest,
				Checks: []theorem.Check{
					{Name: "triality Yukawa input", Passed: a.Triality.GenerationCount == 3 && a.Triality.FullMixingMapCount == 72, Detail: fmt.Sprintf("generations=%d, full flavor maps=%d", a.Triality.GenerationCount, a.Triality.FullMixingMapCount)},
					{Name: "per-kind 3x3 texture spaces", Passed: len(a.KindSummaries) == 4, Detail: FormatKindSummaries(a.KindSummaries)},
					{Name: "texture-space taxonomy", Passed: len(a.TextureSpaces) == 5, Detail: FormatTextureSpaces(a.TextureSpaces)},
					{Name: "exact triality invariant no-go", Passed: !a.ExactTrialityCanBreakAllThree && a.TrialityInvariantTextureDim == 2, Detail: fmt.Sprintf("triality-invariant symmetric dim=%d gives eigen pattern 1+2; example eig=[%.3f, %.3f, %.3f], doublet residual=%.3e", a.TrialityInvariantTextureDim, eig[0], eig[1], eig[2], residual)},
					{Name: "full hierarchy not selected", Passed: !a.ExactTrialitySelectsTexture && !a.CouplingsDerived, Detail: "current finite gates expose texture spaces but do not select entries or coupling strengths"},
					{Name: "mixing matrices not derived", Passed: !a.CKMDerived && !a.PMNSDerived, Detail: "CKM/PMNS require at least two non-commuting generation-breaking texture operators"},
					{Name: "candidate operator absent", Passed: !a.GenerationBreakingOperatorFound && !a.CandidateIsCanonical, Detail: fmt.Sprintf("%s; %s", a.CandidateBreakingOperatorName, a.CandidateEigenPattern)},
					{Name: "remaining unknowns", Passed: len(a.RemainingUnknowns) > 0, Detail: FormatUnknowns(a.RemainingUnknowns)},
				},
				Notes: []string{
					a.NoGoStatement,
					"Gate 28 is a no-go/search gate: exact triality explains three copies, but generation hierarchy requires a new finite symmetry-breaking operator.",
				},
			}
		},
	}
}
