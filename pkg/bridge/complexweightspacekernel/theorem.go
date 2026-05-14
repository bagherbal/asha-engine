package complexweightspacekernel

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ComplexWeightSpaceDecompositionNeutralKernelAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-COMPLEX-WEIGHT-SPACE-8VC-NEUTRAL-KERNEL-AUDIT"
	const name = "Complex Weight-Space Decomposition / 8vC Neutral Kernel Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 251 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 250 real-bivector obstruction inherited", Passed: a.PreviousGate250.Carrier8VKnown && a.PreviousGate250.CliffordAdjointAvailable && a.PreviousGate250.CandidateMatricesComputable && !a.PreviousGate250.EWBivectorsRetrieved && !a.PreviousGate250.Q8VMatrixDerived && !a.PreviousGate250.RealBivector3KernelPossible && !a.PreviousGate250.VTauConstructed, Detail: FormatInherited(a.PreviousGate250)},
			{Name: "8v complexification is lawful and native", Passed: a.ComplexCarrier.ComplexificationNative && a.ComplexCarrier.RealDimension == 8 && a.ComplexCarrier.ComplexDimension == 8 && a.ComplexCarrier.UnderlyingRealDimension == 16 && a.ComplexCarrier.EvenRankObstructionLift, Detail: FormatComplexification(a.ComplexCarrier)},
			{Name: "Hermitian weight-space route permits odd multiplicities", Passed: a.Hermitian.RealSkewGeneratorAvailable && a.Hermitian.HermitianOperatorsHaveRealSpectrum && a.Hermitian.OddWeightSpacesAllowed && a.Hermitian.CandidateSimpleBladeKernelComplexDim == 6 && !a.Hermitian.PhysicalQHermitianDerived && !a.Hermitian.PhysicalZHermitianDerived, Detail: FormatHermitian(a.Hermitian)},
			{Name: "native Hermitian Q8vC/Z8vC matrices remain unavailable", Passed: !a.Cartan.CartanCommutingPairDerived && !a.Cartan.Q8vCMatrixDerived && !a.Cartan.Z8vCMatrixDerived && !a.Cartan.SimultaneouslyDiagonal && !a.Cartan.WeightSpectrumDerived && a.Cartan.ManualChargeAssignmentRejected, Detail: FormatCartan(a.Cartan)},
			{Name: "complex neutral 3-plane is capacity only, not derived", Passed: !a.NeutralKernel.Computed && !a.NeutralKernel.DimensionKnown && !a.NeutralKernel.ExactlyThreeComplexDim && a.NeutralKernel.OddDimAllowedInPrinciple && a.NeutralKernel.DependsOnMissingQ8vC, Detail: FormatNeutralKernel(a.NeutralKernel)},
			{Name: "complex triality arena exists but no canonical J-compatible map", Passed: a.Triality.Spin8TrialityOverC && a.Triality.SameComplexDimension && a.Triality.OuterAutomorphismRequired && !a.Triality.CanonicalUntwistedIsomorphism && !a.Triality.MapNeutralKernelToSpinor && !a.Triality.RealStructureCompatibilityChecked && !a.Triality.CompatibleWithJ, Detail: FormatTriality(a.Triality)},
			{Name: "v_tau and Yukawa texture remain blocked", Passed: a.VTau.NeedsNeutral3Plane && !a.VTau.Neutral3PlaneAvailable && a.VTau.NeedsScalarSlotFrame && !a.VTau.ScalarSlotFrameDerived && !a.VTau.Constructed && !a.VTau.WouldFeedTriality, Detail: FormatVTau(a.VTau)},
			{Name: "firewall preserved: no complex weights or triality map invented", Passed: !a.Firewall.InventedQ8vC && !a.Firewall.InventedZ8vC && !a.Firewall.AssignedComplexWeightsByHand && !a.Firewall.ForcedKernelDim3 && !a.Firewall.InventedTrialityIsomorphism && !a.Firewall.IgnoredRealStructure && !a.Firewall.ConstructedVTauByHand && !a.Firewall.InsertedYukawaTexture && !a.Firewall.ClaimedCKMPMNS && !a.Firewall.PollutedFiniteCore, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records complex route capacity with explicit obstruction", Passed: a.Summary.Complex8VKnown && a.Summary.HermitianWeightCapacity && a.Summary.OddComplexKernelCapacity && !a.Summary.NativeHermitianMatrices && !a.Summary.ComplexNeutralKernelDerived && !a.Summary.NeutralKernelDim3 && a.Summary.ComplexTrialityArena && !a.Summary.CanonicalTrialityMap && !a.Summary.RealStructureCompatible && !a.Summary.VTauConstructed && !a.Summary.TrialityUnblocked && !a.Summary.YukawaTextureDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}
		notes := []string{
			"Gate 251 confirms the correct pivot from real skew-adjoint kernels to complex Hermitian weight spaces.",
			"Complexification removes the even-rank obstruction only in principle; it does not derive Q_8vC, Z_8vC, or their simultaneous weight decomposition.",
			"Complex Spin(8) triality is treated as an outer-automorphism arena, not as a canonical vector-to-spinor type-cast; real-structure compatibility remains a separate unproven theorem.",
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
