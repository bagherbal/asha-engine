package tauetaspatialtagging

import "github.com/bagherbal/asha-engine/pkg/theorem"

func TauEtaSpatialTaggingGenerationBreakingAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-TAU-ETA-SPATIAL-TAGGING-GENERATION-BREAKING-AUDIT"
	const name = "scalar fundamental class tau_eta spatial tagging and generation breaking audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 242 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 241 Reeb route is inherited as blocked", Passed: a.PreviousReeb.Contact.ContactProjectorExists && !a.PreviousReeb.Summary.ReebVectorDerived && !a.PreviousReeb.Summary.UniqueWeakPlaneDerived, Detail: a.PreviousReeb.TruthStatement},
			{Name: "tau_eta sequence is retrieved exactly from the scalar fundamental class", Passed: a.TauEta.StableNativeDegrees && a.TauEta.ThreeComponentSignature && a.TauEta.Sequence[0] == 2 && a.TauEta.Sequence[1] == -2 && a.TauEta.Sequence[2] == 1 && a.TauEta.ScalarTraceFunctionalOnly, Detail: FormatTauEta(a.TauEta)},
			{Name: "tau_eta magnitudes expose a 2+1 spatial-selector capacity", Passed: a.TauEta.TwoPlusOneMagnitudeSelector && a.Spatial.DimensionCompatible && a.Spatial.WeakPlaneConditionallySeen && a.Spatial.UniqueAxisIfMapped == "a†_3" && a.Spatial.ComplementPlaneIfMapped == "U={a†_1,a†_2}", Detail: FormatSpatial(a.Spatial)},
			{Name: "tau_eta is not yet pulled back to Fock spatial modes, so no weak plane is derived", Passed: !a.Spatial.NativePullbackDerived && !a.Spatial.TauEtaActsOnFockW && !a.Spatial.ManualAxisAssignment && !a.Spatial.WeakPlaneDerived && !a.Spatial.S3DegeneracyActuallyBroken, Detail: FormatSpatial(a.Spatial) + " :: " + FormatPlanes(a.Planes)},
			{Name: "tau_eta signs expose generation-breaking capacity beyond exact triality", Passed: a.Generation.TrialityCarrierDimension == 3 && a.Generation.SignedSpectrumBreaksAllThree && a.Generation.ExactTrialityKnownTooSymmetric && a.Generation.CapacitySupported, Detail: FormatGeneration(a.Generation)},
			{Name: "generation texture and mixing remain un-derived without a tau_eta-to-triality pullback", Passed: !a.Generation.TauEtaToGenerationPullback && !a.Generation.CanonicalTrialityOperatorDerived && !a.Generation.MixingOperatorDerived && !a.Generation.CKMPMNSDerived && !a.Generation.TextureDerived, Detail: FormatGeneration(a.Generation)},
			{Name: "physical weak chirality/global H/order-one calculus remain blocked", Passed: a.Weak.InheritsGate241ReebFailure && a.Weak.TauMagnitudeCanSelectAxis && !a.Weak.TauToSpatialPullbackDerived && a.Weak.UniqueWeakPlaneConditionallySeen && !a.Weak.UniqueWeakPlaneDerived && !a.Weak.PhysicalLeftHandedDerived && !a.Weak.GlobalHSummandDerived && !a.Weak.OrderOneReady, Detail: FormatWeak(a.Weak)},
			{Name: "firewall preserved: no trace functional is promoted to a spinor/Fock/generation operator", Passed: !a.Firewall.ForcedTauToSpatialMap && !a.Firewall.ForcedAxisAssignment && !a.Firewall.ImportedSMWeakPlane && !a.Firewall.ImportedGenerationTexture && !a.Firewall.PromotedTraceToSpinorMatrix && !a.Firewall.ClaimedPhysicalChirality && !a.Firewall.ClaimedGlobalH && !a.Firewall.ClaimedCKMPMNS && !a.Firewall.ClaimedMasses && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records selector capacity but not completed selection", Passed: a.Summary.TauEtaRetrieved && a.Summary.MagnitudeSelectorCapacity && !a.Summary.SpatialPullbackDerived && a.Summary.WeakPlaneConditionallySeen && !a.Summary.WeakPlaneDerived && a.Summary.GenerationBreakingCapacity && !a.Summary.GenerationTextureDerived && !a.Summary.GlobalHDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}
		notes := []string{
			"Gate 242 audits tau_eta=(2,-2,1) as an exact finite scalar fundamental-class signature, not as an imported weak-plane or generation operator.",
			"The magnitudes |tau_eta|=(2,2,1) have exactly the shape needed to tag one spatial axis and select the complementary pure-spatial two-plane, but the tau_eta -> W_spatial pullback is missing.",
			"The signed spectrum (2,-2,1) has exactly the 1+1+1 capacity missing from exact triality, but the tau_eta -> triality-generation pullback and non-commuting texture structure remain missing.",
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
