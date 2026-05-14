package neutraleigenspacekernel

import "github.com/bagherbal/asha-engine/pkg/theorem"

func NeutralEigenspaceKernelInvariant3PlaneAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-NEUTRAL-EIGENSPACE-KERNEL-INVARIANT-3PLANE-AUDIT"
	const name = "Neutral Eigenspace Kernel / Invariant 3-Plane Isomorphism Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 249 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 248 8_v obstruction inherited", Passed: a.PreviousGate248.Basis8VKnown && a.PreviousGate248.ScalarTraceOriginKnown && a.PreviousGate248.DimensionallyEmbeddable && !a.PreviousGate248.ScalarTo8VMapDerived && !a.PreviousGate248.VTauConstructed && !a.PreviousGate248.TrialityUnblocked, Detail: FormatInherited(a.PreviousGate248)},
			{Name: "native 8_v carrier is available for a neutral-kernel strategy", Passed: a.VectorCarrier.NativeCarrierKnown && a.VectorCarrier.Dimension == 8 && len(a.VectorCarrier.BasisLabels) == 8 && a.VectorCarrier.RealOctonionicSplitKnown, Detail: FormatVectorCarrier(a.VectorCarrier)},
			{Name: "Q/Z derivation action on 8_v is not derived", Passed: a.EWAction.QKnownAsScalarObservable && a.EWAction.ZKnownAsScalarObservable && !a.EWAction.QMatrixOn8VDerived && !a.EWAction.ZMatrixOn8VDerived && !a.EWAction.ChargeSpectrumKnown && a.EWAction.ManualRepresentationRejected && a.EWAction.Obstruction != "", Detail: FormatEWAction(a.EWAction)},
			{Name: "neutral kernel cannot be computed without Q_8v", Passed: !a.NeutralKernel.Computed && !a.NeutralKernel.DimensionKnown && a.NeutralKernel.Dimension == -1 && !a.NeutralKernel.ExactlyThreeDimensional && !a.NeutralKernel.InvariantSubspaceDerived && a.NeutralKernel.DependsOnMissingQMatrix, Detail: FormatNeutralKernel(a.NeutralKernel)},
			{Name: "scalar trace triple cannot be paired with an uncomputed kernel", Passed: a.ScalarPlane.ScalarSlotCount == 3 && a.ScalarPlane.NeutralKernelDimension == -1 && !a.ScalarPlane.DimensionMatch && !a.ScalarPlane.CanonicalIsomorphism && !a.ScalarPlane.BasisIndependentPairing && !a.ScalarPlane.QZTYToNeutralBasisDerived && a.ScalarPlane.Obstruction != "", Detail: FormatScalarPlane(a.ScalarPlane)},
			{Name: "v_tau is not constructed from a neutral 3-plane", Passed: !a.VTau.Constructed && !a.VTau.LawfulRepresentative && a.VTau.WouldHaveNormSquared == 9 && !a.VTau.WouldFeedTriality && a.VTau.RejectedBecause != "", Detail: FormatVTau(a.VTau)},
			{Name: "triality and Yukawa texture remain blocked", Passed: a.Triality.Requires8VVector && !a.Triality.Neutral8VVectorAvailable && !a.Triality.TrialityCanRun && !a.Triality.DiagonalTextureConstructed && a.Triality.GenerationBreakingCapacity && a.Triality.NonCommutingTextureCapacity && !a.Triality.CKMPMNSDerived && !a.Triality.FermionMassesDerived, Detail: FormatTriality(a.Triality)},
			{Name: "firewall preserved: no neutral kernel or v_tau forced", Passed: !a.Firewall.InventedQActionOn8V && !a.Firewall.ForcedNeutralKernelDim3 && !a.Firewall.AssignedScalarSlotsToBasis && !a.Firewall.ConstructedVTauByHand && !a.Firewall.InventedTrialityMatrices && !a.Firewall.InsertedYukawaTexture && !a.Firewall.ImportedObservedMasses && !a.Firewall.ImportedCKMPMNS && !a.Firewall.ClaimedFiniteFlavorTheorem && !a.Firewall.PollutedFiniteCore, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records a well-typed strategy but no derived 3-plane", Passed: a.Summary.Basis8VKnown && !a.Summary.EWDerivationActionDerived && !a.Summary.NeutralKernelDerived && !a.Summary.NeutralKernelDim3 && !a.Summary.ScalarNeutralIsomorphism && !a.Summary.VTauConstructed && !a.Summary.TrialityUnblocked && !a.Summary.YukawaTextureDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}
		notes := []string{
			"Gate 249 tests a coordinate-free route to v_tau: derive ker(Q_8v) inside the Spin(8) vector carrier and identify it with the three neutral scalar trace slots.",
			"The strategy is mathematically well typed, but the project does not yet derive Q_8v or Z_8v matrices on 8_v, so the neutral kernel is not computable.",
			"Therefore v_tau, Spin(8) triality pullback, and Yukawa texture derivation remain blocked without hand-selecting a basis.",
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
