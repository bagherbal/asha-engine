package characteristicpullback

import "github.com/bagherbal/asha-engine/pkg/theorem"

func CharacteristicClassOperatorToModePullbackAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-CHARACTERISTIC-CLASS-OPERATOR-TO-MODE-PULLBACK-AUDIT"
	const name = "Characteristic Class / Operator-to-Mode Pullback Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 244 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 243 obstruction is inherited", Passed: a.PreviousCliffordPullback.CliffordActionAvailable && !a.PreviousCliffordPullback.TauEtaInCliffordDomain && !a.PreviousCliffordPullback.EndomorphismConstructed, Detail: a.PreviousCliffordPullback.TruthStatement},
			{Name: "tau_eta operator origins are traced exactly", Passed: a.Origin.ExactOperatorOriginsRecovered && a.Origin.StableNativeDegrees && len(a.Origin.Sequence) == 3 && a.Origin.Sequence[0] == 2 && a.Origin.Sequence[1] == -2 && a.Origin.Sequence[2] == 1 && a.Origin.OperatorsAreCurvatureObservables, Detail: FormatOrigin(a.Origin)},
			{Name: "source operators are not spatial Fock-mode projectors", Passed: !a.Origin.OperatorsAreSpatialModeProjectors && !a.Origin.OperatorsAreBasisBlades && !a.SpatialAlignment.NativeOperatorToModeMapDerived && !a.SpatialAlignment.QZT3YInherentlyLinkToSpatialAxes && !a.SpatialAlignment.OperatorDefinitionsUseFockModes && !a.SpatialAlignment.ScalarBundleToFockProjectionDerived && a.SpatialAlignment.ManualMapRejected, Detail: FormatSpatialAlignment(a.SpatialAlignment)},
			{Name: "exterior characteristic representative remains un-derived", Passed: a.CharacteristicRep.CharacteristicClassLanguageAvailable && a.CharacteristicRep.FiniteEtaTraceFunctional && !a.CharacteristicRep.ChernCharacterRepresentativeDerived && !a.CharacteristicRep.PontryaginFormRepresentativeDerived && !a.CharacteristicRep.ExteriorGradeKnown && !a.CharacteristicRep.BasisBladeLabelsKnown && !a.CharacteristicRep.RepresentativeConstructed && a.CharacteristicRep.HypotheticalRepresentativeRejected, Detail: FormatCharacteristic(a.CharacteristicRep)},
			{Name: "weak plane remains conditional only", Passed: a.WeakPlane.InheritedConditionalAxis == "a†_3" && a.WeakPlane.InheritedConditionalPlane == "U={a†_1,a†_2}" && !a.WeakPlane.ExteriorRepresentativeDerived && !a.WeakPlane.SpatialModeAlignmentDerived && !a.WeakPlane.S3DegeneracyBroken && !a.WeakPlane.PhysicalWeakPlaneDerived && !a.WeakPlane.GlobalHSummandDerived, Detail: FormatWeak(a.WeakPlane)},
			{Name: "generation breaking remains capacity only", Passed: len(a.Generation.Sequence) == 3 && a.Generation.DistinctEigenvalueCapacity && !a.Generation.CharacteristicRepresentative && !a.Generation.TrialityCarrierMapDerived && !a.Generation.GenerationOperatorDerived && !a.Generation.GenerationTextureDerived && !a.Generation.CKMPMNSDerived, Detail: FormatGeneration(a.Generation)},
			{Name: "firewall preserved: no hand-labelled representative", Passed: !a.Firewall.ForcedOperatorToModeMap && !a.Firewall.ForcedExteriorRepresentative && !a.Firewall.ForcedCharacteristicClass && !a.Firewall.ForcedTrialityMap && !a.Firewall.ImportedWeakPlane && !a.Firewall.ImportedGenerationTexture && !a.Firewall.PromotedScalarTraceToMatrix && !a.Firewall.ClaimedPhysicalChirality && !a.Firewall.ClaimedGlobalH && !a.Firewall.ClaimedCKMPMNS && !a.Firewall.ClaimedFermionMasses && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records origin support but failed pullback", Passed: a.Summary.TauEtaOriginTraced && a.Summary.NativeSequenceStable && !a.Summary.OperatorModeAlignmentDerived && !a.Summary.ExteriorRepresentativeDerived && !a.Summary.CharacteristicClassDerived && a.Summary.WeakPlaneConditionallyVisible && !a.Summary.WeakPlaneDerived && a.Summary.GenerationBreakingCapacity && !a.Summary.GenerationTextureDerived && !a.Summary.GlobalHDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}
		notes := []string{
			"Gate 244 recovers the exact source expressions behind tau_eta=(2,-2,1): tau_eta(Q^TQ), tau_eta(Z^TZ), and tau_eta(T3L^T Y_phi).",
			"Those source labels are scalar-bundle curvature observables, not spatial Fock-mode projectors or exterior basis blades.",
			"The tempting representative omega_tau=2e1-2e2+e3 remains rejected until a scalar-bundle-to-Fock characteristic-class carrier theorem is derived.",
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
