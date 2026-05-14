package sourcepotential

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func SymmetryBreakingSourceActionTheorem() theorem.Theorem {
	const id = "MATTER-SYMMETRY-BREAKING-SOURCE-ACTION"
	const name = "symmetry-breaking source action search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.OpenTest, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct symmetry-breaking source action audit", Passed: false, Detail: err.Error()}}}
		}
		status := theorem.OpenTest
		if a.NonzeroRadiusDerived && a.TensorOrientationFound {
			status = theorem.BridgeRequired
		}
		candidateDetails := make([]string, 0, len(a.Candidates))
		for _, c := range a.Candidates {
			candidateDetails = append(candidateDetails, FormatCandidate(c))
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: status, Checks: []theorem.Check{
			{Name: "source tensor domain", Passed: a.GenerationDimension == 3 && a.ActiveDimension == 4 && a.TensorDimension == 12, Detail: fmt.Sprintf("M:H_active→H_generation has dimensions %dx%d; Hom dimension=%d", a.GenerationDimension, a.ActiveDimension, a.TensorDimension)},
			{Name: "finite scalar invariants available", Passed: a.ScalarInvariantsFound, Detail: fmt.Sprintf("τ_H=%.10f, Tr(M_K²)=%.10f, activeBFTrace=%.10f, L_BG²=%.10f", a.HiggsOrderParameter, a.HiggsQuarticTrace, a.ActiveBFScalarTrace, a.ContactLeakageSq)},
			{Name: "positive stabilizing data", Passed: a.PositiveQuarticFound, Detail: fmt.Sprintf("Higgs/contact normalized quartic shape=%.10f", a.Potential.NormalizedQuarticShape)},
			{Name: "tachyonic source-tensor sign", Passed: a.TachyonicSignDerived, Detail: "requires a derived negative quadratic coefficient for M, not just positive scalar norms"},
			{Name: "nonzero source-tensor radius", Passed: a.NonzeroRadiusDerived, Detail: fmt.Sprintf("current finite source action stationary norm=%.6e", a.SourceAction.NaturalStationaryNorm)},
			{Name: "source-tensor orientation", Passed: a.TensorOrientationFound, Detail: "requires a canonical direction in Hom(R⁴,R³); current active-generation maps vanish"},
			{Name: "zero-map stability persists", Passed: a.StableZeroPersists, Detail: "minimal positive action still selects M=0"},
			{Name: "arbitrary radius rejected", Passed: a.ArbitraryRadiusRejected, Detail: "choosing a Mexican-hat radius/orientation by hand would be fitting, not derivation"},
		}, Notes: []string{
			a.TruthStatement,
			"Finite Higgs/contact and active-BF curvature belong to the scalar sector unless a canonical source tensor couples them into generations.",
			fmt.Sprintf("candidate audit: %s", strings.Join(candidateDetails, " | ")),
			fmt.Sprintf("remaining unknowns: %v", a.RemainingUnknowns),
		}}
	}}
}
