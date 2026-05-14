package nonunitaryprojectortexture

import "github.com/bagherbal/asha-engine/pkg/theorem"

func NonUnitaryProjectorKineticSafeFlavorTextureSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-NON-UNITARY-PROJECTOR-KINETIC-SAFE-FLAVOR-TEXTURE-SIEVE"
	const name = "Non-Unitary Projector / Kinetic-Safe Flavor Texture Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 357 non-unitary projector texture audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "span inherits Gate 356 without adding fit", Passed: a.Span.InheritedGate == 356 && !a.Span.AddsFit, Detail: FormatSpan(a.Span)},
			{Name: "projector search includes signed tau candidates", Passed: a.Search.Formalized && len(a.Search.Candidates) == 3, Detail: FormatSearch(a.Search)},
			{Name: "tau ray projector is rank one and non-kinetic-safe", Passed: a.Search.Candidates[0].Rank == 1 && a.Search.Candidates[0].RankDefect && !a.Search.Candidates[0].KineticSafe, Detail: FormatCandidate(a.Search.Candidates[0])},
			{Name: "tau null complement is rank two and splits only by rank defect", Passed: a.Search.Candidates[1].Rank == 2 && a.Search.Candidates[1].RankDefect && a.Search.Candidates[1].FirstSecondSplit > 0, Detail: FormatCandidate(a.Search.Candidates[1])},
			{Name: "kinetic safety requires native positive wavefunction repair", Passed: a.Safety.Audited && !a.Safety.NativeRepairDerived, Detail: FormatSafety(a.Safety)},
			{Name: "no kinetic-safe hierarchy is derived", Passed: !a.Search.AnyKineticSafeHierarchy && !a.Summary.KineticSafeHierarchy, Detail: FormatSummary(a.Summary)},
			{Name: "parameter census remains unreduced", Passed: a.Census.TotalReduction == 0 && a.Census.RemainingInputs == 15 && !a.Census.SevenSealReached, Detail: FormatCensus(a.Census)},
		}
		passed := 0
		for _, c := range checks {
			if c.Passed {
				passed++
			}
		}
		status := theorem.BridgeRequired
		if passed != len(checks) {
			status = theorem.FailedRoute
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks}
	}}
}
