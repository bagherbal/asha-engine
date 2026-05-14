package generationbreak

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteGenerationBreakingSearchTheorem() theorem.Theorem {
	const id = "MATTER-FINITE-GENERATION-BREAKING-SEARCH"
	const name = "finite generation-breaking candidate search"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerMatter,
		Status: theorem.OpenTest,
		Verify: func() theorem.Result {
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct generation-breaking search", Passed: false, Detail: err.Error()}},
				}
			}

			checks := []theorem.Check{
				{
					Name:   "three-dimensional carrier exposed",
					Passed: a.GenerationCarrierDimension == 3 && a.ProtectedContactDimension == 3,
					Detail: fmt.Sprintf("triality generations=%d, protected contact directions=%d", a.GenerationCarrierDimension, a.ProtectedContactDimension),
				},
				{
					Name:   "exact triality no-go preserved",
					Passed: !a.Texture.ExactTrialityCanBreakAllThree && !a.Texture.ExactTrialitySelectsTexture,
					Detail: "exact triality still gives at most a 1+2 symmetric pattern, so generation breaking must come from an additional finite source",
				},
				{
					Name:   "Higgs/contact anisotropy yields diagonal spurion",
					Passed: a.DiagonalSpurionFound,
					Detail: FormatCandidate(a.BestCandidate),
				},
				{
					Name:   "second-fundamental curvature source is nonzero",
					Passed: a.SecondFundamentalSize > 0,
					Detail: fmt.Sprintf("max second-fundamental curvature norm = %.10g", a.SecondFundamentalSize),
				},
				{
					Name:   "contact leakage is real but overcomplete",
					Passed: len(a.PartialOverlapSpectrum) > a.GenerationCarrierDimension,
					Detail: fmt.Sprintf("partial-overlap spectrum has %d modes %s; no canonical 3-mode reduction yet", len(a.PartialOverlapSpectrum), FormatFloatSlice(a.PartialOverlapSpectrum)),
				},
				{
					Name:   "mixing operator not yet selected",
					Passed: !a.MixingOperatorFound && !a.CKMDerived && !a.PMNSDerived,
					Detail: "a diagonal spurion can split three generations, but CKM/PMNS require at least two non-commuting finite texture operators",
				},
				{
					Name:   "canonical generation-breaking theorem remains open",
					Passed: a.CanonicalOperatorFound,
					Detail: "requires a canonical 3x3 operator on the protected contact/triality generation carrier; current best object is a bridge-level diagonal spurion",
				},
			}
			return theorem.Result{
				ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.OpenTest,
				Checks: checks,
				Notes: []string{
					a.TruthStatement,
					"This gate searches finite sources from Higgs/contact anisotropy, second-fundamental curvature, contact leakage, and the still-open BF boundary curvature route.",
					"No observed fermion masses, CKM entries, PMNS entries, or fitted physical constants are used.",
					fmt.Sprintf("candidate inventory: %s", FormatCandidates(a.Candidates)),
				},
			}
		},
	}
}
