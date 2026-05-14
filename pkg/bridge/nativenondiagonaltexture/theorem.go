package nativenondiagonaltexture

import "github.com/bagherbal/asha-engine/pkg/theorem"

func NativeNonDiagonalTextureFlavorOrientationSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-NATIVE-NON-DIAGONAL-TEXTURE-FLAVOR-ORIENTATION-SIEVE"
	const name = "Native Non-Diagonal Texture / Flavor Orientation Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 356 native non-diagonal texture audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "span inherits Gate 355 without adding a fit", Passed: a.Span.InheritedGate == 355 && !a.Span.AddsFit, Detail: FormatSpan(a.Span)},
			{Name: "rotation search includes DFT/cyclic/identity candidates", Passed: a.Search.Formalized && len(a.Search.Candidates) == 3, Detail: FormatSearch(a.Search)},
			{Name: "DFT candidate is off-diagonal and exposes sign interference", Passed: a.Search.Candidates[0].Name == "normalized DFT3 flavor rotation" && a.Search.Candidates[0].OffDiagonal && a.Search.Candidates[0].VisibleSignInterference, Detail: FormatCandidate(a.Search.Candidates[0])},
			{Name: "unitary singular-value invariance is proven", Passed: a.Invariance.Proved && a.Invariance.SeedSpectrum == [3]float64{2, 2, 1}, Detail: FormatInvariance(a.Invariance)},
			{Name: "no unitary candidate breaks the 2:2:1 singular hierarchy", Passed: !a.Search.AnyHierarchyBroken && a.Search.BestHighLowRatio == 2, Detail: FormatSearch(a.Search)},
			{Name: "non-unitary/projected texture requirement is formalized", Passed: a.Requirement.Formalized && a.Requirement.NeedsNonUnitaryProjector && a.Requirement.NeedsAdditionalOperator, Detail: FormatRequirement(a.Requirement)},
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
