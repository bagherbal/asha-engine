package generation2etarecordtransferranktraceobstructionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2EtaRecordTransferRankTraceObstructionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 eta-record transfer rank/trace obstruction audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate559 eta-record transfer audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit sealed Gate 558 A_eta_rec=span{I,eta} with 2+2 source ranks", Passed: a.Inherited.AlgebraConstructed && a.Inherited.SourcePlusRank == 2 && a.Inherited.SourceMinusRank == 2 && a.Inherited.TauEtaTraceValuesOnly && a.Inherited.NoPreviousTransferFunctor, Detail: FormatInherited(a.Inherited)},
			{Name: "classify formal unital representations on a 3D target", Passed: a.FormalReps.UnitalRepresentationsExist && a.FormalReps.EquivalentToComplementaryIdempotents && a.FormalReps.Exhaustive && len(a.FormalReps.RankSplits) == 4 && a.FormalReps.AnyTwoPlusOneFormal && !a.FormalReps.AnyCanonicalTwoPlusOne, Detail: FormatFormalReps(a.FormalReps)},
			{Name: "reject canonical 2+1 choice on W_spatial or generation carrier", Passed: a.CanonicalChoice.WSpatialCarrierAvailable && a.CanonicalChoice.GenerationCarrierCapacityVisible && !a.CanonicalChoice.NativeBasisIndependentReasonFor2Plus1 && !a.CanonicalChoice.NativeReasonForU12 && !a.CanonicalChoice.NativeReasonForGenerationLabels, Detail: FormatCanonicalChoice(a.CanonicalChoice)},
			{Name: "obstruct ordinary trace/rank-preserving transfer from 2+2 to dimension 3", Passed: !a.TraceRank.OrdinaryTracePreservingPossible && a.TraceRank.TargetCarrierDimension == 3, Detail: FormatTraceRank(a.TraceRank)},
			{Name: "obstruct normalized-trace-preserving transfer to dimension 3", Passed: !a.NormalizedTrace.IntegralRanksPossible && len(a.NormalizedTrace.RequiredTargetRanks) == 2 && a.NormalizedTrace.RequiredTargetRanks[0] == 1.5 && a.NormalizedTrace.RequiredTargetRanks[1] == 1.5, Detail: FormatNormalizedTrace(a.NormalizedTrace)},
			{Name: "allow formal B-L commutation but reject B-L canonicalization", Passed: a.BL.AnyFormalTransferCommutesWithBL && !a.BL.BLSuppliesRankSplit && !a.BL.BLSuppliesBasisLabels && !a.BL.BLSuppliesCanonicalU12, Detail: FormatBL(a.BL)},
			{Name: "mark spectral-triple compatibility unavailable without candidate transfer", Passed: !a.SpectralTriple.CandidateTransferExists && !a.SpectralTriple.CompatibilityPassed && !a.SpectralTriple.GradingCheckAvailable && !a.SpectralTriple.JCheckAvailable && !a.SpectralTriple.DCheckAvailable && !a.SpectralTriple.FirstOrderCheckAvailable, Detail: FormatSpectralTriple(a.SpectralTriple)},
			{Name: "reject native generation-carrier functor", Passed: a.Generation.FormalDim3GenerationCapacityVisible && !a.Generation.NativeBasisIndependentLabels && !a.Generation.FunctorFromAetaRec && !a.Generation.ProducesGenerationHierarchy && !a.Generation.ProducesYukawaOrCKMPMNS, Detail: FormatGeneration(a.Generation)},
			{Name: "preserve all physical-identification firewalls", Passed: a.Firewall.Preserved && !a.Firewall.WeakPlaneSelectionClaimed && !a.Firewall.WeakIsospinIdentificationClaimed && !a.Firewall.HiggsRadialGoldstoneClaimed && !a.Firewall.GenerationHierarchyClaimed && !a.Firewall.YukawaTextureClaimed && !a.Firewall.CKMPMNSClaimed && !a.Firewall.ObservedFlavorImported, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(Statuses(), a.Truth, a.Final.MissingNextTheorem)}
	}}
}
