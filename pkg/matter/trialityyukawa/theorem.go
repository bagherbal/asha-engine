package trialityyukawa

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func GenerationTrialityYukawaTheorem() theorem.Theorem {
	const id = "MATTER-TRIALITY-YUKAWA-EXTENSION"
	const name = "generation/triality extension of gauge-compatible Yukawa channels"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerMatter,
		Status: theorem.BridgeRequired,
		Verify: func() theorem.Result {
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct triality Yukawa extension", Passed: false, Detail: err.Error()}},
				}
			}
			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerMatter,
				Status: theorem.BridgeRequired,
				Checks: []theorem.Check{
					{Name: "triality sector count", Passed: a.GenerationCount == 3, Detail: FormatSectors(a.TrialitySectors)},
					{Name: "one-generation seed", Passed: a.OneGenerationChannels == 8 && a.OneGenerationFiberEntries == 16, Detail: fmt.Sprintf("one-generation channels=%d, scalar-fiber entries=%d", a.OneGenerationChannels, a.OneGenerationFiberEntries)},
					{Name: "generation-diagonal lift", Passed: a.TrialityCopiesChannelPattern && a.DiagonalChannelCount == 24, Detail: fmt.Sprintf("diagonal triality channels=%d = 3×%d; sample %s", a.DiagonalChannelCount, a.OneGenerationChannels, FormatChannelSample(a.DiagonalChannels, 4))},
					{Name: "full flavor-mixing selection space", Passed: a.GenerationMixingAllowedByCharges && a.FullMixingMapCount == 72, Detail: fmt.Sprintf("charge rules allow %d maps = one-generation channels × 3×3", a.FullMixingMapCount)},
					{Name: "scalar-fiber extension", Passed: a.DiagonalFiberEntries == 48 && a.FullMixingFiberEntries == 144, Detail: fmt.Sprintf("diagonal fiber=%d, full-mixing fiber=%d", a.DiagonalFiberEntries, a.FullMixingFiberEntries)},
					{Name: "per-kind texture blocks", Passed: len(a.KindSummaries) == 4, Detail: FormatKindSummaries(a.KindSummaries)},
					{Name: "triality does not select texture", Passed: !a.TextureSelectedByFiniteData && !a.CouplingsDerived, Detail: "this gate replicates allowed channels and exposes 3×3 texture spaces; it does not derive coupling strengths"},
					{Name: "mixing matrices not yet derived", Passed: !a.CKMDerived && !a.PMNSDerived, Detail: "CKM/PMNS require finite generation-breaking Yukawa operators, not only triality bookkeeping"},
					{Name: "remaining bridge unknowns", Passed: len(a.RemainingUnknowns) > 0, Detail: FormatUnknowns(a.RemainingUnknowns)},
				},
				Notes: []string{
					"Gate 26 lifts the one-generation gauge-compatible Yukawa channels into three triality sectors.",
					"The charge rules allow 3×3 generation mixing for each fermion kind, but the current finite data do not yet choose the numerical texture.",
				},
			}
		},
	}
}
