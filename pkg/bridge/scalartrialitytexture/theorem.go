package scalartrialitytexture

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ScalarBundleToTrialityPullbackYukawaGenerationTextureAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-SCALAR-BUNDLE-TO-TRIALITY-PULLBACK-YUKAWA-GENERATION-TEXTURE-AUDIT"
	const name = "Scalar Bundle to Triality Pullback / Yukawa Generation Texture Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 246 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 245 neutral-scalar origin is inherited", Passed: a.PreviousGate245.TauEtaOriginTraced && a.PreviousGate245.SourceIsNeutralScalarEWPlane && !a.PreviousGate245.TauSlotsAreSU2Basis && !a.PreviousGate245.CarrierProjectionDerived && !a.PreviousGate245.WeakPlaneDerived, Detail: a.PreviousGate245.TruthStatement},
			{Name: "scalar-to-flavor destination is type-plausible but not derived", Passed: a.ScalarFlavor.ScalarOriginKnown && a.ScalarFlavor.NativeHiggsSectorObservable && a.ScalarFlavor.TrialityCarrierDimension == 3 && a.ScalarFlavor.TrialityCarrierDerivedEarlier && a.ScalarFlavor.MapWouldBeTypeCorrect && !a.ScalarFlavor.ScalarToTrialityFunctorDerived && !a.ScalarFlavor.MapActuallyDerived, Detail: FormatScalarFlavor(a.ScalarFlavor)},
			{Name: "tau_eta supplies exact 1+1+1 generation-breaking capacity", Passed: a.GenerationTexture.DistinctEigenvalues == 3 && a.GenerationTexture.BreaksS3Degeneracy && a.GenerationTexture.HermitianDiagonalCapacity && a.GenerationTexture.Trace == 1 && a.GenerationTexture.Determinant == -4 && a.GenerationTexture.FrobeniusNormSquared == 9 && !a.GenerationTexture.GenerationOperatorDerived && !a.GenerationTexture.YukawaTextureDerived, Detail: FormatGenerationTexture(a.GenerationTexture)},
			{Name: "conditional tau texture would solve Gate 173 non-commuting-capacity target", Passed: a.NonCommutingTexture.Gate173NeedsNonCommutingPair && !a.NonCommutingTexture.Gate173FoundQualifiedPair && a.NonCommutingTexture.RawNonCommutingWithTriality && a.NonCommutingTexture.CycleCommutatorNorm > 0 && a.NonCommutingTexture.ReflectionCommutatorNorm > 0 && a.NonCommutingTexture.PairWouldBeQualifiedIfPullbackHeld && !a.NonCommutingTexture.PairActuallyQualified && !a.NonCommutingTexture.CKMDerived && !a.NonCommutingTexture.PMNSDerived, Detail: FormatNonCommuting(a.NonCommutingTexture)},
			{Name: "scalar-to-triality pullback remains the binding obstruction", Passed: !a.PullbackObstruction.PullbackDerived && a.PullbackObstruction.ManualInsertionRejected && a.PullbackObstruction.MissingFunctor != "" && len(a.PullbackObstruction.MissingCompatibility) >= 4, Detail: FormatPullback(a.PullbackObstruction)},
			{Name: "firewall preserved: no masses or mixing matrices inserted", Passed: !a.Firewall.ForcedScalarToGenerationMap && !a.Firewall.ForcedTauDiagonalTexture && !a.Firewall.ImportedYukawaMasses && !a.Firewall.ImportedCKM && !a.Firewall.ImportedPMNS && !a.Firewall.InsertedObservedMasses && !a.Firewall.ClaimedFermionMasses && !a.Firewall.ClaimedFiniteFlavorTheorem && !a.Firewall.ClaimedWeakPlane && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records flavor capacity but not derivation", Passed: a.Summary.ScalarOriginKnown && !a.Summary.ScalarToTrialityFunctorDerived && a.Summary.TauGenerationCapacity && !a.Summary.GenerationTextureDerived && a.Summary.RawNonCommutingCapacity && !a.Summary.QualifiedTexturePairDerived && !a.Summary.CKMPMNSDerived && !a.Summary.FermionMassesDerived && !a.Summary.WeakPlaneDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}
		notes := []string{
			"Gate 246 accepts Gate 245's correction: tau_eta is a neutral scalar/Higgs-sector trace invariant, not a spatial gauge-axis selector.",
			"The signed triple (2,-2,1) has exact 1+1+1 generation-breaking capacity. If lawfully pulled back, D_tau=diag(2,-2,1) would not commute with triality permutations and would supply the missing kind of texture source Gate 173 asked for.",
			"The scalar-to-triality functor is not derived. Therefore no Yukawa matrices, fermion masses, CKM, PMNS, or finite flavor theorem is claimed.",
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
