package tauetayukawasourcemap

import "github.com/bagherbal/asha-engine/pkg/theorem"

func DirectTauEtaYukawaSourceMapGenerationBilinearCarrierAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-DIRECT-TAU-ETA-YUKAWA-SOURCE-MAP-GENERATION-BILINEAR-CARRIER-AUDIT"
	const name = "Direct tau_eta Yukawa Source Map / Generation Bilinear Carrier Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{{Name: "build Gate 261 tau_eta Yukawa source-map audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 260 direct tau_eta route is inherited and the 8_v route remains closed", Passed: a.Inheritance.EightVRouteClosed && a.Inheritance.DirectGenerationCarrierOpened && a.Inheritance.TauEtaYukawaSourceCandidate && !a.Inheritance.DirectYukawaTextureAlreadyDerived, Detail: FormatInheritance(a.Inheritance)},
			{Name: "generation bilinear carrier Hom(G_R,G_L) is lawfully defined without using an 8_v kernel", Passed: a.BilinearCarrier.BilinearCarrierLawful && a.BilinearCarrier.TextureAlgebraDimension == 9 && !a.BilinearCarrier.UsesEightVKernel && a.BilinearCarrier.OperatorValuedCarrier && !a.BilinearCarrier.YukawaAmplitudeInserted, Detail: FormatBilinear(a.BilinearCarrier)},
			{Name: "tau_eta acts directly on generation indices as a signed 1+1+1 diagonal source", Passed: a.TauEtaAction.ActsWithoutEightVKernel && a.TauEtaAction.SelfAdjoint && a.TauEtaAction.DiagonalInTauBasis && a.TauEtaAction.SignedDistinctEigenvalueCount == 3 && a.TauEtaAction.Trace == 1 && a.TauEtaAction.Determinant == -4, Detail: FormatTauEta(a.TauEtaAction)},
			{Name: "ad_tau_eta decomposes the 3x3 texture algebra into commutant and off-diagonal complement", Passed: a.TextureAlgebra.TextureAlgebraDimension == 9 && a.TextureAlgebra.CommutantDimension == 3 && a.TextureAlgebra.OffDiagonalComplementDimension == 6 && a.TextureAlgebra.NonCommutingDirectionsExist && len(a.TextureAlgebra.DistinctAbsMixingGaps) == 3, Detail: FormatTextureAlgebra(a.TextureAlgebra)},
			{Name: "tau_eta opens a diagonal Yukawa source map but does not produce mixing by itself", Passed: a.YukawaSourceMap.ActsOnBilinearCarrier && a.YukawaSourceMap.DiagonalGenerationTextureSeed && a.YukawaSourceMap.CanSplitThreeGenerations && !a.YukawaSourceMap.CanProduceMixingByItself && a.YukawaSourceMap.RequiresSecondNonCommutingOperator, Detail: FormatYukawa(a.YukawaSourceMap)},
			{Name: "finite/sealed ledgers remain separated and EmpiricalYukawaSeal is not activated", Passed: len(a.SealLedger.FiniteDerivedItems) >= 4 && len(a.SealLedger.StillMissingItems) >= 4 && !a.SealLedger.EmpiricalMassDataUsed && !a.SealLedger.ObservedMixingDataUsed && !a.SealLedger.EmpiricalYukawaSealActivated, Detail: FormatSealLedger(a.SealLedger)},
			{Name: "physical Yukawa texture, CKM/PMNS, and masses remain blocked", Passed: !a.YukawaSourceMap.PhysicalYukawaTextureDerived && a.YukawaSourceMap.RequiresFiniteActionFunctional && a.YukawaSourceMap.RequiresKineticNormalization && a.YukawaSourceMap.RequiresScalarVEVAmplitude && !a.YukawaSourceMap.CKMPMNSDerived && !a.YukawaSourceMap.FermionMassesDerived, Detail: FormatYukawa(a.YukawaSourceMap)},
			{Name: "firewall prevents reopening 8_v route or inventing a non-commuting partner", Passed: a.Firewall.Gate260EightVNoGoPreserved && a.Firewall.DoesNotReopenEightVKernelRoute && a.Firewall.DoesNotRewriteTauEtaAsFockVector && a.Firewall.DoesNotUseObservedMasses && a.Firewall.DoesNotUseObservedMixingAngles && a.Firewall.DoesNotPromoteDiagonalSeedToCKM && a.Firewall.DoesNotInventNonCommutingPartner && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{
			a.TruthStatement,
			"Gate 261 derives a lawful diagonal generation source map from tau_eta, and it exposes the exact off-diagonal commutator complement where mixing would have to live.",
			"It does not derive a physical Yukawa texture: a canonical non-commuting finite partner, a spectral/action functional, normalization, and amplitudes remain open.",
		}}
	}}
}
