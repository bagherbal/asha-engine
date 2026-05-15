package generation2topologicaldeformation

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2FiniteAlgebraicDeformationOperatorSearchTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 finite algebraic deformation operator search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate483 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits null-baseline perturbation frontier", Passed: a.Inheritance.Executed && a.Inheritance.Gate480NullBaseline && a.Inheritance.Gate481Cancellation && a.Inheritance.Gate482SourceAbsent && a.Inheritance.AlphaVac == 1 && a.Inheritance.IKVac == 0.5 && a.Inheritance.SectorPerturbationsUnsolved && a.Inheritance.NativeRegistryClean, Detail: "Gate480/481/482 frontier inherited: null baseline is native, but sector perturbations remain unsolved"},
			{Name: "audits finite topological candidates", Passed: a.TopologicalAudit.Executed && a.TopologicalAudit.SectorSeparatorsFound && a.TopologicalAudit.QuarkLeptonOnly && !a.TopologicalAudit.NativeFullSourceFound && a.TopologicalAudit.BridgeLikeSourcesPresent, Detail: a.TopologicalAudit.Reason},
			{Name: "fails generation-awareness requirement", Passed: a.GenerationAwareness.Executed && a.GenerationAwareness.ColorDistinguishesQuarkLepton && !a.GenerationAwareness.ColorDistinguishesGenerations && !a.GenerationAwareness.ColorDistinguishesUpCharmTop && a.GenerationAwareness.WindingDistinguishesQuarkLepton && !a.GenerationAwareness.WindingDistinguishesGenerations && a.GenerationAwareness.CandidatesPassingGenerationAwareness == 0, Detail: a.GenerationAwareness.Reason},
			{Name: "rejects native deformation-coordinate map", Passed: a.DeformationMap.Executed && a.DeformationMap.TopologicalStressNative && !a.DeformationMap.DeltaAlphaMapNative && !a.DeformationMap.DeltaPhiMapNative && !a.DeformationMap.NumericCoordinateMapNative && a.DeformationMap.TraceCompatible && a.DeformationMap.CKMPMNSIndependent && a.DeformationMap.AllZeroDistance == 0, Detail: a.DeformationMap.Reason},
			{Name: "preserves topological bridge slot", Passed: a.BridgeSlot.Executed && a.BridgeSlot.RequiresAirlock && a.BridgeSlot.RequiresProvenance && a.BridgeSlot.RequiresBranchTags && a.BridgeSlot.RequiresUncertainty && a.BridgeSlot.AllowsTopologicalLabels && a.BridgeSlot.RejectsCKMPMNSAsInput && a.BridgeSlot.RejectsNativePromotion && a.BridgeSlot.CanComputeSyntheticResidual && !a.BridgeSlot.CanComputePhysicalResidual, Detail: a.BridgeSlot.Reason},
			{Name: "preserves 13-moduli firewall", Passed: a.Firewall.Executed && !a.Firewall.ObservedMassImported && !a.Firewall.CKMImported && !a.Firewall.PMNSImported && a.Firewall.VacuumIKNativeBaseline && !a.Firewall.TopologicalNativeSourceFound && !a.Firewall.SectorPerturbationsNative && !a.Firewall.PhysicalDUDComputed && !a.Firewall.PhysicalDENuComputed && !a.Firewall.CKMMatrixConstructed && !a.Firewall.PMNSMatrixConstructed && !a.Firewall.NativeRegistryWritten && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: a.Firewall.Reason},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusAuditCompleted, StatusInheritedNullPerturbationFrontier, StatusTopologicalSectorSeparationAudited, StatusQuarkLeptonSeparationOnly, StatusFailedNativeTopologicalSourceAbsent, StatusFailedColorWindingGenerationBlind, StatusFailedHolonomyNoDeltaMap, StatusFailedBettiLedgerNotPresent, StatusFailedSingleElectronNotFiniteOperator, StatusFailedGaugeRepresentationGenerationBlind, StatusFailedYukawaEnvironmental, StatusFailedCKMPMNSRejected, StatusBridgeSlotPreserved, StatusFailedNativePromotion, StatusFirewallPreserved, a.Truth}}
	}}
}
