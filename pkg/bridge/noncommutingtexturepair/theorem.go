package noncommutingtexturepair

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FiniteNonCommutingTexturePairSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-NONCOMMUTING-TEXTURE-PAIR-SEARCH"
	const name = "finite non-commuting texture-pair search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build finite non-commuting texture-pair audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "operator inventory covers currently derived generation sources", Passed: a.Inventory.OperatorCount == 9 && a.Inventory.CanonicalOperators >= 6 && a.Inventory.LinearTextureCandidates >= 5, Detail: FormatInventory(a.Inventory) + " :: " + FormatOperators(a.Operators)},
			{Name: "raw non-commuting maps are detected but quarantined", Passed: a.Inventory.RawNonCommutingPairs > 0 && a.NoGo.TrialityRawNoncommutationSeen && a.NoGo.TrialityRawMapsAreSymmetries, Detail: FormatTopRawPairs(a.Pairs, 4)},
			{Name: "triality-invariant texture algebra remains too symmetric", Passed: a.NoGo.TrialityInvariantTexturesTooSmall && a.Previous.Texture.TrialityInvariantTextureDim == 2 && !a.Previous.Texture.ExactTrialityCanBreakAllThree && !a.Previous.Texture.CKMDerived && !a.Previous.Texture.PMNSDerived, Detail: a.Previous.Texture.NoGoStatement},
			{Name: "curvature/source routes remain zero or bridge-required", Passed: a.NoGo.DiagonalSpurionRequiresBridge && a.NoGo.BFResidualZero && a.NoGo.SourceTensorMinimumZero && a.Inventory.CanonicalNonzeroCurvatureMaps == 0, Detail: FormatNoGo(a.NoGo)},
			{Name: "scalar-shape and real-structure lifts do not create generation textures", Passed: a.NoGo.ScalarShapeProjectorGenerationBlind && a.NoGo.RealStructureGenerationBlind && a.Previous.AxisAudit.ContactKindAssignmentsSurvive == 6, Detail: "Gate 171 kind branches survive; J_gen and scalar-shape lifts are identity-like on generation space"},
			{Name: "no qualified non-commuting Yukawa texture pair exists", Passed: a.Inventory.QualifiedTextureOperators == 0 && a.Inventory.QualifiedNonCommutingPairs == 0 && a.Inventory.CanonicalBreakingTextures == 0 && a.Inventory.CanonicalMixingSources == 0 && a.NoGo.NoQualifiedTexturePair, Detail: FormatInventory(a.Inventory)},
			{Name: "mass-generation problem is sealed as structurally open at current stage", Passed: a.NoGo.MassGenerationSealedAtCurrentStage && a.Firewall.MassProblemLocalizedToYukawaMatrix && a.Firewall.NonCommutingTexturePairRequired && !a.Firewall.NonCommutingTexturePairFound && !a.Firewall.YukawaAmplitudesDerived && !a.Firewall.FermionMassesDerived && !a.Firewall.CKMPMNSDerived, Detail: FormatFirewall(a.Firewall) + " :: " + a.TruthStatement},
			{Name: "absolute-coupling nullity is unchanged and handed to Gate 174", Passed: a.Firewall.ResidualNullityBefore == 3 && a.Firewall.ResidualNullityAfter == 3 && a.Firewall.RecommendedNextGate != "", Detail: a.Firewall.RecommendedNextGate},
		}, Notes: []string{
			"Gate 173 does not deny raw non-commutation inside the triality permutation representation; it denies a qualified non-commuting Yukawa texture source pair.",
			"The mass problem is now localized but not solved: the missing input is a new finite non-commuting generation-breaking source.",
			"The next independent attack is absolute gauge-coupling normalization through the topological action seal.",
		}}
	}}
}
