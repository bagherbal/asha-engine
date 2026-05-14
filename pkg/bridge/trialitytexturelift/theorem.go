package trialitytexturelift

import "github.com/bagherbal/asha-engine/pkg/theorem"

func TrialityLiftedYukawaTextureOperatorSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-TRIALITY-LIFTED-YUKAWA-TEXTURE-OPERATOR-SEARCH"
	const name = "triality-lifted Yukawa texture operator search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build triality-lifted texture audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "triality lift exposes four 3x3 Yukawa matrices", Passed: a.LiftAudit.GenerationCount == 3 && a.LiftAudit.FermionKindBlocks == 4 && a.LiftAudit.YukawaTextureMatrices == 4 && a.LiftAudit.TotalGeneralRealEntries == 36 && a.LiftAudit.FullMixingMaps == 72, Detail: FormatLiftAudit(a.LiftAudit)},
			{Name: "exact triality remains too symmetric", Passed: a.Texture.TrialityInvariantTextureDim == 2 && !a.Texture.ExactTrialityCanBreakAllThree && !a.Texture.ExactTrialitySelectsTexture, Detail: "triality invariant pattern=" + a.Texture.CandidateEigenPattern + " :: " + FormatCandidates(a.Candidates[:1])},
			{Name: "generation spurion exists but is not a canonical mixing texture", Passed: a.Generation.DiagonalSpurionFound && !a.Generation.BestCandidate.Canonical && !a.Generation.MixingOperatorFound && !a.Generation.CKMDerived && !a.Generation.PMNSDerived, Detail: FormatCandidate(a.Candidates[1])},
			{Name: "contact scalar-shape products remain branch-selected and aligned", Passed: a.OperatorAudit.ScalarShapeConditionalCandidates == 2 && a.OperatorAudit.BranchChoiceCandidates >= 4 && a.AxisAudit.ContactKindAssignmentsSurvive == 6 && a.AxisAudit.SeparableAnsatzOnly && a.AxisAudit.AllKindMatricesAligned, Detail: FormatOperatorAudit(a.OperatorAudit) + " :: " + FormatAxisAudit(a.AxisAudit)},
			{Name: "no canonical non-commuting texture pair is selected", Passed: !a.OperatorAudit.UniqueTextureSelected && a.OperatorAudit.CanonicalBreakingOperatorsFound == 0 && a.OperatorAudit.NonCommutingPairsFound == 0, Detail: FormatOperatorAudit(a.OperatorAudit)},
			{Name: "mass and mixing problem is now a finite matrix problem, not solved", Passed: a.MassAudit.FourYukawaMatricesRecognized && a.MassAudit.MassesAreSingularValues && a.MassAudit.MixingNeedsRelativeLeftEigenbasis && a.MassAudit.AtLeastTwoNoncommutingNeeded && !a.MassAudit.AtLeastTwoNoncommutingFound && !a.MassAudit.YukawaMatricesDerived && !a.MassAudit.FermionMassesDerived && !a.MassAudit.CKMPMNSDerived, Detail: FormatMassAudit(a.MassAudit)},
			{Name: "physical firewall remains closed", Passed: a.Firewall.GaugeRatioClosed && a.Firewall.ScalarShapeTargetAvailable && a.Firewall.TrialityLiftPerformed && !a.Firewall.CanonicalTextureOperatorSelected && !a.Firewall.YukawaAmplitudesDerived && !a.Firewall.GenerationHierarchyDerived && !a.Firewall.FermionMassesDerived && !a.Firewall.CKMPMNSDerived && !a.Firewall.PhysicalConstantsDerived && a.Firewall.ResidualNullityBefore == 3 && a.Firewall.ResidualNullityAfter == 3, Detail: FormatFirewall(a.Firewall) + " :: " + a.TruthStatement},
		}, Notes: []string{
			"Gate 172 moves the scalar-shape target from one-generation kind labels into the triality Yukawa texture arena.",
			"The lawful mass object is now explicit: four finite 3x3 Yukawa matrices Y_u, Y_d, Y_ν, Y_e.",
			"Current finite data provide exact triality, a diagonal generation spurion, and a conditional scalar-shape kind target, but no canonical non-commuting texture pair.",
		}}
	}}
}
