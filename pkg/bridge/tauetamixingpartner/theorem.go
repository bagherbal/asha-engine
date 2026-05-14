package tauetamixingpartner

import "github.com/bagherbal/asha-engine/pkg/theorem"

func TauEtaNonCommutingPartnerFinitePhaseMixingSourceAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-TAU-ETA-NON-COMMUTING-PARTNER-FINITE-PHASE-MIXING-SOURCE-AUDIT"
	const name = "TauEta Non-Commuting Partner / Finite Phase-Mixing Source Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{{Name: "build Gate 262 tau_eta mixing-partner audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 261 tau_eta bilinear carrier and 6D off-diagonal complement are inherited", Passed: a.Inheritance.BilinearCarrierDefined && a.Inheritance.TauEtaActionDerived && a.Inheritance.TextureAlgebraDecomposed && a.Inheritance.OffDiagonalComplementDimension == 6 && !a.Inheritance.PreviousCanonicalPartnerFound, Detail: FormatInheritance(a.Inheritance)},
			{Name: "triality operators populate the ad_tau off-diagonal complement", Passed: a.Inventory.RawNonCommutingCandidates >= 4 && a.TrialityPartner.RawComplementDirectionsTouched == 6 && a.TrialityPartner.PermutationCycleNonCommuting && a.TrialityPartner.ReflectionNonCommuting, Detail: FormatTriality(a.TrialityPartner)},
			{Name: "Hermitian real and phase triality bases are exposed but not promoted to amplitudes", Passed: a.TrialityPartner.HermitianRealPartNonCommuting && a.TrialityPartner.HermitianImaginaryPartNonCommuting && a.Inventory.SelfAdjointRawCandidates >= 3 && !a.TrialityPartner.AnyTrialityMapQualifiedAsAmplitude, Detail: FormatInventory(a.Inventory)},
			{Name: "B_gap remains a scalar gap without generation-endomorphism support", Passed: a.FinitePhaseGap.BGapAvailableAsPositiveScale && !a.FinitePhaseGap.BGapHasGenerationMatrix && !a.FinitePhaseGap.BGapCanPopulateOffDiagonalComplement, Detail: FormatFinitePhaseGap(a.FinitePhaseGap)},
			{Name: "Hopf phase residuals remain representation-free for M3(C) texture purposes", Passed: a.FinitePhaseGap.HopfPhasesAvailableAsPhaseLedger && !a.FinitePhaseGap.HopfPhaseGenerationMapDerived && !a.FinitePhaseGap.HopfCanPopulateOffDiagonalComplement, Detail: FormatFinitePhaseGap(a.FinitePhaseGap)},
			{Name: "no qualified finite Yukawa mixing partner is identified", Passed: a.PartnerVerdict.RawNonCommutingPartnerExists && a.PartnerVerdict.RawSelfAdjointOffDiagonalBasisExists && !a.PartnerVerdict.QualifiedFiniteMixingPartnerFound && !a.PartnerVerdict.PhysicalYukawaTextureDerived, Detail: FormatPartnerVerdict(a.PartnerVerdict)},
			{Name: "CKM/PMNS, fermion masses, and empirical Yukawa seal remain blocked", Passed: !a.PartnerVerdict.CKMPMNSDerived && !a.PartnerVerdict.FermionMassesDerived && a.Firewall.EmpiricalYukawaSealInactive && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "firewall prevents promoting symmetry maps, B_gap, or Hopf residuals into amplitudes", Passed: a.Firewall.DoesNotPromoteSymmetryToAmplitude && a.Firewall.DoesNotUseBGapAsTextureWithoutMap && a.Firewall.DoesNotUseHopfPhaseWithoutMap && a.Firewall.DoesNotClaimFiniteActionFunctional && a.Firewall.DoesNotReopenEightVKernelRoute, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{
			a.TruthStatement,
			"Gate 262 finds raw finite off-diagonal triality algebra, including exact Hermitian real/phase bases in the 6D complement exposed by Gate 261.",
			"It does not identify a qualified finite Yukawa amplitude source: symmetry maps lack action coefficients, B_gap is scalar-only, and Hopf phases remain representation-free.",
		}}
	}}
}
