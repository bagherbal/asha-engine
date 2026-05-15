package generation2sectordeformationsource

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2NullBaselineSectorDeformationSourceSearchTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 null-baseline sector deformation source search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate482 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits null baseline and cancellation", Passed: a.Inheritance.Executed && a.Inheritance.Gate480NullBaseline && a.Inheritance.Gate481BaselineCancellation && a.Inheritance.AlphaVac == 1 && a.Inheritance.IKVac == 0.5 && a.Inheritance.PhysicalDUDUnresolved && a.Inheritance.PhysicalDENuUnresolved && a.Inheritance.NativeRegistryClean, Detail: "alpha_vac=1 I_K,vac=1/2 inherited; shared baseline cancels from relative distances"},
			{Name: "audits native deformation candidates", Passed: a.Candidates.Executed && !a.Candidates.NativeSourceFound && a.Candidates.BridgeSourcesPresent && len(a.Candidates.Candidates) >= 5, Detail: a.Candidates.Reason},
			{Name: "proves no native perturbation source in current atlas", Passed: a.Sieve.Executed && a.Sieve.CandidatesPassingNativeSource == 0 && !a.Sieve.DeltaAlphaNative && !a.Sieve.DeltaPhiNative && a.Sieve.AllZeroPerturbationDistance == 0 && a.Sieve.AllZeroWouldPredictNoMixing, Detail: a.Sieve.Reason},
			{Name: "preserves bridge-only perturbation slot", Passed: a.BridgeSlot.Executed && a.BridgeSlot.RequiresAirlock && a.BridgeSlot.RequiresProvenance && a.BridgeSlot.RequiresBranchTags && a.BridgeSlot.RequiresUncertainty && a.BridgeSlot.RejectsCKMPMNSAsInput && a.BridgeSlot.RejectsNativePromotion && a.BridgeSlot.CanComputeSyntheticResidual && !a.BridgeSlot.CanComputePhysicalResidual, Detail: a.BridgeSlot.Reason},
			{Name: "preserves 13-moduli firewall", Passed: a.Firewall.Executed && !a.Firewall.ObservedMassImported && !a.Firewall.CKMImported && !a.Firewall.PMNSImported && a.Firewall.VacuumIKNativeBaseline && !a.Firewall.SectorPerturbationsNative && !a.Firewall.SectorPerturbationsSolved && !a.Firewall.PhysicalDUDComputed && !a.Firewall.PhysicalDENuComputed && !a.Firewall.CKMMatrixConstructed && !a.Firewall.PMNSMatrixConstructed && !a.Firewall.NativeRegistryWritten && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: a.Firewall.Reason},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusAuditCompleted, StatusBaselineInherited, StatusCandidateLedgerAudited, StatusFailedNativeSourceAbsent, StatusFailedOrientationGenerationOnly, StatusFailedChiralityGenerationBlind, StatusFailedHiggsEdgeScaleOnly, StatusFailedGaugeChargesGenerationBlind, StatusFailedYukawaSealed, StatusFailedPMNSCKMRejected, StatusEmpiricalSlotPreserved, StatusFailedNativePromotion, StatusFirewallPreserved, a.Truth}}
	}}
}
