package generation2etarecordtransferranktraceobstructionaudit

import "testing"

func TestGate559EtaRecordTransferRankTraceObstructionAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.AlgebraConstructed || a.Inherited.SourcePlusRank != 2 || a.Inherited.SourceMinusRank != 2 || !a.Inherited.NoPreviousTransferFunctor {
		t.Fatalf("Gate558 inheritance failed: %s", FormatInherited(a.Inherited))
	}
	if !a.FormalReps.UnitalRepresentationsExist || !a.FormalReps.Exhaustive || len(a.FormalReps.RankSplits) != 4 || !a.FormalReps.AnyTwoPlusOneFormal || a.FormalReps.AnyCanonicalTwoPlusOne {
		t.Fatalf("formal representation classification failed: %s", FormatFormalReps(a.FormalReps))
	}
	if a.CanonicalChoice.NativeBasisIndependentReasonFor2Plus1 || a.CanonicalChoice.NativeReasonForU12 || a.CanonicalChoice.NativeReasonForGenerationLabels {
		t.Fatalf("canonical choice unexpectedly available: %s", FormatCanonicalChoice(a.CanonicalChoice))
	}
	if a.TraceRank.OrdinaryTracePreservingPossible || a.NormalizedTrace.IntegralRanksPossible {
		t.Fatalf("trace/rank preservation should be obstructed: %s | %s", FormatTraceRank(a.TraceRank), FormatNormalizedTrace(a.NormalizedTrace))
	}
	if !a.BL.AnyFormalTransferCommutesWithBL || a.BL.BLSuppliesRankSplit || a.BL.BLSuppliesBasisLabels || a.BL.BLSuppliesCanonicalU12 {
		t.Fatalf("B-L audit failed: %s", FormatBL(a.BL))
	}
	if a.SpectralTriple.CandidateTransferExists || a.SpectralTriple.CompatibilityPassed || a.SpectralTriple.FirstOrderCheckAvailable {
		t.Fatalf("spectral triple compatibility should be unavailable: %s", FormatSpectralTriple(a.SpectralTriple))
	}
	if a.Generation.FunctorFromAetaRec || a.Generation.NativeBasisIndependentLabels || a.Generation.ProducesGenerationHierarchy || a.Generation.ProducesYukawaOrCKMPMNS {
		t.Fatalf("generation firewall failed: %s", FormatGeneration(a.Generation))
	}
	if !a.Firewall.Preserved || a.Firewall.WeakPlaneSelectionClaimed || a.Firewall.HiggsRadialGoldstoneClaimed || a.Firewall.YukawaTextureClaimed || a.Firewall.CKMPMNSClaimed {
		t.Fatalf("physical firewall failed: %s", FormatFirewall(a.Firewall))
	}
	if !a.Final.FormalRepresentationsExist || a.Final.CanonicalInASHA || a.Final.TraceRankPreservingTransfer || a.Final.BMinusLCanonicalizesTransfer || a.Final.LawfulTransferAvailable {
		t.Fatalf("final verdict failed: %s", FormatFinal(a.Final))
	}
}

func TestGate559Theorem(t *testing.T) {
	res := Generation2EtaRecordTransferRankTraceObstructionAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
