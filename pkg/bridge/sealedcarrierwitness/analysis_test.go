package sealedcarrierwitness

import "testing"

func TestGate257SealedCarrierWitnessAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Summary.Gate256SealInherited || !a.PreviousGate256.SpontaneousSealRecorded {
		t.Fatalf("expected Gate 256 seal inheritance: %s", FormatSummary(a.Summary))
	}
	if !a.Charges.ChargeEigenvalueTableDerived || a.Charges.ExternalChargeInputUsed {
		t.Fatalf("expected native charge extraction without external input: %s", FormatCharges(a.Charges))
	}
	if a.Embedding.WeakFrameCount != 12 || a.Embedding.ScalarEmbeddingCount != 8 || a.Embedding.TotalCombinedWitnesses != 96 {
		t.Fatalf("unexpected witness inventory: %s", FormatEmbedding(a.Embedding))
	}
	if !a.SO8.AllTranslated || a.SO8.WitnessCount != 96 {
		t.Fatalf("expected all witnesses translated to so(8): %s", FormatSO8(a.SO8))
	}
	if !a.TrialityScan.AllBranchesScanned || a.TrialityScan.ResultCount != 288 || a.TrialityScan.BranchCount != 3 {
		t.Fatalf("expected all branch scan: %s", FormatTrialityScan(a.TrialityScan))
	}
	if a.TrialityScan.ExactPolarized3PlaneResults != 0 || a.TrialityScan.ExactFull3KernelResults != 0 {
		t.Fatalf("Gate 257 should not force a three-plane: %s", FormatTrialityScan(a.TrialityScan))
	}
	if a.TrialityScan.MaxPolarizedKernelComplexDim != 2 || a.TrialityScan.MaxFullQ8vCKernelComplexDim != 4 {
		t.Fatalf("unexpected kernel maxima: %s", FormatTrialityScan(a.TrialityScan))
	}
	if !a.TrialityScan.YOnly.WouldGivePolarizedThreeSlot || !a.TrialityScan.YOnly.RejectedBecauseMissingT3L {
		t.Fatalf("expected scalar-only diagnostic to be rejected: %s", FormatYOnly(a.TrialityScan.YOnly))
	}
	if a.Firewall.ForcedWeakPlane || a.Firewall.SelectedTrialityByHand || a.Firewall.ForcedKernelDim3 || a.Firewall.AcceptedYOnlyAsQ || a.Firewall.PollutedFiniteCore {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}
