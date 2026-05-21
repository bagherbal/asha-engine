package generation2minimalfinitetriplerepresentationdatasealambientactivecarrieraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_MINIMAL_FINITE_TRIPLE_REPRESENTATION_DATA_SEAL_AMBIENT_ACTIVE_CARRIER_AUDIT"
	theoremName = "Gate 851 — Minimal FiniteTriple Representation DataSeal and Ambient/Active Carrier Audit"
)

func Generation2MinimalFiniteTripleRepresentationDataSealAmbientActiveCarrierAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		checks := []theorem.Check{}
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks = append(checks,
			theorem.Check{Name: "separate ambient 16/32 carrier from active 15/30 carrier", Passed: a.Carrier.AmbientActiveSeparated && a.Carrier.HPartAmbientRank == 16 && a.Carrier.HPartMinRank == 15 && a.Carrier.HFAmbientRank == 32 && a.Carrier.HFMinRank == 30 && a.Carrier.Puncture.Name == "e_+ tensor P_1" && a.Carrier.Puncture.Absent && !a.Carrier.Puncture.InMinimalCarrier && containsAll(a.Carrier.Supports, []string{SupportAmbientToActiveForkTyped}), Detail: FormatCarrier(a.Carrier)},
			theorem.Check{Name: "define minimal particle and real carrier ranks", Passed: a.Carrier.WRank == WDim && a.Carrier.HLRank == HLRank && a.Carrier.HRMinRank == HRMinRank && a.Carrier.HPartMinRank == HPartMinRank && a.J.ParticleRank == HPartMinRank && a.J.OppositeRank == HPartMinRank && a.J.TotalRank == HFMinRank && a.J.AntiunitaryExchangeSealed && !a.J.OppositeActionCertified, Detail: FormatCarrier(a.Carrier) + " | " + FormatJ(a.J)},
			theorem.Check{Name: "seal rho_F action on minimal carrier without native proof", Passed: a.RhoF.PreservesMinimalCarrier && a.RhoF.RightSocketsSealed && a.RhoF.LeptoColorBlocksPreserved && a.RhoF.WeakDoubletPreserved && a.RhoF.AbsentCellClosureSafe && !a.RhoF.CompleteActionLedger && !a.RhoF.NativeProof && containsAll(a.RhoF.Supports, []string{SupportRhoFPreservesMinimalCarrier, SupportRightCharacterProjectorsSealed, SupportLeptoColorActionSealed}), Detail: FormatRhoF(a.RhoF)},
			theorem.Check{Name: "preserve kernel-stability and weak-socket firewalls", Passed: a.Carrier.LeftKernel.Name == "h_+ tensor P_1" && a.Carrier.LeftKernel.Kernel && a.Carrier.LeftKernel.InMinimalCarrier && !a.Carrier.LeftKernel.PhysicalName && !a.RhoF.KernelStableUnderFullAction && containsAll(a.RhoF.Failures, []string{FailureKernelStabilityNotCertified, FailureWeakSocketSplitNotNative}), Detail: FormatCarrier(a.Carrier) + " | " + FormatRhoF(a.RhoF)},
			theorem.Check{Name: "seal gamma_F and J_F support data only", Passed: a.Gamma.SupportLevel && !a.Gamma.NativeGammaMatrix && !a.Gamma.KOExtensionSet && a.Gamma.LeftSign == 1 && a.Gamma.RightSign == -1 && a.J.AntiunitaryExchangeSealed && !a.J.OppositeActionCertified && !a.J.KODataCertified && containsAll(a.Gamma.Failures, []string{FailureDataSealNotNativeFiniteTripleProof, FailureNoJKOCompatibilityProof}) && containsAll(a.J.Failures, []string{FailureNoJKOCompatibilityProof}), Detail: FormatGamma(a.Gamma) + " | " + FormatJ(a.J)},
			theorem.Check{Name: "extend symbolic D_F support to minimal H_F without magnitude promotion", Passed: a.DFSym.ExtendedToJCopy && a.DFSym.SupportOnly && !a.DFSym.OperatorValued && !a.DFSym.YukawaMagnitudeSource && a.DFSym.YRank == SymbolicYRank && a.DFSym.DFRank == SymbolicDFRank && a.DFSym.KernelDim == SymbolicKernelDim && containsAll(a.DFSym.Failures, []string{FailureNoOperatorValuedDFMatrix, FailureDFSymNotYukawaMagnitudeSource, FailureNoNumericalYukawaValues}), Detail: FormatDFSym(a.DFSym)},
			theorem.Check{Name: "prepare first-order target without certifying calculation", Passed: a.FirstOrder.ObjectsPrepared && !a.FirstOrder.CanCalculateFirstOrderNow && a.FirstOrder.HasRhoSeal && a.FirstOrder.HasJSeal && a.FirstOrder.HasGammaSeal && a.FirstOrder.HasDFSymSeal && !a.FirstOrder.NativeRhoF && !a.FirstOrder.NativeJ && !a.FirstOrder.NativeGamma && !a.FirstOrder.NativeDF && len(a.FirstOrder.MissingForProof) >= 5 && containsAll(a.FirstOrder.Failures, []string{FailureNoFullFirstOrderConditionProof, FailureNoJKOCompatibilityProof, FailureNoBimoduleCommutantProof}), Detail: FormatFirstOrder(a.FirstOrder)},
			theorem.Check{Name: "preserve ledgers and R3/R4 firewalls", Passed: a.Ledger.OfficialFrozen && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing && !a.Impact.NativeFiniteTriple && !a.Impact.FirstOrderCertified && !a.Impact.JOppositeCertified && !a.Impact.KernelStable && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && !a.Impact.PhysicalNeutrinoTheorem && !a.Impact.MasslessTheorem, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			theorem.Check{Name: "preserve Gate 851 data-seal firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.DataSealNotNative && a.Firewalls.NoFirstOrderProof && a.Firewalls.NoJKOProof && a.Firewalls.NoBimoduleProof && a.Firewalls.NoOperatorDF && a.Firewalls.DFSymNotMagnitudeSource && a.Firewalls.NoNumericalYukawas && a.Firewalls.KernelStabilityNotCertified && a.Firewalls.WeakSocketSealOnly && a.Firewalls.PunctureAbsenceSealOnly && a.Firewalls.NoPhysicalNeutrino && a.Firewalls.NoRightNeutrino && a.Firewalls.NoMasslessness && a.Firewalls.AlphaStillSealed && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.NoOfficialNEffUpdate && a.Firewalls.NoCYukawaCHiggsUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.Verdict == StatusDataSealVerdict, Detail: FormatFirewalls(a.Firewalls)},
		)
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatCarrier(a.Carrier), FormatRhoF(a.RhoF), FormatGamma(a.Gamma), FormatJ(a.J), FormatDFSym(a.DFSym), FormatFirstOrder(a.FirstOrder), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
