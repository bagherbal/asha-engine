package yukawaamplitudesource

import (
	"github.com/bagherbal/asha-engine/pkg/bridge/scalaryukawasupport"
	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteYukawaTextureOperatorAmplitudeSourceObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-YUKAWA-TEXTURE-OPERATOR-AMPLITUDE-SOURCE-OBSTRUCTION-AUDIT"
	const name = "finite Yukawa texture operator / amplitude-source obstruction audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Yukawa amplitude-source audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "Gate 194 support theorem is inherited but remains support-only", Passed: a.Summary.Gate194SupportInherited && a.Support.BilinearSupport.SupportedChannels == 8 && !a.Support.Firewall.PhysicalYukawaAmplitudesDerived && !a.Support.Firewall.FermionMassesDerived, Detail: scalaryukawasupport.FormatBilinear(a.Support.BilinearSupport)},
			{Name: "tensor-lifted fundamental class is generation-blind", Passed: a.Generation.GenerationDimension == 3 && a.Generation.PermutationInvariant && a.Generation.ProjectsToIdentity && !a.Generation.OffDiagonalEntriesSelected && a.Generation.GenerationBlind, Detail: FormatGeneration(a.Generation)},
			{Name: "exact triality still does not select a 3x3 Yukawa texture", Passed: a.Triality.GenerationDimension == 3 && a.Triality.FermionKindBlocks == 4 && a.Triality.GeneralEntriesPerKind == 9 && a.Triality.TrialityInvariantDimPerKind == 2 && a.Triality.FullMixingMapsAllowedByCharges == 72 && !a.Triality.ExactTrialitySelectsTexture && !a.Triality.ExactTrialityCanBreakAllThree && !a.Triality.TextureOperatorFound && !a.Triality.CouplingsDerived && !a.Triality.CKMDerived && !a.Triality.PMNSDerived, Detail: FormatTriality(a.Triality)},
			{Name: "sealed scalar/weak curvature induces no flavor texture pair", Passed: a.Curvature.GaugeGeneratorsActOnScalarWeakFactor && !a.Curvature.GaugeGeneratorsActOnGenerationFactor && a.Curvature.T1T2ScalarOffDiagonal && !a.Curvature.T1T2FlavorOffDiagonal && a.Curvature.CommutatorWithGenerationIdentityNorm == 0 && !a.Curvature.NonCommutingTexturePairInduced, Detail: FormatCurvature(a.Curvature)},
			{Name: "all amplitude-source candidates are obstructed or external", Passed: a.SourceSearch.CandidateCount == 8 && a.SourceSearch.SelectedAmplitudeSources == 0 && a.SourceSearch.InsertedFreeParameterSpaces == 1 && a.SourceSearch.ForbiddenObservedTargetsRejected == 1 && a.SourceSearch.NoCanonicalAmplitudeSource && a.SourceSearch.NoNonCommutingTexturePair && a.SourceSearch.SupportGeometryDoesNotFixTexture, Detail: FormatSourceSearch(a.SourceSearch) + " candidates=" + FormatCandidates(a.Candidates)},
			{Name: "firewall keeps masses, textures, mixings, observed ratios, VEV, thresholds, couplings, and constants sealed", Passed: a.Firewall.SupportGeometryDerived && !a.Firewall.YukawaTextureMatricesDerived && !a.Firewall.YukawaAmplitudesDerived && !a.Firewall.FermionMassesDerived && !a.Firewall.GenerationHierarchyDerived && !a.Firewall.CKMMatrixDerived && !a.Firewall.PMNSMatrixDerived && !a.Firewall.ObservedMassRatiosImported && !a.Firewall.CabibboAngleImported && !a.Firewall.HiggsVEVAmplitudeInserted && a.Firewall.FreeParameterInsertionNeeded && a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records positive obstruction theorem", Passed: a.Summary.TestsAudited == 5 && a.Summary.GenerationFunctionalBlind && a.Summary.TrialityTextureStillUnselected && a.Summary.CurvaturePullbackGenerationBlind && a.Summary.NoAmplitudeSourceFound && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}, Notes: []string{
			a.TruthStatement,
			"Gate 195 is FAILED_ROUTE only as a constructive amplitude-source route. As an obstruction theorem, all checks pass: support geometry does not determine texture.",
			"The lawful next step is a quarantined Yukawa-amplitude seal or a genuinely new finite generation-breaking source, not an untracked mass fit.",
		}}
	}}
}
